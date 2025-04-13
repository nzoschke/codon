package db_test

import (
	"testing"
	"time"

	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/stretchr/testify/assert"
	"zombiezen.com/go/sqlite"
)

type Contact struct {
	Email string
	Name  string
}

func TestCRUDExec(t *testing.T) {
	ctx := t.Context()
	a := assert.New(t)

	db, err := db.New(ctx, "file::memory:?mode=memory&cache=shared")
	a.NoError(err)

	// create
	err = db.Exec(ctx, "INSERT INTO contacts (email, name) VALUES (?, ?)", []any{"user@example.com", "user"}, nil)
	a.NoError(err)

	// list
	users := []Contact{}
	err = db.Exec(ctx, "SELECT email, name FROM contacts", nil, func(stmt *sqlite.Stmt) error {
		users = append(users, Contact{
			Email: stmt.ColumnText(0),
			Name:  stmt.ColumnText(1),
		})
		return nil
	})
	a.NoError(err)

	a.Equal([]Contact{{
		Email: "user@example.com",
		Name:  "user",
	}}, users)

	// read
	user := Contact{}
	err = db.Exec(ctx, "SELECT email, name FROM contacts WHERE name = ?", []any{"user"}, func(stmt *sqlite.Stmt) error {
		user = Contact{
			Email: stmt.ColumnText(0),
			Name:  stmt.ColumnText(1),
		}
		return nil
	})
	a.NoError(err)

	a.Equal(Contact{
		Email: "user@example.com",
		Name:  "user",
	}, user)

	// delete
	err = db.Exec(ctx, "DELETE FROM contacts WHERE name = ?", []any{"user"}, func(stmt *sqlite.Stmt) error {
		return nil
	})
	a.NoError(err)

	exists := false
	err = db.Exec(ctx, "SELECT email, name FROM contacts WHERE name = ?", []any{"user"}, func(stmt *sqlite.Stmt) error {
		exists = true
		return nil
	})
	a.NoError(err)
	a.False(exists)
}

func TestCRUDQ(t *testing.T) {
	ctx := t.Context()
	a := assert.New(t)

	d, err := db.New(ctx, "file::memory:?mode=memory&cache=shared")
	a.NoError(err)

	conn, put, err := d.Take(ctx)
	a.NoError(err)
	defer put()

	out, err := q.ContactCreate(conn).Run(q.ContactCreateParams{
		Email: db.P("user@example.com"),
		Name:  "user",
	})
	a.NoError(err)

	a.Equal(time.Now().Format("2006-01-02"), out.CreatedAt.Format("2006-01-02"))

	a.Equal(&q.ContactCreateRes{
		CreatedAt: out.CreatedAt,
		Email:     db.P("user@example.com"),
		Id:        1,
		Name:      "user",
	}, out)

	rout, err := q.ContactRead(conn).Run(1)
	a.NoError(err)

	a.Equal(&q.ContactReadRes{
		CreatedAt: out.CreatedAt,
		Email:     db.P("user@example.com"),
		Id:        1,
		Name:      "user",
	}, rout)

	err = q.ContactUpdate(conn).Run(q.ContactUpdateParams{
		Email: db.P("user@new.com"),
		Name:  "user",
		Id:    1,
	})
	a.NoError(err)

	rout, err = q.ContactRead(conn).Run(1)
	a.NoError(err)

	a.Equal(&q.ContactReadRes{
		CreatedAt: out.CreatedAt,
		Email:     db.P("user@new.com"),
		Id:        1,
		Name:      "user",
	}, rout)

	err = q.ContactDelete(conn).Run(1)
	a.NoError(err)

	rout, err = q.ContactRead(conn).Run(1)
	a.NoError(err)
	a.Nil(rout)
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
