package main

import (
	"fmt"
	"strings"
)

func isEven(x int) bool {
	return x%2 == 0
}

func main() {
	for true {
		var num int
		var choice string
		fmt.Print("Enter a number to check: ")
		fmt.Scanf("%d\n", &num)
		if isEven(num) {
			fmt.Printf("Number %d is even\n", num)
		} else {
			fmt.Printf("Number %d is odd\n", num)
		}

		fmt.Println("Check again?(y/n): ")
		fmt.Scanln(&choice)
		if strings.ToLower(choice) != "y" {
			break
		}
		fmt.Println()
	}
}
