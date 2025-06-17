package product

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
func (r *repositorySQL) Insert(p models.Product) (int64, error) {
	columns := []string{}
	values := []interface{}{}

	if p.Title != "" {
		columns = append(columns, "Prod_Title")
		values = append(values, p.Title)
	}
	if p.Description != "" {
		columns = append(columns, "Prod_Description")
		values = append(values, p.Description)
	}
	if p.Price != 0 {
		columns = append(columns, "Prod_Price")
		values = append(values, p.Price)
	}
	if p.Stock != 0 {
		columns = append(columns, "Prod_Stock")
		values = append(values, p.Stock)
	}
	if p.CategId != 0 {
		columns = append(columns, "Prod_CategId")
		values = append(values, p.CategId)
	}
	if p.Path != "" {
		columns = append(columns, "Prod_Path")
		values = append(values, p.Path)
	}

	query, args, err := squirrel.
		Insert("product").
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

func (r *repositorySQL) Update(c models.Product) error {
	// Build a safe SQL UPDATE query using the squirrel package
	query, args, err := squirrel.
		Update("product").
		Set("Prod_Title", c.Title).
		Where(squirrel.Eq{"Categ_Id": c.Id}).
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

func (r *repositorySQL) Delete(id int) error {
	// Build a safe SQL UPDATE query using the squirrel package
	query, args, err := squirrel.
		Delete("product").
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

func (r *repositorySQL) GetById(id int) (models.Product, error) {
	// Build a safe SQL UPDATE query using the squirrel package
	query, args, err := squirrel.
		Select("Categ_Name", "Categ_Path").
		From("product").
		Where(squirrel.Eq{"Categ_Id": id}).
		ToSql()

	if err != nil {
		return models.Product{}, err
	}

	var name, path string
	// Execute the query with the generated SQL and arguments
	row := r.db.QueryRow(query, args...)
	if err := row.Scan(&name, &path); err != nil {
		return models.Product{}, err
	}

	return models.Product{
		Id:    id,
		Title: name,
	}, nil

}
func (r *repositorySQL) GetAll() ([]models.Product, error) {
	return []models.Product{}, nil
	// Build a safe SQL UPDATE query using the squirrel package
	// query, args, err := squirrel.
	// 	Select("Categ_Id", "Categ_Name", "Categ_Path").
	// 	From("category").
	// 	ToSql()

	// if err != nil {
	// 	return []models.Product{}, err
	// }

	// var categories []models.Product
	// var id int
	// var name, path string
	// // Execute the query with the generated SQL and arguments
	// rows, err := r.db.Query(query, args...)

	// if err != nil {
	// 	return []models.Product{}, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	if err := rows.Scan(&id, &name, &path); err != nil {
	// 		return []models.Product{}, err
	// 	}
	// 	categories = append(categories, models.Category{
	// 		CategID:   id,
	// 		CategName: name,
	// 		CategPath: path,
	// 	})

	// }

	// return categories, nil

}

func (r *repositorySQL) GetBySlug(slug string) ([]models.Product, error) {
	return []models.Product{}, nil

	// Build a safe SQL UPDATE query using the squirrel package
	// query, args, err := squirrel.
	// 	Select("Categ_Id", "Categ_Name", "Categ_Path").
	// 	From("category").
	// 	Where(squirrel.Like{"Categ_Path": "%" + slug + "%"}).
	// 	ToSql()

	// if err != nil {
	// 	return []models.Category{}, err
	// }

	// var categories []models.Category
	// var id int
	// var name, path string
	// // Execute the query with the generated SQL and arguments
	// rows, err := r.db.Query(query, args...)

	// if err != nil {
	// 	return []models.Category{}, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	if err := rows.Scan(&id, &name, &path); err != nil {
	// 		return []models.Category{}, err
	// 	}
	// 	categories = append(categories, models.Category{
	// 		CategID:   id,
	// 		CategName: name,
	// 		CategPath: path,
	// 	})

	// }

	// return categories, nil

}
