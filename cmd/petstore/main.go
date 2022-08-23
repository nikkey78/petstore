package main

import (
	"console"
	"fmt"
	_ "petstore/internal/database/sqlite"
	"petstore/internal/petstore/repo/sqlite"
)

func main() {
	defer console.GreenReset()()

	// test create category
	var category sqlite.Category
	category.CategoryName = "category_3"
	

	id, err := category.CreateCategory()
	if err != nil {
		console.ColorPrintln(console.Color_Red, err)
		return
	} else {
		fmt.Printf("Category created successed. id: %d", id)
	}

}
