package main

import (
	"fmt"
	"log"
	"net/http"
)

const mainPort = "8080"

type Config struct{}

func main() {
	log.Printf("Starting front end service on port %s\n", mainPort)

	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", mainPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
