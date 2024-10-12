package manager

import (
	"context"
)

type Brew struct{}

func (m *Brew) InstallPackage(ctx context.Context, pkg string, sudo bool) error {
	return installPackageWithManager(ctx, "brew", pkg, sudo)
}
