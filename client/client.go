package client

import (
	"io"
	"net/http"
	"net/http/httptest"
)

const url = "http://localhost:8080"

func WelcomeMessage(handler http.HandlerFunc) (io.ReadCloser, error) {
	path := "/"

	return TestGETHTTPFunc(handler, path)
}

func CheckKubectl(handler http.HandlerFunc) (io.ReadCloser, error) {
	path := "/check"

	return TestGETHTTPFunc(handler, path)
}

func GetUser(handler http.HandlerFunc) (io.ReadCloser, error) {
	path := "/user"

	return TestGETHTTPFunc(handler, path)
}

func TestGETHTTPFunc(handler http.HandlerFunc, path string) (io.ReadCloser, error) {
	uri := url + path
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	rec := httptest.NewRecorder()
	handler(rec, req)
	res := rec.Result()

	return res.Body, nil
}
