package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// DBConfig is the database connect params
type DBConfig struct {
	URL      string `json:"url"`
	DB       string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// DefaultPath is the default path
const DefaultPath = "/run/master/config"

// ReadToConfig create the DBConfig
func ReadToConfig(path string) (*DBConfig, error) {
	var configPath string
	if len(path) == 0 {
		configPath = DefaultPath
	} else if _, err := os.Stat(path); err != nil {
		configPath = DefaultPath
	} else {
		configPath = path
	}

	if _, err := os.Stat(configPath); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config = new(DBConfig)
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
