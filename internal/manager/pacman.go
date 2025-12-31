package manager

import (
	"context"
	"fmt"
)

type Pacman struct {
	Packages []string
	Sudo     bool
}

func (m *Pacman) InstallPackages(ctx context.Context, verbose bool) error {
	for _, pkg := range m.Packages {
		fmt.Printf("Installing %s with pacman...\n", pkg)

		err := runCommand(ctx, []string{"pacman", "-y", "-S", pkg}, m.Sudo, verbose)
		if err != nil {
			return fmt.Errorf("failed to install %s: %w", pkg, err)
		}
	}

	return nil
}

func (m *Pacman) UpdatePackages(ctx context.Context, verbose bool) error {
	if len(m.Packages) == 0 {
		return nil
	}

	fmt.Println("Updating packages with pacman...")

	return runCommand(ctx, append([]string{"pacman", "-Syu"}, m.Packages...), m.Sudo, verbose)
}
