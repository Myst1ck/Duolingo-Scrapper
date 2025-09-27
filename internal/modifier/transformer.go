package modifier

import (
	"spotify-scrapper/internal/shared"
	"strings"
)

func TransformLexemes(lexemes *[]shared.Lexem) *[][]string {
	modifiedLexemes := make([][]string, len(*lexemes))

	for index, lexem := range *lexemes {
		modifiedLexem := TransformedLexem{
			Text:         lexem.Text,
			Translations: strings.Join(lexem.Translations, ", "),
		}

		modifiedLexemes[index] = modifiedLexem.FormatToCSV()
	}

	return &modifiedLexemes
}

func (modifiedLexem *TransformedLexem) FormatToCSV() []string {
	return []string{modifiedLexem.Text, modifiedLexem.Translations}
}
