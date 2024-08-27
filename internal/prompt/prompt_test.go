package prompt

import (
	"errors"
	"testing"
)

// MockInputReader é uma implementação de InputReader para testes.
type MockInputReader struct {
	Prompts map[string]string
}

// ReadString usa os prompts simulados para retornar entradas.
func (m *MockInputReader) ReadString(prompt string) (string, error) {
	if response, ok := m.Prompts[prompt]; ok {
		return response, nil
	}
	return "", errors.New("unexpected prompt")
}

// MockTagFetcher é uma implementação de TagFetcher para testes.
type MockTagFetcher struct {
	Tags []string
	Err  error
}

// GetAvailableTags retorna tags simuladas ou um erro.
func (m *MockTagFetcher) GetAvailableTags(image string) ([]string, error) {
	return m.Tags, m.Err
}

func TestReadServices(t *testing.T) {
	mockReader := &MockInputReader{
		Prompts: map[string]string{
			"How many services do you want to configure? ": "1",
			"Service name: ":                     "web",
			"Service image: ":                    "nginx",
			"Service version: ":                  "latest",
			"Ports (e.g., 8080:80, 5432:5432): ": "8080:80",
			"Environment variables (e.g., KEY=VALUE,KEY2=VALUE2): ": "KEY=VALUE",
		},
	}

	mockFetcher := &MockTagFetcher{
		Tags: []string{"latest", "1.0"},
		Err:  nil,
	}

	services, err := ReadServices(mockReader, mockFetcher)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(services) != 1 {
		t.Fatalf("expected 1 service, got %d", len(services))
	}

	service := services[0]
	if service.Name != "web" || service.Image != "nginx" || service.Version != "latest" {
		t.Fatalf("unexpected service details: %+v", service)
	}

	if len(service.Ports) != 1 || service.Ports[0] != "8080:80" {
		t.Fatalf("unexpected ports: %v", service.Ports)
	}

	if len(service.EnvVars) != 1 || service.EnvVars["KEY"] != "VALUE" {
		t.Fatalf("unexpected environment variables: %v", service.EnvVars)
	}
}
