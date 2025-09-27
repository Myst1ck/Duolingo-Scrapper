package main

import (
	"spotify-scrapper/internal/modifier"
	"spotify-scrapper/internal/reader"
	"spotify-scrapper/internal/shared"
	"spotify-scrapper/internal/state"
	"spotify-scrapper/internal/writer"
)

const INPUT_PATH = "./data/input.json"
const OUTPUT_PATH = "./data/output.csv"

const SHOULD_USE_STATE = false
const STATE_PATH = "./data/state.json"

const LATEST_WORD = "Kochen"

func getLatestWord() string {
	if SHOULD_USE_STATE {
		state := state.GetState(STATE_PATH)
		return state.LatestWord
	}

	return LATEST_WORD
}

func saveLatestState(modifiedLexemes *[][]string, latestWord string) {
	var newState string

	if len(*modifiedLexemes) > 0 {
		newState = (*modifiedLexemes)[0][0]
	} else {
		newState = latestWord
	}

	state.SaveState(STATE_PATH, &state.ProgramState{
		LatestWord: newState,
	})
}

func main() {
	latestWord := getLatestWord()

	lexemes := reader.ReadInput(INPUT_PATH)
	relevantLexemes := modifier.GetRelevantLexemes(lexemes, latestWord)

	if len(*relevantLexemes) == len(*lexemes) {
		shared.Log.LogData("Didn't arrive to the latest word, try to use a bigger batch.")
		shared.Log.LogData("Exiting program.")
		return
	}

	modifiedLexemes := modifier.TransformLexemes(relevantLexemes)

	writer.WriteResult(OUTPUT_PATH, *modifiedLexemes)

	saveLatestState(modifiedLexemes, latestWord)
}
