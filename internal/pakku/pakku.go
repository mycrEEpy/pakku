package pakku

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/mycreepy/pakku/internal/manager"
	"gopkg.in/yaml.v3"
)

type Pakku struct {
	configPath  string
	config      *Config
	AptManager  manager.Manager
	BrewManager manager.Manager
	DnfManager  manager.Manager
	PkgxManager manager.Manager
}

func New() (*Pakku, error) {
	absConfigPath, err := resolveAbsoluteConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve config path: %w", err)
	}

	return &Pakku{
		configPath:  absConfigPath,
		AptManager:  &manager.Apt{},
		BrewManager: &manager.Brew{},
		DnfManager:  &manager.Dnf{},
		PkgxManager: &manager.Pkgx{},
	}, nil
}

func (p *Pakku) Run(ctx context.Context) error {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		return p.printHelp()
	}

	command := os.Args[1]

	if command == "init" {
		return p.initConfig()
	}

	err := p.readConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	switch command {
	case "config":
		return p.printConfig()
	case "add", "remove":
		mgr, pkg := manager.ParseManagerAndPackage(os.Args)
		return p.handlePackage(command, mgr, pkg)
	case "apply":
		return p.applyPackages(ctx)
	case "plan":
		return p.planPackages(ctx)
	case "import":
		return p.importPackages(ctx, manager.ParseManager(os.Args))
	default:
		return fmt.Errorf("unknown command: %s", command)
	}
}

func (p *Pakku) printHelp() error {
	fmt.Println("Usage: pakku <command>")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("	help				Show this help message")
	fmt.Println("	init				Initialize a new pakku configuration")
	fmt.Println("	config				Show current configuration")
	fmt.Println("	add	<manager> <package>	Add a new package to the configuration")
	fmt.Println("	remove	<manager> <package>	Remove a package from the configuration")
	fmt.Println("	import	<manager>		Import all packages from the package manager to the configuration")
	fmt.Println("	plan				Show the differences between the configuration and the system")
	fmt.Println("	apply				Apply the configuration to the system")
	fmt.Println()
	fmt.Println("Supported package managers:")
	fmt.Println("	apt")
	fmt.Println("	brew")
	fmt.Println("	dnf")
	fmt.Println("	pkgx")
	fmt.Println()

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

	p.config = &Config{
		ConfigVersion: ConfigVersion{Version: 1},
	}

	err = p.writeConfigToDisk()
	if err != nil {
		return fmt.Errorf("failed to write default config: %w", err)
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
		return p.addPackageToConfig(manager, pkg)
	case "remove":
		return p.removePackageFromConfig(manager, pkg)
	default:
		return fmt.Errorf("unknown package action: %s", action)
	}
}

func (p *Pakku) addPackageToConfig(manager, pkg string) error {
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

	return p.writeConfigToDisk()
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

		idx := slices.Index(p.config.Pkgx.Packages, pkg)

		p.config.Pkgx.Packages = slices.Delete(p.config.Pkgx.Packages, idx, idx+1)
	default:
		return fmt.Errorf("unsupported package manager: %s", manager)
	}

	return p.writeConfigToDisk()
}

func (p *Pakku) writeConfigToDisk() error {
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

func (p *Pakku) applyPackages(ctx context.Context) error {
	fs := flag.NewFlagSet("apply", flag.ExitOnError)

	verbose := fs.Bool("verbose", false, "Show package manager output")

	err := fs.Parse(os.Args[2:])
	if err != nil {
		return fmt.Errorf("failed to parse apply flags: %w", err)
	}

	for _, pkg := range p.config.Apt.Packages {
		err := p.AptManager.InstallPackage(ctx, pkg, p.config.Apt.Sudo, *verbose)
		if err != nil {
			return fmt.Errorf("failed to install package %s: %w", pkg, err)
		}
	}

	for _, pkg := range p.config.Brew.Packages {
		err := p.BrewManager.InstallPackage(ctx, pkg, p.config.Brew.Sudo, *verbose)
		if err != nil {
			return fmt.Errorf("failed to install package %s: %w", pkg, err)
		}
	}

	for _, pkg := range p.config.Dnf.Packages {
		err := p.DnfManager.InstallPackage(ctx, pkg, p.config.Dnf.Sudo, *verbose)
		if err != nil {
			return fmt.Errorf("failed to install package %s: %w", pkg, err)
		}
	}

	for _, pkg := range p.config.Pkgx.Packages {
		err := p.PkgxManager.InstallPackage(ctx, pkg, p.config.Pkgx.Sudo, *verbose)
		if err != nil {
			return fmt.Errorf("failed to install package %s: %w", pkg, err)
		}
	}

	return nil
}

func (p *Pakku) planPackages(ctx context.Context) error {
	return errors.New("plan not implemented")
}

func (p *Pakku) importPackages(ctx context.Context, mgr string) error {
	if mgr == "" {
		return errors.New("no package manager specified")
	}

	return errors.New("import not implemented")
}
