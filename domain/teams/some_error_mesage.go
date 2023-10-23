package teams

func NewSomeErrorMessage(msg string, err error, options map[string]string) *ChatMessage {
	chatMessage := &ChatMessage{
		Type:       "MessageCard",
		Context:    "http://schema.org/extensions",
		Summary:    msg,
		ThemeColor: "0076D7",
		Sections: []section{
			{
				Title: msg,
				Text:  err.Error(),
			},
		},
	}
	if len(options) > 0 {
		facts := make([]fact, 0, len(options))
		for k := range options {
			facts = append(facts, fact{
				Name:  k,
				Value: options[k],
			})
		}
		chatMessage.Sections[0].Facts = facts
	}
	return chatMessage
}
