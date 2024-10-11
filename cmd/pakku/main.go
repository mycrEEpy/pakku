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
	configPath         = flag.String("config", "", "Path to pakku config file (defaults to $HOME/.config/pakku/config.yml)")
	shouldPrintVersion = flag.Bool("version", false, "Show version")

	version = "develop"
	commit  = "HEAD"
	date    = "just now"
)

func main() {
	flag.Parse()

	if *shouldPrintVersion {
		printVersion()
		os.Exit(0)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	p, err := pakku.New(*configPath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	err = p.Run(ctx)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func printVersion() {
	fmt.Printf("pakku version %s (commit %s) built at %s\n", version, commit, date)
}
