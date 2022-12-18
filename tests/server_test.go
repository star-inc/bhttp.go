// Butterfly HTTP - The easiest HTTP client for GoLang from Butterfly.
// Copyright(c) 2022 Star Inc. All Rights Reserved.
// The software licensed under Mozilla Public License Version 2.0

package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func Test_HttpServer(t *testing.T) {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		message, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}
		_, err = fmt.Fprintln(w, string(message))
		if err != nil {
			t.Error(err)
		}
	})

	t.Fatal(http.ListenAndServe(":8080", nil))
}
