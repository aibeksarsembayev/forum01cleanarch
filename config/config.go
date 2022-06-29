package config

import (
	"encoding/json"
	"os"
)

// Config struct ...
type Config struct {
	Server struct {
		Address string `json:"address"`
	} `json:"server"`
	Database struct {
		DSN  string `json:"dsn"`
		Path string `json:"path"`
	} `json:"database"`
	Context struct {
		Timeout int `json:"timeout"`
	} `json:"context"`
}

func LoadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}
