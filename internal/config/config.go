package config

import (
	"context"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1beta1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1beta1"
)

const (
	dbUserURI     = "projects/628113837053/secrets/DB_USER/versions/1"
	dbPasswordURI = "projects/628113837053/secrets/DB_PASSWORD/versions/1"
	dbHostURI     = "projects/628113837053/secrets/DB_HOST/versions/2"

	cookieDomain = ".siyan.io"
)

// DatabaseConfig : Config for database settings
type DatabaseConfig struct {
	Host     string
	User     string
	Password string
}

// CookieConfig : Config for cookie settings
type CookieConfig struct {
	Domain string
}

// Config : Wrapper config that holds all settings for application
type Config struct {
	DB     DatabaseConfig `json:"database"`
	Cookie CookieConfig   `json:"cookie"`
}

// GetConfig : Builds config struct based depending production/development environment.
func GetConfig() *Config {
	config := &Config{}

	if err := config.populateConfig(); err != nil {
		log.Fatal(err)
	}

	return config
}

func (c *Config) populateConfig() error {
	dbUser, err := accessSecretVersion(dbUserURI)
	dbPass, err := accessSecretVersion(dbPasswordURI)
	dbHost, err := accessSecretVersion(dbHostURI)

	if err != nil {
		return err
	}

	c.DB.User = string(dbUser)
	c.DB.Password = string(dbPass)
	c.DB.Host = string(dbHost)

	c.Cookie.Domain = cookieDomain
	return nil
}

// UTILS

func accessSecretVersion(name string) ([]byte, error) {
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
