package main

import (
	"github.com/Ressetkk/nino/internal/db"
	"github.com/Ressetkk/nino/pkg/downloader"
	"github.com/Ressetkk/nino/pkg/logging"
	"github.com/Ressetkk/nino/pkg/music"
	"github.com/clevergo/jsend"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	err := db.InitDBConnection()
	if err != nil {
		log.Fatalf("Could not establish connection to database: %v\n", err)
	} else {
		log.Info("Successfully connected to database.")
	}

	r := NewRouter()

	srv := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Error(err)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
}

func NewRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/v1").Subrouter()
	r.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		if err := jsend.Success(writer, map[string]string{"status": "ok"}); err != nil {
			log.Error(err)
		}
	})
	downloader.AddRouter(r)
	music.AddRouter(r)
	r.NotFoundHandler = r.NewRoute().HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		jsend.Error(writer, "not found", http.StatusNotFound)
	}).GetHandler()

	r.MethodNotAllowedHandler = r.NewRoute().HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		jsend.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
	}).GetHandler()

	r.Use(logging.Middleware)

	return r
}
