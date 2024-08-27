package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ang-costa-neto/docker-compose-generator/internal/docker"
)

// InputReader é uma interface para abstrair a leitura de entrada.
type InputReader interface {
	ReadString(prompt string) (string, error)
}

// TagFetcher é uma interface para buscar tags de imagens Docker.
type TagFetcher interface {
	GetAvailableTags(image string) ([]string, error)
}

// ConsoleReader é a implementação padrão de InputReader usando bufio.Reader.
type ConsoleReader struct {
	reader *bufio.Reader
}

// ReadString lê uma string da entrada padrão.
func (r *ConsoleReader) ReadString(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := r.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// NewConsoleReader cria uma nova instância de ConsoleReader.
func NewConsoleReader() *ConsoleReader {
	return &ConsoleReader{reader: bufio.NewReader(os.Stdin)}
}

// readEnvVars lê variáveis de ambiente do usuário.
func readEnvVars(prompt string, reader InputReader) (map[string]string, error) {
	envVarsStr, err := reader.ReadString(prompt)
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

// readPorts lê portas do usuário.
func readPorts(prompt string, reader InputReader) ([]string, error) {
	ports, err := reader.ReadString(prompt)
	if err != nil {
		return nil, err
	}
	return strings.Split(ports, ","), nil
}

// ReadServices lê detalhes dos serviços do usuário.
func ReadServices(reader InputReader, tagFetcher TagFetcher) ([]docker.Service, error) {
	numServicesStr, err := reader.ReadString("How many services do you want to configure? ")
	if err != nil {
		return nil, err
	}
	numServicesStr = strings.TrimSpace(numServicesStr)
	var numServices int
	_, err = fmt.Sscanf(numServicesStr, "%d", &numServices)
	if err != nil {
		return nil, err
	}

	services := make([]docker.Service, numServices)

	for i := 0; i < numServices; i++ {
		fmt.Printf("Configuring service %d\n", i+1)

		name, err := reader.ReadString("Service name: ")
		if err != nil {
			return nil, err
		}

		image, err := reader.ReadString("Service image: ")
		if err != nil {
			return nil, err
		}

		tags, err := tagFetcher.GetAvailableTags(image)
		if err != nil {
			fmt.Printf("Error fetching tags for image %s: %v\n", image, err)
			continue
		}

		fmt.Printf("Available tags for %s: %s\n", image, strings.Join(tags, ", "))

		version, err := reader.ReadString("Service version: ")
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
