package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi_authSession(t *testing.T) {
	data := authInfo{
		User: "admin",
		Pass: "admin",
	}

	payload, err := json.Marshal(data)
	if err != nil {
		t.Error("JSON marshal error")
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/authSession", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()

	api.Router.ServeHTTP(rr, req)
	session, _ := api.Store.Get(req, "session-cookie")

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
		fmt.Println("Session cookie: ", session.Values)
	}
}
