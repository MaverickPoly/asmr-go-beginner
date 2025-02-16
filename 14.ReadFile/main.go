package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile() {
	// Open File
	file, err := os.Open("14.ReadFile/text.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// New Reader
	reader := bufio.NewReader(file)
	fmt.Println("Reading content of a file:")

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		fmt.Print(line)
	}

	// Close File
	defer file.Close()
}

// ?Not very practical
func fileRead() {
	file, err := os.Open("14.ReadFile/text.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	text := "edsakleldjsamiqewkjdlasmwks"
	b := []byte(text)
	file.Read(b)
	fmt.Println(string(b))
	defer file.Close()
}

func readFile2() {
	data, err := os.ReadFile("14.ReadFile/text.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(data))
}

func main() {
	fileRead()
	readFile2()
}
