package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct to hold bot and Transmission configuration
type Config struct {
	Transmission struct {
		URL      string `yaml:"url"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"transmission"`

	Solarman struct {
		AppId     string `yaml:"appId"`
		AppSecret string `yaml:"appSecret"`
		Email     string `yaml:"email"`
		Password  string `yaml:"password"`
	} `yaml:"solarman"`

	API struct {
		AuthURL string `yaml:"authURL"`
		ApiURL  string `yaml:"apiURL"`
	} `yaml:"api"`

	Telegram struct {
		BotToken string `yaml:"botToken"`
		ChatID   string `yaml:"chatID"`
	} `yaml:"telegram"`

	Device struct {
		DeviceSn string `yaml:"deviceSn"`
	} `yaml:"device"`
}

// ReadConfig loads configuration from a YAML file
func ReadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
