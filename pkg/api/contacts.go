package api

import (
	"database/sql"

	"github.com/go-fuego/fuego"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/models"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
)

func Contacts(s *fuego.Server, db db.DB) {
	g := fuego.Group(s, "/contacts")

	fuego.Get(g, "", func(c fuego.ContextNoBody) ([]models.Contact, error) {
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

		out := []models.Contact{}
		for _, r := range rows {
			out = append(out, models.Contact{
				CreatedAt: r.CreatedAt,
				Email:     r.Email,
				ID:        int(r.Id),
				Info:      models.Info(r.Info),
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

	fuego.Post(g, "", func(c fuego.ContextWithBody[models.ContactCreateIn]) (models.Contact, error) {
		ctx := c.Context()

		in, err := c.Body()
		if err != nil {
			return models.Contact{}, err
		}

		conn, put, err := db.Take(ctx)
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
		}
		defer put()

		r, err := q.ContactCreate(conn, q.ContactCreateIn{
			Email: in.Email,
			Info:  models.Info(in.Info),
			Name:  in.Name,
			Phone: in.Phone,
		})
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
		}

		out := models.Contact{
			CreatedAt: r.CreatedAt,
			Email:     r.Email,
			ID:        int(r.Id),
			Info:      r.Info,
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

	fuego.Get(g, "/{id}", func(c fuego.ContextNoBody) (models.Contact, error) {
		ctx := c.Context()

		conn, put, err := db.Take(ctx)
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
		}
		defer put()

		r, err := q.ContactRead(conn, int64(c.PathParamInt("id")))
		if err != nil {
			if err == sql.ErrNoRows {
				return models.Contact{}, fuego.NotFoundError{}
			}
			return models.Contact{}, errors.WithStack(err)
		}

		out := models.Contact{
			CreatedAt: r.CreatedAt,
			Email:     r.Email,
			ID:        int(r.Id),
			Info:      r.Info,
			Name:      r.Name,
			Phone:     r.Phone,
			UpdatedAt: r.UpdatedAt,
		}

		return out, nil
	},
		fuego.OptionOverrideDescription(""),
		fuego.OptionSummary("get"),
	)

	fuego.Put(g, "/{id}", func(c fuego.ContextWithBody[models.ContactUpdateIn]) (models.Contact, error) {
		ctx := c.Request().Context()

		id := int64(c.PathParamInt("id"))

		in, err := c.Body()
		if err != nil {
			return models.Contact{}, err
		}

		conn, put, err := db.Take(ctx)
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
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
			return models.Contact{}, errors.WithStack(err)
		}

		r, err := q.ContactRead(conn, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return models.Contact{}, fuego.NotFoundError{}
			}
			return models.Contact{}, errors.WithStack(err)
		}

		out := models.Contact{
			CreatedAt: r.CreatedAt,
			Email:     r.Email,
			ID:        int(r.Id),
			Info:      r.Info,
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
