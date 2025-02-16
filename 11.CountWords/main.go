package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func countWords(s string) int {
	return len(strings.Fields(s))
}

func main() {
	var text string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a text to count its length:")
	text, _ = reader.ReadString('\n')
	fmt.Printf("Number of characters in text: %d\n", len(text))
	fmt.Printf("Number of words in text: %d\n", countWords(text))

	fmt.Println()
	fmt.Println("Single Quotes:", reflect.TypeOf('h'))
	fmt.Println("Double quotes:", reflect.TypeOf("Hello"))
}
