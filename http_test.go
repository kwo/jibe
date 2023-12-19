package jibe

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	handler1 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if GetID(r.Context()) == "" {
			http.Error(w, "no id", http.StatusInternalServerError)
		}
	})

	handler := HTTPHandler(NewID, handler1)
	httpd := httptest.NewServer(handler)

	req, err := http.NewRequestWithContext(context.Background(), "GET", httpd.URL, nil)
	if err != nil {
		t.Error(err)
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	_ = rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		t.Errorf("rsp.StatusCode != http.StatusOK; rsp.StatusCode == %d", rsp.StatusCode)
	}
}
