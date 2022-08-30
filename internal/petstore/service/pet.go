package service

import (
	db "petstore/internal/petstore/repo/sqlite"
	"petstore/pkg/models"
)

var ps db.PetService
var pet db.Pet

func CreatePet(p *db.Pet) (int, error) {
	ps = p
	id, err := ps.CreatePet()
	if err != nil {
		return id, err
	}
	return id, nil
}

func GetPetsByCategory(catogoryID int) ([]*models.Pet, error) {
	ps = &pet
	pets, err := ps.GetPetsByCategory(catogoryID)
	if err != nil {
		return nil, err
	}
	return pets, nil
}

func GetPet(id int) (*models.Pet, error) {
	ps = &pet
	pet, err := ps.GetPet(id)
	if err != nil {
		return nil, err
	}
	return pet, nil
}

func DeletePet(id int) (int, error) {
	ps = &pet
	id, err := ps.DeletePet(id)
	if err != nil {
		return id, err
	}
	return id, nil
}
