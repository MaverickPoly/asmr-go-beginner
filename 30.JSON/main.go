package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
	City string
}

func writeFile(path string, person Person) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}

func readFile(path string) Person {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return Person{}
	}
	defer file.Close()

	var person Person
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding file:", err)
	}
	return person
}

func main() {
	// Write File
	person := Person{Name: "Jamal", Age: 30, City: "Morocco"}
	writeFile("30.JSON/person.json", person)
	fmt.Println("JSON data written successfully to person.json")

	// Read File
	readPerson := readFile("30.JSON/person.json")
	fmt.Println("Read from JSON File:", readPerson)
}
