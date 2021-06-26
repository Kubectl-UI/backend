package handlers

import (
	"encoding/json"
	"net/http"
)

type Handlers struct {
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) WelcomeMessage(w http.ResponseWriter, r *http.Request) {
	sendJson(w, "Welcome to kubectl UI")
}

func sendJson(w http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}

func getJson(r *http.Request, data interface{}) {
	json.NewDecoder(r.Body).Decode(data)
}
