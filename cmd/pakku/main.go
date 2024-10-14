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
	shouldPrintVersion bool

	version = "develop"
	commit  = "HEAD"
	date    = "just now"
)

func main() {
	flag.BoolVar(&shouldPrintVersion, "version", false, "Show version")

	flag.Parse()

	if shouldPrintVersion {
		printVersion()
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	p, err := pakku.New()
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
