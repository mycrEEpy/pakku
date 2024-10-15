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

func (m *Pkgx) UpdatePackages(_ context.Context, _ []string, _, _ bool) error {
	return nil
}
