package test_http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

const (
	PORT          = "23456"
	IMAGE_NAME    = "gof-server"
	IMAGE_VERSION = "latest"
)

func TestPingEndpoint(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        fmt.Sprintf("%s:%s", IMAGE_NAME, IMAGE_VERSION),
		ExposedPorts: []string{fmt.Sprintf("%s/tcp", PORT)},
	}

	container, err := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Error(err)
		}
	}()

	ip, err := container.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	mappedPort, err := container.MappedPort(ctx, PORT)
	if err != nil {
		t.Fatal(err)
	}

	url := "http://" + ip + ":" + mappedPort.Port() + "/ping"

	client := &http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "pong", string(body))
}
