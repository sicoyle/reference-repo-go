package server

import (
	"encoding/json"
	"net/http"

	"github.com/reference-repo-go/api/internal/calculator"
)

const (
	HeaderKeyContentType = "Content-Type"
	ContentTypeJSON      = "application/json; charset=utf-8"
)

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (s *server) handleAddition() http.HandlerFunc {
	type response struct {
		total int
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var digitsToAdd calculator.Calculator
		err := s.Decode(w, r, &digitsToAdd)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		response := digitsToAdd.Total()

		w.Header().Set(HeaderKeyContentType, ContentTypeJSON)
		if err := s.Encode(w, response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// Decode is a helper to decode JSON into a Go value
func (s *server) Decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(&v)
}

// Encode is a helper to encode raw bytes into JSON response
func (s *server) Encode(w http.ResponseWriter, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
