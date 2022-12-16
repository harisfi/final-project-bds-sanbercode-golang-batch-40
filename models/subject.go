package models

import "time"

type Subject struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" form:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
