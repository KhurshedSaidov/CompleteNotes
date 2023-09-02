package config

import (
	"encoding/json"
	"os"
)

type ServerConfig struct {
	Ip   string `json: "ip"`
	Port string `json: "port`
}

func InitConfigs() (*ServerConfig, error) {
	bytes, err := os.ReadFile("./config/config.json")
	if err != nil {
		return nil, err
	}

	var config ServerConfig
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
