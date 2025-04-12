package run

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

func Run(ctx context.Context, args []string, getenv func(string) string, stdout io.Writer) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)

	var (
		dbf  = flags.String("db", "codon.sqlite", "database file")
		dev  = flags.Bool("dev", false, "dev mode")
		port = flags.Int("port", 1234, "port")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return errors.WithStack(err)
	}

	slog.Info("flag", "dev", *dev)

	log.SetDefault(getenv, stdout)
	slog.Debug("run", "args", args)

	db, err := db.New(ctx, *dbf)
	if err != nil {
		return errors.WithStack(err)
	}

	if *dev {
		if err := bun.Dev(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	if err := api.New(ctx, fmt.Sprintf(":%d", *port), db, *dev); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
