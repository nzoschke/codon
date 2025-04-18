package models

import "time"

type Contact struct {
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	Id        int64     `json:"id"`
	Info      Info      `json:"info"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Info struct {
	Age int `json:"age"`
}
