package main

import (
	"console"
	"encoding/json"
	"fmt"
	"petstore/internal/petstore/repo/sqlite"
)

func main() {
	defer console.GreenReset()()

	// var id int
	var err error

	//	// ----------------------- TEST Create Category ----=--------------------
	// var category sqlite.Category = sqlite.Category{CategoryName: "cat"}

	// id, err := category.CreateCategory()
	// if err != nil {
	// 	console.ColorPrintln(console.Color_Red, err)
	// 	return
	// } else {
	// 	fmt.Printf("Category created successed. id: %d\n", id)
	// }

	//	// ----------------------- TEST GetCategory by Id --------------------
	// id := 1

	// cat, err := category.GetCategory(id)
	// if err != nil {
	// 	console.ColorPrintln(console.Color_Red, err)
	// 	return
	// }
	// fmt.Printf("category: id=%d, category_name=%s\n", cat.ID, cat.CategoryName)

	//	// ----------------------- TEST Get All Categories --------------------------
	// categories, err := category.GetAllCategory()
	// if err != nil {
	// 	console.ColorPrintln(console.Color_Red, err)
	// }

	// for _, cat := range categories {
	// 	fmt.Printf("category: id=%d, category_name=%s\n", cat.ID, cat.CategoryName)
	// }

	//	// ----------------------- TEST Create Breed --------------------------
	// var breed sqlite.Breed = sqlite.Breed{BreedName: "spitz", CategoryID: 1}

	// id, err = breed.CreateBreed()
	// if err != nil {
	// 	console.ColorPrintln(console.Color_Red, err)
	// }
	// fmt.Printf("Category created successed. id: %d\n", id)

	//	// ----------------------- TEST Create Location --------------------------
	// var location sqlite.Location = sqlite.Location{LocationName: "St Petersburg"}
	// id, err = location.CreateLocation()
	// if err != nil {
	// 	console.ColorPrintln(console.Color_Red, err)
	// }
	// fmt.Printf("Category created successed. id: %d\n", id)

	//	// ----------------------- TEST Create Location --------------------------
	// var pet1 sqlite.Pet = sqlite.Pet{Name: "heavy", Age: 0.5, ImageURL: "picture",
	// 	Description: "funny shiba inu", PetCategory: models.Category{ID: 1}, PetBreed: models.Breed{ID: 2}, PetLocation: models.Location{ID: 2}}

	// var pet2 sqlite.Pet = sqlite.Pet{Name: "rock", Age: 2.0, ImageURL: "picture",
	// 	Description: "funny spitz", PetCategory: models.Category{ID: 1}, PetBreed: models.Breed{ID: 3}, PetLocation: models.Location{ID: 2}}

	// id, err = pet.CreatePet()
	// if err != nil {
	// 	console.ColorPrintln(console.Color_Red, err)
	// 	return
	// }
	// fmt.Printf("Category created successed. id: %d\n", id)

	//	// ----------------------- TEST GetPetsByCategory --------------------------

	var pet sqlite.Pet
	categoryId := 1

	pets, err := pet.GetPetsByCategory(categoryId)
	if err != nil {
		console.ColorPrintln(console.Color_Red, err)
		return
	}

	// for _, p := range pets {
	// 	fmt.Printf("%+v\n", p)
	// }

	data, err := json.MarshalIndent(&pets, "", "   ")
	if err != nil {
		console.ColorPrintln(console.Color_Red, err)
		return
	}
	fmt.Println(string(data))

}
