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

	// Initialize Firebase with the credentials file path
	// This sets up Firebase services using the provided service account JSON file
	firebaseCredentialsPath := "path/to/your/firebase-service-account.json"
	firebase.InitFirebase(firebaseCredentialsPath)

	// Set the SendGrid API key as an environment variable
	// This is required for sending emails using the SendGrid service
	err := os.Setenv("SENDGRID_API_KEY", "your-sendgrid-api-key")
	if err != nil {
		log.Fatalf("Failed to set SENDGRID_API_KEY: %v", err)
	}

	// Initialize PostgreSQL connection
	// The DSN (Data Source Name) contains the database connection details
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=mydb port=5432 sslmode=disable"
	store, err := store.NewPostgresUserStore(dsn)
	// Uncomment the line below if using a different user store implementation
	// store, err := store.NewUserStore()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Create a new UserHandler with the initialized store
	// This handler will manage user-related operations
	userHandler := &handler.UserHandler{Store: store}

	// Initialize the router with the UserHandler
	// The router will handle incoming HTTP requests and route them to the appropriate handlers
	r := router.NewRouter(userHandler)

	// Start the HTTP server on port 8080
	// Log a message indicating the server has started and handle any fatal errors
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
