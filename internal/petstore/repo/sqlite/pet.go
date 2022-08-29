package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	client "petstore/internal/database/sqlite"
	"petstore/pkg/models"
)

type Pet models.Pet
type PetService models.PetService

func (p Pet) CreatePet() (int, error) {
	result, err := client.DbClient.Exec("INSERT INTO pet (name, age, image_url, description, category_id, breed_id, location_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		p.Name, p.Age, p.ImageURL, p.Description, p.PetCategory.ID, p.PetBreed.ID, p.PetLocation.ID)

	if err != nil {
		return p.ID, fmt.Errorf("Inserted failed. error: %v", err)
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return p.ID, fmt.Errorf("Inserted ID failed. error: %v", err)
	}
	p.ID = int(insertedId)
	return p.ID, nil
}

func (p Pet) GetPetsByCategory(catogoryID int) ([]*models.Pet, error) {
	var pets []*models.Pet

	// select pet.id,pet.pet_name,pet.age,pet.description,pet.category_id,category.category_name,pet.breed_id,breed.breed_name,pet.location_id,location.location_name from pet inner join category on pet.category_id = category.id inner join breed on pet.breed_id=breed.id inner join location on pet.location_id = location.id where category.id=$1;", categoryID)

	rows, err := client.DbClient.Query(`SELECT pet.id, pet.name, pet.age, pet.image_url, pet.description, pet.category_id, category.category_name, pet.breed_id, breed.breed_name, breed.category_id, pet.location_id, location.location_name 
									      FROM pet 
									INNER JOIN category ON pet.category_id=category.id 
									INNER JOIN breed ON pet.breed_id=breed.id 
									INNER JOIN location ON pet.location_id=location.id 
									WHERE pet.category_id=$1;`, catogoryID)

	if err != nil {
		return nil, fmt.Errorf("Query to database failed. error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&p.ID, &p.Name, &p.Age, &p.ImageURL, &p.Description, &p.PetCategory.ID, &p.PetCategory.CategoryName, &p.PetBreed.ID, &p.PetBreed.BreedName,
			&p.PetBreed.CategoryID, &p.PetLocation.ID, &p.PetLocation.LocationName)
		newPet := models.Pet(p)
		pets = append(pets, &newPet)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	return pets, nil
}

func (p Pet) GetPet(id int) (*models.Pet, error) {
	row := client.DbClient.QueryRow(`SELECT pet.id, pet.name, pet.age, pet.image_url, pet.description, pet.category_id, category.category_name, pet.breed_id, breed.breed_name, breed.category_id, pet.location_id, location.location_name 
							   	       FROM pet 
							     INNER JOIN category ON pet.category_id=category.id 
							     INNER JOIN breed ON pet.breed_id=breed.id 
							     INNER JOIN location ON pet.location_id=location.id 
							     WHERE pet.id=$1;`, id)

	if err := row.Scan(&p.ID, &p.Name, &p.Age, &p.ImageURL, &p.Description, &p.PetCategory.ID, &p.PetCategory.CategoryName, &p.PetBreed.ID, &p.PetBreed.BreedName,
		&p.PetBreed.CategoryID, &p.PetLocation.ID, &p.PetLocation.LocationName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("No records found")
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	pet := models.Pet(p)

	client.DbClient.Close()
	return &pet, nil
}

func (p Pet) DeletePet(id int) (int, error) {
	var count int64
	var err error
	result, err := client.DbClient.Exec("DELETE FROM pet WHERE id=$1", id)

	if err != nil {
		return int(count), fmt.Errorf("error query delete. error: %v", err)
	}

	count, err = result.RowsAffected()
	if err != nil {
		return int(count), fmt.Errorf("error parse delete result. error: %v", err)
	}
	
	return int(count), nil
}
