package db

import (
	"context"
	"log/slog"

	"github.com/olekukonko/errors"
	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

func New(ctx context.Context) error {
	conn, err := sqlite.OpenConn(":memory:", sqlite.OpenReadWrite)
	if err != nil {
		return errors.WithStack(err)
	}
	defer conn.Close()

	conn.SetInterrupt(ctx.Done())

	err = sqlitex.ExecuteTransient(conn, "SELECT 'hello, world';", &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			slog.Info("db", "col", stmt.ColumnText(0))
			return nil
		},
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
