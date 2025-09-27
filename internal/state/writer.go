package state

import (
	"encoding/json"
	"os"
	"spotify-scrapper/internal/shared"
)

func writeFile(filePath string, state *ProgramState) {
	data, marshalErr := json.MarshalIndent(state, "", " ")

	if marshalErr != nil {
		shared.Log.LogData("There was an error when marshalling the state: " + marshalErr.Error())
	}

	writingErr := os.WriteFile(filePath, data, 0644)

	if writingErr != nil {
		shared.Log.LogData("There was an error reading the input file: " + writingErr.Error())
	}
}

func SaveState(fileName string, state *ProgramState) {
	writeFile(fileName, state)
}
