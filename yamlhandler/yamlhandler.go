package yamlhandler

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Feeds  []Feed       `yaml:"feeds"`
	Server ServerConfig `yaml:"server"`
	Login  LoginConfig  `yaml:"login"`
}

type Feed struct {
	URL          string `yaml:"url"`
	DownloadPath string `yaml:"download_path"`
}

type ServerConfig struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	RPCPath string `yaml:"rpc_path"`
}

type LoginConfig struct {
	User string `yaml:"username"`
	Pass string `yaml:"password"`
}

const filePath = "rss/rss.conf"

func AddFeedToYAML(url, downloadPath string) error {
	// Read existing YAML content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %v", err)
	}

	var config Config
	// Unmarshal YAML content into Config struct
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return fmt.Errorf("error unmarshalling YAML: %v", err)
	}

	// Add new feed to Config
	newFeed := Feed{
		URL:          url,
		DownloadPath: downloadPath,
	}
	config.Feeds = append(config.Feeds, newFeed)

	// Marshal updated Config back to YAML
	updatedYAML, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("error marshalling YAML: %v", err)
	}

	// Write updated YAML content to file
	err = os.WriteFile(filePath, updatedYAML, 0644)
	if err != nil {
		return fmt.Errorf("error writing YAML file: %v", err)
	}

	return nil
}
