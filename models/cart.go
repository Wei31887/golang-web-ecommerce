package models

import (
	"SideProject/store/data"
	"time"
)

type Cart struct{
	Id int
	Uuid string
	UserId int
	CartItems []*CartItem
	TotalCount int
	TotalAmount float64
	CreatedAt time.Time
}

// GetTotalAmount 
func (cart *Cart) GetTotalCount() (totalCount int) {
	totalCount = 0
	for _, item := range cart.CartItems {
		totalCount += item.Count
	}
	cart.TotalCount = totalCount
	return 
}

func (cart *Cart) GetTotalAmount() (totalAmount float64) {
	totalAmount = 0.
	for _, item := range cart.CartItems {
		totalAmount += item.GetCartItemAmount()
	}
	cart.TotalAmount = totalAmount
	return 	
}

// Create : create cart to database
func (cart *Cart) Create() (err error) {
	statement := "INSERT INTO carts(uuid, user_id, total_amount, total_count, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id, uuid, created_at"
	stmt, err := data.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		createUuid(),
		cart.UserId,
		cart.GetTotalAmount(),
		cart.GetTotalCount(),
		time.Now(),
	).Scan(&cart.Id, &cart.Uuid, &cart.CreatedAt)
	if err != nil {
		return
	}

	// create cart item
	for _, item := range cart.CartItems {
		item.CartId = cart.Id
		err = item.Create()
		if err != nil {
			return
		}
	}
	return
}

// CartDeleteAll
func CartDeleteAll() (err error) {
	statement := "DELETE FROM carts"
	_, err = data.Db.Exec(statement)
	return
}