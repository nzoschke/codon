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

func contacts(a huma.API, db db.DB) {
	g := NewGroup(a, "/contacts")

	List(g, "", func(ctx context.Context, in struct{}) (ContactListOut, error) {
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
			out.Contacts = append(out.Contacts, models.Contact{
				CreatedAt: r.CreatedAt,
				Email:     r.Email,
				ID:        int(r.Id),
				Info:      r.Info,
				Name:      r.Name,
				Phone:     r.Phone,
				UpdatedAt: r.UpdatedAt,
			})
		}

		return out, nil
	})

	Post(g, "", func(ctx context.Context, in q.ContactCreateIn) (models.Contact, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
		}
		defer put()

		r, err := q.ContactCreate(conn, in)
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
	})

	Delete(g, "/{id}", func(ctx context.Context, in struct {
		ID int `path:"id"`
	}) error {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		_, err = q.ContactRead(conn, int64(in.ID))
		if err != nil {
			if err == sql.ErrNoRows {
				return huma.Error404NotFound("")
			}
			return errors.WithStack(err)
		}

		err = q.ContactDelete(conn, int64(in.ID))
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	})

	Get(g, "/{id}", func(ctx context.Context, in struct {
		ID int `path:"id"`
	}) (models.Contact, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
		}
		defer put()

		r, err := q.ContactRead(conn, int64(in.ID))
		if err != nil {
			if err == sql.ErrNoRows {
				return models.Contact{}, huma.Error404NotFound("")
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
	})

	Put(g, "/{id}", func(ctx context.Context, in struct {
		Body q.ContactUpdateIn
		ID   int64 `path:"id"`
	}) (models.Contact, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
		}
		defer put()

		err = q.ContactUpdate(conn, q.ContactUpdateIn{
			Email: in.Body.Email,
			Id:    in.ID,
			Info:  in.Body.Info,
			Name:  in.Body.Name,
			Phone: in.Body.Phone,
		})
		if err != nil {
			return models.Contact{}, errors.WithStack(err)
		}

		r, err := q.ContactRead(conn, in.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return models.Contact{}, huma.Error404NotFound("")
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
	})
}
