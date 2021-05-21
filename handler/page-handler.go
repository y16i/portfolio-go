package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/y16i/backend-go/models"
)

type BaseHandler struct {
	pageRepository models.PageRepository
}

func NewBaseHandler(pageRepository models.PageRepository) *BaseHandler {
	return &BaseHandler{
		pageRepository: pageRepository,
	}
}

// uri: /pages?slug=portfolio-about
func (h *BaseHandler) PageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	query := r.URL.Query()
	slug, present := query["slug"]
	if !present {
		notFound(w, r)
		return
	}

	// support only GET
	if r.Method == "GET" {
		page, pageErr := h.pageRepository.FindBySlug(slug[0])
		if pageErr != nil {
			notFound(w, r)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			var pages [1]models.Page
			pages[0] = *page
			marshaled, marshalErr := json.Marshal(pages)
			if marshalErr != nil {
				log.Println(marshalErr)
			}
			w.Write(marshaled)
		}
	} else {
		// PUT, POST, PATCH and DELETE will be 404
		notFound(w, r)
		return
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
