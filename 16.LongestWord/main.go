package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func longestWord(s string) string {
	words := strings.Fields(s)
	fmt.Println("Split Words:", words)
	if len(words) == 0 {
		return ""
	}

	longest := words[0]
	for _, word := range words[1:] {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a sentence:")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)
	fmt.Println("Sentence:", sentence+"\n")
	fmt.Println("\nLongest word:", longestWord(sentence))
	fmt.Println("Reversed:", reverseString(sentence))
}
