package models

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetProducts() []Product {
	db, err := sql.Open("mysql", "root:0519@tcp(127.0.0.1:3306)/test")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM product")
	
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	products := []Product{}
	for results.Next() {
		var prod Product
		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, prod)
	}

	return products
}

func GetProduct(code string) *Product {
	db, err := sql.Open("mysql", "root:0519@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	prod := &Product{}
	results, err := db.Query("SELECT * FROM product where code=?", code)
	
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	return prod
}

func AddProduct(product Product) {
	db, err := sql.Open("mysql", "root:0519@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO product (code, name, qty, last_updated) VALUES (?, ?, ?, now())",
		product.Code, product.Name, product.Qty)
	
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}