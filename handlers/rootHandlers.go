package handlers

import (
	"net/http"
	"fmt"
)


func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Modziki API")
}