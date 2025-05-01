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
	m := http.NewServeMux()
	NewAPI(m, db, dev)

	s := &http.Server{
		Addr:    addr,
		Handler: m,
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

func NewAPI(m *http.ServeMux, db db.DB, dev bool) huma.API {
	cfg := huma.DefaultConfig("Codon", "1.0.0")
	cfg.DocsPath = "/spec"

	a := humago.New(m, cfg)

	gs := huma.GenerateSummary
	huma.GenerateSummary = func(method, path string, response any) string {
		s := gs(method, path, response)
		return strings.Replace(s, "API ", "", 1)
	}

	a.UseMiddleware(func(ctx huma.Context, next func(huma.Context)) {
		next(ctx)
		slog.Info("api", "method", ctx.Method(), "path", ctx.URL().Path, "status", ctx.Status())
	})

	g := huma.NewGroup(a, "/api")

	dist(m, dev)
	contacts(g, db)
	health(g)

	return a
}
