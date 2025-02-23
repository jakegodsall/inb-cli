package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	NotionApiToken string
	DatabaseID     string
}

func CreateConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not access the user's home directory: %v", err)
	}

	config := filepath.Join(home, ".inb-cli")
	log.Println("Config path:", config)

	err = os.MkdirAll(config, 0755)
	if err != nil {
		return "", fmt.Errorf("could not create config directory: %v", err)
	}

	return config, nil
}