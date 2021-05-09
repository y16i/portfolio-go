package main

import (
	"log"
	"net/http"

	"bitbucket.org/y16i/backend-go/database"
	"bitbucket.org/y16i/backend-go/handler"
)

const (
	port         string = ":8080"
	wpConfigPath string = "/var/www/portfolio/wp/wp-config.php"
)

func main() {
	database.InitDatabase(wpConfigPath)

	http.HandleFunc("/api/v1.0/pages", handler.PageHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
