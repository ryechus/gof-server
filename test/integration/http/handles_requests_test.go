package test_http

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/testcontainers/testcontainers-go"
// )

// const (
// 	PORT          = "23456"
// 	IMAGE_NAME    = "gof-server"
// 	IMAGE_VERSION = "latest"
// )

// type respValueType struct {
// 	Value any `json:"value"`
// }

// func TestEndpoints(t *testing.T) {
// 	ctx := context.Background()

// 	req := testcontainers.ContainerRequest{
// 		Image:        fmt.Sprintf("%s:%s", IMAGE_NAME, IMAGE_VERSION),
// 		ExposedPorts: []string{fmt.Sprintf("%s/tcp", PORT)},
// 	}

// 	container, err := testcontainers.GenericContainer(ctx,
// 		testcontainers.GenericContainerRequest{
// 			ContainerRequest: req,
// 			Started:          true,
// 		},
// 	)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer func() {
// 		if err := container.Terminate(ctx); err != nil {
// 			t.Error(err)
// 		}
// 	}()

// 	ip, err := container.Host(ctx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	mappedPort, err := container.MappedPort(ctx, PORT)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	type ex struct {
// 		name           string
// 		endpoint       string
// 		expectedResult any
// 	}

// 	examples := []ex{
// 		{
// 			name:           "valid bool lookup",
// 			endpoint:       "/bool/grant_soil_access",
// 			expectedResult: false,
// 		},
// 		{
// 			name:           "valid string lookup",
// 			endpoint:       "/string/dataplane_generation",
// 			expectedResult: "metal.v1",
// 		},
// 		{
// 			name:           "valid float lookup",
// 			endpoint:       "/float/special_ability_buff_perc",
// 			expectedResult: 0.23456,
// 		},
// 		{
// 			name:           "valid int lookup",
// 			endpoint:       "/int/num_of_special_abilities",
// 			expectedResult: float64(12), // json.Unmarshal is turning the int into a float64. why?
// 		},
// 		{
// 			name:           "ping works",
// 			endpoint:       "/ping",
// 			expectedResult: "pong",
// 		},
// 	}
// 	for _, ex := range examples {
// 		t.Run(ex.name, func(t *testing.T) {
// 			// http://host:port/string/<keyname>
// 			url := "http://" + ip + ":" + mappedPort.Port() + ex.endpoint

// 			client := &http.Client{Timeout: 1 * time.Second}
// 			resp, err := client.Get(url)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			defer resp.Body.Close()

// 			body, err := io.ReadAll(resp.Body)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			respValue := &respValueType{}
// 			assert.NoError(t, json.Unmarshal(body, respValue))

// 			switch actual := respValue.Value.(type) {
// 			case string:
// 				assert.Equal(t, actual, ex.expectedResult.(string))
// 			case int64:
// 				assert.Equal(t, actual, ex.expectedResult.(int64))
// 			case float64:
// 				assert.Equal(t, actual, ex.expectedResult.(float64))
// 			case bool:
// 				assert.Equal(t, actual, ex.expectedResult.(bool))
// 			default:
// 				t.Error("unexpected type")
// 			}
// 		})
// 	}
// }
