package models

import "time"

type Book struct {
	Id            int       `json:"id"`
	Title         string    `json:"title" form:"title" validate:"required"`
	Author        string    `json:"author" form:"author" validate:"required"`
	YearPublished int       `json:"year_published" form:"year_published" validate:"numeric"`
	PublisherId   int       `json:"publisher_id" form:"publisher_id" validate:"required,numeric"`
	SubjectId     int       `json:"subject_id" form:"subject_id" validate:"required,numeric"`
	IsBorrowed    bool      `json:"is_borrowed"`
	BorrowedBy    int       `json:"borrowed_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
