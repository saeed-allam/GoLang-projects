package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	First_Name      *string            `bson:"first_name" json:"first_name" validate:"required,min=2,max=30"`
	Last_Name       *string            `bson:"last_name" json:"last_name" validate:"required,min=2,max=30"`
	Password        *string            `bson:"password" json:"password" validate:"required,min=6"`
	Email           *string            `bson:"email" json:"email" validate:"required,email"`
	Phone           *string            `bson:"phone" json:"phone" validate:"required,min=10,max=15"`
	Token           *string            `bson:"token" json:"token"`
	Refresh_Token   *string            `bson:"refresh_token" json:"refresh_token"`
	Created_At      time.Time          `bson:"created_at" json:"created_at"`
	Updated_At      time.Time          `bson:"updated_at" json:"updated_at"`
	User_ID         string             `bson:"user_id" json:"user_id"`
	UserCart        []ProdctUser       `bson:"usercart" json:"usercart"`
	Address_Details []Address          `bson:"address" json:"address"`
	Orders_Status   []Order            `bson:"orders" json:"orders"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Product_Name *string            `bson:"product_name" json:"product_name"`
	Price        *uint64            `bson:"price" json:"price"`
	Rating       *uint8             `bson:"rating" json:"rating"`
	Image        *string            `bson:"image" json:"image"`
}

type ProdctUser struct {
	Product_ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Product_Name *string            `bson:"product_name" json:"product_name"`
	Price        int                `bson:"price" json:"price"`
	Rating       *uint8             `bson:"rating" json:"rating"`
	Image        *string            `bson:"image" json:"image"`
}

type Address struct {
	Address_ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	House      *string            `bson:"house_name" json:"house_name"`
	Street     *string            `bson:"street_name" json:"street_name"`
	City       *string            `bson:"city_name" json:"city_name"`
	Pincode    *string            `bson:"pin_code" json:"pin_code"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Order_Cart     []ProdctUser       `bson:"order_list" json:"order_list"`
	Oreder_At      time.Time          `bson:"order_at" json:"order_at"`
	Price          int                `bson:"total_price" json:"total_price"`
	Discount       *int               `bson:"discount" json:"discount"`
	Payment_Method Payment            `bson:"payment_method" json:"payment_method"`
}

type Payment struct {
	Digital bool `bson:"digital" json:"digital"`
	COD     bool `bson:"cod" json:"cod"`
}
