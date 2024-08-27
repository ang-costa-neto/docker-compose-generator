package main

import (
	"fmt"
	"net/http"

	"github.com/ang-costa-neto/docker-compose-generator/internal/docker"
	"github.com/ang-costa-neto/docker-compose-generator/internal/prompt"
	"github.com/ang-costa-neto/docker-compose-generator/internal/utils"
)

// DockerHubTagFetcher é uma implementação concreta de TagFetcher que usa a API Docker Hub.
type DockerHubTagFetcher struct {
	client *http.Client
}

// NewDockerHubTagFetcher cria uma nova instância de DockerHubTagFetcher com um cliente HTTP.
func NewDockerHubTagFetcher(client *http.Client) *DockerHubTagFetcher {
	return &DockerHubTagFetcher{client: client}
}

// GetAvailableTags busca tags disponíveis para a imagem especificada.
func (f *DockerHubTagFetcher) GetAvailableTags(image string) ([]string, error) {
	return docker.GetAvailableTags(image, f.client)
}

func main() {
	// Cria uma instância do cliente HTTP.
	client := &http.Client{}

	// Cria instâncias dos leitores e fetchers.
	consoleReader := prompt.NewConsoleReader()
	tagFetcher := NewDockerHubTagFetcher(client)

	// Lê os serviços do usuário.
	services, err := prompt.ReadServices(consoleReader, tagFetcher)
	if err != nil {
		fmt.Println("Error reading services:", err)
		return
	}

	// Gera o conteúdo do arquivo docker-compose.yml.
	composeContent := docker.GenerateComposeFile(services)

	// Escreve o conteúdo em um arquivo.
	err = utils.WriteToFile("docker-compose.yml", composeContent)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}

	fmt.Println("docker-compose.yml file generated successfully!")
}
