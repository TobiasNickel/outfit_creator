package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TobiasNickel/outfit_creator/frontend"
	"github.com/TobiasNickel/outfit_creator/internal/server"
)

func TestPing(t *testing.T) {
	staticfs := frontend.BuildHTTPFS()
	// Start the server
	server := server.New(staticfs)

	// Create a request to the /ping endpoint
	request, err := http.NewRequest("GET", "/api/outfit/random", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Send the request
	responseRecorder := httptest.NewRecorder()
	server.ServeHTTP(responseRecorder, request)

	// Check the response
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, responseRecorder.Code)
	}

	var data map[string]interface{}

	jsonData := responseRecorder.Body.Bytes()
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		t.Error("does not return json")
	}
	_, ok := data["upper"]
	if ok == false {
		t.Error("upper is not defined")
	}
}
