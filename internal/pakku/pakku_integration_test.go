package pakku

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	ctx = context.Background()

	pakkuFile = testcontainers.ContainerFile{
		HostFilePath:      "../../dist/pakku",
		ContainerFilePath: "/usr/local/bin/pakku",
		FileMode:          0o755,
	}
)

func mustSucceed(t *testing.T, container testcontainers.Container, cmd []string) {
	rc, data, err := container.Exec(ctx, cmd)
	require.NoError(t, err)
	require.Zerof(t, rc, "expected return code 0 but got %d: %s", rc, mustReadAll(data))
}

func mustReadAll(reader io.Reader) []byte {
	data, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	return data
}

func TestApt(t *testing.T) {
	req := testcontainers.ContainerRequest{
		Image:      "debian:12",
		Entrypoint: []string{"bash", "-c", "echo ready && sleep 300"},
		Files:      []testcontainers.ContainerFile{pakkuFile},
		WaitingFor: wait.ForLog("ready"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	defer testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)

	mustSucceed(t, container, []string{"pakku", "init"})
	mustSucceed(t, container, []string{"pakku", "config"})
	mustSucceed(t, container, []string{"pakku", "add", "apt", "curl"})
	mustSucceed(t, container, []string{"pakku", "apply", "-verbose"})
	mustSucceed(t, container, []string{"pakku", "update", "-verbose"})
	mustSucceed(t, container, []string{"pakku", "remove", "apt", "curl"})
}

func TestDnf(t *testing.T) {
	req := testcontainers.ContainerRequest{
		Image:      "fedora:41",
		Entrypoint: []string{"bash", "-c", "echo ready && sleep 300"},
		Files:      []testcontainers.ContainerFile{pakkuFile},
		WaitingFor: wait.ForLog("ready"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	defer testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)

	mustSucceed(t, container, []string{"pakku", "init"})
	mustSucceed(t, container, []string{"pakku", "config"})
	mustSucceed(t, container, []string{"pakku", "add", "dnf", "curl"})
	mustSucceed(t, container, []string{"pakku", "apply", "-verbose"})
	mustSucceed(t, container, []string{"pakku", "update", "-verbose"})
	mustSucceed(t, container, []string{"pakku", "remove", "dnf", "curl"})
}

func TestBrew(t *testing.T) {
	req := testcontainers.ContainerRequest{
		Image:      "homebrew/brew:latest",
		Entrypoint: []string{"bash", "-c", "echo ready && sleep 300"},
		Files:      []testcontainers.ContainerFile{pakkuFile},
		WaitingFor: wait.ForLog("ready"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	defer testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)

	mustSucceed(t, container, []string{"pakku", "init"})
	mustSucceed(t, container, []string{"pakku", "config"})
	mustSucceed(t, container, []string{"pakku", "add", "brew", "curl"})
	mustSucceed(t, container, []string{"pakku", "apply", "-verbose"})
	mustSucceed(t, container, []string{"pakku", "update", "-verbose"})
	mustSucceed(t, container, []string{"pakku", "remove", "brew", "curl"})
}

func TestPkgx(t *testing.T) {
	req := testcontainers.ContainerRequest{
		Image:      "pkgxdev/pkgx:v2",
		Entrypoint: []string{"bash", "-c", "echo ready && sleep 300"},
		Files:      []testcontainers.ContainerFile{pakkuFile},
		WaitingFor: wait.ForLog("ready"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	defer testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)

	mustSucceed(t, container, []string{"pakku", "init"})
	mustSucceed(t, container, []string{"pakku", "config"})
	mustSucceed(t, container, []string{"pakku", "add", "pkgx", "curl.se^8"})
	mustSucceed(t, container, []string{"pakku", "apply", "-verbose"})
	mustSucceed(t, container, []string{"pakku", "update", "-verbose"})
	mustSucceed(t, container, []string{"pakku", "remove", "pkgx", "curl.se^8"})
}
