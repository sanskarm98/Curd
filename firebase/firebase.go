package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

// FirebaseApp is a global variable to hold the initialized Firebase app instance.
var FirebaseApp *firebase.App

// FCMClient is a global variable to hold the initialized Firebase Cloud Messaging client.
var FCMClient *messaging.Client

// InitFirebase initializes the Firebase app and FCM client using the provided credentials file path.
func InitFirebase(credentialsFilePath string) {
	// Create an option to specify the credentials file for Firebase initialization.
	opt := option.WithCredentialsFile(credentialsFilePath)

	// Initialize the Firebase app with the provided credentials.
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		// Log a fatal error and terminate the program if Firebase app initialization fails.
		log.Fatalf("error initializing Firebase app: %v", err)
	}
	// Assign the initialized app to the global FirebaseApp variable.
	FirebaseApp = app

	// Initialize the Firebase Cloud Messaging (FCM) client from the app instance.
	client, err := app.Messaging(context.Background())
	if err != nil {
		// Log a fatal error and terminate the program if FCM client initialization fails.
		log.Fatalf("error initializing FCM client: %v", err)
	}
	// Assign the initialized client to the global FCMClient variable.
	FCMClient = client
}
