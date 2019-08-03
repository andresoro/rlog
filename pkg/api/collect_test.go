package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCollect(t *testing.T) {

	t.Run("test good url", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/collect/ID/path/test", nil)
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
	})

	t.Run("test bad URL", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/collect/", nil)
		if err != nil {
			t.Fatal(err)
		}

		api := &API{}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Collect)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Error()
		}

		log.Println(rr.Body.String())
	})

	t.Run("test only siteID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/collect/ID/", nil)
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
	})
}
