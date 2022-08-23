package models

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

type CategoryService interface {
	CreateCategory() error
	GetCategory(id int) (*Category, error)
	GetAllCategories() ([]*Category, error)
}
