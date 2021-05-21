package main

import (
	"log"
	"net/http"

	"bitbucket.org/y16i/backend-go/database"
	"bitbucket.org/y16i/backend-go/handler"
	"bitbucket.org/y16i/backend-go/repositories"
)

const (
	port         string = ":8080"
	wpConfigPath string = "/var/www/portfolio/wp/wp-config.php"
)

func main() {
	db := database.InitDatabase(wpConfigPath)

	pageRepo := repositories.NewPageRepository(db)
	defer pageRepo.Close()

	h := handler.NewBaseHandler(pageRepo)

	http.HandleFunc("/api/v1.0/pages", h.PageHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
