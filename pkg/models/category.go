package models

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

type CategoryService interface {
	CreateCategory() (int, error)
	GetCategory(id int) (*Category, error)
	GetAllCategory() ([]*Category, error)
}
