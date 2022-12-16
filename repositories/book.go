package repositories

import (
	"database/sql"
	"time"

	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/models"
)

func GetAllBook(db *sql.DB) (results []models.Book, err error) {
	query := `SELECT * FROM books`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var book = models.Book{}
		var borrower sql.NullInt64

		err = rows.Scan(
			&book.Id, &book.Title, &book.Author, &book.YearPublished,
			&book.PublisherId, &book.SubjectId, &book.IsBorrowed,
			&borrower, &book.CreatedAt, &book.UpdatedAt,
		)
		book.BorrowedBy = int(borrower.Int64)

		if err != nil {
			return nil, err
		}

		results = append(results, book)
	}

	return results, nil
}

func GetBookById(db *sql.DB, id int) (result models.Book, err error) {
	var book = models.Book{}
	query := `SELECT * FROM books WHERE id = $1`

	var borrower sql.NullInt64

	errs := db.QueryRow(query, id).Scan(
		&book.Id, &book.Title, &book.Author, &book.YearPublished,
		&book.PublisherId, &book.SubjectId, &book.IsBorrowed,
		&borrower, &book.CreatedAt, &book.UpdatedAt,
	)
	book.BorrowedBy = int(borrower.Int64)

	if errs != nil {
		return book, errs
	}

	return book, nil
}

func InsertBook(db *sql.DB, book models.Book) (data interface{}, err error) {
	query := `INSERT INTO books 
	(
		title, author, year_published, publisher_id, subject_id, created_at, updated_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	var borrower sql.NullInt64

	errs := db.QueryRow(query,
		book.Title, book.Author, book.YearPublished, book.PublisherId,
		book.SubjectId, book.CreatedAt, book.UpdatedAt,
	).Scan(
		&book.Id, &book.Title, &book.Author, &book.YearPublished,
		&book.PublisherId, &book.SubjectId, &book.IsBorrowed,
		&borrower, &book.CreatedAt, &book.UpdatedAt,
	)

	if errs != nil {
		return nil, errs
	}

	return book, nil
}

func UpdateBook(db *sql.DB, book models.Book) (err error) {
	sql := `UPDATE books
	SET title = $2, author = $3, year_published = $4, publisher_id = $5,
	subject_id = $6, is_borrowed = $7, borrowed_by = $8, updated_at = $9
	WHERE id = $1`

	book.UpdatedAt = time.Now()
	errs := db.QueryRow(sql,
		book.Id, book.Title, book.Author, book.YearPublished, book.PublisherId,
		book.SubjectId, book.IsBorrowed, book.BorrowedBy, book.UpdatedAt,
	)

	return errs.Err()
}

func DeleteBook(db *sql.DB, book models.Book) (err error) {
	sql := `DELETE FROM books WHERE id = $1`

	errs := db.QueryRow(sql, book.Id)

	return errs.Err()
}
