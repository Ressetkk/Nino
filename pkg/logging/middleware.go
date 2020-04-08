package logging

import (
	log "github.com/sirupsen/logrus"
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
	ApacheFormatPattern = "%s - - [%s %s %s] \"%s %d %d\"\n"
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := Wrap(w)
		h.ServeHTTP(sw, r)

		log.Infof(ApacheFormatPattern, r.RemoteAddr, time.Now().UTC(), r.Method, r.URL, r.Proto, sw.Status, sw.ResponseSize)
	})
}
