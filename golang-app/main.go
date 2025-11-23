package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintln(w, `<html>
<head>
	<title>Received Headers</title>
	<link rel="preconnect" href="https://fonts.googleapis.com">
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
	<link href="https://fonts.googleapis.com/css2?family=JetBrains+Mono&display=swap" rel="stylesheet">
	<style>
		body {
			font-family: 'JetBrains Mono', monospace;
			font-size: 10px;
			margin: 20px;
		}
		.highlight {
			color: blue;
		}
	</style>
</head>
<body>`)
	fmt.Fprintln(w, "<h1>Received HTTP Headers</h1>")

	for name, values := range r.Header {
		for _, value := range values {
			if strings.HasPrefix(name, "Remote-") {
				fmt.Fprintf(w, "<p class=\"highlight\"><strong>%s:</strong> %s</p>\n", name, value)
			} else {
				fmt.Fprintf(w, "<p><strong>%s:</strong> %s</p>\n", name, value)
			}
		}
	}

	fmt.Fprintln(w, "</body></html>")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/", handler)

	log.Printf("🚀 Golang app starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
