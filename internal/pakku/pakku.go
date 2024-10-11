package pakku

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"gopkg.in/yaml.v3"
)

type Pakku struct {
	configPath string
	config     *Config
}

func New(configPath string) (*Pakku, error) {
	absConfigPath, err := resolveAbsoluteConfigPath(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve config path: %w", err)
	}

	return &Pakku{configPath: absConfigPath}, nil
}

func (p *Pakku) Run(ctx context.Context) error {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		return p.printHelp()
	}

	if os.Args[1] == "init" {
		return p.initConfig()
	}

	err := p.readConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	switch os.Args[1] {
	case "config":
		return p.printConfig()
	case "add", "remove":
		manager, pkg := parseManagerAndPackage()
		return p.handlePackage(os.Args[1], manager, pkg)
	case "apply":
		fmt.Println("apply not implemented")
	default:
		return fmt.Errorf("unknown command: %s", os.Args[1])
	}

	return nil
}

func parseManagerAndPackage() (string, string) {
	if len(os.Args) < 4 {
		return "", ""
	}

	return os.Args[2], os.Args[3]
}

func (p *Pakku) printHelp() error {
	fmt.Println("Usage: pakku <command>")
	fmt.Println("Available commands:")
	fmt.Println("	help				Show this help message")
	fmt.Println("	init				Initialize a new pakku config")
	fmt.Println("	config				Show current config")
	fmt.Println("	add	<manager> <package>	Add a new package")
	fmt.Println("	remove	<manager> <package>	Remove a package")
	fmt.Println("	apply				Install & remove packages")

	return nil
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

func (p *Pakku) handlePackage(action, manager, pkg string) error {
	if manager == "" || pkg == "" {
		return errors.New("manager and package are required")
	}

	switch action {
	case "add":
		return p.writePackageToConfig(manager, pkg)
	case "remove":
		return p.removePackageFromConfig(manager, pkg)
	default:
		return fmt.Errorf("unknown package action: %s", action)
	}
}

func (p *Pakku) writePackageToConfig(manager, pkg string) error {
	switch manager {
	case "apt":
		if slices.Contains(p.config.Apt.Packages, pkg) {
			return fmt.Errorf("package %s has already been added for %s", pkg, manager)
		}

		p.config.Apt.Packages = append(p.config.Apt.Packages, pkg)
	case "brew":
		if slices.Contains(p.config.Brew.Packages, pkg) {
			return fmt.Errorf("package %s has already been added for %s", pkg, manager)
		}

		p.config.Brew.Packages = append(p.config.Brew.Packages, pkg)
	case "dnf":
		if slices.Contains(p.config.Dnf.Packages, pkg) {
			return fmt.Errorf("package %s has already been added for %s", pkg, manager)
		}

		p.config.Dnf.Packages = append(p.config.Dnf.Packages, pkg)
	case "pkgx":
		if slices.Contains(p.config.Pkgx.Packages, pkg) {
			return fmt.Errorf("package %s has already been added for %s", pkg, manager)
		}

		p.config.Pkgx.Packages = append(p.config.Pkgx.Packages, pkg)
	default:
		return fmt.Errorf("unsupported package manager: %s", manager)
	}

	file, err := os.Create(p.configPath)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}

	defer file.Close()

	return yaml.NewEncoder(file).Encode(p.config)
}

func (p *Pakku) removePackageFromConfig(manager, pkg string) error {
	switch manager {
	case "apt":
		if !slices.Contains(p.config.Apt.Packages, pkg) {
			return fmt.Errorf("package %s has not been added for %s", pkg, manager)
		}

		idx := slices.Index(p.config.Apt.Packages, pkg)

		p.config.Apt.Packages = slices.Delete(p.config.Apt.Packages, idx, idx+1)
	case "brew":
		if !slices.Contains(p.config.Brew.Packages, pkg) {
			return fmt.Errorf("package %s has not been added for %s", pkg, manager)
		}

		idx := slices.Index(p.config.Brew.Packages, pkg)

		p.config.Brew.Packages = slices.Delete(p.config.Brew.Packages, idx, idx+1)
	case "dnf":
		if !slices.Contains(p.config.Dnf.Packages, pkg) {
			return fmt.Errorf("package %s has not been added for %s", pkg, manager)
		}

		idx := slices.Index(p.config.Dnf.Packages, pkg)

		p.config.Dnf.Packages = slices.Delete(p.config.Dnf.Packages, idx, idx+1)
	case "pkgx":
		if !slices.Contains(p.config.Pkgx.Packages, pkg) {
			return fmt.Errorf("package %s has not been added for %s", pkg, manager)
		}

		idx := slices.Index(p.config.Brew.Packages, pkg)

		p.config.Pkgx.Packages = slices.Delete(p.config.Pkgx.Packages, idx, idx+1)
	default:
		return fmt.Errorf("unsupported package manager: %s", manager)
	}

	file, err := os.Create(p.configPath)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}

	defer file.Close()

	return yaml.NewEncoder(file).Encode(p.config)
}

func (p *Pakku) printConfig() error {
	encoder := yaml.NewEncoder(os.Stdout)
	encoder.SetIndent(2)
	return encoder.Encode(p.config)
}
