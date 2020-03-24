package config

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1beta1"
	"context"
	"encoding/json"
	"fmt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1beta1"
	"io/ioutil"
	"log"
)

const DB_PASSWORD_URI = "projects/628113837053/secrets/DB_PASSWORD/versions/1"

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
}

type Config struct {
	DB DatabaseConfig `json:"database"`
}

func GetConfig() Config {
	file, err := ioutil.ReadFile("secrets.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	json.Unmarshal(file, &config)

	err = config.populateSecrets()

	if err != nil {
		log.Fatal(err)
	}
	return config
}

func (c *Config) populateSecrets() error {
	dbPass, err := accessSecretVersion(DB_PASSWORD_URI)

	if err != nil {
		return err
	}

	c.DB.Password = string(dbPass)
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
