package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // List of allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},      // List of allowed HTTP methods
		AllowedHeaders:   []string{"*"}, // List of allowed headers
		ExposedHeaders:   []string{"Link"},                                         // List of headers exposed to the client
		AllowCredentials: true,                                                     // Allow sending credentials (cookies, authorization headers)
		MaxAge:           300,                                                      
	}))

	v1Router := chi.NewRouter()
	// v1Router.HandleFunc("/healthz", handlerReadiness)
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err",handlerErr)

	router.Mount("/v1",v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server Start On Port %v", portString)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
