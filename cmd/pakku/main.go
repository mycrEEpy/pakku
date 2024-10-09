package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mycreepy/pakku/internal/pakku"
)

var (
	configPath = flag.String("config", "", "path to pakku config file (defaults to $HOME/.config/pakku/config.yml")
)

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	err := pakku.Run(ctx, *configPath)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
