package stock

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
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

func (r *repositorySQL) UpdateStock(productId, delta int) error {
	query, args, err := squirrel.
		Update("products").
		PlaceholderFormat(squirrel.Question).
		Set("Prod_Updated", squirrel.Expr("NOW()")).
		Set("Prod_Stock", squirrel.Expr("Prod_Stock + ?", delta)).
		Where(squirrel.Eq{"Prod_Id": productId}).
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
