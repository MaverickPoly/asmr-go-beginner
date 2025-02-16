package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("Number guessing game!")
	fmt.Println("Guess the random number between 0 and 100 to win!")
	const min, max = 0, 100
	running := true
	for running {
		var again string
		var randNumber = rand.Intn(max-min) + min
		var attempts = 0
		var guess int

		for true {
			fmt.Print("Take a guess: ")
			fmt.Scanf("%d\n", &guess)
			attempts++
			if guess == randNumber {
				break
			} else if guess > randNumber {
				fmt.Println("Too high!")
			} else {
				fmt.Println("Too low!")
			}
		}
		fmt.Printf("Correct! The answer was %d, it took you %d attempts to find!\n", randNumber, attempts)

		fmt.Print("Play again?(y/n): ")
		fmt.Scanln(&again)
		running = strings.ToLower(again) != "n"
		fmt.Println()
	}
	fmt.Println("Thanks for playing!")
}
