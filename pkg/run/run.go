package run

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

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

// Health waits for health check or returns an error
func Health(ctx context.Context, timeout time.Duration, port string) error {
	client := http.Client{}
	startTime := time.Now().UTC()
	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://localhost:%s/health", port), nil)
		if err != nil {
			return errors.WithStack(err)
		}

		res, err := client.Do(req)
		if err == nil && res.StatusCode == http.StatusOK {
			res.Body.Close()
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if time.Since(startTime) >= timeout {
				return fmt.Errorf("timeout reached while waiting for endpoint")
			}

			time.Sleep(25 * time.Millisecond)
		}
	}
}
