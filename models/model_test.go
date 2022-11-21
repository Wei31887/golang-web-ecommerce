package models

import (
	"fmt"
	"testing"
)

// Set up for testing
var testProducts = []Product {
	{
		Title: "ABC",
		Prices: 100.0,
		Stocks: 100,
		Sales: 1,
	},
	{
		Title: "ㄅㄆㄇ",
		Prices: 150.0,
		Stocks: 50,
		Sales: 2,
	},
}

func TestSetup(t *testing.T) {
	fmt.Println("Testing set up ... ")
	err = setup()
	if err != nil {
		t.Error(err, "Set up fail")
	}
}

func TestProduct(t * testing.T) {
	fmt.Println("Testing product ... ")
	t.Run("Testing for creating product", TestProductCreate)
}

func TestProductCreate(t *testing.T) {
	setup()
	for _, product := range testProducts {
		err := product.Create()
		if err != nil {
			t.Error(err, "Can not create product")
		}
	}
}

// Cart test
func TestCart(t *testing.T) {
	fmt.Println("Testing cart relative ... ")
	// t.Run("Testing for creating cart item", TestCartItemCreate)
	t.Run("Testing cart create", TestCartCreate)
}

func TestCartCreate(t *testing.T) {
	setup()
	fmt.Println("Testing cart create ... ")
	// create products and cart items
	cartItems := []*CartItem{}
	for i := range testProducts {
		testProducts[i].Create()
		item := CartItem{
			Product: &testProducts[i],
			Count: i+1,
		}
		// item.GetCartItemCount()
		fmt.Println("cart item product id", item.Product.Id)
		cartItems = append(cartItems, &item)
	}

	// create user
	if err := testUsers[0].Create(); err != nil {
		t.Error(err, "can not create user")
	}

	// create cart
	cart := Cart{
		UserId: testUsers[0].Id,
		Uuid: createUuid(),
		CartItems: cartItems,
	}
	err := cart.Create(); 
	if err != nil {
		t.Error(err, "Can not create cart")
	}
	fmt.Println("print cart info: ", cart)
}



func TestImage(t *testing.T) {
	fmt.Println("Testing image relative ... ")
	t.Run("Testing image create", TestImageCreate)
}

func TestImageCreate(t *testing.T) {
	setup()
	
}