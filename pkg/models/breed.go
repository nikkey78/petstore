package models

type Breed struct{
	ID int `json:"id"`
	BreedName string `json:"breedName"`
	CategoryID int `json:"categoryId"`
}

type BreedService interface{
	CreateBreed() error
	GetBreedsByCategory(categoryId int)([]*Breed, error)
}