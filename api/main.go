package main

import (
	"fmt"
	"net/http"
	"github.com/blueshirts/madziki/api/logger"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	port := 3000
	logger.Info("main", "main", fmt.Sprintf("Starting server on port: %d", port))

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
