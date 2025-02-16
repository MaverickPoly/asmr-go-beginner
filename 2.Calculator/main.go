package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Welcome to our calculator!")
	running := true
	for running {
		fmt.Println()
		var num1, num2, result float32
		var operator string
		fmt.Print("1. Addition\n2.Subtraction\n3.Division\n4.Multiplication\n5.Exit\nChoose an operator: ")
		fmt.Scanln(&operator)
		if !slices.Contains([]string{"1", "2", "3", "4", "5"}, operator) {
			fmt.Println("Invalid option! Please try again!")
			continue
		} else if operator == "5" {
			goto end
		}
		fmt.Print("Enter the first and second number: ")
		fmt.Scanf("%f %f", &num1, &num2)
		switch operator {
		case "1":
			result = num1 + num2
		case "2":
			result = num1 - num2
		case "3":
			result = num1 / num2
		case "4":
			result = num1 * num2
		}
		fmt.Println("Result:", result)
		fmt.Scanln()
		// Play Again?
		fmt.Print("\nDo you want to calculate again? (y/n): ")
		var again string
		fmt.Scanln(&again)
		running = strings.ToLower(again) == "y"
	}
end:
	fmt.Println("Thanks for using our calculator!")
}
