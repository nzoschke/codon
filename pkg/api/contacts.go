package api

import (
	"context"
	"database/sql"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/models"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
)

type ContactListOut struct {
	Contacts []models.Contact `json:"contacts"`
}

type ContactUpdateIn struct {
	Email string             `json:"email"`
	Info  models.ContactInfo `json:"info"`
	Name  string             `json:"name"`
	Phone string             `json:"phone"`
}

func contacts(a huma.API, db db.DB) {
	g := NewGroup(a, "/contacts")

	List(g, func(ctx context.Context, in struct{}) (ContactListOut, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return ContactListOut{}, errors.WithStack(err)
		}
		defer put()

		rows, err := q.ContactList(conn, 10)
		if err != nil {
			return ContactListOut{}, errors.WithStack(err)
		}

		out := ContactListOut{
			Contacts: []models.Contact{},
		}

		for _, r := range rows {
			out.Contacts = append(out.Contacts, models.Contact(r))
		}

		return out, nil
	})

	Post(g, func(ctx context.Context, in q.ContactCreateIn) (models.Contact, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
		}
		defer put()

		r, err := q.ContactCreate(conn, in)
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
		}

		return models.Contact(*r), nil
	})

	Delete(g, "/{id}", func(ctx context.Context, id int64) error {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		_, err = q.ContactRead(conn, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return huma.Error404NotFound("")
			}
			return errors.WithStack(err)
		}

		err = q.ContactDelete(conn, id)
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	})

	Get(g, "/{id}", func(ctx context.Context, id int64) (models.Contact, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
		}
		defer put()

		r, err := q.ContactRead(conn, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return models.Contact{}, huma.Error404NotFound("")
			}
			return models.Contact{}, errors.WithStack(err)
		}

		return models.Contact(*r), nil
	})

	Put(g, "/{id}", func(ctx context.Context, id int64, in ContactUpdateIn) (models.Contact, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
		}
		defer put()

		err = q.ContactUpdate(conn, q.ContactUpdateIn{
			Email: in.Email,
			ID:    id,
			Info:  in.Info,
			Name:  in.Name,
			Phone: in.Phone,
		})
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
		}

		r, err := q.ContactRead(conn, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return models.Contact{}, huma.Error404NotFound("")
			}
			return models.Contact{}, errors.WithStack(err)
		}

		return models.Contact(*r), nil
	})
}
