package main

import "fmt"

// Return the sum of all digits in the given string
func sumOfDigits(number string) int {
	sum := 0
	for _, digit := range number {
		if digit >= '0' && digit <= '9' {
			sum += int(digit - '0')
		}
	}
	return sum
}

func main() {
	fmt.Print("Enter a number: ")
	var number string
	fmt.Scanf("%s\n", &number)
	fmt.Println("Number:", number)
	fmt.Printf("Type: %T\n", number)
	fmt.Println('9' - '0')
	fmt.Println('9')

	sum := sumOfDigits(number)
	fmt.Println("Sum of digits:", sum)
}
