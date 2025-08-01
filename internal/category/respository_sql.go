package category

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
func (r *repositorySQL) InsertCategory(c models.Category) (int64, error) {
	// Build a safe SQL INSERT query using the squirrel package
	query, args, err := squirrel.
		Insert("category").
		Columns("Categ_Name", "Categ_Path").
		Values(c.CategName, c.CategPath).
		ToSql()

	if err != nil {
		return 0, err
	}

	// Execute the query with the generated SQL and arguments
	result, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	// Return the last inserted ID
	return result.LastInsertId()
}

func (r *repositorySQL) UpdateCategory(c models.Category) error {
	// Build a safe SQL UPDATE query using the squirrel package
	query, args, err := squirrel.
		Update("category").
		Set("Categ_Name", c.CategName).
		Set("Categ_Path", c.CategPath).
		Where(squirrel.Eq{"Categ_Id": c.CategID}).
		ToSql()

	if err != nil {
		return err
	}

	// Execute the query with the generated SQL and arguments
	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	// Return the updated ID
	return nil
}

func (r *repositorySQL) DeleteCategory(id int) error {
	// Build a safe SQL UPDATE query using the squirrel package
	query, args, err := squirrel.
		Delete("category").
		Where(squirrel.Eq{"Categ_Id": id}).
		ToSql()

	if err != nil {
		return err
	}

	// Execute the query with the generated SQL and arguments
	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	// Return the updated ID
	return nil
}

func (r *repositorySQL) GetCategory(id int) (models.Category, error) {
	// Build a safe SQL UPDATE query using the squirrel package
	query, args, err := squirrel.
		Select("Categ_Name", "Categ_Path").
		From("category").
		Where(squirrel.Eq{"Categ_Id": id}).
		ToSql()

	if err != nil {
		return models.Category{}, err
	}

	var name, path string
	// Execute the query with the generated SQL and arguments
	row := r.db.QueryRow(query, args...)
	if err := row.Scan(&name, &path); err != nil {
		return models.Category{}, err
	}

	return models.Category{
		CategID:   id,
		CategName: name,
		CategPath: path,
	}, nil

}
func (r *repositorySQL) GetCategories() ([]models.Category, error) {
	// Build a safe SQL UPDATE query using the squirrel package
	query, args, err := squirrel.
		Select("Categ_Id", "Categ_Name", "Categ_Path").
		From("category").
		ToSql()

	if err != nil {
		return []models.Category{}, err
	}

	var categories []models.Category
	var id int
	var name, path string
	// Execute the query with the generated SQL and arguments
	rows, err := r.db.Query(query, args...)

	if err != nil {
		return []models.Category{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id, &name, &path); err != nil {
			return []models.Category{}, err
		}
		categories = append(categories, models.Category{
			CategID:   id,
			CategName: name,
			CategPath: path,
		})

	}

	return categories, nil

}

func (r *repositorySQL) GetCategoryBySlug(slug string) ([]models.Category, error) {
	// Build a safe SQL UPDATE query using the squirrel package
	query, args, err := squirrel.
		Select("Categ_Id", "Categ_Name", "Categ_Path").
		From("category").
		Where(squirrel.Like{"Categ_Path": "%" + slug + "%"}).
		ToSql()

	if err != nil {
		return []models.Category{}, err
	}

	var categories []models.Category
	var id int
	var name, path string
	// Execute the query with the generated SQL and arguments
	rows, err := r.db.Query(query, args...)

	if err != nil {
		return []models.Category{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id, &name, &path); err != nil {
			return []models.Category{}, err
		}
		categories = append(categories, models.Category{
			CategID:   id,
			CategName: name,
			CategPath: path,
		})

	}

	return categories, nil

}
