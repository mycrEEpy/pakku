package manager

import (
	"context"
	"fmt"
)

type Apt struct {
	Packages []string
	Sudo     bool
}

func (m *Apt) InstallPackages(ctx context.Context, verbose bool) error {
	for _, pkg := range m.Packages {
		fmt.Printf("Installing %s with apt...\n", pkg)

		err := runCommand(ctx, []string{"apt-get", "--yes", "install", pkg}, m.Sudo, verbose)
		if err != nil {
			return fmt.Errorf("failed to install %s: %w", pkg, err)
		}
	}

	return nil
}

func (m *Apt) UpdatePackages(ctx context.Context, verbose bool) error {
	if len(m.Packages) == 0 {
		return nil
	}

	fmt.Println("Updating packages with apt...")

	err := runCommand(ctx, []string{"apt-get", "--yes", "update"}, m.Sudo, verbose)
	if err != nil {
		return err
	}

	return runCommand(ctx, append([]string{"apt-get", "--yes", "upgrade"}, m.Packages...), m.Sudo, verbose)
}
