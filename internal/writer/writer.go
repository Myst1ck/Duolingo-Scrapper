package writer

import (
	"encoding/csv"
	"os"
	"spotify-scrapper/internal/shared"
)

func openFile(filePath string) *os.File {
	file, err := os.Create(filePath)

	if err != nil {
		shared.Log.LogData("Error when opening result file: " + err.Error())
	}

	return file
}

func writeData(file *os.File, data [][]string) {
	writer := csv.NewWriter(file)

	err := writer.WriteAll(data)

	if err != nil {
		shared.Log.LogData("Error when tried to start writing to the result file: " + err.Error())
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		shared.Log.LogData("Error when writing data to the result file: " + err.Error())
	}
}

func WriteResult(filePath string, data [][]string) {
	file := openFile(filePath)

	defer file.Close()

	writeData(file, data)
}
