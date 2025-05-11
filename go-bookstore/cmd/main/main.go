package main

import (
	"github.com/gorilla/mux" //Gorilla Mux router
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/saeed-allam/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))

}
