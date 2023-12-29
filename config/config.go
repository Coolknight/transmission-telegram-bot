package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config struct to hold bot and Transmission configuration
type Config struct {
	BotToken             string `yaml:"bot_token"`
	TransmissionURL      string `yaml:"transmission_url"`
	TransmissionUser     string `yaml:"transmission_user"`
	TransmissionPassword string `yaml:"transmission_password"`
}

// LoadConfig loads configuration from a YAML file
func LoadConfig(filename string) (Config, error) {
	var config Config
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
