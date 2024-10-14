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
