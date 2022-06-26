package server

import (
	"net/http"

	"github.com/reference-repo-go/api/internal/calculator"
	"github.com/reference-repo-go/pkg/handlers"
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
		err := handlers.Decode(w, r, &digitsToAdd)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		response := digitsToAdd.Total()

		w.Header().Set(HeaderKeyContentType, ContentTypeJSON)
		if err := handlers.Encode(w, response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
