package main // import "github.com/blueshirts/madziki"

import (
	"net/http"
	"fmt"
	"github.com/blueshirts/madziki/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	gh "github.com/gorilla/handlers"
	"os"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	// root
	r.HandleFunc("/", handlers.RootHandler)
	// movements
	r.HandleFunc("/movements", handlers.PostMovementHandler).Methods(http.MethodPost)
	r.HandleFunc("/movements/{id}", handlers.GetMovementHandler).Methods(http.MethodGet)
	r.HandleFunc("/movements", handlers.UpdateMovementHandler).Methods(http.MethodPut)
	r.HandleFunc("/movements/{id}", handlers.DeleteMovementHandler).Methods(http.MethodDelete)
	r.HandleFunc("/movements", handlers.GetMovementsListHandler).Methods(http.MethodGet)

	// add handler logging
	routed := gh.LoggingHandler(os.Stdout, gh.RecoveryHandler()(r))

	port := 3000
	log.Info(fmt.Sprintf("Starting server on port: %d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), routed))
}
