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

func WriteConfig(filename string, data []byte) error {
	configDir, err := GetOrCreateConfigDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(configDir, filename)
	return os.WriteFile(configPath, data, 0644)
}

func ReadConfig(filename string) ([]byte, error) {
	configDir, err := GetOrCreateConfigDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(configDir, filename)
	return os.ReadFile(configPath)
}

func DeleteConfig(filename string) error {
	configDir, err := GetOrCreateConfigDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(configDir, filename)
	return os.Remove(configPath)
}