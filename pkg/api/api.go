package api

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/olekukonko/errors"
)

func New(ctx context.Context, addr string, dev bool) error {
	mux := http.NewServeMux()

	dist(dev, mux)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		slog.Info("api", "serve", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("api", "err", err)
		}
	}()

	<-ctx.Done()

	sctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(sctx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
