package main

import (
	"Curd/firebase"
	"Curd/handler"
	"Curd/router"
	"Curd/store"
	"log"
	"net/http"
	"os"
)

func main() {

	// Initialize Firebase with credentials file path
	firebaseCredentialsPath := "path/to/your/firebase-service-account.json"
	firebase.InitFirebase(firebaseCredentialsPath)

	// Initialize SendGrid with API key
	err := os.Setenv("SENDGRID_API_KEY", "your-sendgrid-api-key")
	if err != nil {
		log.Fatalf("Failed to set SENDGRID_API_KEY: %v", err)
	}

	// Initialize PostgreSQL connection
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
