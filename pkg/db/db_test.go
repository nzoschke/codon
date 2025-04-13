package db_test

import (
	"encoding/json"
	"testing"

	"github.com/nzoschke/codon/pkg/db"
	"github.com/stretchr/testify/assert"
	"zombiezen.com/go/sqlite"
)

type Contact struct {
	ID    int
	Email string
	Name  string
}

func TestCRUD(t *testing.T) {
	ctx := t.Context()
	a := assert.New(t)

	db, err := db.New(ctx, "file::memory:?mode=memory&cache=shared")
	a.NoError(err)

	// create
	err = db.Exec(ctx, "INSERT INTO contacts (email, name) VALUES (?, ?)", []any{"a@example.com", "Ann"}, nil)
	a.NoError(err)

	// list
	users := []Contact{}
	err = db.Exec(ctx, "SELECT email, id, name FROM contacts", nil, func(stmt *sqlite.Stmt) error {
		users = append(users, Contact{
			Email: stmt.ColumnText(0),
			ID:    stmt.ColumnInt(1),
			Name:  stmt.ColumnText(2),
		})
		return nil
	})
	a.NoError(err)

	a.Equal([]Contact{{
		Email: "a@example.com",
		ID:    1,
		Name:  "Ann",
	}}, users)

	// read
	user := Contact{}
	err = db.Exec(ctx, "SELECT email, id, name FROM contacts WHERE id = ?", []any{1}, func(stmt *sqlite.Stmt) error {
		user = Contact{
			Email: stmt.ColumnText(0),
			ID:    stmt.ColumnInt(1),
			Name:  stmt.ColumnText(2),
		}
		return nil
	})
	a.NoError(err)

	a.Equal(Contact{
		Email: "a@example.com",
		ID:    1,
		Name:  "Ann",
	}, user)

	// delete
	err = db.Exec(ctx, "DELETE FROM contacts WHERE id = ?", []any{1}, func(stmt *sqlite.Stmt) error {
		return nil
	})
	a.NoError(err)

	exists := false
	err = db.Exec(ctx, "SELECT id FROM contacts WHERE id = ?", []any{1}, func(stmt *sqlite.Stmt) error {
		exists = true
		return nil
	})
	a.NoError(err)
	a.False(exists)
}
func TestJSON(t *testing.T) {
	ctx := t.Context()
	a := assert.New(t)

	db, err := db.New(ctx, "file::memory:?mode=memory&cache=shared")
	a.NoError(err)

	meta := map[string]any{
		"age": 21,
	}

	bs, err := json.Marshal(meta)
	a.NoError(err)

	err = db.Exec(ctx, "INSERT INTO contacts (name, meta) VALUES (?, ?)", []any{"Ann", bs}, nil)
	a.NoError(err)

	bs = []byte{}
	err = db.Exec(ctx, "SELECT meta FROM contacts WHERE id = ?", []any{1}, func(stmt *sqlite.Stmt) error {
		bs = []byte(stmt.ColumnText(0))
		return nil
	})
	a.NoError(err)
	a.Equal(`{"age":21}`, string(bs))

	age := 0
	err = db.Exec(ctx, "SELECT meta->>'$.age' AS age FROM contacts WHERE id = ?", []any{1}, func(stmt *sqlite.Stmt) error {
		age = stmt.ColumnInt(0)
		return nil
	})
	a.NoError(err)
	a.Equal(21, age)

	err = db.Exec(ctx, "DELETE FROM contacts WHERE id = ?", []any{1}, func(stmt *sqlite.Stmt) error {
		return nil
	})
	a.NoError(err)
}

func TestMigrate(t *testing.T) {
	ctx := t.Context()
	a := assert.New(t)

	db, err := db.New(ctx, "file::memory:?mode=memory&cache=shared")
	a.NoError(err)

	v, err := db.Version(ctx)
	a.NoError(err)
	a.Equal("3.46.1", v)

	ts, err := db.Schema(ctx)
	a.NoError(err)
	a.Equal([]string{"table/contacts"}, ts)
}
