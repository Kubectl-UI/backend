package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strconv"

	utils "github.com/phirmware/kubectl-gui/utils"
)

type DeploymentBody struct {
	Name       string
	Image      string
	Replicas   int
	NameSpace  string
	FilePath   string
	Execute    bool
	ExposePort int
}

/**
* get deployments
 */
func (h *Handlers) GetDeployments(w http.ResponseWriter, r *http.Request) {
	ns := r.URL.Query().Get("ns")

	fmt.Println(ns)

	var cmdGetDeployments *exec.Cmd

	if ns == "" {
		cmdGetDeployments = &exec.Cmd{
			Path: h.ExecPath,
			Args: []string{h.ExecPath, "get", "deployments"},
		}
	} else {
		cmdGetDeployments = &exec.Cmd{
			Path: h.ExecPath,
			Args: []string{h.ExecPath, "get", "deployments", "-n=" + ns},
		}
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

/**
* create deployments
 */
func (h *Handlers) CreateDeployments(w http.ResponseWriter, r *http.Request) {

	var data DeploymentBody
	getJson(r, &data)

	fmt.Println(data.Name)

	command := []string{h.ExecPath, "create", "deployment"}

	if data.Name != "" {
		command = append(command, data.Name)
	}

	if data.Image != "" {
		command = append(command, "--image="+data.Image)
	}

	if data.NameSpace != "" {
		getNamespace := &exec.Cmd{
			Path: h.ExecPath,
			Args: []string{h.ExecPath, "get", "-n", data.NameSpace},
		}

		_, err := getNamespace.Output()

		if err != nil {
			sendJson(w, http.StatusBadRequest, Message{Message: "name space does not exist"})
			return
		}

		command = append(command, "-n="+data.NameSpace)
	}

	if data.Replicas != 0 {
		command = append(command, "-r="+strconv.Itoa(data.Replicas))
	}

	if data.Execute {
		command = append(command, "--run-dry -o=yaml")
	}

	cmdCreatePod := &exec.Cmd{
		Path: h.ExecPath,
		Args: command,
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

func (h *Handlers) DeleteDeployments(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	ns := r.URL.Query().Get("ns")
	if name == "" {
		sendJson(w, http.StatusBadRequest, Message{Message: "name is required"})
		return
	}
	var command []string

	if ns != "" {
		command = append(command, h.ExecPath, "delete", "deployments", name, "-n=", ns)
	} else {
		command = append(command, h.ExecPath, "delete", "deployments", name)
	}
	cmdDeleteDeployments := &exec.Cmd{
		Path: h.ExecPath,
		Args: command,
	}
	result, err := cmdDeleteDeployments.Output()
	if err != nil {
		log.Println(err)
		sendJson(w, http.StatusInternalServerError, Message{Message: "Could not execute stated command"})
		return
	}
	sendJson(w, http.StatusOK, Message{Message: string(result)})
}

func (h *Handlers) UpdateDeployments(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query()

	var data DeploymentBody
	getJson(r, &data)

	fmt.Println(data, name)

}

func (h *Handlers) UploadDeploymentFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(100000)

	file, handler, err := r.FormFile("deploymentFile")

	if err != nil {
		fmt.Println("error getting the file")
		sendJson(w, http.StatusBadRequest, "failed to retrieve file")
		return
	}

	print(handler)

	filename, err := utils.FileUpload(file)

	if err != nil {
		sendJson(w, http.StatusInternalServerError, err)
	}

	createDeploymentsFromFile := &exec.Cmd{
		Path: h.ExecPath,
		Args: []string{h.ExecPath, "apply", "-f", filepath.Dir("../files/" + filename)},
	}

	result, err := createDeploymentsFromFile.Output()

	if err != nil {
		log.Println(err)
		sendJson(w, http.StatusInternalServerError, Message{Message: "Could not execute stated command"})
		return
	}
	sendJson(w, http.StatusOK, Message{Message: string(result)})

}

func (h *Handlers) DescribeDeploymentScripts(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		sendJson(w, 400, Message{Message: "Could not execute stated command"})
		return
	}

	cmdDescripeDeployments := &exec.Cmd{
		Path: h.ExecPath,
		Args: []string{" h.ExecPath", "descripe", name},
	}

	result, err := cmdDescripeDeployments.Output()

	if err != nil {
		sendJson(w, 500, Message{Message: "Could not execute stated command"})
		return
	}

	fmt.Println(result)

	sendJson(w, 400, Message{Message: string(result)})

}
