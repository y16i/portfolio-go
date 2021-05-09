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
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}

	// support only GET
	if r.Method == "GET" {
		page, pageErr := database.GetWordPressPage(slug[0])
		if pageErr != nil {
			log.Printf("%v", pageErr)
		}
		log.Printf("%v", page)
		w.WriteHeader(http.StatusOK)
		marshaled, marshalErr := json.Marshal(page)
		if marshalErr != nil {
			log.Fatal(marshalErr)
		}
		w.Write([]byte(marshaled))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
