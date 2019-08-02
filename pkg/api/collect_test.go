package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCollect(t *testing.T) {
	req, err := http.NewRequest("GET", "/collect/url/path/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	api := &API{}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.Collect)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Error()
	}

	log.Println(rr.Body.String())
}
