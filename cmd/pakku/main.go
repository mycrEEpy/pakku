package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mycreepy/pakku/internal/pakku"
)

var (
	configPath         string
	shouldPrintVersion bool

	version = "develop"
	commit  = "HEAD"
	date    = "just now"
)

func main() {
	mustParseFlags()

	if shouldPrintVersion {
		printVersion()
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	p, err := pakku.New(configPath)
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

func mustParseFlags() {
	fs := flag.NewFlagSet("pakku", flag.ExitOnError)

	fs.StringVar(&configPath, "config", "", "Path to pakku config file (defaults to $HOME/.config/pakku/config.yml)")
	fs.BoolVar(&shouldPrintVersion, "version", false, "Show version")

	err := fs.Parse(os.Args[flag.NArg()+1:])
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			fs.PrintDefaults()
			os.Exit(0)
		}

		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func printVersion() {
	fmt.Printf("pakku version %s (commit %s) built at %s\n", version, commit, date)
}
