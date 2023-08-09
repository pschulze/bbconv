package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// Ask for filename and retrieve files
	inFileName := fileNamePrompt()
	records := readCsv(inFileName)

	csvOutFile, err := os.Create(outFileName(inFileName))
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer csvOutFile.Close()

	csvWriter := csv.NewWriter(csvOutFile)
	defer csvWriter.Flush()

	// Read headers, get index of Date column
	headers := records[0]
	dateIndex, _ := dateIndex(headers)

	// Write headers to new file
	csvWriter.Write(headers)

	// Line by line: read old file, update date column, write to new file
	for _, record := range records[1:] {
		record[dateIndex] = convertDate(record[dateIndex])
		_ = csvWriter.Write(record)
	}
}
