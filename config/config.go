package config

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1beta1"
	"context"
	"fmt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1beta1"
	"log"
)

const (
	DB_USER_URI     = "projects/628113837053/secrets/DB_USER/versions/1"
	DB_PASSWORD_URI = "projects/628113837053/secrets/DB_PASSWORD/versions/1"
	DB_HOST_URI     = "projects/628113837053/secrets/DB_HOST/versions/2"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
}

type Config struct {
	DB DatabaseConfig `json:"database"`
}

func GetConfig() *Config {
	config := &Config{}

	err := config.populateSecrets()

	if err != nil {
		log.Fatal(err)
	}
	return config
}

func (c *Config) populateSecrets() error {
	dbUser, err := accessSecretVersion(DB_USER_URI)
	dbPass, err := accessSecretVersion(DB_PASSWORD_URI)
	dbHost, err := accessSecretVersion(DB_HOST_URI)

	if err != nil {
		return err
	}

	c.DB.User = string(dbUser)
	c.DB.Password = string(dbPass)
	c.DB.Host = string(dbHost)
	return nil
}

func accessSecretVersion(name string) ([]byte, error) {
	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create secretmanager client: %v", err)
	}

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to access secret version: %v", err)
	}

	return result.Payload.Data, nil
}
