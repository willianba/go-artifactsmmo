package actions

import (
	"net/http"
	"sync"

	"github.com/0xN0x/go-artifactsmmo/internal/client"
)

const (
	apiUrl = "https://api.artifactsmmo.com"
)

type ArtifactsMMO struct {
	mu     sync.Mutex
	Config *client.ArtifactsConfig
}

// NewClient creates a new client to access the ArtifactsMMO API
func NewClient(token string, username string) *ArtifactsMMO {
	return &ArtifactsMMO{
		mu:     sync.Mutex{},
		Config: client.NewConfig(&http.Client{}, apiUrl, token, username),
	}
}

// NewClientWithCustomHttpClient creates a new client with a custom http.Client, mainly used for testing
func NewClientWithCustomHttpClient(token string, username string, httpClient *http.Client) *ArtifactsMMO {
	return &ArtifactsMMO{
		mu:     sync.Mutex{},
		Config: client.NewConfig(httpClient, apiUrl, token, username),
	}
}
