package main

import (
	"fmt"
	"time"
)

func main() {
	var number int
	fmt.Print("Enter the number to count down: ")
	fmt.Scanf("%d\n", &number)
	fmt.Println()

	for i := number; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Duration(1) * time.Second)
	}
	fmt.Println("Blast off!")
}
