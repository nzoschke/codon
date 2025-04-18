package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/sql/models"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
)

func contacts(g *echo.Group, d db.DB) {
	g.GET("/contacts", func(c echo.Context) error {
		ctx := c.Request().Context()

		conn, put, err := d.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		out, err := q.ContactList(conn, 10)
		if err != nil {
			return errors.WithStack(err)
		}

		return c.JSON(http.StatusOK, out)
	})

	g.DELETE("/contacts/:id", func(c echo.Context) error {
		ctx := c.Request().Context()

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return errors.WithStack(err)
		}

		conn, put, err := d.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		err = q.ContactDelete(conn, id)
		if err != nil {
			return errors.WithStack(err)
		}

		return c.JSON(http.StatusOK, nil)
	})

	g.GET("/contacts/:id", func(c echo.Context) error {
		ctx := c.Request().Context()

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return errors.WithStack(err)
		}

		conn, put, err := d.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		out, err := q.ContactRead(conn, id)
		if err != nil {
			return errors.WithStack(err)
		}

		return c.JSON(http.StatusOK, out)
	})

	g.POST("/contacts", func(c echo.Context) error {
		ctx := c.Request().Context()

		in := struct {
			Email string `form:"email" json:"email"`
			Name  string `form:"name" json:"name"`
			Phone string `form:"phone" json:"phone"`
		}{}
		if err := c.Bind(&in); err != nil {
			return errors.WithStack(err)
		}

		conn, put, err := d.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		out, err := q.ContactCreate(conn, q.ContactCreateIn{
			Email: in.Email,
			Meta:  models.Meta{},
			Name:  in.Name,
			Phone: in.Phone,
		})
		if err != nil {
			return errors.WithStack(err)
		}

		if c.Request().Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
			return c.Redirect(http.StatusSeeOther, "/#/contacts")
		}

		return c.JSON(http.StatusOK, out)
	})

	g.POST("/contacts/:id", func(c echo.Context) error {
		ctx := c.Request().Context()

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return errors.WithStack(err)
		}

		in := struct {
			Email string `form:"email"`
			Name  string `form:"name"`
			Phone string `form:"phone"`
		}{}
		if err := c.Bind(&in); err != nil {
			return errors.WithStack(err)
		}

		conn, put, err := d.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		err = q.ContactUpdate(conn, q.ContactUpdateIn{
			Email: in.Email,
			Id:    id,
			Meta:  models.Meta{},
			Name:  in.Name,
			Phone: in.Phone,
		})
		if err != nil {
			return errors.WithStack(err)
		}

		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?id=%d#/contacts/read", id))
	})

	g.PUT("/contacts/:id", func(c echo.Context) error {
		ctx := c.Request().Context()

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return errors.WithStack(err)
		}

		in := q.ContactCreateIn{}
		if err := c.Bind(&in); err != nil {
			return errors.WithStack(err)
		}

		conn, put, err := d.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		err = q.ContactUpdate(conn, q.ContactUpdateIn{
			Email: in.Email,
			Meta:  models.Meta{},
			Name:  in.Name,
			Id:    id,
		})
		if err != nil {
			return errors.WithStack(err)
		}

		out, err := q.ContactRead(conn, id)
		if err != nil {
			return errors.WithStack(err)
		}

		return c.JSON(http.StatusOK, out)
	})
}
