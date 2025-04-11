package db

import (
	"testing"

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

	found := false
	err = db.Exec(ctx, "SELECT email, name FROM users WHERE name = ?", []any{"user"}, func(stmt *sqlite.Stmt) error {
		found = true
		return nil
	})
	a.NoError(err)
	a.False(found)
}

func TestMigrate(t *testing.T) {
	ctx := t.Context()
	a := assert.New(t)

	db, err := New(ctx, "file::memory:?mode=memory&cache=shared")
	a.NoError(err)

	v, err := db.Version(ctx)
	a.NoError(err)
	a.Equal("3.46.0", v)

	ts, err := db.Schema(ctx)
	a.NoError(err)
	a.Equal([]string{"table/users"}, ts)
}
