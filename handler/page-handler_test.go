package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/y16i/backend-go/models"
	"github.com/stretchr/testify/assert"
)

type MockHandler struct {
	FindBySlugFunc func(slug string) (*models.Page, error)
}

func (m *MockHandler) FindBySlug(slug string) (*models.Page, error) {
	return m.FindBySlugFunc("some-slug")
}

func TestPageHandler(t *testing.T) {
	t.Run("can get page array", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/pages", nil)
		if err != nil {
			t.Fatal(err)
		}
		// add query param
		q := req.URL.Query()
		q.Add("slug", "portfolio-about")
		req.URL.RawQuery = q.Encode()

		expectedReturn := &models.Page{}
		repo := &MockHandler{
			FindBySlugFunc: func(slug string) (*models.Page, error) {
				return expectedReturn, nil
			},
		}
		handler := NewBaseHandler(repo)
		// Record the response
		rr := httptest.NewRecorder()
		handler.PageHandler(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("returned a wrong status: got %v, expected %v", status, http.StatusOK)
		}

		var pages [1]models.Page
		pages[0] = *expectedReturn
		marshaled, marshalErr := json.Marshal(pages)
		if marshalErr != nil {
			t.Fatal(marshalErr)
		}

		if rr.Body.String() != string(marshaled) {
			t.Errorf("handler returned a wrong body: got %v, expected %v", rr.Body.String(), string(marshaled))
		}
	})

	t.Run("no slug should return 404", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/pages", nil)
		if err != nil {
			t.Fatal(err)
		}

		handler := NewBaseHandler(nil)
		rr := httptest.NewRecorder()
		handler.PageHandler(rr, req)
		assert.Equal(t, 404, rr.Result().StatusCode, "should be 404")
	})

	t.Run("POST return 404", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/pages", nil)
		if err != nil {
			t.Fatal(err)
		}
		// add query param
		q := req.URL.Query()
		q.Add("slug", "portfolio-about")
		req.URL.RawQuery = q.Encode()

		handler := NewBaseHandler(nil)
		rr := httptest.NewRecorder()
		handler.PageHandler(rr, req)
		assert.Equal(t, 404, rr.Result().StatusCode, "should be 404")
	})

	t.Run("PUT return 404", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/pages", nil)
		if err != nil {
			t.Fatal(err)
		}
		// add query param
		q := req.URL.Query()
		q.Add("slug", "portfolio-about")
		req.URL.RawQuery = q.Encode()

		handler := NewBaseHandler(nil)
		rr := httptest.NewRecorder()
		handler.PageHandler(rr, req)
		assert.Equal(t, 404, rr.Result().StatusCode, "should be 404")
	})

	t.Run("PATCH return 404", func(t *testing.T) {
		req, err := http.NewRequest("PATCH", "/pages", nil)
		if err != nil {
			t.Fatal(err)
		}
		// add query param
		q := req.URL.Query()
		q.Add("slug", "portfolio-about")
		req.URL.RawQuery = q.Encode()

		handler := NewBaseHandler(nil)
		rr := httptest.NewRecorder()
		handler.PageHandler(rr, req)
		assert.Equal(t, 404, rr.Result().StatusCode, "should be 404")
	})

	t.Run("DELETE return 404", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/pages", nil)
		if err != nil {
			t.Fatal(err)
		}
		// add query param
		q := req.URL.Query()
		q.Add("slug", "portfolio-about")
		req.URL.RawQuery = q.Encode()

		handler := NewBaseHandler(nil)
		rr := httptest.NewRecorder()
		handler.PageHandler(rr, req)
		assert.Equal(t, 404, rr.Result().StatusCode, "should be 404")
	})
}
