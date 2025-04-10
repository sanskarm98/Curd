package notification

import (
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendEmail sends an email using the SendGrid API.
// Parameters:
// - toEmail: The recipient's email address.
// - subject: The subject of the email.
// - body: The content of the email.
// Returns an error if the email fails to send.
func SendEmail(toEmail, subject, body string) error {
	// Create the sender's email information.
	from := mail.NewEmail("Your App Name", "your-email@example.com")

	// Create the recipient's email information.
	to := mail.NewEmail("Recipient", toEmail)

	// Create a single email message with the provided subject and body.
	message := mail.NewSingleEmail(from, subject, to, body, body)

	// Initialize the SendGrid client using the API key from environment variables.
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	// Send the email and capture the response or error.
	response, err := client.Send(message)
	if err != nil {
		// Log the error if the email fails to send.
		log.Printf("Failed to send email: %v", err)
		return err
	}

	// Log the status code of the email response.
	log.Printf("Email sent. Status Code: %d", response.StatusCode)
	return nil
}
