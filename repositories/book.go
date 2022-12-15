package repositories

import (
	"database/sql"
	"time"

	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/models"
)

func GetAllBook(db *sql.DB) (results []models.Book, err error) {
	sql := `SELECT * FROM books`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var book = models.Book{}

		err = rows.Scan(
			&book.Id, &book.Title, &book.YearPublished,
			&book.PublisherId, &book.SubjectId, &book.IsBorrowed,
			&book.BorrowedBy, &book.CreatedAt, &book.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, book)
	}

	return results, nil
}

func InsertBook(db *sql.DB, book models.Book) (data interface{}, err error) {
	sql := `INSERT INTO books (
		title, year_published, publisher_id, subject_id,
		is_borrowed, borrowed_by, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *`

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	errs := db.QueryRow(sql,
		book.Title, book.YearPublished, book.PublisherId,
		book.SubjectId, book.IsBorrowed, book.BorrowedBy,
		book.CreatedAt, book.UpdatedAt,
	).Scan(
		&book.Id, &book.Title, &book.YearPublished,
		&book.PublisherId, &book.SubjectId, &book.IsBorrowed,
		&book.BorrowedBy, &book.CreatedAt, &book.UpdatedAt,
	)

	if errs != nil {
		return nil, errs
	}

	return book, nil
}

func UpdateBook(db *sql.DB, book models.Book) (err error) {
	sql := `UPDATE books
	SET title = $2, year_published = $3, publisher_id = $4,
	subject_id = $5, is_borrowed = $6, borrowed_by = $7, updated_at = $8
	WHERE id = $1`

	book.UpdatedAt = time.Now()
	errs := db.QueryRow(sql,
		book.Id, book.Title, book.YearPublished, book.PublisherId,
		book.SubjectId, book.IsBorrowed, book.BorrowedBy, book.UpdatedAt,
	)

	return errs.Err()
}

func DeleteBook(db *sql.DB, book models.Book) (err error) {
	sql := `DELETE FROM books WHERE id = $1`

	errs := db.QueryRow(sql, book.Id)

	return errs.Err()
}
