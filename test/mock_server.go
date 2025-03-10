package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/0xN0x/go-artifactsmmo"
	"github.com/0xN0x/go-artifactsmmo/internal/client"
)

func CreateMockedClient(responseData interface{}) *artifactsmmo.ArtifactsMMO {
	response := map[string]interface{}{
		"data": responseData,
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}))

	customConfig := client.NewConfig(&http.Client{}, server.URL, "token", "username")
	client := &artifactsmmo.ArtifactsMMO{Config: customConfig}
	return client
}

func CreateInvalidMockedClient(errorStatusCode int, errorMessage string) *artifactsmmo.ArtifactsMMO {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(errorStatusCode)
		w.Write([]byte(errorMessage))
	}))

	customConfig := client.NewConfig(&http.Client{}, server.URL, "token", "username")
	client := &artifactsmmo.ArtifactsMMO{Config: customConfig}
	return client
}
