package models

import "time"

type Contact struct {
	CreatedAt time.Time   `json:"created_at"`
	Email     string      `json:"email"`
	Id        int64       `json:"id"`
	Info      ContactInfo `json:"info"`
	Name      string      `json:"name"`
	Phone     string      `json:"phone"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type ContactInfo struct {
	Age int64 `json:"age"`
}
