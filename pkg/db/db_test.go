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
	err = db.Exec(ctx, "INSERT INTO contacts (email, info, name, phone) VALUES (?, ?, ?, ?)", []any{"a@example.com", []byte("{}"), "Ann", ""}, nil)
	a.NoError(err)

	// list
	rows := []Contact{}
	err = db.Exec(ctx, "SELECT email, id, name FROM contacts", nil, func(stmt *sqlite.Stmt) error {
		rows = append(rows, Contact{
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
	}}, rows)

	// read
	row := Contact{}
	err = db.Exec(ctx, "SELECT email, id, name FROM contacts WHERE id = ?", []any{1}, func(stmt *sqlite.Stmt) error {
		row = Contact{
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
	}, row)

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

	info := map[string]any{
		"age": 21,
	}

	bs, err := json.Marshal(info)
	a.NoError(err)

	err = db.Exec(ctx, "INSERT INTO contacts (email, info, name, phone) VALUES (?, ?, ?, ?)", []any{"a@example.com", bs, "Ann", ""}, nil)
	a.NoError(err)

	bs = []byte{}
	err = db.Exec(ctx, "SELECT info FROM contacts WHERE id = ?", []any{1}, func(stmt *sqlite.Stmt) error {
		bs = []byte(stmt.ColumnText(0))
		return nil
	})
	a.NoError(err)
	a.Equal(`{"age":21}`, string(bs))

	age := 0
	err = db.Exec(ctx, "SELECT info->>'$.age' AS age FROM contacts WHERE id = ?", []any{1}, func(stmt *sqlite.Stmt) error {
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
	a.Equal([]string{"index/sqlite_autoindex_sessions_1", "index/sqlite_autoindex_users_1", "table/contacts", "table/sessions", "table/users"}, ts)
}
