package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"os/user"
)

type Message struct {
	Message string
}

type Handlers struct {
	ExecPath string
}

type IncomingData struct {
	PodName  string
	FilePath string
	FromPort string
	ToPort   string
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

func (h *Handlers) GetPods(w http.ResponseWriter, r *http.Request) {
	cmdGetPods := &exec.Cmd{
		Path: h.ExecPath,
		Args: []string{h.ExecPath, "get", "pods"},
	}

	log.Println(cmdGetPods.String())

	result, err := cmdGetPods.Output()
	if err != nil {
		log.Println(err)
		sendJson(w, http.StatusInternalServerError, Message{Message: "Could not execute stated command"})
		return
	}
	sendJson(w, http.StatusOK, Message{Message: string(result)})
}

func (h *Handlers) DescribePod(w http.ResponseWriter, r *http.Request) {
	var data IncomingData
	getJson(r, &data)
	if data.PodName == "" {
		log.Println("NO POD NAME PASSED")
		sendJson(w, http.StatusInternalServerError, Message{Message: "No pod name passed"})
		return
	}

	cmdDescribePod := &exec.Cmd{
		Path: h.ExecPath,
		Args: []string{h.ExecPath, "describe", "pod", data.PodName},
	}

	log.Println(cmdDescribePod.String())

	result, err := cmdDescribePod.Output()
	if err != nil {
		log.Println(err)
		sendJson(w, http.StatusInternalServerError, Message{Message: "Could not execute stated command"})
		return
	}

	sendJson(w, http.StatusOK, Message{Message: string(result)})
}

func (h *Handlers) CreatePod(w http.ResponseWriter, r *http.Request) {
	var data IncomingData
	getJson(r, &data)
	fmt.Println(data)
	if data.FilePath == "" {
		log.Println("NO FILE NAME PASSED")
		sendJson(w, http.StatusInternalServerError, Message{Message: "No file name passed"})
		return
	}

	cmdCreatePod := &exec.Cmd{
		Path: h.ExecPath,
		Args: []string{h.ExecPath, "apply", "-f", data.FilePath},
	}

	log.Println(cmdCreatePod.String())

	result, err := cmdCreatePod.Output()
	if err != nil {
		log.Println(err)
		sendJson(w, http.StatusInternalServerError, Message{Message: "Could not execute stated command"})
		return
	}

	sendJson(w, http.StatusOK, Message{Message: string(result)})
}

func (h *Handlers) GetUser(w http.ResponseWriter, r *http.Request) {
	currentUser, err := user.Current()
	if err != nil {
		sendJson(w, http.StatusNotFound, "User not found")
		return
	}

	sendJson(w, http.StatusOK, currentUser)
}

func (h *Handlers) DeletePod(w http.ResponseWriter, r *http.Request) {
	var data IncomingData
	getJson(r, &data)
	if data.PodName == "" {
		log.Println("NO POD NAME PASSED")
		sendJson(w, http.StatusInternalServerError, Message{Message: "No pod name passed"})
		return
	}

	cmdDeletePod := &exec.Cmd{
		Path: h.ExecPath,
		Args: []string{h.ExecPath, "delete", "pod", data.PodName},
	}

	log.Println(cmdDeletePod.String())

	result, err := cmdDeletePod.Output()
	if err != nil {
		log.Println(err)
		sendJson(w, http.StatusInternalServerError, Message{Message: "Could not execute stated command"})
		return
	}

	sendJson(w, http.StatusOK, Message{Message: string(result)})
}

func (h *Handlers) PortForwadPod(w http.ResponseWriter, r *http.Request) {
	var data IncomingData
	getJson(r, &data)

	if data.ToPort == "" || data.FromPort == "" {
		log.Println("Need to pass in a toport and fromport value")
		sendJson(w, http.StatusInternalServerError, Message{Message: "Port values missing"})
		return
	}

	cmdPortFoward := &exec.Cmd{
		Path: h.ExecPath,
		Args: []string{h.ExecPath, "port-forward", data.PodName, data.FromPort + ":" + data.ToPort},
	}

	log.Println(cmdPortFoward.String())

	result, err := cmdPortFoward.Output()
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
