package docker

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

// MockHTTPClient é uma implementação de HTTPClient para testes.
type MockHTTPClient struct {
	MockGet func(url string) (*http.Response, error)
}

// Get usa o mock para retornar uma resposta.
func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return m.MockGet(url)
}

func TestGetAvailableTags(t *testing.T) {
	mockClient := &MockHTTPClient{
		MockGet: func(url string) (*http.Response, error) {
			// Simula uma resposta da API do Docker Hub
			responseBody := `{
				"count": 1,
				"next": "",
				"previous": "",
				"results": [
					{"name": "latest"}
				]
			}`
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(responseBody)),
			}, nil
		},
	}

	tags, err := GetAvailableTags("nginx", mockClient)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(tags) != 1 || tags[0] != "latest" {
		t.Fatalf("expected [latest], got %v", tags)
	}
}
