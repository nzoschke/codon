package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"

	"github.com/nzoschke/codon/pkg/api"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/log"
	"github.com/pkg/errors"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, getenv func(string) string, stdout io.Writer) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	log.SetDefault(getenv, stdout)
	slog.Debug("run", "args", args)

	if err := db.New(ctx); err != nil {
		return errors.WithStack(err)
	}

	if err := api.New(":1234", ":3000"); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
