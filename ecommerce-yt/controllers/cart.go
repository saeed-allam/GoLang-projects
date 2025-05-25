package controllers

import (
	"context"
	"ecommerce-yt/database"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	productCollection *mongo.Collection
	userCollection    *mongo.Collection
}

func NewApplication(productCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		productCollection: productCollection,
		userCollection:    userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Print("Product ID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product ID is required"))
			return
		}
		userQueryID := c.Query("userId")
		if userQueryID == "" {
			log.Print("user ID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("userId ID is required"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel() // releases resources if slowOperation completes before timeout elapses

		err = database.AddProductToCart(ctx, app.productCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			log.Println("Error adding product to cart:", err)
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "Product added to cart successfully")

	}
}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Print("Product ID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product ID is required"))
			return
		}
		userQueryID := c.Query("userId")
		if userQueryID == "" {
			log.Print("user ID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("userId ID is required"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel() // releases resources if slowOperation completes before timeout elapses
 
		err = database.RemoveCartItem(ctx,app.productCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			log.Println("Error removing item from cart:", err)
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "Item removed from cart successfully")
	}
}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("userId")
		if userQueryID == "" {
			log.Panicln("user ID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("userId ID is empty"))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel() // releases resources if slowOperation completes before timeout elapses

		err := database.BuyItemFromCart(ctx, app.userCollection, userQueryID)
		if err != nil {
			log.Println("Error buying items from cart:", err)
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "Items bought successfully")
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Print("Product ID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product ID is required"))
			return
		}
		userQueryID := c.Query("userId")
		if userQueryID == "" {
			log.Print("user ID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("userId ID is required"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel() // releases resources if slowOperation completes before timeout elapses

		err = database.InstantBuy(ctx, app.productCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			log.Println("Error in instant buy:", err)
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "Instant buy successful")
	}
}
