package db

import (
	"context"
	"fmt"

	"github.com/olekukonko/errors"
	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

type DB struct {
	path string
	pool *sqlitex.Pool
}

func New(ctx context.Context, path string) (DB, error) {
	pool, err := sqlitex.NewPool(path, sqlitex.PoolOptions{
		PrepareConn: func(conn *sqlite.Conn) error {
			return sqlitex.ExecuteTransient(conn, "PRAGMA foreign_keys = ON;", nil)
		},
	})
	if err != nil {
		return DB{}, errors.WithStack(err)
	}

	db := DB{
		path: path,
		pool: pool,
	}

	if err := db.migrate(ctx); err != nil {
		return DB{}, errors.WithStack(err)
	}

	return db, nil
}

func (d *DB) Exec(ctx context.Context, query string, args []any, fn func(stmt *sqlite.Stmt) error) error {
	conn, err := d.pool.Take(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	defer d.pool.Put(conn)

	err = sqlitex.ExecuteTransient(conn,
		query,
		&sqlitex.ExecOptions{
			Args:       args,
			ResultFunc: fn,
		})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d *DB) Schema(ctx context.Context) ([]string, error) {
	conn, err := d.pool.Take(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer d.pool.Put(conn)

	tables := []string{}
	if err := sqlitex.ExecuteTransient(conn,
		`SELECT "type", "name" FROM sqlite_master ORDER BY 1, 2`,
		&sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				tables = append(tables, fmt.Sprintf("%s/%s", stmt.ColumnText(0), stmt.ColumnText(1)))
				return nil
			},
		},
	); err != nil {
		return nil, errors.WithStack(err)
	}

	return tables, nil
}

func (d *DB) Version(ctx context.Context) (string, error) {
	conn, err := d.pool.Take(ctx)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer d.pool.Put(conn)

	version := ""
	if err := sqlitex.ExecuteTransient(conn,
		"select sqlite_version()",
		&sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				version = stmt.ColumnText(0)
				return nil
			},
		},
	); err != nil {
		return "", errors.WithStack(err)
	}

	return version, nil
}
