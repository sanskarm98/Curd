package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App
var FCMClient *messaging.Client

func InitFirebase(credentialsFilePath string) {
	opt := option.WithCredentialsFile(credentialsFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v", err)
	}
	FirebaseApp = app

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error initializing FCM client: %v", err)
	}
	FCMClient = client
}
