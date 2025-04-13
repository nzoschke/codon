// Code generated by "sqlc-gen-zombiezen". DO NOT EDIT.

package q

import (
	"fmt"
	"zombiezen.com/go/sqlite"

	"time"
)

type ContactListRes struct {
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	Id        int64     `json:"id"`
	Meta      []byte    `json:"meta"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
}

type ContactListStmt struct {
	stmt *sqlite.Stmt
}

func ContactList(tx *sqlite.Conn) *ContactListStmt {
	// Prepare the statement into connection cache
	stmt := tx.Prep(`
SELECT
  created_at, email, id, meta, name, phone
FROM
  contacts
ORDER BY
  created_at DESC
LIMIT
  ?
    `)
	ps := &ContactListStmt{stmt: stmt}
	return ps
}

func (ps *ContactListStmt) Run(
	limit int64,
) (
	res []ContactListRes,
	err error,
) {
	defer ps.stmt.Reset()

	// Bind parameters
	ps.stmt.BindInt64(1, limit)

	// Execute the query
	for {
		if hasRow, err := ps.stmt.Step(); err != nil {
			return res, fmt.Errorf("failed to execute {{.Name.Lower}} SQL: %w", err)
		} else if !hasRow {
			break
		}

		row := ContactListRes{}
		row.CreatedAt = JulianDayToTime(ps.stmt.ColumnFloat(0))
		row.Email = ps.stmt.ColumnText(1)
		row.Id = ps.stmt.ColumnInt64(2)
		row.Meta = StmtBytesByCol(ps.stmt, 3)
		row.Name = ps.stmt.ColumnText(4)
		row.Phone = ps.stmt.ColumnText(5)
		res = append(res, row)
	}

	return res, nil
}

func OnceContactList(
	tx *sqlite.Conn,
	limit int64,
) (
	res []ContactListRes,
	err error,
) {
	ps := ContactList(tx)

	return ps.Run(
		limit,
	)
}
