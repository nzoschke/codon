package contacts

import "github.com/nzoschke/codon/pkg/sql/models"

type Contact struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ContactCreateIn struct {
	Email string      `form:"email" json:"email"`
	Info  models.Info `form:"info" json:"info"`
	Name  string      `form:"name" json:"name"`
	Phone string      `form:"phone" json:"phone"`
}

type ContactUpdateIn struct {
	Email string      `json:"email"`
	Info  models.Info `json:"info"`
	Name  string      `json:"name"`
	Phone string      `json:"phone"`
}

type Contacts interface {
	Get(id string) (Contact, error)
	Create(in ContactCreateIn) (Contact, error)
	List() ([]Contact, error)
	Update(id string, in ContactUpdateIn) (Contact, error)
	Delete(id string) (any, error)
}
