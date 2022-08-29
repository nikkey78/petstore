package models

type Breed struct{
	ID int `json:"id"`
	BreedName string `json:"breed_name"`
	CategoryID int `json:"category_id"`
}

type BreedService interface{
	CreateBreed() error
	GetBreedsByCategory(categoryId int)([]*Breed, error)
}