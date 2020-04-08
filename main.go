package main

import (
	"encoding/json"
	"github.com/Ressetkk/nino/config"
	"github.com/Ressetkk/nino/downloader"
	"github.com/Ressetkk/nino/internal/db"
	"github.com/Ressetkk/nino/music"
	"github.com/Ressetkk/nino/pkg/logging"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	i := config.Config{SessionID: "asd", UserID: "asdads"}
	_ = json.NewEncoder(w).Encode(i)
}

func main() {
	err := db.InitDBConnection()
	if err != nil {
		log.Fatalf("Could not establish connection to database: %v\n", err)
	} else {
		log.Info("Successfully connected to database.")
	}
	r := NewRouter()
	r.Use(logging.Middleware)
	log.Infof("Server listening on %s\n", ":8080")
	_ = http.ListenAndServe(":8080", r)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/v1").Subrouter()
	downloader.AddRouter(r)
	music.AddRouter(r)
	return r
}
