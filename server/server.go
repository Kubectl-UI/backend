package server

import (
	"fmt"
	"kubectl-gui/config"
	handler "kubectl-gui/server/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setup(cfg *config.KubectlUI) *mux.Router {
	r := mux.NewRouter()
	exec := cfg.Get(config.ConfigCommand)
	if exec == nil {
		log.Fatalf("FATAL: No execution path found")
	}

	h := handler.NewHandlers(exec.(string))

	r.HandleFunc("/", h.WelcomeMessage).Methods("GET")
	r.HandleFunc("/check", h.CheckKubectl).Methods("GET")
	r.HandleFunc("/user", h.GetUser).Methods("GET")
	r.HandleFunc("/version", h.GetVersion)
	r.HandleFunc("/get-pods", h.GetPods)
	r.HandleFunc("/describe-pod", h.DescribePod).Methods("POST")
	r.HandleFunc("/create-pod", h.CreatePod).Methods("POST")
	r.HandleFunc("/delete-pod", h.DeletePod).Methods("POST")
	r.HandleFunc("/port-forward", h.PortForwadPod).Methods("POST")

	r.HandleFunc("/get/{resource}", h.Get).Methods("GET")
	r.HandleFunc("/delete/{resource}", h.Delete).Methods("POST")
	r.HandleFunc("/custom", h.Custom).Methods("POST")

	return r
}

func Start(cfg *config.KubectlUI) {
	r := setup(cfg)
	port := cfg.Get(config.ApplicationPort)
	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", port),
	}

	log.Printf("Listening on port %s", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("Something went wrong, could not start server", err)
	}
}
