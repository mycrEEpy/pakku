package manager

import (
	"context"
	"fmt"
)

type Dnf struct {
	Packages []string
	Sudo     bool
}

func (m *Dnf) InstallPackages(ctx context.Context, verbose bool) error {
	for _, pkg := range m.Packages {
		fmt.Printf("Installing %s with dnf...\n", pkg)

		err := runCommand(ctx, []string{"dnf", "--assumeyes", "install", pkg}, m.Sudo, verbose)
		if err != nil {
			return fmt.Errorf("failed to install %s: %w", pkg, err)
		}
	}

	return nil
}

func (m *Dnf) UpdatePackages(ctx context.Context, verbose bool) error {
	if len(m.Packages) == 0 {
		return nil
	}

	fmt.Println("Updating packages with dnf...")

	return runCommand(ctx, append([]string{"dnf", "--assumeyes", "upgrade"}, m.Packages...), m.Sudo, verbose)
}
