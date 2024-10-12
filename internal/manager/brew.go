package manager

import (
	"context"
)

type Brew struct{}

func (m *Brew) InstallPackage(ctx context.Context, pkg string, sudo, verbose bool) error {
	return installPackageWithManager(ctx, "brew", pkg, sudo, verbose)
}
