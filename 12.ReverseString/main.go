package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string to reverse: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	reversed := reverseString(text)
	fmt.Println("Original:", text)
	fmt.Println("Reversed:", reversed)

}
