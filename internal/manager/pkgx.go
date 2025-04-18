package manager

import (
	"context"
	"fmt"
)

const pkgmMinVersion = "pkgm^0.11.0"

type Pkgx struct {
	Packages []string
	Sudo     bool
}

func (m *Pkgx) InstallPackages(ctx context.Context, verbose bool) error {
	for _, pkg := range m.Packages {
		fmt.Printf("Installing %s with pkgx...\n", pkg)

		err := runCommand(ctx, []string{"pkgx", pkgmMinVersion, "install", pkg}, m.Sudo, verbose)
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

	fmt.Println("Updating packages with pkgx...")

	for _, pkg := range m.Packages {
		err := runCommand(ctx, []string{"pkgx", pkgmMinVersion, "update", pkg}, m.Sudo, verbose)
		if err != nil {
			return fmt.Errorf("failed to update %s: %w", pkg, err)
		}
	}

	return nil
}
