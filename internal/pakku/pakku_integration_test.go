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
		HostFilePath:      "../../pakku",
		ContainerFilePath: "/usr/local/bin/pakku",
		FileMode:          0o755,
	}
)

func mustReadAll(reader io.Reader) []byte {
	data, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	return data
}

func TestApt(t *testing.T) {
	req := testcontainers.ContainerRequest{
		Image:      "public.ecr.aws/docker/library/debian:12",
		Entrypoint: []string{"bash", "-c", "echo ready && sleep 60"},
		Files:      []testcontainers.ContainerFile{pakkuFile},
		WaitingFor: wait.ForLog("ready"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	defer testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)

	rc, data, err := container.Exec(ctx, []string{"pakku", "init"})
	require.NoError(t, err)
	require.Zerof(t, rc, "expected return code 0 but got %d: %s", rc, mustReadAll(data))

	rc, data, err = container.Exec(ctx, []string{"pakku", "config"})
	require.NoError(t, err)
	require.Zerof(t, rc, "expected return code 0 but got %d: %s", rc, mustReadAll(data))

	rc, data, err = container.Exec(ctx, []string{"pakku", "add", "apt", "vim"})
	require.NoError(t, err)
	require.Zerof(t, rc, "expected return code 0 but got %d: %s", rc, mustReadAll(data))

	rc, data, err = container.Exec(ctx, []string{"pakku", "apply"})
	require.NoError(t, err)
	require.Zerof(t, rc, "expected return code 0 but got %d: %s", rc, mustReadAll(data))

	rc, data, err = container.Exec(ctx, []string{"pakku", "update"})
	require.NoError(t, err)
	require.Zerof(t, rc, "expected return code 0 but got %d: %s", rc, mustReadAll(data))

	rc, data, err = container.Exec(ctx, []string{"pakku", "remove", "apt", "vim"})
	require.NoError(t, err)
	require.Zerof(t, rc, "expected return code 0 but got %d: %s", rc, mustReadAll(data))
}
