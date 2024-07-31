package main

import (
	"testing"
)

func TestGenerateComposeFile(t *testing.T) {
	tests := []struct {
		name     string
		services []Service
		expected string
	}{
		{
			name: "single service with ports and env vars",
			services: []Service{
				{
					Name:  "web",
					Image: "myapp:latest",
					Ports: []string{"8080:8080"},
					EnvVars: map[string]string{
						"KEY": "value",
					},
				},
			},
			expected: `version: '3.8'

services:
  web:
    image: myapp:latest
    ports:
      - "8080:8080"
    environment:
      - KEY=value

`,
		},
		{
			name: "multiple services",
			services: []Service{
				{
					Name:  "db",
					Image: "postgres:13",
					Ports: []string{"5432:5432"},
					EnvVars: map[string]string{
						"POSTGRES_USER":     "user",
						"POSTGRES_PASSWORD": "password",
						"POSTGRES_DB":       "mydb",
					},
				},
				{
					Name:  "web",
					Image: "myapp:latest",
					Ports: []string{"8080:8080"},
				},
			},
			expected: `version: '3.8'

services:
  db:
    image: postgres:13
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=password
      - POSTGRES_DB=mydb

  web:
    image: myapp:latest
    ports:
      - "8080:8080"

`,
		},
		{
			name: "service without ports and env vars",
			services: []Service{
				{
					Name:  "web",
					Image: "myapp:latest",
				},
			},
			expected: `version: '3.8'

services:
  web:
    image: myapp:latest

`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateComposeFile(tt.services)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
