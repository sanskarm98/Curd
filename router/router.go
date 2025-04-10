package router

import (
	"Curd/handler" // Importing the handler package for handling user-related requests
	"net/http"     // Importing the net/http package for HTTP server and routing
)

// NewRouter initializes and returns a new HTTP router.
// It takes a UserHandler as a parameter to handle user-related routes.
func NewRouter(userHandler *handler.UserHandler) http.Handler {
	mux := http.NewServeMux() // Create a new HTTP request multiplexer (router)

	// Register the userHandler to handle requests to "/users" and "/users/"
	mux.Handle("/users", userHandler)  // Handles requests to "/users"
	mux.Handle("/users/", userHandler) // Handles requests to "/users/" and subpaths

	return mux // Return the configured router
}
