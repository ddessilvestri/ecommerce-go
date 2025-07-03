package address

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/ddessilvestri/ecommerce-go/models"
)

// This struct acts like a "class" in Go.
// It implements the Storage interface for SQL-based storage.
type repositorySQL struct {
	db *sql.DB // Dependency to the database connection
}

// Constructor-like function (Go does not support constructors like C# or Java).
// By convention, we use New<Name>() to instantiate and return the interface type.
func NewSQLRepository(db *sql.DB) Storage {
	// We return a pointer to the struct instance
	return &repositorySQL{db: db}
}

// Method bound to the repositorySQL struct.
// The receiver is a pointer (*repositorySQL), which allows modifying internal state
// and avoids copying the struct on each method call.
func (r *repositorySQL) Insert(a models.Address, userUUID string) (int64, error) {

	columns := []string{
		"Add_UserID",
		"Add_Name",
		"Add_Title",
		"Add_Address",
		"Add_City",
		"Add_State",
		"Add_PostalCode",
		"Add_Phone",
	}

	values := []interface{}{
		userUUID,
		a.Name,
		a.Title,
		a.Address,
		a.City,
		a.State,
		a.PostalCode,
		a.Phone,
	}

	query, args, err := squirrel.
		Insert("addresses").
		Columns(columns...).
		Values(values...).
		PlaceholderFormat(squirrel.Question).
		ToSql()

	if err != nil {
		return 0, err
	}

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()

}

func (r *repositorySQL) Exists(id int) bool {
	query, args, err := squirrel.
		Select("1").
		From("addresses").
		Where(squirrel.Eq{"Add_Id": id}).
		Limit(1).
		ToSql()

	if err != nil {
		return false
	}

	var exists int
	err = r.db.QueryRow(query, args...).Scan(&exists)
	if err != nil {
		return false
	}

	return true
}

func (r *repositorySQL) Update(a models.Address) error {
	columns := []string{
		"Add_Name",
		"Add_Title",
		"Add_Address",
		"Add_City",
		"Add_State",
		"Add_PostalCode",
		"Add_Phone",
		"Add_Updated", // campo para fecha/hora de actualización
	}

	values := []interface{}{
		a.Name,
		a.Title,
		a.Address,
		a.City,
		a.State,
		a.PostalCode,
		a.Phone,
		squirrel.Expr("NOW()"),
	}

	// Armamos el builder
	builder := squirrel.
		Update("addresses").
		PlaceholderFormat(squirrel.Question)

	// agregamos cada columna dinámicamente
	for i, col := range columns {
		builder = builder.Set(col, values[i])
	}

	// Agregamos WHERE
	builder = builder.Where(squirrel.Eq{"Add_Id": a.Id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositorySQL) Delete(id int) error {
	query, args, err := squirrel.
		Delete("addresses").
		Where(squirrel.Eq{"Add_Id": id}).
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

func (r *repositorySQL) GetAllByUserUUID(userUUID string) ([]models.Address, error) {
	queryBuilder := squirrel.
		Select("*").
		From("addresses").
		Where(squirrel.Eq{"Add_UserID": userUUID}).
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

	var addresses []models.Address
	for rows.Next() {
		var a models.Address
		err = rows.Scan(
			&a.Id,
			&a.Name,
			&a.Title,
			&a.Address,
			&a.City,
			&a.State,
			&a.PostalCode,
			&a.Phone,
		)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, a)
	}

	return addresses, nil
}
