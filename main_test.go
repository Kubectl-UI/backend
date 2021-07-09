package main

import (
	"encoding/json"
	"kubectl-gui/client"
	handler "kubectl-gui/handlers"
	"testing"
)

var h = handler.NewHandlers(exec)

func TestWelcomeMessage(t *testing.T) {
	handler := h.WelcomeMessage
	want := "Welcome to kubectl UI"

	resp, err := client.WelcomeMessage(handler)
	if err != nil {
		t.Errorf("Error %s", err)
	}

	var data struct{ Message string }
	json.NewDecoder(resp).Decode(&data)

	got := data.Message

	if want != got {
		t.Errorf("Got wrong info from endpoint, want %s got %s", want, got)
	}
}

func TestCheckKubectl(t *testing.T) {
	handler := h.CheckKubectl
	want := "Your kubectl path: /usr/local/bin/kubectl"

	resp, err := client.CheckKubectl(handler)
	if err != nil {
		t.Errorf("Error %s", err)
	}

	var data struct{ Message string }
	json.NewDecoder(resp).Decode(&data)

	got := data.Message

	if got != want {
		t.Errorf("Error testing welcome, want %s, got %s", want, got)
	}
}

func TestGetUser(t *testing.T) {
	handler := h.GetUser
	want := ""

	resp, err := client.GetUser(handler)
	if err != nil {
		t.Errorf("Error %s", err)
	}

	var data struct{ Message string }
	json.NewDecoder(resp).Decode(&data)

	got := data.Message

	if got != want {
		t.Errorf("Error testing get user, want %s got %s", want, got)
	}
}
