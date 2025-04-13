package sql_test

import (
	"encoding/json"
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

	out, err := q.ContactCreate(conn).Run(q.ContactCreateParams{
		Email: "a@example.com",
		Name:  "Ann",
	})
	a.NoError(err)

	a.Equal(time.Now().Format("2006-01-02"), out.CreatedAt.Format("2006-01-02"))

	a.Equal(&q.ContactCreateRes{
		CreatedAt: out.CreatedAt,
		Email:     "a@example.com",
		Id:        1,
		Name:      "Ann",
		UpdatedAt: out.UpdatedAt,
	}, out)

	rout, err := q.ContactRead(conn).Run(1)
	a.NoError(err)

	a.Equal(&q.ContactReadRes{
		CreatedAt: out.CreatedAt,
		Email:     "a@example.com",
		Id:        1,
		Name:      "Ann",
		UpdatedAt: out.UpdatedAt,
	}, rout)

	// wait for CURRENT_TIMESTAMP to advance
	time.Sleep(2 * time.Second)

	err = q.ContactUpdate(conn).Run(q.ContactUpdateParams{
		Email: "a@new.com",
		Name:  "Ann",
		Id:    1,
	})
	a.NoError(err)

	rout, err = q.ContactRead(conn).Run(1)
	a.NoError(err)

	a.Equal(&q.ContactReadRes{
		CreatedAt: rout.CreatedAt,
		Email:     "a@new.com",
		Id:        1,
		Name:      "Ann",
		UpdatedAt: rout.UpdatedAt,
	}, rout)

	a.True(rout.UpdatedAt.After(out.UpdatedAt))

	err = q.ContactDelete(conn).Run(1)
	a.NoError(err)

	rout, err = q.ContactRead(conn).Run(1)
	a.NoError(err)
	a.Nil(rout)
}

func TestJSON(t *testing.T) {
	ctx := t.Context()
	a := assert.New(t)

	d, err := db.New(ctx, "file::memory:?mode=memory&cache=shared")
	a.NoError(err)

	conn, put, err := d.Take(ctx)
	a.NoError(err)
	defer put()

	meta := map[string]any{
		"age": 21,
	}

	bs, err := json.Marshal(meta)
	a.NoError(err)

	res, err := q.ContactCreate(conn).Run(q.ContactCreateParams{
		Email: "a@example.com",
		Meta:  bs,
		Name:  "Ann",
	})
	a.NoError(err)

	a.Equal(&q.ContactCreateRes{
		CreatedAt: res.CreatedAt,
		Email:     "a@example.com",
		Id:        1,
		Meta:      []byte(`{"age":21}`),
		Name:      "Ann",
		UpdatedAt: res.UpdatedAt,
	}, res)

	out, err := models.ToContact(*res)
	a.NoError(err)

	a.Equal(models.Contact{
		CreatedAt: res.CreatedAt,
		Email:     "a@example.com",
		Id:        1,
		Meta: map[string]any{
			"age": float64(21),
		},
		Name:      "Ann",
		UpdatedAt: res.UpdatedAt,
	}, out)

	age, err := q.ContactAge(conn).Run(1)
	a.NoError(err)

	a.Equal(int64(21), age)
}
