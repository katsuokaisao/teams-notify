package main

import (
	"fmt"
	"log"

	"github.com/katsuokaisao/teams-notify/infra/teams"
)

func main() {
	repo := teams.NewChatMessageRepository("https://outlook.office.com/webhook/xxxxxx")
	err := repo.SendSomeErrorMessage("some error occurred", fmt.Errorf("some error"), map[string]string{
		"key1": "value1",
		"key2": "value2",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")
}
