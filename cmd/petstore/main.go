package main

import (
	"console"
	"fmt"
	"petstore/internal/petstore/repo/sqlite"
)

func main() {
	defer console.GreenReset()()

	// ----------------------- TEST Create Category ----=--------------------

	var category sqlite.Category

	// category.CategoryName = "category_3"

	// id, err := category.CreateCategory()
	// if err != nil {
	// 	console.ColorPrintln(console.Color_Red, err)
	// 	return
	// } else {
	// 	fmt.Printf("Category created successed. id: %d", id)
	// }

	// ----------------------- TEST GetCategory by Id --------------------
	// id := 3
	// cat, err := category.GetCategory(id)
	// if err != nil {
	// 	console.ColorPrintln(console.Color_Red, err)
	// 	return
	// }
	// fmt.Printf("category: id=%d, category_name=%s\n", cat.ID, cat.CategoryName)

	// ----------------------- TEST Get All Categories --------------------------
	categories, err := category.GetAllCategory()
	if err != nil {
		console.ColorPrintln(console.Color_Red, err)
	}

	for _, cat := range categories {
		fmt.Printf("category: id=%d, category_name=%s\n", cat.ID, cat.CategoryName)
	}

}
