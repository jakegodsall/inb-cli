package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	NotionApiToken string `json:"notion_api_token"`
	DatabaseID     string `json:"database_id"`
}

func GetOrCreateConfigDir() (string, error) {
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

func ConfigFileExists(filename string) (bool, error) {
	configDir, err := GetOrCreateConfigDir()
	if err != nil {
		return false, err
	}
	configPath := filepath.Join(configDir, filename)
	_, err = os.Stat(configPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadConfig(filename string) (Config, error) {
	var config Config
	configDir, err := GetOrCreateConfigDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(configDir, filename)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("could not unmarshal JSON config: %v", err)
	}

	return config, nil
}

func WriteConfig(filename string, config Config) error {
	configDir, err := GetOrCreateConfigDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(configDir, filename)
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return fmt.Errorf("could not marshal config to JSON: %v", err)
	}
	return os.WriteFile(configPath, data, 0644)
}

func DeleteConfig(filename string) error {
	configDir, err := GetOrCreateConfigDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(configDir, filename)
	return os.Remove(configPath)
}