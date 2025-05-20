package db

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/nzoschke/codon/pkg/sql"
	"github.com/olekukonko/errors"
	"zombiezen.com/go/sqlite/sqlitemigration"
)

func (d *DB) migrate(ctx context.Context) error {
	dir := "schema"
	entries, err := fs.ReadDir(sql.SQL, dir)
	if err != nil {
		return errors.WithStack(err)
	}

	migrations := []string{}
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}

		bs, err := fs.ReadFile(sql.SQL, filepath.Join(dir, entry.Name()))
		if err != nil {
			return errors.WithStack(err)
		}

		migrations = append(migrations, string(bs))
	}

	conn, err := d.pool.Take(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	defer d.pool.Put(conn)

	if err := sqlitemigration.Migrate(ctx, conn,
		sqlitemigration.Schema{
			AppID:      0xc0d09, // codon
			Migrations: migrations,
		},
	); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
