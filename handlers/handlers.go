package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
)

type Message struct {
	Message string
}

type Handlers struct {
	ExecPath string
}

func NewHandlers(executable string) *Handlers {
	execPath, err := exec.LookPath(executable)
	if err != nil {
		log.Fatal("could not find executable")
	}
	return &Handlers{
		ExecPath: execPath,
	}
}

func (h *Handlers) WelcomeMessage(w http.ResponseWriter, r *http.Request) {
	sendJson(w, http.StatusOK, Message{Message: "Welcome to kubectl UI"})
}

func (h *Handlers) CheckKubectl(w http.ResponseWriter, r *http.Request) {
	sendJson(w, http.StatusOK, Message{Message: "Your kubectl path: " + h.ExecPath})
}

func (h *Handlers) GetVersion(w http.ResponseWriter, r *http.Request) {
	cmdKubectlVersion := &exec.Cmd{
		Path: h.ExecPath,
		Args: []string{h.ExecPath, "version"},
	}

	result, err := cmdKubectlVersion.Output()
	if err != nil {
		log.Println(err)
		sendJson(w, http.StatusInternalServerError, Message{Message: "Could not execute stated command"})
		return
	}
	sendJson(w, http.StatusOK, Message{Message: string(result)})
}

func sendJson(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func getJson(r *http.Request, data interface{}) {
	json.NewDecoder(r.Body).Decode(data)
}
