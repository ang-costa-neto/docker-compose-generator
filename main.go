package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Service struct {
	Name    string
	Image   string
	Ports   []string
	EnvVars map[string]string
}

// generateComposeFile generates the content of the docker-compose.yml file
func generateComposeFile(services []Service) string {
	composeContent := "version: '3.8'\n\nservices:\n"
	for _, service := range services {
		composeContent += fmt.Sprintf("  %s:\n", service.Name)
		composeContent += fmt.Sprintf("    image: %s\n", service.Image)

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

// readServices reads service details from the user
func readServices() []Service {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many services do you want to configure? ")
	var numServices int
	fmt.Scanln(&numServices)

	services := make([]Service, numServices)

	for i := 0; i < numServices; i++ {
		fmt.Printf("Configuring service %d\n", i+1)

		fmt.Print("Service name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Print("Service image: ")
		image, _ := reader.ReadString('\n')
		image = strings.TrimSpace(image)

		fmt.Print("Ports (e.g., 8080:80, 5432:5432): ")
		ports, _ := reader.ReadString('\n')
		ports = strings.TrimSpace(ports)

		fmt.Print("Environment variables (e.g., KEY=VALUE,KEY2=VALUE2): ")
		envVarsStr, _ := reader.ReadString('\n')
		envVarsStr = strings.TrimSpace(envVarsStr)

		envVars := make(map[string]string)
		if envVarsStr != "" {
			for _, envVar := range strings.Split(envVarsStr, ",") {
				kv := strings.SplitN(envVar, "=", 2)
				if len(kv) == 2 {
					envVars[kv[0]] = kv[1]
				}
			}
		}

		services[i] = Service{
			Name:    name,
			Image:   image,
			Ports:   strings.Split(ports, ","),
			EnvVars: envVars,
		}
	}

	return services
}

// writeToFile writes the content to a file
func writeToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func main() {
	services := readServices()
	composeContent := generateComposeFile(services)

	err := writeToFile("docker-compose.yml", composeContent)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}

	fmt.Println("docker-compose.yml file generated successfully!")
}
