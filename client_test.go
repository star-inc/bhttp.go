// Butterfly HTTP - The easiest HTTP client for GoLang from Butterfly.
// Copyright(c) 2022 Star Inc. All Rights Reserved.
// The software licensed under Mozilla Public License Version 2.0

package bHttp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

const testEndPoint = "https://httpbin.org"

func Test_HttpGet(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	if status, response := client.GET("/get"); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPost(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	if status, response := client.POST("/post", nil); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPostWithBodyForm(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	formData := url.Values{
		"method": []string{"post"},
		"from":   []string{"test@butterfly"},
	}
	if status, response := client.POST("/post", strings.NewReader(formData.Encode())); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPostWithBodyJSON(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	data, err := json.Marshal(map[string]string{"method": "post", "from": "test@butterfly"})
	if err != nil {
		t.Error(err)
	}
	if status, response := client.POST("/post", bytes.NewReader(data)); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPut(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	if status, response := client.PUT("/put", nil); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPutWithBodyForm(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	formData := url.Values{
		"method": []string{"put"},
		"from":   []string{"test@butterfly"},
	}
	if status, response := client.PUT("/put", strings.NewReader(formData.Encode())); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPutWithBodyJSON(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	data, err := json.Marshal(map[string]string{"method": "put", "from": "test@butterfly"})
	if err != nil {
		t.Error(err)
	}
	if status, response := client.PUT("/put", bytes.NewReader(data)); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpDelete(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	if status, response := client.DELETE("/delete", nil); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPatch(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	if status, response := client.PATCH("/patch", nil); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}
