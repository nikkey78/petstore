package main

import (
	"petstore/internal/petstore/repo/sqlite"
	"testing"
)


// func TestCreateCategory(t *testing.T) {
// 	var category sqlite.Category
// 	category.CategoryName = "dog"

// 	var err error
// 	id, err := category.CreateCategory()
// 	if err != nil {
// 		t.Errorf("inner database error: %v", err)
// 	}
// 	if id <= 0 {
// 		t.Errorf("want id > 0, got id: %d", id)
// 	}
// }

func TestGetCategoryById(t *testing.T) {
	var category sqlite.Category

	wantCategoryName := "dog"
	id := 1

	cat, err := category.GetCategory(id)
	if err != nil {
		t.Errorf("inner database error: %v", err)
	}
	if wantCategoryName != cat.CategoryName {
		t.Errorf("want: %s, got: %s", wantCategoryName, cat.CategoryName)
	}
}

func TestGetAllCategory(t *testing.T) {
	var category sqlite.Category

	cats, err := category.GetAllCategory()
	if err != nil {
		t.Errorf("inner database error: %v", err)
	}
	if len(cats) == 0 {
		t.Errorf("want categories length more than 0, got: %d", len(cats))
	}
}
