package manager

import (
	"context"
)

type Dnf struct{}

func (m *Dnf) InstallPackage(ctx context.Context, pkg string, sudo bool) error {
	return installPackageWithManager(ctx, "dnf", pkg, sudo)
}
