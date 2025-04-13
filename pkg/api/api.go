package api

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/olekukonko/errors"
)

func New(ctx context.Context, addr string, db db.DB, dev bool) error {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	dist(e, dev)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	api := e.Group("/api")
	contacts(api, db)

	go func() {
		slog.Info("api", "serve", addr)
		if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
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
