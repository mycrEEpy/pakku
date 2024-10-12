package manager

import (
	"context"
)

type Apt struct{}

func (m *Apt) InstallPackage(ctx context.Context, pkg string, sudo, verbose bool) error {
	return installPackageWithManager(ctx, "apt", pkg, sudo, verbose)
}
