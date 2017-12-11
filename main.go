package main // import "github.com/blueshirts/madziki"

import (
	"net/http"
	"fmt"
	"github.com/blueshirts/madziki/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
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
	r.HandleFunc("/movements", handlers.PostMovementHandler).Methods("POST")
	r.HandleFunc("/movements/{id}", handlers.GetMovementHandler).Methods("GET")

	port := 3000
	log.Info(fmt.Sprintf("Starting server on port: %d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

//s := &http.Server{
//Addr:           ":8080",
//Handler:        myHandler,
//ReadTimeout:    10 * time.Second,
//WriteTimeout:   10 * time.Second,
//MaxHeaderBytes: 1 << 20,
//}
//log.Fatal(s.ListenAndServe())
