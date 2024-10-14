package manager

import (
	"context"
	"fmt"
)

type Brew struct{}

func (m *Brew) InstallPackage(ctx context.Context, pkg string, sudo, verbose bool) error {
	fmt.Printf("Installing %s with brew...\n", pkg)

	return runCommand(ctx, []string{"brew", "install", pkg}, sudo, verbose)
}

func (m *Brew) UpdatePackages(ctx context.Context, sudo, verbose bool) error {
	fmt.Println("Updating packages with brew...")

	return runCommand(ctx, []string{"brew", "upgrade"}, sudo, verbose)
}
