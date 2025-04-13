package models

import (
	"encoding/json"
	"time"

	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
)

type Contact struct {
	CreatedAt time.Time      `json:"created_at"`
	Email     string         `json:"email"`
	Id        int64          `json:"id"`
	Meta      map[string]any `json:"meta"`
	Name      string         `json:"name"`
	Phone     string         `json:"phone"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func ToContact(r q.ContactCreateRes) (Contact, error) {
	c := Contact{
		CreatedAt: r.CreatedAt,
		Email:     r.Email,
		Id:        r.Id,
		Meta:      map[string]any{},
		Name:      r.Name,
		Phone:     r.Phone,
		UpdatedAt: r.UpdatedAt,
	}

	if err := json.Unmarshal(r.Meta, &c.Meta); err != nil {
		return Contact{}, errors.WithStack(err)
	}

	return c, nil
}
