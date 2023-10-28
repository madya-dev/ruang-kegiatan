package firebase

import (
	"context"
	"fmt"
	"log"
	"madyasantosa/ruangkegiatan/config"
	"madyasantosa/ruangkegiatan/helper"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func SendNotification(targetRegistration string, title string, body string) {

	credential, err := helper.GetDecodedFireBaseKey(*config.InitConfig())
	if err != nil {
		fmt.Errorf("Failed to get credential: %v", err)
	}

	// Konfigurasi Firebase Admin SDK
	opt := option.WithCredentialsJSON(credential)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase app: %v", err)
	}

	// Inisialisasi FCM client
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Failed to create FCM client: %v", err)
	}

	// Data pesan yang akan dikirim
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: targetRegistration, // Token perangkat klien yang akan menerima notifikasi
	}

	// Kirim notifikasi
	response, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	}

	fmt.Println("Successfully sent notification:", response)
}
