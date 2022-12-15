package models

import "time"

type Book struct {
	Id            int       `json:"id"`
	Title         string    `json:"title" form:"title"`
	Author        string    `json:"author" form:"author"`
	YearPublished int       `json:"year_published" form:"year_published"`
	PublisherId   int       `json:"publisher_id" form:"publisher_id"`
	SubjectId     int       `json:"subject_id" form:"subject_id"`
	IsBorrowed    int       `json:"is_borrowed"`
	BorrowedBy    int       `json:"borrowed_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
