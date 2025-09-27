package shared

type Lexem struct {
	Text         string   `json:"text"`
	Translations []string `json:"translations"`
	AudioUrl     string   `json:"audioURL"`
	IsNew        bool     `json:"isNew"`
}
