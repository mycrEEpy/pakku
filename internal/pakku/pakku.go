package pakku

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Run(ctx context.Context, configPath string) error {
	if len(os.Args) < 2 {
		fmt.Printf("usage: pakku <command>\n")
		return nil
	}

	var err error

	configPath, err = resolveAbsoluteConfigPath(configPath)
	if err != nil {
		return fmt.Errorf("failed to resolve config path: %w", err)
	}

	command := os.Args[1]

	switch command {
	case "init":
		return initConfig(configPath)
	case "help":
		fmt.Println("help not implemented")
		return nil
	}

	version, err := parseConfigVersion(configPath)
	if err != nil {
		return fmt.Errorf("failed to parse config version: %w", err)
	}

	var _ *Config

	switch version.Version {
	case 0, 1:
		_, err = parseConfig(configPath)
		if err != nil {
			return fmt.Errorf("failed to parse config: %w", err)
		}
	default:
		return fmt.Errorf("unknown pakku config version %d", version.Version)
	}

	switch command {
	case "add":
		fmt.Println("add not implemented")
	case "remove":
		fmt.Println("remove not implemented")
	default:
		return fmt.Errorf("unknown command: %s", os.Args[1])
	}

	return nil
}

func initConfig(configPath string) error {
	err := os.MkdirAll(filepath.Dir(configPath), 0755)
	if err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	_, err = os.Stat(configPath)
	if err == nil {
		return errors.New("config file already exists")
	}

	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}

	defaultConfig := Config{
		ConfigVersion: ConfigVersion{Version: 1},
	}

	return yaml.NewEncoder(file).Encode(defaultConfig)
}
