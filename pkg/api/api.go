package api

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/olekukonko/errors"
)

//go:generate ./oapi.sh

func NewServer(addr string, db db.DB, options ...func(*fuego.Server)) *fuego.Server {
	s := fuego.NewServer(append(
		options,
		fuego.WithAddr("localhost"+addr),
	)...)

	fuego.Get(s, "/api/health", func(c fuego.ContextNoBody) (string, error) {
		return "ok", nil
	},
		option.Summary("health"),
	)

	Contacts(s, db)

	return s
}

func New(ctx context.Context, addr string, db db.DB, dev bool) error {
	s := NewServer(addr, db, fuego.WithEngineOptions(
		fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
			DisableLocalSave: true,
		}),
	),
	)

	dist(s.Mux, dev)

	go func() {
		slog.Info("api", "serve", addr)
		if err := s.Run(); err != nil {
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
