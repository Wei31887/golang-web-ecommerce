package models

import (
	"SideProject/store/data"
	"log"
)

type Image struct {
	Id 			int
	ImagePath 	string
	ImageName 	string
	ProductId 	int
}

// Create
func (image *Image) Create() (err error) {
	statemment := "INSERT INTO images (image_path, image_name, product_id) VALUES ($1, $2, $3) RETURNING id"
	stml, err := data.Db.Prepare(statemment)
	if err != nil {
		log.Println(err)
		return
	}
	defer stml.Close()

	err = stml.QueryRow(image.ImagePath, image.ImageName, image.ProductId).Scan(&image.Id)
	return
}

// DeleteImageAll
func DeleteImageAll() (err error) {
	statement := "DELETE FROM images"
	_, err = data.Db.Exec(statement)
	return
}