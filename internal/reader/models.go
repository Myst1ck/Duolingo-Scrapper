package reader

import "spotify-scrapper/internal/shared"

type Input struct {
	LearnedLexemes []shared.Lexem `json:"learnedLexemes"`
}
