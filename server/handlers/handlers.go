package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"os/user"

	"github.com/gorilla/mux"
)

type Message struct {
	Message string
}

type Handlers struct {
	ExecPath string
}

type IncomingData struct {
	Commands                            []string
	PodName, FilePath, FromPort, ToPort string
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

func (h *Handlers) Custom(w http.ResponseWriter, r *http.Request) {
	var data IncomingData
	getJson(r, &data)

	query := r.URL.Query()

	namespace := query.Get("namespace")
	if namespace == "" {
		err := sendJson(w, http.StatusBadRequest, Message{Message: "Namespace missing"})
		if err != nil {
			return
		}
		panic(err)
	}

	commands := data.Commands
	commands = append(commands, "-n", namespace)

	args := []string{h.ExecPath}

	customArgs := append(args, commands...)
	customCommand := &exec.Cmd{
		Path: h.ExecPath,
		Args: customArgs,
	}

	log.Println(customCommand.String())

	result, err := customCommand.Output()
	if err != nil {
		log.Printf("Something went wrong : %s", err)
		return
	}

	err = sendJson(w, http.StatusOK, Message{Message: string(result)})
	if err != nil {
		return
	}
	panic(err)
}

func (h *Handlers) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	resource := params["resource"]
	namespace := query.Get("namespace")

	getResource := &exec.Cmd{
		Path: h.ExecPath,
		Args: []string{h.ExecPath, "get", resource},
	}

	if namespace != "" {
		getResource.Args = append(getResource.Args, "-n", namespace)
	}

	log.Println(getResource.String())

	result, err := getResource.Output()
	if err != nil {
		log.Panicf("Error GETTING resourse, error - %s", err)
		return
	}

	sendJson(w, http.StatusOK, Message{Message: string(result)})
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	resource := params["resource"]
	query := r.URL.Query()

	name := query.Get("name")
	namespace := query.Get("namespace")

	var args []string

	if name == "" {
		sendJson(w, http.StatusBadRequest, Message{Message: "Missing object name of resource"})
		return
	}

	if namespace != "" {
		args = []string{h.ExecPath, "delete", resource, name, "-n", namespace}
	} else {
		args = []string{h.ExecPath, "delete", resource, name}
	}

	deleteResource := &exec.Cmd{
		Path: h.ExecPath,
		Args: args,
	}

	log.Println(deleteResource.String())

	result, err := deleteResource.Output()
	if err != nil {
		log.Panicf("Error deleting resource, err - %s", err)
	}

	sendJson(w, http.StatusOK, Message{Message: string(result)})
}

func (h *Handlers) WelcomeMessage(w http.ResponseWriter, _ *http.Request) {
	sendJson(w, http.StatusOK, Message{Message: "Welcome to kubectl UI"})
}

func (h *Handlers) CheckKubectl(w http.ResponseWriter, _ *http.Request) {
	sendJson(w, http.StatusOK, Message{Message: "Your kubectl path: " + h.ExecPath})
}

func (h *Handlers) GetVersion(w http.ResponseWriter, _ *http.Request) {
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

func (h *Handlers) GetPods(w http.ResponseWriter, _ *http.Request) {
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

func (h *Handlers) GetDeployments(w http.ResponseWriter, _ *http.Request) {
	cmdGetDeployments := &exec.Cmd{
		Path: h.ExecPath,
		Args: []string{h.ExecPath, "get", "deployments"},
	}

	log.Println(cmdGetDeployments.String())

	result, err := cmdGetDeployments.Output()
	if err != nil {
		log.Println(err)
		sendJson(w, http.StatusInternalServerError, Message{Message: "Could not execute stated command"})
		return
	}
	sendJson(w, http.StatusOK, Message{Message: string(result)})
}
func (h *Handlers) DescribeDeployment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deployment := params["deployment"]
	query := r.URL.Query()
	namespace := query.Get("namespace")

	var args []string

	if namespace != "" {
		args = []string{h.ExecPath, "describe", "deployment", deployment, "-n", namespace}
	} else {
		args = []string{h.ExecPath, "describe", "deployment", deployment}
	}

	describeDeployment := &exec.Cmd{
		Path: h.ExecPath,
		Args: args,
	}

	log.Println(describeDeployment.String())

	result, err := describeDeployment.Output()
	if err != nil {
		log.Println(err)
		sendJson(w, http.StatusInternalServerError, Message{Message: "Could not execute stated command"})
		return
	}
	sendJson(w, http.StatusOK, Message{Message: string(result)})
}

func (h *Handlers) GetServices(w http.ResponseWriter, _ *http.Request) {
	cmdGetServices := &exec.Cmd{
		Path: h.ExecPath,
		Args: []string{h.ExecPath, "get", "services"},
	}

	log.Println(cmdGetServices.String())

	result, err := cmdGetServices.Output()
	if err != nil {
		log.Println(err)
	}
	// display result in a structured format
	sendJson(w, http.StatusOK, Message{Message: string(result)})
}
func (h *Handlers) DescribeService(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := params["service"]
	query := r.URL.Query()
	namespace := query.Get("namespace")

	var args []string

	if namespace != "" {
		args = []string{h.ExecPath, "describe", "service", service, "-n", namespace}
	} else {
		args = []string{h.ExecPath, "describe", "service", service}
	}

	describeService := &exec.Cmd{
		Path: h.ExecPath,
		Args: args,
	}
	log.Println(describeService.String())
	result, err := describeService.Output()
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

func (h *Handlers) GetUser(w http.ResponseWriter, _ *http.Request) {
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
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return
	}
}
