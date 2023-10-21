package teams

type ChatMessageRepository interface {
	SendSomeErrorMessage(msg string, err error, options map[string]string) error
}
