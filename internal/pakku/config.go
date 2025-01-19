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
	Sudo     bool     `yaml:"sudo"`
}

type Config struct {
	ConfigVersion `yaml:",inline"`

	Apt  ConfigPackageManager `yaml:"apt"`
	Brew ConfigPackageManager `yaml:"brew"`
	Dnf  ConfigPackageManager `yaml:"dnf"`
	Pkgm ConfigPackageManager `yaml:"pkgm"`
}

func resolveAbsoluteConfigPath() (string, error) {
	if _, found := os.LookupEnv("PAKKU_CONFIG"); !found {
		cfgDir, err := os.UserConfigDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user config directory: %w", err)
		}

		return filepath.Join(cfgDir, "pakku", "config.yml"), nil
	}

	return filepath.Abs(os.Getenv("PAKKU_CONFIG"))
}

func parseConfigVersion(path string) (*ConfigVersion, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}

	defer file.Close()

	var configVersion ConfigVersion

	err = yaml.NewDecoder(file).Decode(&configVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return &configVersion, nil
}

func parseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}

	defer file.Close()

	var config Config

	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)

	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return &config, nil
}
