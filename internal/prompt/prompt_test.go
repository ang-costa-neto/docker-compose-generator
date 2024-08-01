package prompt

import (
	"io"
	"os"
	"testing"

	"github.com/ang-costa-neto/docker-compose-generator/internal/docker"
)

func TestReadServices(t *testing.T) {
	// Mock input
	input := "2\nservice1\nimage1\nlatest\n8080:80\nKEY1=VALUE1,KEY2=VALUE2\nservice2\nimage2\nstable\n5432:5432\nKEY3=VALUE3\n"
	expectedServices := []docker.Service{
		{
			Name:    "service1",
			Image:   "image1",
			Version: "latest",
			Ports:   []string{"8080:80"},
			EnvVars: map[string]string{"KEY1": "VALUE1", "KEY2": "VALUE2"},
		},
		{
			Name:    "service2",
			Image:   "image2",
			Version: "stable",
			Ports:   []string{"5432:5432"},
			EnvVars: map[string]string{"KEY3": "VALUE3"},
		},
	}

	// Create a pipe to simulate stdin
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	// Write the input to the writer
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write input to pipe: %v", err)
	}
	w.Close()

	// Backup the original os.Stdin and restore it after the test
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = r

	// Execute the function
	services, err := ReadServices()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Validate the result
	if len(services) != len(expectedServices) {
		t.Fatalf("expected %d services, got %d", len(expectedServices), len(services))
	}

	for i, service := range services {
		expected := expectedServices[i]
		if service.Name != expected.Name ||
			service.Image != expected.Image ||
			service.Version != expected.Version ||
			!equalSlices(service.Ports, expected.Ports) ||
			!equalMaps(service.EnvVars, expected.EnvVars) {
			t.Errorf("expected service %v, got %v", expected, service)
		}
	}
}

// Helper function to compare slices
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Helper function to compare maps
func equalMaps(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
