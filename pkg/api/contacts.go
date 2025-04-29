package api

import (
	"database/sql"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/sql/models"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
)

type Contact struct {
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	ID        int       `json:"id"`
	Info      Info      `json:"info"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ContactCreateIn struct {
	Email string `form:"email" json:"email"`
	Info  Info   `form:"info" json:"info"`
	Name  string `form:"name" json:"name"`
	Phone string `form:"phone" json:"phone"`
}

type ContactUpdateIn struct {
	Email string `json:"email"`
	Info  Info   `json:"info"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Info struct {
	Age int `json:"age"`
}

func Contacts(s *fuego.Server, db db.DB) {
	g := fuego.Group(s, "/contacts")

	fuego.Get(g, "", func(c fuego.ContextNoBody) ([]Contact, error) {
		ctx := c.Context()

		conn, put, err := db.Take(ctx)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		defer put()

		rows, err := q.ContactList(conn, 10)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		out := []Contact{}
		for _, r := range rows {
			out = append(out, Contact{
				CreatedAt: r.CreatedAt,
				Email:     r.Email,
				ID:        int(r.Id),
				Info:      Info(r.Info),
				Name:      r.Name,
				Phone:     r.Phone,
				UpdatedAt: r.UpdatedAt,
			})
		}

		return out, nil
	},
		fuego.OptionOverrideDescription(""),
		fuego.OptionSummary("list"),
	)

	fuego.Post(g, "", func(c fuego.ContextWithBody[ContactCreateIn]) (Contact, error) {
		ctx := c.Context()

		in, err := c.Body()
		if err != nil {
			return Contact{}, err
		}

		conn, put, err := db.Take(ctx)
		if err != nil {
			return Contact{}, errors.WithStack(err)
		}
		defer put()

		r, err := q.ContactCreate(conn, q.ContactCreateIn{
			Email: in.Email,
			Info:  models.Info(in.Info),
			Name:  in.Name,
			Phone: in.Phone,
		})
		if err != nil {
			return Contact{}, errors.WithStack(err)
		}

		out := Contact{
			CreatedAt: r.CreatedAt,
			Email:     r.Email,
			ID:        int(r.Id),
			Info:      Info(r.Info),
			Name:      r.Name,
			Phone:     r.Phone,
			UpdatedAt: r.UpdatedAt,
		}

		return out, nil
	},
		fuego.OptionOverrideDescription(""),
		fuego.OptionSummary("create"),
	)

	fuego.Delete(g, "/{id}", func(c fuego.ContextNoBody) (string, error) {
		ctx := c.Request().Context()

		id := int64(c.PathParamInt("id"))

		conn, put, err := db.Take(ctx)
		if err != nil {
			return "", errors.WithStack(err)
		}
		defer put()

		err = q.ContactDelete(conn, id)
		if err != nil {
			return "", errors.WithStack(err)
		}

		return "ok", nil
	},
		fuego.OptionOverrideDescription(""),
		fuego.OptionSummary("delete"),
	)

	fuego.Get(g, "/{id}", func(c fuego.ContextNoBody) (Contact, error) {
		ctx := c.Context()

		conn, put, err := db.Take(ctx)
		if err != nil {
			return Contact{}, errors.WithStack(err)
		}
		defer put()

		r, err := q.ContactRead(conn, int64(c.PathParamInt("id")))
		if err != nil {
			if err == sql.ErrNoRows {
				return Contact{}, fuego.NotFoundError{}
			}
			return Contact{}, errors.WithStack(err)
		}

		out := Contact{
			CreatedAt: r.CreatedAt,
			Email:     r.Email,
			ID:        int(r.Id),
			Info:      Info(r.Info),
			Name:      r.Name,
			Phone:     r.Phone,
			UpdatedAt: r.UpdatedAt,
		}

		return out, nil
	},
		fuego.OptionOverrideDescription(""),
		fuego.OptionSummary("get"),
	)

	fuego.Put(g, "/{id}", func(c fuego.ContextWithBody[ContactUpdateIn]) (Contact, error) {
		ctx := c.Request().Context()

		id := int64(c.PathParamInt("id"))

		in, err := c.Body()
		if err != nil {
			return Contact{}, err
		}

		conn, put, err := db.Take(ctx)
		if err != nil {
			return Contact{}, errors.WithStack(err)
		}
		defer put()

		err = q.ContactUpdate(conn, q.ContactUpdateIn{
			Email: in.Email,
			Id:    id,
			Info:  models.Info(in.Info),
			Name:  in.Name,
			Phone: in.Phone,
		})
		if err != nil {
			return Contact{}, errors.WithStack(err)
		}

		r, err := q.ContactRead(conn, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return Contact{}, fuego.NotFoundError{}
			}
			return Contact{}, errors.WithStack(err)
		}

		out := Contact{
			CreatedAt: r.CreatedAt,
			Email:     r.Email,
			ID:        int(r.Id),
			Info:      Info(r.Info),
			Name:      r.Name,
			Phone:     r.Phone,
			UpdatedAt: r.UpdatedAt,
		}

		return out, nil
	},
		fuego.OptionOverrideDescription(""),
		fuego.OptionSummary("update"),
	)
}
