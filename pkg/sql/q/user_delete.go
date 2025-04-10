// Code generated by "sqlc-gen-zombiezen". DO NOT EDIT.

package q

import (
	"fmt"
	"zombiezen.com/go/sqlite"
)

type UserDeleteStmt struct {
	stmt *sqlite.Stmt
}

func UserDelete(tx *sqlite.Conn) *UserDeleteStmt {
	// Prepare the statement into connection cache
	stmt := tx.Prep(`
DELETE FROM
  users
WHERE
  id = ?
    `)
	ps := &UserDeleteStmt{stmt: stmt}
	return ps
}

func (ps *UserDeleteStmt) Run(
	id int64,
) (
	err error,
) {
	defer ps.stmt.Reset()

	// Bind parameters
	ps.stmt.BindInt64(1, id)

	// Execute the query
	if _, err := ps.stmt.Step(); err != nil {
		return fmt.Errorf("failed to execute userdelete SQL: %w", err)
	}

	return nil
}

func OnceUserDelete(
	tx *sqlite.Conn,
	id int64,
) (
	err error,
) {
	ps := UserDelete(tx)

	return ps.Run(
		id,
	)
}
