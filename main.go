package main

import (
	"hello_world1/handler"
	"hello_world1/router"
	"hello_world1/store"
	"log"
	"net/http"
)

func main() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=mydb port=5432 sslmode=disable"
	store, err := store.NewPostgresUserStore(dsn)
	//store, err := store.NewUserStore()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	userHandler := &handler.UserHandler{Store: store}
	r := router.NewRouter(userHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
