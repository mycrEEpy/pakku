package pakku

import (
	"context"
	"fmt"
	"os/exec"
)

func parseManagerAndPackage(args []string) (string, string) {
	if len(args) < 4 {
		return "", ""
	}

	return args[2], args[3]
}

func installAptPackage(ctx context.Context, pkg string, sudo bool) error {
	fmt.Printf("Installing %s with apt...\n", pkg)

	var cmd *exec.Cmd

	if sudo {
		cmd = exec.CommandContext(ctx, "sudo", "apt-get", "install", pkg)
	} else {
		cmd = exec.CommandContext(ctx, "apt-get", "install", pkg)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to install %s: %s\n", pkg, string(out))
		return err
	}

	return nil
}
