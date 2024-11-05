package tests

import (
	"Background-Worker/src/go/handlers"
	"Background-Worker/src/go/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	user := models.User{
		Username: "testuser",
		Password: "testpassword",
	}
	payload, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.RegisterUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestLoginUser(t *testing.T) {
	user := models.User{
		Username: "testuser",
		Password: "testpassword",
	}
	payload, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.LoginUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Login successful"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGenerateCaptcha(t *testing.T) {
	req, err := http.NewRequest("GET", "/captcha", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GenerateCaptcha)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var captcha map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &captcha)
	if err != nil {
		t.Fatal(err)
	}

	if _, exists := captcha["id"]; !exists {
		t.Error("captcha response missing 'id' field")
	}

	if _, exists := captcha["content"]; !exists {
		t.Error("captcha response missing 'content' field")
	}
}
