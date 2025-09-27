package reader

import (
	"encoding/json"
	"os"

	"spotify-scrapper/internal/shared"
)

func readFile(filePath string) *Input {
	fileBytes, readingErr := os.ReadFile(filePath)

	if readingErr != nil {
		shared.Log.LogData("There was an error reading the input file: " + readingErr.Error())
	}

	var input Input

	unmarshalErr := json.Unmarshal(fileBytes, &input)

	if unmarshalErr != nil {
		shared.Log.LogData("There was an error unmarshaling the input file: " + unmarshalErr.Error())
	}

	return &input
}

func ReadInput(filePath string) *[]shared.Lexem {
	input := readFile(filePath)

	return &input.LearnedLexemes
}
