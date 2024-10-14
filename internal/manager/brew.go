package manager

import (
	"context"
	"fmt"
)

type Brew struct{}

func (m *Brew) InstallPackage(ctx context.Context, pkg string, sudo, verbose bool) error {
	fmt.Printf("Installing %s with brew...\n", pkg)

	return runCommand(ctx, []string{"brew", "install", pkg}, sudo, verbose)
}
