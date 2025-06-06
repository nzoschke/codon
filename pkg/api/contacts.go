package api

import (
	"context"
	"database/sql"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/sql/models"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
	"zombiezen.com/go/sqlite"
)

type ContactListIn struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type ContactListOut struct {
	Contacts []q.Contact `json:"contacts"`
}

type ContactUpdateIn struct {
	Email string             `json:"email"`
	Info  models.ContactInfo `json:"info"`
	Name  string             `json:"name"`
	Phone string             `json:"phone"`
}

func contacts(a huma.API, db db.DB) {
	g := NewGroup(a, "/contacts")

	DeleteID(g, "/{id}", func(ctx context.Context, id int64) error {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		_, err = contactRead(conn, id)
		if err != nil {
			return errors.WithStack(err)
		}

		err = q.ContactDelete(conn, id)
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	})

	GetID(g, "/{id}", func(ctx context.Context, id int64) (q.Contact, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return q.Contact{}, errors.WithStack(err)
		}
		defer put()

		c, err := contactRead(conn, id)
		if err != nil {
			return q.Contact{}, errors.WithStack(err)
		}

		return c, nil
	})

	List(g, "/", func(ctx context.Context, in ContactListIn) (ContactListOut, error) {
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
			Contacts: []q.Contact{},
		}

		for _, r := range rows {
			out.Contacts = append(out.Contacts, q.Contact(r))
		}

		return out, nil
	})

	PostBody(g, "/", func(ctx context.Context, in q.ContactCreateIn) (q.Contact, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return q.Contact{}, errors.WithStack(err)
		}
		defer put()

		r, err := q.ContactCreate(conn, in)
		if err != nil {
			return q.Contact{}, errors.WithStack(err)
		}

		return q.Contact(*r), nil
	})

	PutID(g, "/{id}", func(ctx context.Context, id int64, in ContactUpdateIn) (q.Contact, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return q.Contact{}, errors.WithStack(err)
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
			return q.Contact{}, errors.WithStack(err)
		}

		c, err := contactRead(conn, id)
		if err != nil {
			return q.Contact{}, errors.WithStack(err)
		}

		return c, nil
	})
}

func contactRead(conn *sqlite.Conn, id int64) (q.Contact, error) {
	r, err := q.ContactRead(conn, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return q.Contact{}, huma.Error404NotFound("")
		}
		return q.Contact{}, errors.WithStack(err)
	}

	return q.Contact(*r), nil
}
