package handlers

import (
	"net/http"
	"fmt"
)


func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Modziki API")

	//id := r.URL.Query().Get("id")
	//var movement *api.Movement
	//err := api.GetMovement(bson.ObjectId(id), movement)
	//
	//w.Header().Set("Content-Type", "application/json")
	//encoder := json.NewEncoder(w)
	//if err != nil {
	//	encoder.Encode(movement)
	//} else {
	//	encoder.Encode(HandlerError{
	//		errors: []string{err.Error()},
	//	})
	//}
}