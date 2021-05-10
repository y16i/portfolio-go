package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/y16i/backend-go/database"
)

// uri: /pages?slug=portfolio-about
func PageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	query := r.URL.Query()
	slug, present := query["slug"]
	if !present {
		notFound(w, r)
	}

	// support only GET
	if r.Method == "GET" {
		page, pageErr := database.GetWordPressPage(slug[0])
		if pageErr != nil {
			notFound(w, r)
		} else {
			w.WriteHeader(http.StatusOK)
			var pages [1]database.WordPressPage
			pages[0] = *page
			marshaled, marshalErr := json.Marshal(pages)
			if marshalErr != nil {
				log.Println(marshalErr)
			}
			w.Write([]byte(marshaled))
		}
	} else {
		notFound(w, r)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
