package state

import (
	"encoding/json"
	"os"
	"spotify-scrapper/internal/shared"
)

func readFile(filePath string) *ProgramState {
	fileBytes, readingErr := os.ReadFile(filePath)

	if readingErr != nil {
		shared.Log.LogData("There was an error reading the input file: " + readingErr.Error())
	}

	var state ProgramState

	unmarshalErr := json.Unmarshal(fileBytes, &state)

	if unmarshalErr != nil {
		shared.Log.LogData("There was an error unmarshaling the input file: " + unmarshalErr.Error())
	}

	return &state
}

func GetState(filePath string) *ProgramState {
	state := readFile(filePath)

	return state
}
