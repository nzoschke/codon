package db

import (
	"testing"

	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/stretchr/testify/assert"
	"zombiezen.com/go/sqlite"
)

type User struct {
	Email string
	Name  string
}

func TestCRUD(t *testing.T) {
	ctx := t.Context()
	a := assert.New(t)

	db, err := New(ctx, "file::memory:?mode=memory&cache=shared")
	a.NoError(err)

	// create
	err = db.Exec(ctx, "INSERT INTO users (email, name) VALUES (?, ?)", []any{"user@example.com", "user"}, nil)
	a.NoError(err)

	// list
	users := []User{}
	err = db.Exec(ctx, "SELECT email, name FROM users", nil, func(stmt *sqlite.Stmt) error {
		users = append(users, User{
			Email: stmt.ColumnText(0),
			Name:  stmt.ColumnText(1),
		})
		return nil
	})
	a.NoError(err)

	a.Equal([]User{{
		Email: "user@example.com",
		Name:  "user",
	}}, users)

	// read
	user := User{}
	err = db.Exec(ctx, "SELECT email, name FROM users WHERE name = ?", []any{"user"}, func(stmt *sqlite.Stmt) error {
		user = User{
			Email: stmt.ColumnText(0),
			Name:  stmt.ColumnText(1),
		}
		return nil
	})
	a.NoError(err)

	a.Equal(User{
		Email: "user@example.com",
		Name:  "user",
	}, user)

	// delete
	err = db.Exec(ctx, "DELETE FROM users WHERE name = ?", []any{"user"}, func(stmt *sqlite.Stmt) error {
		return nil
	})
	a.NoError(err)

	exists := false
	err = db.Exec(ctx, "SELECT email, name FROM users WHERE name = ?", []any{"user"}, func(stmt *sqlite.Stmt) error {
		exists = true
		return nil
	})
	a.NoError(err)
	a.False(exists)
}

func TestCRUDQ(t *testing.T) {
	ctx := t.Context()
	a := assert.New(t)

	db, err := New(ctx, "file::memory:?mode=memory&cache=shared")
	a.NoError(err)

	conn, put, err := db.Take(ctx)
	a.NoError(err)
	defer put()

	out, err := q.UserCreate(conn).Run(q.UserCreateParams{
		Email: "user@example.com",
		Name:  "user",
	})
	a.NoError(err)

	a.Equal(&q.UserCreateRes{
		CreatedAt: out.CreatedAt,
		Email:     "user@example.com",
		Id:        1,
		Name:      "user",
	}, out)

	rout, err := q.UserRead(conn).Run(1)
	a.NoError(err)

	a.Equal(&q.UserReadRes{
		CreatedAt: out.CreatedAt,
		Email:     "user@example.com",
		Id:        1,
		Name:      "user",
	}, rout)

	err = q.UserUpdate(conn).Run(q.UserUpdateParams{
		Email: "user@new.com",
		Name:  "user",
		Id:    1,
	})
	a.NoError(err)

	rout, err = q.UserRead(conn).Run(1)
	a.NoError(err)

	a.Equal(&q.UserReadRes{
		CreatedAt: out.CreatedAt,
		Email:     "user@new.com",
		Id:        1,
		Name:      "user",
	}, rout)

	err = q.UserDelete(conn).Run(1)
	a.NoError(err)

	rout, err = q.UserRead(conn).Run(1)
	a.NoError(err)
	a.Nil(rout)
}

func TestMigrate(t *testing.T) {
	ctx := t.Context()
	a := assert.New(t)

	db, err := New(ctx, "file::memory:?mode=memory&cache=shared")
	a.NoError(err)

	v, err := db.Version(ctx)
	a.NoError(err)
	a.Equal("3.46.1", v)

	ts, err := db.Schema(ctx)
	a.NoError(err)
	a.Equal([]string{"table/users"}, ts)
}
