package api

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
)

func users(g *echo.Group, db db.DB) {
	g.POST("/users", func(c echo.Context) error {
		ctx := c.Request().Context()

		in := q.UserCreateParams{}
		if err := c.Bind(&in); err != nil {
			return errors.WithStack(err)
		}
		slog.Info("users", "in", in)

		conn, put, err := db.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		out, err := q.UserCreate(conn).Run(in)
		if err != nil {
			slog.Error("users", "err", err)
			return errors.WithStack(err)
		}

		slog.Error("users", "out", out)

		return c.JSON(http.StatusOK, out)
	})
}
