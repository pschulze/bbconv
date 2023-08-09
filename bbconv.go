package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func fileNamePrompt() string {
	var s string
	prompt := "Enter the name of the file you'd like to convert: "
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, prompt)
		s, _ = reader.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func outFileName(inFileName string) string {
	return strings.TrimSuffix(filepath.Base(inFileName), filepath.Ext(inFileName)) +
		"-fixed" +
		filepath.Ext(inFileName)
}

func readCsv(fileName string) [][]string {
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Unable to read input file "+fileName, err)
	}

	sanitizedFileContents := strings.Replace(string(fileContents), " \r\n", "\n", -1)

	csvReader := csv.NewReader(strings.NewReader(sanitizedFileContents))
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+fileName, err)
	}

	return records
}

func dateIndex(record []string) (int, error) {
	index := -1

	for i, val := range record {
		if val == "Date" && index == -1 {
			index = i
		} else if val == "Date" {
			return -1, errors.New("ambiguous Date column")
		}
	}

	if index == -1 {
		return -1, errors.New("no Date column found")
	} else {
		return index, nil
	}
}

func convertDate(date string) string {
	return "<Fixed Date Here>"
}
