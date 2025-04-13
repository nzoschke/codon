// Code generated by "sqlc-gen-zombiezen". DO NOT EDIT.

package q

import (
	"fmt"
	"zombiezen.com/go/sqlite"

	"time"
)

type ContactCreateRes struct {
	CreatedAt *time.Time `json:"created_at"`
	Email     *string    `json:"email"`
	Id        int64      `json:"id"`
	Meta      *[]byte    `json:"meta"`
	Name      string     `json:"name"`
	Phone     *string    `json:"phone"`
}

type ContactCreateParams struct {
	Email *string `json:"email"`
	Name  string  `json:"name"`
	Phone *string `json:"phone"`
	Meta  *[]byte `json:"meta"`
}

type ContactCreateStmt struct {
	stmt *sqlite.Stmt
}

func ContactCreate(tx *sqlite.Conn) *ContactCreateStmt {
	// Prepare the statement into connection cache
	stmt := tx.Prep(`
INSERT INTO
  contacts (email, name, phone, meta)
VALUES
  (?, ?, ?, ?)
RETURNING
  created_at, email, id, meta, name, phone
    `)
	ps := &ContactCreateStmt{stmt: stmt}
	return ps
}

func (ps *ContactCreateStmt) Run(
	params ContactCreateParams,
) (
	res *ContactCreateRes,
	err error,
) {
	defer ps.stmt.Reset()

	// Bind parameters
	if params.Email == nil {
		ps.stmt.BindNull(1)
	} else {
		ps.stmt.BindText(1, *params.Email)
	}
	ps.stmt.BindText(2, params.Name)
	if params.Phone == nil {
		ps.stmt.BindNull(3)
	} else {
		ps.stmt.BindText(3, *params.Phone)
	}
	if params.Meta == nil {
		ps.stmt.BindNull(4)
	} else {
		ps.stmt.BindBytes(4, *params.Meta)
	}

	// Execute the query
	if hasRow, err := ps.stmt.Step(); err != nil {
		return res, fmt.Errorf("failed to execute {{.Name.Lower}} SQL: %w", err)
	} else if hasRow {
		row := ContactCreateRes{}
		isNullCreatedAt := ps.stmt.ColumnIsNull(0)
		if !isNullCreatedAt {
			tmp := JulianDayToTime(ps.stmt.ColumnFloat(0))
			row.CreatedAt = &tmp
		}
		isNullEmail := ps.stmt.ColumnIsNull(1)
		if !isNullEmail {
			tmp := ps.stmt.ColumnText(1)
			row.Email = &tmp
		}
		row.Id = ps.stmt.ColumnInt64(2)
		isNullMeta := ps.stmt.ColumnIsNull(3)
		if !isNullMeta {
			tmp := StmtBytesByCol(ps.stmt, 3)
			row.Meta = &tmp
		}
		row.Name = ps.stmt.ColumnText(4)
		isNullPhone := ps.stmt.ColumnIsNull(5)
		if !isNullPhone {
			tmp := ps.stmt.ColumnText(5)
			row.Phone = &tmp
		}
		res = &row
	}

	return res, nil
}

func OnceContactCreate(
	tx *sqlite.Conn,
	params ContactCreateParams,
) (
	res *ContactCreateRes,
	err error,
) {
	ps := ContactCreate(tx)

	return ps.Run(
		params,
	)
}
