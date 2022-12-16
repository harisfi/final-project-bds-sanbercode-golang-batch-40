package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/middlewares"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/models"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/utils"
)

func GetAllUser(db *sql.DB) (results []models.User, err error) {
	sql := `SELECT * FROM users`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user = models.User{}

		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		results = append(results, user)
	}

	return results, nil
}

func InsertUser(db *sql.DB, user models.User) (data interface{}, err error) {
	sql := `INSERT INTO users (
		username, password, name, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5) RETURNING *`

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	errs := db.QueryRow(sql, user.Username, user.Password, user.Name, user.CreatedAt, user.UpdatedAt).
		Scan(&user.Id, &user.Username, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	if errs != nil {
		return nil, errs
	}

	return user, nil
}

func UpdateUser(db *sql.DB, user models.User) (err error) {
	sql := `UPDATE users SET username = $2, password = $3, name = $4, updated_at = $5 WHERE id = $1`

	user.UpdatedAt = time.Now()
	errs := db.QueryRow(sql, user.Id, user.Username, user.Password, user.Name, user.UpdatedAt)

	return errs.Err()
}

func DeleteUser(db *sql.DB, user models.User) (err error) {
	sql := `DELETE FROM users WHERE id = $1`

	errs := db.QueryRow(sql, user.Id)

	return errs.Err()
}

func LoginUser(db *sql.DB, user models.User) (interface{}, error) {
	dUser := models.User{}
	sql := `SELECT id, password FROM users WHERE username = $1`

	errs := db.QueryRow(sql, user.Username).Scan(&dUser.Id, &dUser.Password)

	if errs != nil {
		return nil, errs
	}

	if utils.CheckPasswordHash(user.Password, dUser.Password) {
		token, e := middlewares.CreateToken(dUser.Id)
		if e != nil {
			return nil, e
		}

		return token, nil
	} else {
		return nil, errors.New("unauthorized")
	}

}
