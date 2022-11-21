package models

import (
	"log"

	"github.com/google/uuid"
)

//createUuid
func createUuid() (string) {
	uuid := uuid.New().String()
	return uuid
}

func setup() (err error) {
	err = CartItemDeleteAll()
	if err != nil {
		return
	}

	err = CartDeleteAll()
	if err != nil {
		return
	}
	
	err =  DeleteImageAll()
	if err != nil {
		log.Println(err)
		return
	}

	err = ProductDeleteAll()
	if err != nil {
		return
	}
	
	err = SessionDeleteAll()
	if err != nil {
		return
	}

	err = UserDeleteAll()
	if err != nil {
		return
	}
	return
}