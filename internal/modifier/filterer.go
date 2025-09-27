package modifier

import (
	"spotify-scrapper/internal/shared"
	"strings"
)

func GetRelevantLexemes(lexemes *[]shared.Lexem, latestWord string) *[]shared.Lexem {
	var relevantLexemes []shared.Lexem

	for _, lexem := range *lexemes {
		if strings.EqualFold(lexem.Text, latestWord) {
			break
		}

		relevantLexemes = append(relevantLexemes, lexem)
	}

	return &relevantLexemes
}
