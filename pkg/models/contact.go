package models

import "time"

type Contact struct {
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	ID        int       `json:"id"`
	Info      Info      `json:"info"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ContactCreateIn struct {
	Email string `form:"email" json:"email"`
	Info  Info   `form:"info" json:"info"`
	Name  string `form:"name" json:"name"`
	Phone string `form:"phone" json:"phone"`
}

type ContactUpdateIn struct {
	Email string `json:"email"`
	Info  Info   `json:"info"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Info struct {
	Age int `json:"age"`
}
