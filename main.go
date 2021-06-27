package main

import (
	"fmt"
	handler "kubectl-gui/handlers"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const port = ":8080"
const exec = "kubectl"

func main() {
	r := mux.NewRouter()

	h := handler.NewHandlers(exec)

	r.HandleFunc("/", h.WelcomeMessage).Methods("GET")
	r.HandleFunc("/check", h.CheckKubectl).Methods("GET")
	r.HandleFunc("/version", h.GetVersion)
	r.HandleFunc("/get-pods", h.GetPods)
	r.HandleFunc("/describe-pod", h.DescribePod).Methods("POST")
	r.HandleFunc("/create-pod", h.CreatePod).Methods("POST")
	r.HandleFunc("/delete-pod", h.DeletePod).Methods("DELETE")

	fmt.Println("Application listening at port " + port)
	http.ListenAndServe(port, handlers.CORS()(r))
}
