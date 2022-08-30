package service

import (
	db "petstore/internal/petstore/repo/sqlite"
	"petstore/pkg/models"
)

var cs db.CategoryService
var category db.Category


func CreateCategory(c *db.Category) (int, error) {
	cs = c
	id, err := cs.CreateCategory()
	if err != nil {
		return id, err
	}
	return id, err
}
func GetCategory(id int) (*models.Category, error) {
	cs = &category
	category, err := cs.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func GetAllCategory() ([]*models.Category, error) {
	cs = &category
	categories, err := cs.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
