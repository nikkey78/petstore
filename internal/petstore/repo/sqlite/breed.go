package sqlite

import (
	"fmt"
	client "petstore/internal/database/sqlite"
	"petstore/pkg/models"
)

type Breed models.Breed
type BreedService models.BreedService

// CreateBreed() error
// 	GetBreedsByCategory(categoryId int)([]*Breed, error)

// insert into category(category_name) values ("category name")
func (b Breed) CreateBreed() (int, error) {
	result, err := client.DbClient.Exec("INSERT INTO breed (breed_name, category_id) VALUES ($1, $2)", b.BreedName, b.CategoryID)
	if err != nil {
		return b.ID, fmt.Errorf("Inserted failed. error: %v", err)
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return b.ID, fmt.Errorf("Inserted ID failed. error: %v", err)
	}
	b.ID = int(insertedId)
	return b.ID, nil
}

func (b Breed) GetBreedsByCategory(categoryId int) ([]*Breed, error) {
	var breeds []*Breed
	rows, err := client.DbClient.Query("SELECT id, breed_name, category_id FROM breed WHERE category_id=$1", categoryId)
	if err != nil {
		return nil, fmt.Errorf("Inserted failed. error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&b.ID, &b.BreedName, &b.CategoryID)
		newBreed := b
		breeds = append(breeds, &newBreed)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}

	return breeds, nil
}
