package handlers_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/placer14/gof-server/internal/handlers"
	"github.com/placer14/gof-server/internal/provider"
	"github.com/stretchr/testify/assert"
)

type respValueType struct {
	Value any `json:"value"`
}

func TestGetStringValueHandler(t *testing.T) {
	m := provider.NewProviderMock()
	m.SetString("hello", "world")
	ctx := context.Background()
	ctx = context.WithValue(ctx, provider.KeyFlagStore, m)
	req := httptest.NewRequestWithContext(ctx, http.MethodGet, "/string/hello", nil)
	w := httptest.NewRecorder()

	handlers.GetStringValue(w, req)
	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	respValue := &respValueType{}
	assert.NoError(t, json.Unmarshal(body, respValue))
	assert.Equal(t, "world", respValue.Value, string(body))
}
