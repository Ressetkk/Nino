package downloader

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

// TODO refactor entirely
func AddRouter(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/downloads").Subrouter()
	s.HandleFunc("/add", helloDownload).Methods("GET")
	return s
}

func helloDownload(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"message": "Hello Download"}`)
}
