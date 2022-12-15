package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username" form:"username"`
	Password  string    `json:"password" form:"password"`
	Name      string    `json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
