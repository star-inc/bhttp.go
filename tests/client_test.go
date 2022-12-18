// Butterfly HTTP - The easiest HTTP client for GoLang from Butterfly.
// Copyright(c) 2022 Star Inc. All Rights Reserved.
// The software licensed under Mozilla Public License Version 2.0

package tests

import (
	"bytes"
	"encoding/json"
	"github.com/star-inc/bhttp.go"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

const testEndPoint = "http://localhost:8080"

func Test_HttpGet(t *testing.T) {
	client := bHttp.NewHttpClient(testEndPoint)
	if status, response := client.GET("/echo"); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPost(t *testing.T) {
	client := bHttp.NewHttpClient(testEndPoint)
	if status, response := client.POST("/echo", nil); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPostWithBodyForm(t *testing.T) {
	client := bHttp.NewHttpClient(testEndPoint)
	formData := url.Values{
		"method": []string{"post"},
		"from":   []string{"test@butterfly"},
	}
	if status, response := client.POST("/echo", strings.NewReader(formData.Encode())); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPostWithBodyJSON(t *testing.T) {
	client := bHttp.NewHttpClient(testEndPoint)
	data, err := json.Marshal(map[string]string{"method": "post", "from": "test@butterfly"})
	if err != nil {
		t.Error(err)
	}
	if status, response := client.POST("/echo", bytes.NewReader(data)); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPut(t *testing.T) {
	client := bHttp.NewHttpClient(testEndPoint)
	if status, response := client.PUT("/echo", nil); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPutWithBodyForm(t *testing.T) {
	client := bHttp.NewHttpClient(testEndPoint)
	formData := url.Values{
		"method": []string{"put"},
		"from":   []string{"test@butterfly"},
	}
	if status, response := client.PUT("/echo", strings.NewReader(formData.Encode())); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPutWithBodyJSON(t *testing.T) {
	client := bHttp.NewHttpClient(testEndPoint)
	data, err := json.Marshal(map[string]string{"method": "put", "from": "test@butterfly"})
	if err != nil {
		t.Error(err)
	}
	if status, response := client.PUT("/echo", bytes.NewReader(data)); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpDelete(t *testing.T) {
	client := bHttp.NewHttpClient(testEndPoint)
	if status, response := client.DELETE("/echo", nil); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPatch(t *testing.T) {
	client := bHttp.NewHttpClient(testEndPoint)
	if status, response := client.PATCH("/echo", nil); status.IsEqual(http.StatusOK) {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}
