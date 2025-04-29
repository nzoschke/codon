package models

import "time"

type Contact struct {
	CreatedAt time.Time   `json:"created_at"`
	Email     string      `json:"email"`
	ID        int         `json:"id"`
	Info      ContactInfo `json:"info"`
	Name      string      `json:"name"`
	Phone     string      `json:"phone"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type ContactInfo struct {
	Age int `json:"age"`
}
