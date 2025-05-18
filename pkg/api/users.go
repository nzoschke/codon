package api

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserCreateIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func users(a huma.API, db db.DB) {
	g := NewGroup(a, "/users")

	PostBody(g, "/", func(ctx context.Context, in UserCreateIn) (q.UserCreateOut, error) {
		bs, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return q.UserCreateOut{}, errors.WithStack(err)
		}

		conn, put, err := db.Take(ctx)
		if err != nil {
			return q.UserCreateOut{}, errors.WithStack(err)
		}
		defer put()

		out, err := q.UserCreate(conn, q.UserCreateIn{
			Email:        in.Email,
			PasswordHash: string(bs),
		})
		if err != nil {
			return q.UserCreateOut{}, errors.WithStack(err)
		}

		return *out, nil
	})

	PostBody(g, "/check", func(ctx context.Context, in UserCreateIn) (q.UserCreateOut, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return q.UserCreateOut{}, errors.WithStack(err)
		}
		defer put()

		out, err := q.UserGet(conn, in.Email)
		if err != nil {
			return q.UserCreateOut{}, errors.WithStack(err)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(out.PasswordHash), []byte(in.Password)); err != nil {
			if err == bcrypt.ErrMismatchedHashAndPassword {
				return q.UserCreateOut{}, huma.Error401Unauthorized("invalid email or password")
			}

			return q.UserCreateOut{}, errors.WithStack(err)
		}

		return q.UserCreateOut{
			Email: out.Email,
			ID:    out.ID,
		}, nil
	})

	// Post(g, "", func(ctx context.Context) (q.SessionCreateOut, error) {
	// 	t, err := token()
	// 	if err != nil {
	// 		return q.SessionCreateOut{}, errors.WithStack(err)
	// 	}

	// 	conn, put, err := db.Take(ctx)
	// 	if err != nil {
	// 		return q.SessionCreateOut{}, errors.WithStack(err)
	// 	}
	// 	defer put()

	// 	out, err := q.SessionCreate(conn, q.SessionCreateIn{
	// 		ID:        t,
	// 		UserId:    1,
	// 		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	// 	})
	// 	if err != nil {
	// 		return q.SessionCreateOut{}, errors.WithStack(err)
	// 	}

	// 	return *out, nil
	// })
}

func token() (string, error) {
	bs := make([]byte, 20)
	_, err := rand.Read(bs)
	if err != nil {
		return "", errors.WithStack(err)
	}

	token := strings.ToLower(strings.TrimRight(base32.StdEncoding.EncodeToString(bs), "="))

	return token, nil
}
