package main

import (
	"fmt"
	"time"
)

func returnAge(year int) int {
	fmt.Println("Current year:", time.Now().Year())
	return time.Now().Year() - year
}

func main() {
	for true {
		var year int
		fmt.Print("Enter the year you were born (-1 to quit): ")
		fmt.Scanf("%d\n", &year)
		if year == -1 {
			break
		}
		if year > time.Now().Year() {
			fmt.Println("Invalid year:", year)
			continue
		}
		fmt.Printf("You are %d years old!\n", returnAge(year))
		time.Sleep(2)
	}
	fmt.Println("Thanks for calculating!")
}
