package service

import (
	db "petstore/internal/petstore/repo/sqlite"
	"petstore/pkg/models"
)

var bs db.BreedService
var breed db.Breed

func CreateBreed(b *db.Breed) (int, error) {
	bs = b
	id, err := bs.CreateBreed()
	if err != nil {
		return id, err
	}
	return id, err
}

func GetBreedsByCategory(categoryId int) ([]*models.Breed, error) {
	bs = &breed
	breeds, err := bs.GetBreedsByCategory(categoryId)
	if err != nil {
		return nil, err
	}
	return breeds, nil
}
