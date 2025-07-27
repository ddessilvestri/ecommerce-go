package user

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

func (r *repositorySQL) Update(u models.User) error {
	update := squirrel.
		Update("users").
		PlaceholderFormat(squirrel.Question)

	if u.FirstName != "" {
		update = update.Set("User_FirstName", u.FirstName)
	}
	if u.LastName != "" {
		update = update.Set("User_LastName", u.LastName)
	}

	query, args, err := update.
		Where(squirrel.Eq{"User_UUID": u.UUID}).
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

func (r *repositorySQL) GetByUUID(uuid string) (models.User, error) {
	query, args, err := squirrel.
		Select("*").
		From("users").
		Where(squirrel.Eq{"User_UUID": uuid}).
		PlaceholderFormat(squirrel.Question).
		ToSql()

	fmt.Println(query)
	if err != nil {
		return models.User{}, err
	}

	row := r.db.QueryRow(query, args...)
	var u models.User
	err = row.Scan(&u.UUID, &u.Email, &u.FirstName, &u.LastName, &u.Status, &u.DateAdd, &u.DateUpg)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}
