package helpers

import (
	"fmt"
	"io"
	"net/http"
)

func WriteSuccessResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	// TODO write success response
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, e error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := io.WriteString(w, fmt.Sprintf(`{"status":"error", "message": "%v"}`, e))
	return err
}

func WriteFailResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	// TODO write error response handling
}
