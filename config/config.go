package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

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

	return config
}
