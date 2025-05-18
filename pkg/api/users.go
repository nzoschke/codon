package api

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"encoding/hex"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
	"golang.org/x/crypto/bcrypt"
)

const sessionDur = 30 * 24 * time.Hour

type UserCreateIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

	DeleteIn(g, "/session", func(ctx context.Context, in SessionGetIn) error {
		value := in.Session.Value
		hash := sha256.Sum256([]byte(value))
		id := hex.EncodeToString(hash[:])

		conn, put, err := db.Take(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		defer put()

		if err := q.SessionDelete(conn, id); err != nil {
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

		u, err := q.UserGetPasswordHash(conn, in.Email)
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

		t, err := token()
		if err != nil {
			return http.Cookie{}, errors.WithStack(err)
		}

		hash := sha256.Sum256([]byte(t))
		id := hex.EncodeToString(hash[:])

		out, err := q.SessionCreate(conn, q.SessionCreateIn{
			ID:        id,
			UserId:    u.ID,
			ExpiresAt: time.Now().Add(sessionDur),
		})
		if err != nil {
			return http.Cookie{}, errors.WithStack(err)
		}

		return http.Cookie{
			Expires: out.ExpiresAt,
			Name:    "session",
			Value:   t,
		}, nil
	})

	GetInOut(g, "/session", func(ctx context.Context, in SessionGetIn) (SessionGetOut, error) {
		value := in.Session.Value
		hash := sha256.Sum256([]byte(value))
		id := hex.EncodeToString(hash[:])

		slog.Info("session", "value", in.Session.Value, "id", id)

		conn, put, err := db.Take(ctx)
		if err != nil {
			return SessionGetOut{}, errors.WithStack(err)
		}
		defer put()

		s, err := q.SessionGet(conn, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return SessionGetOut{}, huma.Error401Unauthorized("invalid session")
			}

			return SessionGetOut{}, errors.WithStack(err)
		}

		if s.ExpiresAt.Before(time.Now()) {
			if err := q.SessionDelete(conn, id); err != nil {
				return SessionGetOut{}, errors.WithStack(err)
			}
			return SessionGetOut{
				Status: http.StatusUnauthorized,
			}, nil
		}

		if s.ExpiresAt.Before(s.ExpiresAt.Add(-sessionDur / 2)) {
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

		slog.Info("session", "value", in.Session.Value, "id", id, "user_id", s.UserId)

		return SessionGetOut{
			Body: *u,
			Session: http.Cookie{
				Expires: s.ExpiresAt,
				Name:    "session",
				Value:   value,
			},
			Status: http.StatusOK,
		}, nil
	})
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
