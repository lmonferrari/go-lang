package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func validateCSVFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Failed to open file: %s", err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("File is empty")
		}
		return fmt.Errorf("Failed to read header: %s", err.Error())
	}

	expectedHeader := []string{"name", "age", "email"}
	if !equalSlices(header, expectedHeader) {
		return fmt.Errorf("Unexpected header: %v", header)
	}

	// valid file
	return nil
}

func main() {
	err := validateCSVFile("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("CSV file is valid!")
}
