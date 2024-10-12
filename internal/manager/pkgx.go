package manager

import (
	"context"
)

type Pkgx struct{}

func (m *Pkgx) InstallPackage(ctx context.Context, pkg string, sudo bool) error {
	return installPackageWithManager(ctx, "pkgx", pkg, sudo)
}
