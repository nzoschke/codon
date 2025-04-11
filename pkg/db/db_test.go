package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCRUD(t *testing.T) {
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
