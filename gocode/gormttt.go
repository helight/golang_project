package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:Xu504600@/gormtt?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})
	db.Create(&Product{Code: "L12123", Price: 10003})

	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	fmt.Println("product: ", product)
	db.First(&product, "code = ?", "L1212") // find product with code l1212
	fmt.Println("product: ", product)

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	fmt.Println("product: ", product.ID)
	db.Delete(&product)
	db.Where("id = 9").Delete(&product)

	db.First(&product, 1)
	fmt.Println("product: ", product)
}
