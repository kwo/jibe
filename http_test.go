package jibe

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	handler1 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := GetID(r.Context())
		if id == "" {
			http.Error(w, "no id", http.StatusInternalServerError)
		}
		xRequestID := w.Header().Get(HeaderRequestID)
		if got, want := xRequestID, id; got != want {
			http.Error(w, "no request id", http.StatusInternalServerError)
		}
	})

	handler := WithResponseHeader(HeaderRequestID, handler1) // must be before WithRequestID
	handler = WithRequestID(newID, handler)
	httpd := httptest.NewServer(handler)

	req, err := http.NewRequestWithContext(context.Background(), "GET", httpd.URL, nil)
	if err != nil {
		t.Error(err)
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatal(err)
	}
	_ = rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		t.Errorf("status code: %d, message: %s", rsp.StatusCode, string(data))
	}
}

func newID() string {
	return "1234"
}
