package manager

import (
	"context"
	"fmt"
)

type Brew struct {
	Packages []string
	Sudo     bool
}

func (m *Brew) InstallPackages(ctx context.Context, verbose bool) error {
	for _, pkg := range m.Packages {
		fmt.Printf("Installing %s with brew...\n", pkg)

		err := runCommand(ctx, []string{"brew", "install", pkg}, m.Sudo, verbose)
		if err != nil {
			return fmt.Errorf("failed to install %s: %w", pkg, err)
		}
	}

	return nil
}

func (m *Brew) UpdatePackages(ctx context.Context, verbose bool) error {
	fmt.Println("Updating packages with brew...")

	return runCommand(ctx, append([]string{"brew", "upgrade"}, m.Packages...), m.Sudo, verbose)
}
