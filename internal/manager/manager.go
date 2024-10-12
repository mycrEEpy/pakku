package manager

import (
	"context"
	"fmt"
	"os/exec"
)

type Manager interface {
	InstallPackage(ctx context.Context, pkg string, sudo bool) error
}

func ParseManager(args []string) string {
	if len(args) < 3 {
		return ""
	}

	return args[2]
}

func ParseManagerAndPackage(args []string) (string, string) {
	if len(args) < 4 {
		return "", ""
	}

	return args[2], args[3]
}

func installPackageWithManager(ctx context.Context, mgr, pkg string, sudo bool) error {
	fmt.Printf("Installing %s with %s...\n", pkg, mgr)

	var cmd *exec.Cmd

	if sudo {
		cmd = exec.CommandContext(ctx, "sudo", mgr, "install", pkg)
	} else {
		cmd = exec.CommandContext(ctx, mgr, "install", pkg)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to install %s: %s\n", pkg, string(out))
		return err
	}

	return nil
}
