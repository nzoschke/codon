package api

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/domains/books"
	"github.com/olekukonko/errors"
)

//go:generate ./oapi.sh

func NewServer(options ...func(*fuego.Server)) *fuego.Server {
	s := fuego.NewServer(options...)
	fuego.Get(s, "/api/health", func(c fuego.ContextNoBody) (string, error) {
		return "ok", nil
	})

	booksr := books.BooksResources{
		BooksService: books.NewBooksService(),
	}

	booksr.Routes(s)

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

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	api := e.Group("/api")
	contacts(api, db)

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
