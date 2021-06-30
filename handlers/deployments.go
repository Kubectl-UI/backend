package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
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

	if data.Replicas != 0 {
		command = append(command, "-r=", string(data.Replicas))
	}

	if data.Execute {
		command = append(command, "--run-dry -o=yaml")
	}

	// get image -name -replicas -port to be exposed

	// kubectl create deployment kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1

	// 	name,
	// 	image,
	// 	ns
	// 	port=-1: The port that this container exposes.
	//   -r, --replicas=1

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
