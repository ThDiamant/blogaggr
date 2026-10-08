package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

const CONFIGPATH string = "/home/spacemonkey/Projects/bootdev-learning/go/.gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (config *Config) SetUser(username string) error {
	config.CurrentUserName = username

	data, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("Error while marshalling config data: %w\n", err)
	}

	err = os.WriteFile(CONFIGPATH, data, 0644)
	if err != nil {
		return fmt.Errorf("Error while writing config data to file: %w\n", err)
	}

	return nil
}

func Read() (Config, error) {
	data, err := os.ReadFile(CONFIGPATH)
	if err != nil {
		return Config{}, fmt.Errorf("Error reading config file: %w\n", err)
	}

	var conf Config
	if err := json.Unmarshal(data, &conf); err != nil {
		return Config{}, fmt.Errorf("Error unmarshalling config file: %w\n", err)
	}

	return conf, nil
}
