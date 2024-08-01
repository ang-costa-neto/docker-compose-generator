package docker

import (
	"fmt"
)

type Service struct {
	Name    string
	Image   string
	Version string
	Ports   []string
	EnvVars map[string]string
}

// GenerateComposeFile generates the content of the docker-compose.yml file
func GenerateComposeFile(services []Service) string {
	composeContent := "version: '3.8'\n\nservices:\n"
	for _, service := range services {
		composeContent += fmt.Sprintf("  %s:\n", service.Name)
		composeContent += fmt.Sprintf("    image: %s:%s\n", service.Image, service.Version)

		if len(service.Ports) > 0 && service.Ports[0] != "" {
			composeContent += "    ports:\n"
			for _, port := range service.Ports {
				composeContent += fmt.Sprintf("      - \"%s\"\n", port)
			}
		}

		if len(service.EnvVars) > 0 {
			composeContent += "    environment:\n"
			for key, value := range service.EnvVars {
				composeContent += fmt.Sprintf("      - %s=%s\n", key, value)
			}
		}
		composeContent += "\n"
	}
	return composeContent
}
