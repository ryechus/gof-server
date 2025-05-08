package handlers_test

import (
	"testing"
)

// type respValueType struct {
// 	Value any `json:"value"`
// }

func TestGetFlagWithVariations(t *testing.T) {

}

func TestEvaluateFlag(t *testing.T) {

}

func TestUpdateFlag(t *testing.T) {

}

func TestCreateFlag(t *testing.T) {

}

func TestGetFlagVariations(t *testing.T) {

}

func TestUpdateFlagVariation(t *testing.T) {

}

// func TestGetStringValueHandler(t *testing.T) {
// 	m := storage.NewInMemoryStorage()
// 	m.SetString("hello", "world")
// 	ctx := context.Background()
// 	ctx = context.WithValue(ctx, storage.KeyInMemoryStorage, m)
// 	req := httptest.NewRequestWithContext(ctx, http.MethodGet, "/string/{flagKey}", nil)
// 	req.SetPathValue("flagKey", "hello")
// 	w := httptest.NewRecorder()

// 	handlers.GetStringValue(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	respValue := &respValueType{}
// 	assert.NoError(t, json.Unmarshal(body, respValue))
// 	assert.Equal(t, "world", respValue.Value, string(body))
// }

// func TestGetBoolValueHandler(t *testing.T) {
// 	m := storage.NewInMemoryStorage()
// 	flagKey := "hello"
// 	flagValue := true
// 	m.SetBool(flagKey, flagValue)
// 	ctx := context.Background()
// 	ctx = context.WithValue(ctx, storage.KeyInMemoryStorage, m)
// 	req := httptest.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/bool/%s", flagKey), nil)
// 	req.SetPathValue("flagKey", "hello")
// 	w := httptest.NewRecorder()

// 	handlers.GetBoolValue(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	respValue := &respValueType{}
// 	assert.NoError(t, json.Unmarshal(body, respValue))
// 	assert.Equal(t, flagValue, respValue.Value, string(body))
// }

// func TestGetIntValueHandler(t *testing.T) {
// 	m := storage.NewInMemoryStorage()
// 	flagKey := "hello"
// 	flagValue := int64(1_000_000)
// 	m.SetInt(flagKey, flagValue)
// 	ctx := context.Background()
// 	ctx = context.WithValue(ctx, storage.KeyInMemoryStorage, m)
// 	req := httptest.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/int/%s", flagKey), nil)
// 	req.SetPathValue("flagKey", "hello")
// 	w := httptest.NewRecorder()

// 	handlers.GetIntValue(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	respValue := &respValueType{}
// 	assert.NoError(t, json.Unmarshal(body, respValue))
// 	assert.Equal(t, float64(flagValue), respValue.Value, string(body))
// }

// func TestGetFloatValueHandler(t *testing.T) {
// 	m := storage.NewInMemoryStorage()
// 	flagKey := "hello"
// 	flagValue := 1.0891
// 	m.SetFloat(flagKey, flagValue)
// 	ctx := context.Background()
// 	ctx = context.WithValue(ctx, storage.KeyInMemoryStorage, m)
// 	req := httptest.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/int/%s", flagKey), nil)
// 	req.SetPathValue("flagKey", "hello")
// 	w := httptest.NewRecorder()

// 	handlers.GetFloatValue(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	respValue := &respValueType{}
// 	assert.NoError(t, json.Unmarshal(body, respValue))
// 	assert.Equal(t, float64(flagValue), respValue.Value, string(body))
// }
