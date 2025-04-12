package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type FileReader struct {
	filePath string
}

func NewFileReader(filePath string) *FileReader {
	return &FileReader{filePath: filePath}
}
func (me FileReader) readCSVToMatrix() ([][]string, error) {
	// Open the CSV file
	file, err := os.Open(me.filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %w", err)
	}

	return records, nil
}
