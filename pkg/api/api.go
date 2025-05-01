package api

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/olekukonko/errors"
)

//go:generate ./oapi.sh

func New(ctx context.Context, addr string, db db.DB, dev bool) error {
	mux := http.NewServeMux()

	cfg := huma.DefaultConfig("Codon", "1.0.0")
	cfg.DocsPath = "/spec"
	a := humago.New(mux, cfg)
	g := huma.NewGroup(a, "/api")

	a.UseMiddleware(func(ctx huma.Context, next func(huma.Context)) {
		next(ctx)
		slog.Info("api", "method", ctx.Method(), "path", ctx.URL().Path, "status", ctx.Status())
	})

	gs := huma.GenerateSummary
	huma.GenerateSummary = func(method, path string, response any) string {
		s := gs(method, path, response)
		return strings.Replace(s, "API ", "", 1)
	}

	dist(mux, dev)
	contacts(g, db)
	health(g)

	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		slog.Info("api", "serve", addr)
		if err := s.ListenAndServe(); err != nil {
			slog.Error("api", "err", err)
		}
	}()

	<-ctx.Done()

	sctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(sctx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
