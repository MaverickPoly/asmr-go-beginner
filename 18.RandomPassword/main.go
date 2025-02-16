package main

import (
	"fmt"
	"math"
	"math/rand"
)

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var digits = "0123456789"
var punctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

func generateRandomPassword(length int) string {
	password := make([]byte, length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}
	return string(password)
}

func main() {
	fmt.Println("PI:", math.Pi)
	fmt.Println("Euclidian Number:", math.E)
	fmt.Println()

	var length int
	chars += digits
	chars += punctuation
	fmt.Print("Enter a length of a random password: ")
	fmt.Scanf("%d\n", &length)
	fmt.Println(generateRandomPassword(length))
}
