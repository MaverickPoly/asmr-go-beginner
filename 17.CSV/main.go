package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Reads as a normal string
func readCSVFake(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func readCSV(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for _, record := range records {
		fmt.Println(record)
	}

}

func writeCSV(data [][]string, path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// writer.WriteAll(data)  // *Alternative to this ⬇️
	for _, record := range data {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error writing record:", err)
			return
		}
	}

	fmt.Println("CSV file written successfully!")
}

func main() {
	data := [][]string{
		{"Name", "Age", "Country"},
		{"Alice", "25", "USA"},
		{"Bob", "30", "UK"},
		{"Charlie", "28", "Canada"},
	}

	writeCSV(data, "17.CSV/data.csv")
	readCSV("17.CSV/data.csv")
}
