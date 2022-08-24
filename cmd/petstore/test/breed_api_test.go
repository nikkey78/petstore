package main

import (
	"petstore/internal/petstore/repo/sqlite"
	"testing"
)

// func TestCreateBreed(t *testing.T) {
// 	var breed sqlite.Breed
// 	breed.BreedName = "doberman"
// 	breed.CategoryID = 1

// 	var err error
// 	id, err := breed.CreateBreed()
// 	if err != nil {
// 		t.Errorf("inner database error: %v", err)
// 	}
// 	if id <= 0 {
// 		t.Errorf("want id > 0, got id: %d", id)
// 	}
// }

func TestGetBreedsByCategory(t *testing.T) {
	var breed sqlite.Breed

	categoryId := 1

	breeds, err := breed.GetBreedsByCategory(categoryId)
	if err != nil {
		t.Errorf("inner database error: %v", err)
	}
	if len(breeds) == 0 {
		t.Errorf("want categories length more than 0, got: %d", len(breeds))
	}

	if len(breeds) > 0 {

		wantBreedName := "doberman"

		if wantBreedName != breeds[0].BreedName {
			t.Errorf("want breeds[0] name: %s, got: %s", wantBreedName, breeds[0].BreedName)
		}
	}
}
