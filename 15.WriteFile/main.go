package main

import (
	"fmt"
	"os"
)

func appendFile(filePath, text string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	_, err = file.WriteString(text + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data appended successfully!")
}

func writeFile(filePath, text string) {
	file, err := os.Create(filePath)
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	data := []byte(text)
	_, err = file.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data written successfully!")
}

func main() {
	// writeFile("15.WriteFile/text.txt", "Hello world")
	appendFile("15.WriteFile/text.txt", "Hello world")
}
