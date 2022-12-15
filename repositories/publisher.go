package repositories

import (
	"database/sql"
	"time"

	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/models"
)

func GetAllPublisher(db *sql.DB) (results []models.Publisher, err error) {
	sql := `SELECT * FROM publishers`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var publisher = models.Publisher{}

		err = rows.Scan(&publisher.Id, &publisher.Name, &publisher.CreatedAt, &publisher.UpdatedAt)
		if err != nil {
			return nil, err
		}

		results = append(results, publisher)
	}

	return results, nil
}

func InsertPublisher(db *sql.DB, publisher models.Publisher) (data interface{}, err error) {
	sql := `INSERT INTO publishers (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING *`

	publisher.CreatedAt = time.Now()
	publisher.UpdatedAt = time.Now()

	errs := db.QueryRow(sql, publisher.Name, publisher.CreatedAt, publisher.UpdatedAt).Scan(
		&publisher.Id, &publisher.Name, &publisher.CreatedAt, &publisher.UpdatedAt)

	if errs != nil {
		return nil, errs
	}

	return publisher, nil
}

func UpdatePublisher(db *sql.DB, publisher models.Publisher) (err error) {
	sql := `UPDATE publishers SET name = $2, updated_at = $3 WHERE id = $1`

	publisher.UpdatedAt = time.Now()
	errs := db.QueryRow(sql, publisher.Id, publisher.Name, publisher.UpdatedAt)

	return errs.Err()
}

func DeletePublisher(db *sql.DB, publisher models.Publisher) (err error) {
	sql := `DELETE FROM publishers WHERE id = $1`

	errs := db.QueryRow(sql, publisher.Id)

	return errs.Err()
}
