package main

import (
	"fmt"
	"kubectl-gui/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":8080"

func main() {
	r := mux.NewRouter()

	h := handlers.NewHandlers()

	r.HandleFunc("/", h.WelcomeMessage).Methods("GET")
	r.HandleFunc("/check", h.CheckKubectl).Methods("GET")
	r.HandleFunc("/version", h.GetVersion)

	fmt.Println("Application listening at port " + port)
	http.ListenAndServe(port, r)
}
