package docker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DockerHubTagResponse struct {
	Count    int `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

// GetAvailableTags fetches the tags for a given image from Docker Hub
func GetAvailableTags(image string) ([]string, error) {
	url := fmt.Sprintf("https://hub.docker.com/v2/repositories/library/%s/tags", image)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tagsResponse DockerHubTagResponse
	if err := json.NewDecoder(resp.Body).Decode(&tagsResponse); err != nil {
		return nil, err
	}

	var tags []string
	for _, result := range tagsResponse.Results {
		tags = append(tags, result.Name)
	}
	return tags, nil
}
