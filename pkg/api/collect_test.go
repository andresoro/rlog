package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/andresoro/rlog/pkg/model"
)

func TestCollect(t *testing.T) {

	t.Run("test good url", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/collect/123/path/test", nil)
		if err != nil {
			t.Fatal(err)
		}

		a := New(&MockStore{})
		a.Routes()

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.Collect)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Error(rr.Code)
		}

		log.Println(rr.Body.String())
	})

	t.Run("test bad URL", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/collect/", nil)
		if err != nil {
			t.Fatal(err)
		}

		a := New(&MockStore{})
		a.Routes()

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.Collect)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Error()
		}

		log.Println(rr.Body.String())
	})

	t.Run("test only siteID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/collect/123/", nil)
		if err != nil {
			t.Fatal(err)
		}

		a := New(&MockStore{})
		a.Routes()

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.Collect)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Error()
		}

		log.Println(rr.Body.String())
	})
}

type MockStore struct{}

func (m *MockStore) InsertEvent(e *model.Event) error             { return nil }
func (m *MockStore) RetrieveEvent(id int64) (*model.Event, error) { return nil, nil }
func (m *MockStore) RetrieveSiteStats(start, end time.Time, siteID int64) (*model.SiteStats, error) {
	return nil, nil
}
func (m *MockStore) RetrieveAll(siteID int64) ([]*model.Event, error) { return nil, nil }
