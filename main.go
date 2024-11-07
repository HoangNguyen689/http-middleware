package main

import (
	"fmt"
	"net/http"
)

func errorHandler(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			if customErr, ok := err.(*Error); ok {
				switch customErr.Code {
				case Invalid:
					http.Error(w, err.Error(), http.StatusBadRequest)
				default:
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", errorHandler(healthCheckHandler))

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return &Error{
			error: fmt.Errorf("method %s not allowed", r.Method),
			Code:  Invalid,
		}
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))

	return err
}
