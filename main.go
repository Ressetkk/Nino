package main

import (
	"encoding/json"
	"github.com/Ressetkk/nino/config"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	i := config.Config{SessionID:"asd", UserID:"asdads"}
	json.NewEncoder(w).Encode(i)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello).Methods("GET")

	log.Infof("Server listening on %s\n",":8080")
	http.ListenAndServe(":8080", r)
}