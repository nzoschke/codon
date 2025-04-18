package sql_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/sql/models"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/stretchr/testify/assert"
)

func TestCRUD(t *testing.T) {
	ctx := t.Context()
	a := assert.New(t)

	d, err := db.New(ctx, "file::memory:?mode=memory&cache=shared")
	a.NoError(err)

	conn, put, err := d.Take(ctx)
	a.NoError(err)
	defer put()

	res, err := q.ContactCreate(conn, q.ContactCreateIn{
		Email: "a@example.com",
		Name:  "Ann",
	})
	a.NoError(err)

	a.Equal(time.Now().UTC().Format("2006-01-02"), res.CreatedAt.Format("2006-01-02"))

	a.Equal(&q.ContactCreateOut{
		CreatedAt: res.CreatedAt,
		Email:     "a@example.com",
		Id:        1,
		Name:      "Ann",
		Meta:      models.Meta{},
		UpdatedAt: res.UpdatedAt,
	}, res)

	rres, err := q.ContactRead(conn, 1)
	a.NoError(err)

	a.Equal(&q.ContactReadOut{
		CreatedAt: res.CreatedAt,
		Email:     "a@example.com",
		Id:        1,
		Name:      "Ann",
		Meta:      models.Meta{},
		UpdatedAt: res.UpdatedAt,
	}, rres)

	// wait for CURRENT_TIMESTAMP to advance
	time.Sleep(1001 * time.Millisecond)

	err = q.ContactUpdate(conn, q.ContactUpdateIn{
		Email: "a@new.com",
		Name:  "Ann",
		Id:    1,
	})
	a.NoError(err)

	rres, err = q.ContactRead(conn, 1)
	a.NoError(err)

	a.Equal(&q.ContactReadOut{
		CreatedAt: res.CreatedAt,
		Email:     "a@new.com",
		Id:        1,
		Name:      "Ann",
		Meta:      models.Meta{},
		UpdatedAt: rres.UpdatedAt,
	}, rres)

	a.Greater(rres.UpdatedAt, res.UpdatedAt)

	err = q.ContactDelete(conn, 1)
	a.NoError(err)

	rres, err = q.ContactRead(conn, 1)
	a.EqualError(err, sql.ErrNoRows.Error())
	a.Nil(rres)
}

func TestJSON(t *testing.T) {
	ctx := t.Context()
	a := assert.New(t)

	d, err := db.New(ctx, "file::memory:?mode=memory&cache=shared")
	a.NoError(err)

	conn, put, err := d.Take(ctx)
	a.NoError(err)
	defer put()

	res, err := q.ContactCreate(conn, q.ContactCreateIn{
		Email: "a@example.com",
		Meta: models.Meta{
			Age: 21,
		},
		Name: "Ann",
	})
	a.NoError(err)

	a.Equal(&q.ContactCreateOut{
		CreatedAt: res.CreatedAt,
		Email:     "a@example.com",
		Id:        1,
		Meta: models.Meta{
			Age: 21,
		},
		Name:      "Ann",
		UpdatedAt: res.UpdatedAt,
	}, res)

	age, err := q.ContactAge(conn, 1)
	a.NoError(err)

	a.Equal(int64(21), age)
}
