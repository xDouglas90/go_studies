package models

import (
	"studies/alura/store/db"
)

type Product struct {
	Id, Quantity      int
	Name, Description string
	Price             float64
}

func GetAllProducts() []Product {
	db := db.DbConnect()

	getAllProducts, err := db.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for getAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = getAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()

	return products
}

func GetProduct(id string) Product {
	db := db.DbConnect()

	getProduct, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	p := Product{}

	for getProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = getProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
	}

	defer db.Close()

	return p
}

func CreateProduct(name, description string, price float64, quantity int) {
	db := db.DbConnect()

	insertProduct, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DbConnect()

	deleteProduct, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)

	defer db.Close()
}

func UpdateProduct(id, name, description string, price float64, quantity int) {
	db := db.DbConnect()

	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)

	defer db.Close()
}
