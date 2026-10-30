// bhttp.go
// (c) 2023 Star Inc.

package bHttp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

var testEndPoint string

func TestMain(m *testing.M) {
	mux := http.NewServeMux()
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = io.Copy(io.Discard, r.Body)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}
	mux.HandleFunc("/get", handler)
	mux.HandleFunc("/post", handler)
	mux.HandleFunc("/put", handler)
	mux.HandleFunc("/delete", handler)
	mux.HandleFunc("/patch", handler)

	ts := httptest.NewServer(mux)
	testEndPoint = ts.URL

	code := m.Run()
	ts.Close()
	os.Exit(code)
}

func Test_HttpGet(t *testing.T) {
	client, err := NewHttpClient(testEndPoint)
	if err != nil {
		t.Fatal(err)
	}
	status, response, err := client.GET("/get")
	if err != nil {
		t.Fatal(err)
	}
	if status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPost(t *testing.T) {
	client, err := NewHttpClient(testEndPoint)
	if err != nil {
		t.Fatal(err)
	}
	status, response, err := client.POST("/post", nil)
	if err != nil {
		t.Fatal(err)
	}
	if status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPostWithBodyForm(t *testing.T) {
	client, err := NewHttpClient(testEndPoint)
	if err != nil {
		t.Fatal(err)
	}
	formData := url.Values{
		"method": []string{"post"},
		"from":   []string{"test@butterfly"},
	}
	status, response, err := client.POST("/post", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	if status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPostWithBodyJSON(t *testing.T) {
	client, err := NewHttpClient(testEndPoint)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(map[string]string{"method": "post", "from": "test@butterfly"})
	if err != nil {
		t.Error(err)
	}
	status, response, err := client.POST("/post", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPut(t *testing.T) {
	client, err := NewHttpClient(testEndPoint)
	if err != nil {
		t.Fatal(err)
	}
	status, response, err := client.PUT("/put", nil)
	if err != nil {
		t.Fatal(err)
	}
	if status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPutWithBodyForm(t *testing.T) {
	client, err := NewHttpClient(testEndPoint)
	if err != nil {
		t.Fatal(err)
	}
	formData := url.Values{
		"method": []string{"put"},
		"from":   []string{"test@butterfly"},
	}
	status, response, err := client.PUT("/put", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	if status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPutWithBodyJSON(t *testing.T) {
	client, err := NewHttpClient(testEndPoint)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(map[string]string{"method": "put", "from": "test@butterfly"})
	if err != nil {
		t.Error(err)
	}
	status, response, err := client.PUT("/put", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpDelete(t *testing.T) {
	client, err := NewHttpClient(testEndPoint)
	if err != nil {
		t.Fatal(err)
	}
	status, response, err := client.DELETE("/delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	if status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPatch(t *testing.T) {
	client, err := NewHttpClient(testEndPoint)
	if err != nil {
		t.Fatal(err)
	}
	status, response, err := client.PATCH("/patch", nil)
	if err != nil {
		t.Fatal(err)
	}
	if status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_NewHttpClient_InvalidURL(t *testing.T) {
	client, err := NewHttpClient(":")
	if err == nil {
		t.Errorf("expected error for invalid URL, got nil (client: %v)", client)
	}
}

func Test_SetBaseURL_InvalidURL(t *testing.T) {
	client := new(Client).Initialize()
	err := client.SetBaseURL(":")
	if err == nil {
		t.Errorf("expected error for invalid URL, got nil")
	}
}

func Test_Do_InvalidURI(t *testing.T) {
	client, err := NewHttpClient(testEndPoint)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = client.GET(":")
	if err == nil {
		t.Errorf("expected error for invalid URI, got nil")
	}
}

func Test_Do_InvalidMethod(t *testing.T) {
	client, err := NewHttpClient(testEndPoint)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = client.Do("INVALID METHOD", "/get", nil)
	if err == nil {
		t.Errorf("expected error for invalid method, got nil")
	}
}

func Test_LocalServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}))
	defer ts.Close()

	client, err := NewHttpClient(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	code, body, err := client.GET("/test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !code.IsEqual(http.StatusOK) {
		t.Errorf("expected 200, got %d", code)
	}
	if string(body) != "ok" {
		t.Errorf("expected 'ok', got %s", body)
	}
}
