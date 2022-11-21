package models

import (
	"SideProject/store/data"
	"log"
	"time"
)


type Product struct {
	Id 			int
	Title 		string
	Prices 		float64
	Stocks 		int
	Sales 		int
	Images 		[]*Image
	CreatedAt 	time.Time
}

// Create
func (product *Product) Create() (err error) {
	statement := "INSERT INTO products(title, prices, stocks, sales, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id, created_at"
	stmt, err := data.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		product.Title,
		product.Prices,
		product.Stocks,
		product.Sales,
		time.Now(),
	).Scan(&product.Id, &product.CreatedAt)

	for _, image := range product.Images {
		image.ProductId = product.Id
		err = image.Create() 
		if err != nil {
			log.Println(err)
			return
		}
	}
	return
}

// ProductDeleteAll
func ProductDeleteAll() (err error) {
	statement := "DELETE FROM products"
	_, err = data.Db.Exec(statement)
	return err
}

