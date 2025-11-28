package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/protected", protectedPageHandler)

	log.Printf("🚀 Golang app starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
