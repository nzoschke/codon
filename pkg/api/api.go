package api

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nzoschke/codon/pkg/api/contacts"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/olekukonko/errors"
)

//go:generate ./oapi.sh

func NewServer(options ...func(*fuego.Server)) *fuego.Server {
	s := fuego.NewServer(options...)
	fuego.Get(s, "/api/health", func(c fuego.ContextNoBody) (string, error) {
		return "ok", nil
	},
		option.Summary("health"),
		option.OverrideDescription("Check if API is healthy"),
	)

	contacts.Resources{
		Contacts: contacts.New(),
	}.Routes(s)

	return s
}

func New(ctx context.Context, addr string, db db.DB, dev bool) error {
	s := NewServer()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	dist(e, dev)

	api := e.Group("/api")
	api.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	go func() {
		slog.Info("api", "serve", addr)
		if err := s.Run(); err != nil {
			slog.Error("api", "err", err)
		}
	}()

	<-ctx.Done()

	sctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(sctx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
