package docker

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
			name: "Single service with ports and environment variables",
			services: []Service{
				{
					Name:    "web",
					Image:   "nginx",
					Version: "latest",
					Ports:   []string{"80:80", "443:443"},
					EnvVars: map[string]string{"ENV": "production"},
				},
			},
			expected: `version: '3.8'

services:
  web:
    image: nginx:latest
    ports:
      - "80:80"
      - "443:443"
    environment:
      - ENV=production

`,
		},
		{
			name: "Service with only ports",
			services: []Service{
				{
					Name:    "db",
					Image:   "postgres",
					Version: "13",
					Ports:   []string{"5432:5432"},
					EnvVars: nil,
				},
			},
			expected: `version: '3.8'

services:
  db:
    image: postgres:13
    ports:
      - "5432:5432"

`,
		},
		{
			name: "Service with only environment variables",
			services: []Service{
				{
					Name:    "cache",
					Image:   "redis",
					Version: "alpine",
					Ports:   nil,
					EnvVars: map[string]string{"REDIS_PASSWORD": "secret"},
				},
			},
			expected: `version: '3.8'

services:
  cache:
    image: redis:alpine
    environment:
      - REDIS_PASSWORD=secret

`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateComposeFile(tt.services)
			if result != tt.expected {
				t.Errorf("GenerateComposeFile() = %v, want %v", result, tt.expected)
			}
		})
	}
}
