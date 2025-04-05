package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type FileReader struct {
	fileName string
}

func NewFileReader(fileName string) FileReader {
	return FileReader{fileName: fileName}
}
func (this FileReader) getMatrix() [][]string {

	file, err := os.Open(this.fileName)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)

	fmt.Println("CSV content:")
	for {
		record, err := reader.Read()
		if err != nil {
			fmt.Println("Error reading record:", err)
			break // Exit on error or EOF
		}
		fmt.Println(record) // Each 'record' is a slice of strings representing a row
	}
	return nil
}
