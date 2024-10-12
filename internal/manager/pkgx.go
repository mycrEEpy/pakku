package manager

import (
	"context"
)

type Pkgx struct{}

func (m *Pkgx) InstallPackage(ctx context.Context, pkg string, sudo, verbose bool) error {
	return installPackageWithManager(ctx, "pkgx", pkg, sudo, verbose)
}
