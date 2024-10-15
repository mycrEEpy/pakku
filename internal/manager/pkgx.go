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

		err := runCommand(ctx, []string{"pkgx", "install", pkg}, m.Sudo, verbose)
		if err != nil {
			return fmt.Errorf("failed to install %s: %w", pkg, err)
		}
	}

	return nil
}

func (m *Pkgx) UpdatePackages(_ context.Context, _ bool) error {
	if len(m.Packages) == 0 {
		return nil
	}

	fmt.Println("Updating packages with pkgx is currently not supported...")

	return nil
}
