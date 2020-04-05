package helpers

import (
	"fmt"
	"io"
	"net/http"
)

func WriteSuccessResponse(w http.ResponseWriter) {
	// TODO write success response
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, e error) error {
	w.WriteHeader(statusCode)
	_, err := io.WriteString(w, fmt.Sprintf(`{"status":"error", "message": "%v"}`, e))
	return err
}

func WriteFailResponse(w http.ResponseWriter) {
	// TODO write error response handling
}
