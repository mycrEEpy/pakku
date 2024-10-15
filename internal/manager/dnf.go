package manager

import (
	"context"
	"fmt"
)

type Dnf struct{}

func (m *Dnf) InstallPackage(ctx context.Context, pkg string, sudo, verbose bool) error {
	fmt.Printf("Installing %s with dnf...\n", pkg)

	return runCommand(ctx, []string{"dnf", "install", "--yes", pkg}, sudo, verbose)
}

func (m *Dnf) UpdatePackages(ctx context.Context, pkgs []string, sudo, verbose bool) error {
	fmt.Println("Updating packages with dnf...")

	return runCommand(ctx, append([]string{"dnf", "upgrade", "--yes"}, pkgs...), sudo, verbose)
}
