package repository

import (
	"context"
	"database/sql"

	"git.01.alem.school/quazar/forum-authentication/models"
)

type sqliteCategoryRepository struct {
	Conn *sql.DB
}

// NewSqliteUserRepository will create an object that represents the category. Repository interface
func NewSqliteCategoryRepository(Conn *sql.DB) models.CategoryRepository {
	return &sqliteCategoryRepository{
		Conn: Conn,
	}
}

// Create category ...
func (c *sqliteCategoryRepository) Create(ctx context.Context, category *models.Category) (id int, err error) {
	return 0, nil
}

// GetAll categories ...
func (c *sqliteCategoryRepository) GetAll(ctx context.Context) (*[]models.CategoryRequestDTO, error) {
	rows, err := c.Conn.Query("SELECT category_id, category_name FROM category ORDER BY category_id DESC")
	if err != nil {
		return nil, err
	}

	var categories []models.CategoryRequestDTO

	for rows.Next() {
		category := &models.CategoryRequestDTO{}

		err = rows.Scan(&category.CategoryID, &category.CategoryName)

		if err != nil {
			return nil, err
		}

		categories = append(categories, *category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &categories, nil
}

// GetbyID category usecase ...
func (c *sqliteCategoryRepository) GetByID(ctx context.Context, id int) (category *models.CategoryRequestDTO, err error) {
	return nil, nil
}

// Update category ...
func (c *sqliteCategoryRepository) Update(ctx context.Context, category *models.Category) (err error) {
	return nil
}

// Delete category ...
func (c *sqliteCategoryRepository) Delete(ctx context.Context, id int) (err error) {
	return nil
}
