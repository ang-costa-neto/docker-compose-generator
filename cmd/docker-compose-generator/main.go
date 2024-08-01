package main

import (
	"fmt"

	"github.com/ang-costa-neto/docker-compose-generator/internal/docker"
	"github.com/ang-costa-neto/docker-compose-generator/internal/prompt"
	"github.com/ang-costa-neto/docker-compose-generator/internal/utils"
)

func main() {
	services, err := prompt.ReadServices()
	if err != nil {
		fmt.Println("Error reading services:", err)
		return
	}

	composeContent := docker.GenerateComposeFile(services)

	err = utils.WriteToFile("docker-compose.yml", composeContent)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}

	fmt.Println("docker-compose.yml file generated successfully!")
}
