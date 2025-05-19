package api

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
	"golang.org/x/crypto/bcrypt"
	"zombiezen.com/go/sqlite"
)

const sessionDur = 30 * 24 * time.Hour

type UserCreateIn struct {
	Email    string `json:"email" format:"email"`
	Password string `json:"password" minLength:"8"`
}

type SessionGetIn struct {
	Session http.Cookie `cookie:"session"`
}

type SessionGetOut struct {
	Body    q.UserGetOut
	Session http.Cookie `cookie:"session"`
	Status  int
}

func users(a huma.API, db db.DB) {
	g := NewGroup(a, "/users")

	PostBody(g, "/", func(ctx context.Context, in UserCreateIn) (q.UserGetOut, error) {
		bs, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return q.UserGetOut{}, errors.WithStack(err)
		}

		conn, put, err := db.Take(ctx)
		if err != nil {
			return q.UserGetOut{}, errors.WithStack(err)
		}
		defer put()

		u, err := q.UserCreate(conn, q.UserCreateIn{
			Email:        in.Email,
			PasswordHash: string(bs),
		})
		if err != nil {
			return q.UserGetOut{}, errors.WithStack(err)
		}

		out, err := q.UserGet(conn, u.ID)
		if err != nil {
			return q.UserGetOut{}, errors.WithStack(err)
		}

		return *out, nil
	})

	DeleteIn(g, "/session", func(ctx context.Context, in SessionGetIn) error {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		s, err := sessionGet(conn, in.Session.Value)
		if err != nil {
			return errors.WithStack(err)
		}

		if err := q.SessionDeleteUser(conn, s.UserId); err != nil {
			return errors.WithStack(err)
		}

		return nil
	})

	PostCookie(g, "/session", func(ctx context.Context, in UserCreateIn) (http.Cookie, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return http.Cookie{}, errors.WithStack(err)
		}
		defer put()

		u, err := q.UserGetByEmail(conn, in.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				return http.Cookie{}, huma.Error401Unauthorized("invalid email or password")
			}

			return http.Cookie{}, errors.WithStack(err)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(in.Password)); err != nil {
			if err == bcrypt.ErrMismatchedHashAndPassword {
				return http.Cookie{}, huma.Error401Unauthorized("invalid email or password")
			}

			return http.Cookie{}, errors.WithStack(err)
		}

		bs := make([]byte, 20)
		rand.Read(bs)
		id := base32.StdEncoding.EncodeToString(bs)

		out, err := q.SessionCreate(conn, q.SessionCreateIn{
			ID:        sessionHash(id),
			UserId:    u.ID,
			ExpiresAt: time.Now().Add(sessionDur),
		})
		if err != nil {
			return http.Cookie{}, errors.WithStack(err)
		}

		return http.Cookie{
			Expires: out.ExpiresAt,
			Name:    "session",
			Value:   id,
		}, nil
	})

	GetInOut(g, "/session", func(ctx context.Context, in SessionGetIn) (SessionGetOut, error) {
		conn, put, err := db.Take(ctx)
		if err != nil {
			return SessionGetOut{}, errors.WithStack(err)
		}
		defer put()

		s, err := sessionGet(conn, in.Session.Value)
		if err != nil {
			return SessionGetOut{}, errors.WithStack(err)
		}

		if time.Now().After(s.ExpiresAt) {
			if err := q.SessionDelete(conn, s.ID); err != nil {
				return SessionGetOut{}, errors.WithStack(err)
			}

			return SessionGetOut{
				Status: http.StatusUnauthorized,
			}, nil
		}

		if time.Now().After(s.ExpiresAt.Add(-sessionDur / 2)) {
			s.ExpiresAt = time.Now().Add(sessionDur)
			if err := q.SessionUpdate(conn, q.SessionUpdateIn{
				ExpiresAt: s.ExpiresAt,
				ID:        s.ID,
			}); err != nil {
				return SessionGetOut{}, errors.WithStack(err)
			}
		}

		u, err := q.UserGet(conn, s.UserId)
		if err != nil {
			return SessionGetOut{}, errors.WithStack(err)
		}

		return SessionGetOut{
			Body: *u,
			Session: http.Cookie{
				Expires: s.ExpiresAt,
				Name:    "session",
				Value:   in.Session.Value,
			},
			Status: http.StatusOK,
		}, nil
	})
}

func sessionGet(conn *sqlite.Conn, v string) (q.SessionGetOut, error) {
	s, err := q.SessionGet(conn, sessionHash(v))
	if err != nil {
		if err == sql.ErrNoRows {
			return q.SessionGetOut{}, huma.Error401Unauthorized("invalid session")
		}

		return q.SessionGetOut{}, errors.WithStack(err)
	}

	return *s, nil
}

func sessionHash(v string) string {
	hash := sha256.Sum256([]byte(v))
	return hex.EncodeToString(hash[:])
}
