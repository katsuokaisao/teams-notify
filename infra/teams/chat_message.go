package teams

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/katsuokaisao/teams-notify/domain/teams"
)

type chatMessageRepositoryImpl struct {
	incomingWebhookURL string
}

func NewChatMessageRepository(incomingWebhookURL string) teams.ChatMessageRepository {
	return &chatMessageRepositoryImpl{
		incomingWebhookURL: incomingWebhookURL,
	}
}

func (c *chatMessageRepositoryImpl) SendSomeErrorMessage(msg string, err error, options map[string]string) error {
	chatMessage := teams.NewSomeErrorMessage(msg, err, options)
	return c.send(chatMessage)
}

func (c *chatMessageRepositoryImpl) send(chatMessage *teams.ChatMessage) error {
	body, err := json.Marshal(chatMessage)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.incomingWebhookURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err == nil {
		fmt.Printf("response: %s\n", string(bodyBytes))
	}

	return nil
}
