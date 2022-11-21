package models

import "SideProject/store/data"


type CartItem struct {
	Id	int
	Product *Product
	Count int
	Amount float64
	CartId int
}

// GetCartItemAmount get the total Amount of the cart item
func (cartItem *CartItem) GetCartItemAmount() (Amount float64) {
	Amount = cartItem.Product.Prices * float64(cartItem.Count)
	cartItem.Amount = Amount
	return
}

// CreateCartItem
func (cartItem *CartItem) Create() (err error) {
	statement := "INSERT INTO cart_items(count, amount, cart_id, product_id) VALUES($1, $2, $3, $4) RETURNING id"
	stmt, err := data.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		cartItem.Count,
		cartItem.Amount, 
		cartItem.CartId,
		cartItem.Product.Id,
		).Scan(&cartItem.Id)
	return
}

// UpdateCartItem
func (cartItem *CartItem) UpdateCartItem() (err error) {
	statement := "UPDATE cart_items SET amount=$1, count=$2, cart_id=$3, product_id=$4 WHERE id=$5"
	stmt, err := data.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		cartItem.GetCartItemAmount(), 
		cartItem.Count, 
		cartItem.CartId,
		cartItem.Product.Id,
		cartItem.Id,
	)
	return
}

// CartItemDeleteAll
func CartItemDeleteAll() error {
	statement := "DELETE FROM cart_items"
	_, err := data.Db.Exec(statement)
	return err
}