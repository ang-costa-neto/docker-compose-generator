package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ang-costa-neto/docker-compose-generator/internal/docker"
)

// readStringInput reads a string input from the user
func readStringInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// readEnvVars reads environment variables input from the user
func readEnvVars(prompt string, reader *bufio.Reader) (map[string]string, error) {
	envVarsStr, err := readStringInput(prompt, reader)
	if err != nil {
		return nil, err
	}

	envVars := make(map[string]string)
	if envVarsStr != "" {
		for _, envVar := range strings.Split(envVarsStr, ",") {
			kv := strings.SplitN(envVar, "=", 2)
			if len(kv) == 2 {
				envVars[kv[0]] = kv[1]
			}
		}
	}
	return envVars, nil
}

// readPorts reads ports input from the user
func readPorts(prompt string, reader *bufio.Reader) ([]string, error) {
	ports, err := readStringInput(prompt, reader)
	if err != nil {
		return nil, err
	}
	return strings.Split(ports, ","), nil
}

// ReadServices reads service details from the user
func ReadServices() ([]docker.Service, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many services do you want to configure? ")
	var numServices int
	_, err := fmt.Scanln(&numServices)
	if err != nil {
		return nil, err
	}

	services := make([]docker.Service, numServices)

	for i := 0; i < numServices; i++ {
		fmt.Printf("Configuring service %d\n", i+1)

		name, err := readStringInput("Service name: ", reader)
		if err != nil {
			return nil, err
		}

		image, err := readStringInput("Service image: ", reader)
		if err != nil {
			return nil, err
		}

		tags, err := docker.GetAvailableTags(image)
		if err != nil {
			fmt.Printf("Error fetching tags for image %s: %v\n", image, err)
			continue
		}

		fmt.Printf("Available tags for %s: %s\n", image, strings.Join(tags, ", "))

		version, err := readStringInput("Service version: ", reader)
		if err != nil {
			return nil, err
		}

		ports, err := readPorts("Ports (e.g., 8080:80, 5432:5432): ", reader)
		if err != nil {
			return nil, err
		}

		envVars, err := readEnvVars("Environment variables (e.g., KEY=VALUE,KEY2=VALUE2): ", reader)
		if err != nil {
			return nil, err
		}

		services[i] = docker.Service{
			Name:    name,
			Image:   image,
			Version: version,
			Ports:   ports,
			EnvVars: envVars,
		}
	}

	return services, nil
}
