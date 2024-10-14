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

func (m *Apt) UpdatePackages(ctx context.Context, sudo, verbose bool) error {
	fmt.Println("Updating packages with apt...")

	err := runCommand(ctx, []string{"apt-get", "update"}, sudo, verbose)
	if err != nil {
		return err
	}

	return runCommand(ctx, []string{"apt-get", "upgrade"}, sudo, verbose)
}
