package manager

import (
	"context"
	"fmt"
)

type Apt struct{}

func (m *Apt) InstallPackage(ctx context.Context, pkg string, sudo, verbose bool) error {
	fmt.Printf("Installing %s with apt...\n", pkg)

	return runCommand(ctx, []string{"apt-get", "install", pkg}, sudo, verbose)
}
