package docker

import (
	"fmt"
	"strings"
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
	var sb strings.Builder
	sb.WriteString("version: '3.8'\n\nservices:\n")

	for _, service := range services {
		sb.WriteString(fmt.Sprintf("  %s:\n", service.Name))
		sb.WriteString(fmt.Sprintf("    image: %s:%s\n", service.Image, service.Version))

		if len(service.Ports) > 0 {
			sb.WriteString("    ports:\n")
			for _, port := range service.Ports {
				sb.WriteString(fmt.Sprintf("      - \"%s\"\n", port))
			}
		}

		if len(service.EnvVars) > 0 {
			sb.WriteString("    environment:\n")
			for key, value := range service.EnvVars {
				sb.WriteString(fmt.Sprintf("      - %s=%s\n", key, value))
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
