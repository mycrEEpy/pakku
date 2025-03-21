package manager

import (
	"context"
	"fmt"
)

type Pkgx struct {
	Packages []string
	Sudo     bool
}

func (m *Pkgx) InstallPackages(ctx context.Context, verbose bool) error {
	for _, pkg := range m.Packages {
		fmt.Printf("Installing %s with pkgx...\n", pkg)

		installCommand := "local-install"

		if m.Sudo {
			installCommand = "install"
		}

		err := runCommand(ctx, []string{"pkgx", "pkgm", installCommand, pkg}, m.Sudo, verbose)
		if err != nil {
			return fmt.Errorf("failed to install %s: %w", pkg, err)
		}
	}

	return nil
}

func (m *Pkgx) UpdatePackages(ctx context.Context, verbose bool) error {
	if len(m.Packages) == 0 {
		return nil
	}

	//fmt.Println("Updating packages with pkgx...")
	//return runCommand(ctx, []string{"pkgx", "pkgm", "update"}, m.Sudo, verbose)

	fmt.Println("Updating packages with pkgx is currently not supported...")

	return nil
}
