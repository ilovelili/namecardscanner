package middleware

import (
	"encoding/json"
	"net/http"
)

// Payload response
type Payload interface{}

// RespondWithError respond with error
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON respond with json
func RespondWithJSON(w http.ResponseWriter, code int, payload Payload) {
	jData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	w.Write(jData)
}
