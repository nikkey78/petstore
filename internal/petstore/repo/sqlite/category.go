package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	client "petstore/internal/database/sqlite"
	"petstore/pkg/models"
)

type Category models.Category
type CategoryService models.CategoryService

// insert into category(category_name) values ("category name")
func (c Category) CreateCategory() (int, error) {
	result, err := client.DbClient.Exec("INSERT INTO category (category_name) VALUES ($1)", c.CategoryName)
	if err != nil {
		return c.ID, fmt.Errorf("Inserted failed. error: %v", err)
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return c.ID, fmt.Errorf("Inserted ID failed. error: %v", err)
	}
	c.ID = int(insertedId)
	return c.ID, nil
}

func (c Category) GetCategory(id int) (*Category, error) {
	row := client.DbClient.QueryRow("SELECT * FROM category WHERE id=$1", id)

	if err := row.Scan(&c.ID, &c.CategoryName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("No records found")
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	return &c, nil
}

func (c Category) GetAllCategory() ([]*Category, error) {
	var categories []*Category

	rows, err := client.DbClient.Query("SELECT * FROM category")
	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&c.ID, &c.CategoryName)
		newCat := c
		categories = append(categories, &newCat)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	return categories, nil
}
