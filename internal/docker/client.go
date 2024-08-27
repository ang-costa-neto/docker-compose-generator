package docker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HTTPClient é uma interface para permitir a injeção de um cliente HTTP.
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// DefaultHTTPClient é a implementação padrão do HTTPClient usando http.Client.
type DefaultHTTPClient struct {
	client *http.Client
}

// Get faz uma requisição GET usando o http.Client.
func (c *DefaultHTTPClient) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}

// DockerHubTagResponse representa a resposta da API do Docker Hub.
type DockerHubTagResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

// GetAvailableTags retorna uma lista de tags disponíveis para a imagem especificada.
func GetAvailableTags(image string, client HTTPClient) ([]string, error) {
	url := fmt.Sprintf("https://hub.docker.com/v2/repositories/library/%s/tags", image)
	var allTags []string

	for url != "" {
		resp, err := client.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch tags from Docker Hub: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch tags: status code %d", resp.StatusCode)
		}

		var tagsResponse DockerHubTagResponse
		if err := json.NewDecoder(resp.Body).Decode(&tagsResponse); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}

		for _, result := range tagsResponse.Results {
			allTags = append(allTags, result.Name)
		}

		url = tagsResponse.Next
	}

	return allTags, nil
}

// NewDefaultHTTPClient cria uma nova instância de DefaultHTTPClient com um timeout padrão.
func NewDefaultHTTPClient() *DefaultHTTPClient {
	return &DefaultHTTPClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
