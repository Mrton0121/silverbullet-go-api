package sbapi

import (
	"net/http"
)

// SBClient holds the necessary data to send http requests to the silverbullet API
type SBClient struct {
	Endpoint   string
	Token      string
	HttpClient *http.Client
}

// NewClient function creates and returns a new client
func NewClient(endpoint string, token string) *SBClient {
	client := &SBClient{}

	client.Endpoint = endpoint
	client.Token = token
	client.HttpClient = &http.Client{}

	return client
}
