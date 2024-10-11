package pakku

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func parseManagerAndPackage() (string, string) {
	if len(os.Args) < 4 {
		return "", ""
	}

	return os.Args[2], os.Args[3]
}

func installAptPackage(ctx context.Context, pkg string) error {
	fmt.Printf("Installing %s with apt...\n", pkg)

	cmd := exec.CommandContext(ctx, "sudo", "apt-get", "install", pkg)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to install %s: %s\n", pkg, string(out))
		return err
	}

	return nil
}
