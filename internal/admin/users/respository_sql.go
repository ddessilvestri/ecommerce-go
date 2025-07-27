package adminusers

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/ddessilvestri/ecommerce-go/models"
)

type repositorySQL struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) Storage {
	return &repositorySQL{db: db}
}

func (r *repositorySQL) Delete(uuid string) error {
	query, args, err := squirrel.
		Delete("users").
		Where(squirrel.Eq{"User_UUID": uuid}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositorySQL) GetAll(offset, limit int, sortBy, order string) ([]models.User, error) {
	allowedSorts := map[string]string{
		"uuid":       "User_UUID",
		"email":      "User_Email",
		"first_name": "User_FirstName",
		"last_name":  "User_LastName",
		"status":     "User_Status",
		"date_add":   "User_DateAdd",
		"date_upg":   "User_DateUpg",
	}

	dbSortBy, ok := allowedSorts[sortBy]
	if !ok {
		dbSortBy = "User_UUID"
	}

	queryBuilder := squirrel.
		Select("User_UUID", "User_Email", "User_FirstName", "User_LastName", "User_Status", "User_DateAdd", "User_DateUpg").
		From("users").
		OrderBy(fmt.Sprintf("%s %s", dbSortBy, order)).
		Offset(uint64(offset)).
		Limit(uint64(limit)).
		PlaceholderFormat(squirrel.Question)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.UUID, &u.Email, &u.FirstName, &u.LastName, &u.Status, &u.DateAdd, &u.DateUpg); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
