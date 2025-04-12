package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
)

func users(g *echo.Group, db db.DB) {
	g.DELETE("/users/:id", func(c echo.Context) error {
		ctx := c.Request().Context()

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return errors.WithStack(err)
		}

		conn, put, err := db.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		err = q.UserDelete(conn).Run(id)
		if err != nil {
			return errors.WithStack(err)
		}

		return c.JSON(http.StatusOK, nil)
	})

	g.GET("/users/:id", func(c echo.Context) error {
		ctx := c.Request().Context()

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return errors.WithStack(err)
		}

		conn, put, err := db.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		out, err := q.UserRead(conn).Run(id)
		if err != nil {
			return errors.WithStack(err)
		}

		return c.JSON(http.StatusOK, out)
	})

	g.POST("/users", func(c echo.Context) error {
		ctx := c.Request().Context()

		in := q.UserCreateParams{}
		if err := c.Bind(&in); err != nil {
			return errors.WithStack(err)
		}

		conn, put, err := db.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		out, err := q.UserCreate(conn).Run(in)
		if err != nil {
			return errors.WithStack(err)
		}

		return c.JSON(http.StatusOK, out)
	})

	g.PUT("/users/:id", func(c echo.Context) error {
		ctx := c.Request().Context()

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return errors.WithStack(err)
		}

		in := q.UserCreateParams{}
		if err := c.Bind(&in); err != nil {
			return errors.WithStack(err)
		}

		conn, put, err := db.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		err = q.UserUpdate(conn).Run(q.UserUpdateParams{
			Email: in.Email,
			Name:  in.Name,
			Id:    id,
		})
		if err != nil {
			return errors.WithStack(err)
		}

		out, err := q.UserRead(conn).Run(id)
		if err != nil {
			return errors.WithStack(err)
		}

		return c.JSON(http.StatusOK, out)
	})
}
