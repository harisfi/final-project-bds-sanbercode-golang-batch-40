package repositories

import (
	"database/sql"
	"time"

	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/models"
)

func GetAllSubject(db *sql.DB) (results []models.Subject, err error) {
	sql := `SELECT * FROM subjects`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var subject = models.Subject{}

		err = rows.Scan(&subject.Id, &subject.Name, &subject.CreatedAt, &subject.UpdatedAt)
		if err != nil {
			return nil, err
		}

		results = append(results, subject)
	}

	return results, nil
}

func InsertSubject(db *sql.DB, subject models.Subject) (data interface{}, err error) {
	sql := `INSERT INTO subjects (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING *`

	subject.CreatedAt = time.Now()
	subject.UpdatedAt = time.Now()

	errs := db.QueryRow(sql, subject.Name, subject.CreatedAt, subject.UpdatedAt).Scan(
		&subject.Id, &subject.Name, &subject.CreatedAt, &subject.UpdatedAt)

	if errs != nil {
		return nil, errs
	}

	return subject, nil
}

func UpdateSubject(db *sql.DB, subject models.Subject) (err error) {
	sql := `UPDATE subjects SET name = $2, updated_at = $3 WHERE id = $1`

	subject.UpdatedAt = time.Now()
	errs := db.QueryRow(sql, subject.Id, subject.Name, subject.UpdatedAt)

	return errs.Err()
}

func DeleteSubject(db *sql.DB, subject models.Subject) (err error) {
	sql := `DELETE FROM subjects WHERE id = $1`

	errs := db.QueryRow(sql, subject.Id)

	return errs.Err()
}
