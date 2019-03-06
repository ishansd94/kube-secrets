package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ishansd94/kube-secrets/internal/app/secret"
)

func req(r http.Handler, method, path string, body string) *httptest.ResponseRecorder {

	reader := strings.NewReader(body)

	req, _ := http.NewRequest(method, path, reader)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func TestGin(t *testing.T) {

	r := Router()

	// #1
	w := req(r.Handler, "GET", "/version", "")
	if w.Code != http.StatusOK {
		t.Errorf("expected statusOK,  got %v", w.Code)
	}

	// #2
	d, err := json.Marshal(secret.SecretRequest{
		Name: "foo",
		Namespace: "secret-test",
		Content: map[string]string{
			"foo": "bar",
		},
	})
	if err != nil{
		t.Errorf("cannot marshall %v", err)
	}

	w = req(r.Handler, "POST", "/", string(d))
	if w.Code != http.StatusCreated {
		t.Errorf("expected statusCreated,  got %v, respones: %v", w.Code, w.Body)
	}

	// #3
	w = req(r.Handler, "POST", "/", "")
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected statusBadRequest,  got %v, respones: %v", w.Code, w.Body)
	}

	// #4
	w = req(r.Handler, "POST", "/?name=foo", "")
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected statusBadRequest,  got %v, respones: %v", w.Code, w.Body)
	}

	// #5
	w = req(r.Handler, "GET", "/?namespace=secret-test", "")
	if w.Code != http.StatusOK {
		t.Errorf("expected statusOK,  got %v, respones: %v", w.Code, w.Body)
	}

	// #6
	w = req(r.Handler, "GET", "/?namespace=secret-test&name=foo", "")
	if w.Code != http.StatusOK {
		t.Errorf("expected statusOK,  got %v, respones: %v", w.Code, w.Body)
	}

	// #7
	w = req(r.Handler, "GET", "/?namespace=secret-test&name=f", "")
	if w.Code != http.StatusNotFound {
		t.Errorf("expected statusNotFound,  got %v, respones: %v", w.Code, w.Body)
	}
}