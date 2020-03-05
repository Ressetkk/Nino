package downloader

import "github.com/gorilla/mux"

func InitDownloaderRouters(r *mux.Router) {
	s := r.PathPrefix("/downloads").Subrouter()
	s.HandleFunc("/add", AddToQueue).Methods("POST")
}
