package main

import (
	"fmt"
	"strings"
)

func celToFahr(cel float32) float32 {
	return (9/5)*cel + 32
}

func fahrToCel(fahr float32) float32 {
	return (fahr - 32) * 5 / 9
}

func main() {
	running := true
	for running {
		var choice, again string
		var input, result float32
		fmt.Print("1. Celsius to Fahrenheit\n2.Fahrenheit to Celsius\nSelect a choice: ")
		fmt.Scanln(&choice)

		fmt.Print("\nEnter the temperature: ")
		fmt.Scanf("%f\n", &input)

		switch choice {
		case "1":
			result = celToFahr(input)
			fmt.Printf("%.2f °C = %.2f °F\n", input, result)
		case "2":
			result = fahrToCel(input)
			fmt.Printf("%.2f °F = %.2f °C\n", input, result)
		}

		fmt.Print("Calculate again?(y/n): ")
		fmt.Scanln(&again)
		running = strings.ToLower(again) != "n"
		fmt.Println()
	}
	fmt.Println("Thanks for calculating!")
}
