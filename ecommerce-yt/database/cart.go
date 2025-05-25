package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCantFindProduct = errors.New("product not found")
	ErrCantDecodeProduct = errors.New("cant find the product")
	ErrUserIdNotValid = errors.New("user id not valid")
	ErrCantUpdateUser = errors.New("cant update user")
	ErrContRemoveItemCart = errors.New("cant remove item from cart")
	ErrCantGetItem = errors.New("cant get item from cart")
	ErrCantBuyCartItem = errors.New("cant buy item from cart")
)

// ctx, app.productCollection, app.userCollection, productID, userQueryID
func AddProductToCart( ) error {
	// Implementation to add a product to the user's cart in the database
	return nil
}

func RemoveCartItem(userID, productID string) error {
	// Implementation to remove a product from the user's cart in the database
	return nil
}

func BuyItemFromCart(userID string) error {}

func InstantBuy(userID, productID string) error {
	// Implementation to buy a product instantly in the database
	return nil
}