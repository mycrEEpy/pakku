package pakku

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Pakku struct {
	configPath    string
	configVersion int
	config        *Config
}

func New(configPath string) (*Pakku, error) {
	absConfigPath, err := resolveAbsoluteConfigPath(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve config path: %w", err)
	}

	return &Pakku{configPath: absConfigPath}, nil
}

func (p *Pakku) Run(ctx context.Context) error {
	if len(os.Args) < 2 {
		printHelp()
		return nil
	}

	switch os.Args[1] {
	case "help":
		printHelp()
		return nil
	case "init":
		return p.initConfig()
	}

	err := p.readConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	switch os.Args[1] {
	case "add":
		fmt.Println("add not implemented")
	case "remove":
		fmt.Println("remove not implemented")
	case "apply":
		fmt.Println("apply not implemented")
	default:
		return fmt.Errorf("unknown command: %s", os.Args[1])
	}

	return nil
}

func printHelp() {
	fmt.Println("Usage: pakku <command>")
	fmt.Println("Available commands:")
	fmt.Println("	init				Initialize a new pakku config")
	fmt.Println("	help				Show this help message")
	fmt.Println("	add	<manager> <package>	Add a new package")
	fmt.Println("	remove	<manager> <package>	Remove a package")
	fmt.Println("	apply				Install & remove packages")
}

func (p *Pakku) readConfig() error {
	version, err := parseConfigVersion(p.configPath)
	if err != nil {
		return fmt.Errorf("failed to parse config version: %w", err)
	}

	switch version.Version {
	case 0, 1:
		p.config, err = parseConfig(p.configPath)
		if err != nil {
			return fmt.Errorf("failed to parse config: %w", err)
		}
	default:
		return fmt.Errorf("unknown pakku config version %d", version.Version)
	}

	return nil
}

func (p *Pakku) initConfig() error {
	err := os.MkdirAll(filepath.Dir(p.configPath), 0755)
	if err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	_, err = os.Stat(p.configPath)
	if err == nil {
		return errors.New("config file already exists")
	}

	file, err := os.Create(p.configPath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}

	defaultConfig := Config{
		ConfigVersion: ConfigVersion{Version: 1},
	}

	err = yaml.NewEncoder(file).Encode(defaultConfig)
	if err != nil {
		return fmt.Errorf("failed to encode default config: %w", err)
	}

	fmt.Printf("Created new pakku config: %s\n", p.configPath)

	return nil
}
