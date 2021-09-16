package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var h = NewHandlers("kubectl")

func TestWelcomeMessage(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Fatalf("Could not create request: err - %s", err)
	}

	rec := httptest.NewRecorder()

	h.WelcomeMessage(rec, req)
	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Invalid status code... Expected %d, got %d", http.StatusOK, res.StatusCode)
	}

	var response struct{ Message string }
	json.NewDecoder(res.Body).Decode(&response)

	want := "Welcome to kubectl UI"
	got := response.Message

	if want != got {
		t.Errorf("Got invalid message from response, want %s, got %s", want, got)
	}
}
