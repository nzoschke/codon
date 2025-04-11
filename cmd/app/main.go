package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"

	"github.com/nzoschke/codon/pkg/api"
	"github.com/nzoschke/codon/pkg/bun"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/log"
	"github.com/olekukonko/errors"
)

var dev = flag.Bool("dev", false, "dev mode")

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, getenv func(string) string, stdout io.Writer) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	flag.Parse()
	slog.Info("flag", "dev", *dev)

	log.SetDefault(getenv, stdout)
	slog.Debug("run", "args", args)

	_, err := db.New(ctx, "codon.sqlite")
	if err != nil {
		return errors.WithStack(err)
	}

	if *dev {
		if err := bun.Dev(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	if err := api.New(ctx, ":1234", *dev); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
