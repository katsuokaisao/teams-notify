package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/katsuokaisao/teams-notify/infra/teams"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	incomingWebhookURL := os.Getenv("INCOMING_WEBHOOK_URL")

	repo := teams.NewChatMessageRepository(incomingWebhookURL)
	err := repo.SendSomeErrorMessage("some error occurred", fmt.Errorf("some error"), map[string]string{
		"key1": "value1",
		"key2": "value2",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")
}
