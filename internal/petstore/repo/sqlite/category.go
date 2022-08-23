package sqlite

import (
	"fmt"
	client "petstore/internal/database/sqlite"
	"petstore/pkg/models"
)

type Category models.Category
type CategoryService models.CategoryService

// insert into category(category_name) values ("category name")
func (c *Category) CreateCategory() (int, error) {
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

func (c *Category) GetCategory(id int) (*Category, error) { return nil, nil }
func (c *Category) GetAllCategory() ([]*Category, error)  { return nil, nil }
