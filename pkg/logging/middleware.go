package logging

import (
	"github.com/clevergo/jsend"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	Status, ResponseSize int
}

func (sw *statusWriter) WriteHeader(statusCode int) {
	sw.Status = statusCode
	sw.ResponseWriter.WriteHeader(statusCode)
}

func (sw *statusWriter) Write(b []byte) (int, error) {
	if sw.Status == 0 {
		sw.Status = http.StatusOK
	}
	sw.ResponseSize = len(b)
	return sw.ResponseWriter.Write(b)
}

func Wrap(w http.ResponseWriter) *statusWriter {
	return &statusWriter{ResponseWriter: w}
}

const (
	ApacheFormatPattern = `%s - - [%s] "%s %s %s" %d %d\n`
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := Wrap(w)
		h.ServeHTTP(sw, r)

		log.Infof(ApacheFormatPattern, r.RemoteAddr, time.Now().UTC(), r.Method, r.URL, r.Proto, sw.Status, sw.ResponseSize)
	})
}

type ErrorHandler func(http.ResponseWriter, *http.Request) error

func (eh ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := eh(w, r)
	if err == nil {
		return
	} else {
		log.Errorf("caught an error: %v", err)
		// TODO error handling
		if merr, ok := err.(mongo.WriteException); ok {
			switch merr.WriteErrors[0].Code {
			case 11000:
				jsend.Error(w, "duplicate found", http.StatusBadRequest)
			}
		}
	}
}
