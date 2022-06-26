package handlers

import (
	"encoding/json"
	"net/http"
)

// Decode is a helper to decode JSON into a Go value
func Decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(&v)
}

// Encode is a helper to encode raw bytes into JSON response
func Encode(w http.ResponseWriter, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
