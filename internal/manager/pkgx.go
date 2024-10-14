package manager

import (
	"context"
	"fmt"
)

type Pkgx struct{}

func (m *Pkgx) InstallPackage(ctx context.Context, pkg string, sudo, verbose bool) error {
	fmt.Printf("Installing %s with pkgx...\n", pkg)

	return runCommand(ctx, []string{"pkgx", "install", pkg}, sudo, verbose)
}

func (m *Pkgx) UpdatePackages(ctx context.Context, sudo, verbose bool) error {
	fmt.Println("Updating packages with pkgx...")

	return runCommand(ctx, []string{"pkgx", "--update"}, sudo, verbose)
}
