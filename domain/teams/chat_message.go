package teams

type ChatMessage struct {
	Type       string    `json:"@type"`
	Context    string    `json:"@context"`
	Summary    string    `json:"summary"`
	ThemeColor string    `json:"themeColor"`
	Sections   []section `json:"sections"`
}

type section struct {
	Title            string `json:"title"`
	Text             string `json:"text"`
	ActivityTitle    string `json:"activityTitle"`
	ActivitySubtitle string `json:"activitySubtitle"`
	ActivityImage    string `json:"activityImage"`
	Facts            []fact `json:"facts"`
}

type fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
