package api

import (
	"context"
	"database/sql"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/models"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
)

type ContactCreateIn struct {
	Email string             `json:"email"`
	Info  models.ContactInfo `json:"info"`
	Name  string             `json:"name"`
	Phone string             `json:"phone"`
}

type ContactUpdateIn struct {
	Email string             `json:"email"`
	Info  models.ContactInfo `json:"info"`
	Name  string             `json:"name"`
	Phone string             `json:"phone"`
}

type GetContactOut struct {
	Body models.Contact `json:"contact"`
}

type ListContactsOut struct {
	Body struct {
		Contacts []models.Contact `json:"contacts"`
	}
}

func contacts(a huma.API, db db.DB) {
	g := huma.NewGroup(a, "/contacts")
	g.UseModifier(func(op *huma.Operation, next func(*huma.Operation)) {
		op.Path = strings.TrimSuffix(op.Path, "/")
		op.Tags = []string{"Contacts"}
		next(op)
	})

	huma.Get(g, "/", func(ctx context.Context, in *struct{}) (*ListContactsOut, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		defer put()

		rows, err := q.ContactList(conn, 10)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		contacts := []models.Contact{}
		for _, r := range rows {
			contacts = append(contacts, models.Contact{
				CreatedAt: r.CreatedAt,
				Email:     r.Email,
				ID:        int(r.Id),
				Info:      r.Info,
				Name:      r.Name,
				Phone:     r.Phone,
				UpdatedAt: r.UpdatedAt,
			})
		}

		out := &ListContactsOut{}
		out.Body.Contacts = contacts
		return out, nil
	})

	huma.Post(g, "/", func(ctx context.Context, in *struct {
		Body ContactCreateIn
	}) (*GetContactOut, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		defer put()

		r, err := q.ContactCreate(conn, q.ContactCreateIn{
			Email: in.Body.Email,
			Info:  in.Body.Info,
			Name:  in.Body.Name,
			Phone: in.Body.Phone,
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}

		out := GetContactOut{}
		out.Body = models.Contact{
			CreatedAt: r.CreatedAt,
			Email:     r.Email,
			ID:        int(r.Id),
			Info:      r.Info,
			Name:      r.Name,
			Phone:     r.Phone,
			UpdatedAt: r.UpdatedAt,
		}

		return &out, nil
	})

	huma.Delete(g, "/{id}", func(ctx context.Context, in *struct {
		ID int `path:"id"`
	}) (*struct{}, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		defer put()

		_, err = q.ContactRead(conn, int64(in.ID))
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, huma.Error404NotFound("")
			}
			return nil, errors.WithStack(err)
		}

		err = q.ContactDelete(conn, int64(in.ID))
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return nil, nil
	})

	huma.Get(g, "/{id}", func(ctx context.Context, in *struct {
		ID int `path:"id"`
	}) (*GetContactOut, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		defer put()

		r, err := q.ContactRead(conn, int64(in.ID))
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, huma.Error404NotFound("")
			}
			return nil, errors.WithStack(err)
		}

		out := &GetContactOut{}
		out.Body = models.Contact{
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

	huma.Put(g, "/{id}", func(ctx context.Context, in *struct {
		Body ContactUpdateIn
		ID   int64 `path:"id"`
	}) (*GetContactOut, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return nil, errors.WithStack(err)
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
			return nil, errors.WithStack(err)
		}

		r, err := q.ContactRead(conn, in.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, huma.Error404NotFound("")
			}
			return nil, errors.WithStack(err)
		}

		out := &GetContactOut{}
		out.Body = models.Contact{
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
