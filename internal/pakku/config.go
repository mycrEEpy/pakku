package pakku

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ConfigVersion struct {
	Version int `yaml:"version"`
}

type ConfigPackageManager struct {
	Packages []string `yaml:"packages"`
}

type Config struct {
	ConfigVersion `yaml:",inline"`

	Apt  ConfigPackageManager `yaml:"apt"`
	Brew ConfigPackageManager `yaml:"brew"`
	Dnf  ConfigPackageManager `yaml:"dnf"`
	Pkgx ConfigPackageManager `yaml:"pkgx"`
	Yum  ConfigPackageManager `yaml:"yum"`
}

func resolveAbsoluteConfigPath(configPath string) (string, error) {
	if len(configPath) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}

		return filepath.Join(homeDir, ".config", "pakku", "config.yml"), nil
	}

	return filepath.Abs(configPath)
}

func parseConfigVersion(path string) (*ConfigVersion, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}

	defer file.Close()

	var configVersion *ConfigVersion

	err = yaml.NewDecoder(file).Decode(configVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return configVersion, nil
}

func parseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}

	defer file.Close()

	var config *Config

	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)

	err = decoder.Decode(config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return config, nil
}
