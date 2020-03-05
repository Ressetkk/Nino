package main

import (
	"encoding/json"
	"github.com/Ressetkk/nino/config"
	"github.com/Ressetkk/nino/downloader"
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
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	downloader.InitDownloaderRouters(api)
	log.Infof("Server listening on %s\n", ":8080")
	_ = http.ListenAndServe(":8080", r)
}
