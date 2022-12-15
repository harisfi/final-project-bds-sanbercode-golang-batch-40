package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username" form:"username" validate:"required,unique"`
	Password  string    `json:"password" form:"password" validate:"required"`
	Name      string    `json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
