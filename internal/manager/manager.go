package manager

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
)

type Manager interface {
	InstallPackages(ctx context.Context, verbose bool) error
	UpdatePackages(ctx context.Context, verbose bool) error
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

func runCommand(ctx context.Context, cmdAndArgs []string, sudo, verbose bool) error {
	var cmd *exec.Cmd

	if sudo {
		cmd = exec.CommandContext(ctx, "sudo", cmdAndArgs...)
	} else {
		cmd = exec.CommandContext(ctx, cmdAndArgs[0], cmdAndArgs[1:]...)
	}

	var buf bytes.Buffer

	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stdout = &buf
		cmd.Stderr = &buf
	}

	err := cmd.Run()
	if err != nil {
		if buf.Len() > 0 {
			fmt.Printf("Command failed: %s\n", buf.String())
		}

		return err
	}

	return nil
}
