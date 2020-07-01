package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type apiError struct {
	Message string `json:"error"`
}

func jsonError(rw http.ResponseWriter, code int, error string) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("X-Content-Type-Options", "nosniff")
	rw.WriteHeader(code)

	msg := apiError{
		Message: error,
	}

	content, err := json.Marshal(msg)
	if err != nil {
		log.Printf("failed to process error %q: %v", error, err)

		_, _ = fmt.Fprintln(rw, `{"error": "internal error"}`)
		return
	}

	_, _ = fmt.Fprintln(rw, string(content))
}

func jsonErrorf(rw http.ResponseWriter, code int, error string, args ...interface{}) {
	jsonError(rw, code, fmt.Sprintf(error, args...))
}
