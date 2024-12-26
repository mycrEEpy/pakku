package pakku

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

var (
	ctx = context.Background()

	pakkuFile = testcontainers.ContainerFile{
		HostFilePath:      "../../pakku",
		ContainerFilePath: "/usr/local/bin/pakku",
		FileMode:          0o755,
	}
)

func TestApt(t *testing.T) {
	req := testcontainers.ContainerRequest{
		Image:      "public.ecr.aws/docker/library/debian:12",
		Entrypoint: []string{"sleep", "60"},
		Files:      []testcontainers.ContainerFile{pakkuFile},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	defer testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)

	rc, _, err := container.Exec(ctx, []string{"pakku", "init"})
	require.NoError(t, err)
	require.Zero(t, rc)

	rc, _, err = container.Exec(ctx, []string{"pakku", "config"})
	require.NoError(t, err)
	require.Zero(t, rc)

	rc, _, err = container.Exec(ctx, []string{"pakku", "add", "apt", "vim"})
	require.NoError(t, err)
	require.Zero(t, rc)

	rc, _, err = container.Exec(ctx, []string{"pakku", "apply"})
	require.NoError(t, err)
	require.Zero(t, rc)

	rc, _, err = container.Exec(ctx, []string{"pakku", "update"})
	require.NoError(t, err)
	require.Zero(t, rc)

	rc, _, err = container.Exec(ctx, []string{"pakku", "remove", "apt", "vim"})
	require.NoError(t, err)
	require.Zero(t, rc)
}
