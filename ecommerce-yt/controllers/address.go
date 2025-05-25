package controllers

import (
	"context"
	"ecommerce-yt/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		 user_id := c.Query("id")

		if user_id == "" {
			c.Header("Context-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid Search Index"})
			c.Abort() 
			return
		}


		addresses := make([]models.Address, 0)

		usert_id, err := primitive.ObjectIDFromHex(user_id)

		if err !=  nil {
			c.IndentedJSON(500, "InternaL Server Error")
		}


		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	
		defer cancel()
	
		filter := bson.D{primitive.E{Key:"_id", Value: usert_id}}
		update := bson.D{Key:"$set", Value: bson.D{primitive.E{Key: "address",Value: addresses }}}
		UserCollection.UpdateOne(ctx, filter,update)
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(404, "wrong commande")
			return 

		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "successfully deleted")
	}
}
