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
	"github.com/pkg/errors"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	l := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(l)

	if err := db.New(ctx); err != nil {
		return errors.WithStack(err)
	}

	if err := api.New(":1234", ":3000"); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
