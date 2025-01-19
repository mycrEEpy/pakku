package manager

import (
	"context"
	"fmt"
)

type Pkgm struct {
	Packages []string
	Sudo     bool
}

func (m *Pkgm) InstallPackages(ctx context.Context, verbose bool) error {
	for _, pkg := range m.Packages {
		fmt.Printf("Installing %s with pkgm...\n", pkg)

		err := runCommand(ctx, []string{"pkgm", "install", pkg}, m.Sudo, verbose)
		if err != nil {
			return fmt.Errorf("failed to install %s: %w", pkg, err)
		}
	}

	return nil
}

func (m *Pkgm) UpdatePackages(ctx context.Context, verbose bool) error {
	if len(m.Packages) == 0 {
		return nil
	}

	fmt.Println("Updating packages with pkgm...")

	return runCommand(ctx, []string{"pkgm", "update"}, m.Sudo, verbose)
}
