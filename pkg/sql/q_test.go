package sql_test

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/nzoschke/codon/pkg/db"
	"github.com/nzoschke/codon/pkg/sql/models"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/stretchr/testify/assert"
)

const (
	secondsInADay      = 86400
	UnixEpochJulianDay = 2440587.5
)

func JulianDayToTime(d float64) time.Time {
	return time.Unix(int64((d-UnixEpochJulianDay)*secondsInADay), 0).UTC()
}

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

	created := JulianDayToTime(res.CreatedAt)

	a.Equal(time.Now().Format("2006-01-02"), created.Format("2006-01-02"))

	a.Equal(&q.ContactCreateOut{
		CreatedAt: res.CreatedAt,
		Email:     "a@example.com",
		Id:        1,
		Name:      "Ann",
		Meta:      []byte{},
		UpdatedAt: res.UpdatedAt,
	}, res)

	rres, err := q.ContactRead(conn, 1)
	a.NoError(err)

	a.Equal(&q.ContactReadOut{
		CreatedAt: res.CreatedAt,
		Email:     "a@example.com",
		Id:        1,
		Name:      "Ann",
		Meta:      []byte{},
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
		Meta:      []byte{},
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

	meta := map[string]any{
		"age": 21,
	}

	bs, err := json.Marshal(meta)
	a.NoError(err)

	res, err := q.ContactCreate(conn, q.ContactCreateIn{
		Email: "a@example.com",
		Meta:  bs,
		Name:  "Ann",
	})
	a.NoError(err)

	a.Equal(&q.ContactCreateOut{
		CreatedAt: res.CreatedAt,
		Email:     "a@example.com",
		Id:        1,
		Meta:      []byte(`{"age":21}`),
		Name:      "Ann",
		UpdatedAt: res.UpdatedAt,
	}, res)

	out, err := models.ToContact(q.Contact(*res))
	a.NoError(err)

	a.Equal(models.Contact{
		CreatedAt: out.CreatedAt,
		Email:     "a@example.com",
		Id:        1,
		Meta: map[string]any{
			"age": float64(21),
		},
		Name:      "Ann",
		UpdatedAt: out.UpdatedAt,
	}, out)

	age, err := q.ContactAge(conn, 1)
	a.NoError(err)

	a.Equal(int64(21), age)
}
