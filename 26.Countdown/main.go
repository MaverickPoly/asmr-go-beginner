package main

import (
	"fmt"
	"time"
)

func main() {
	var seconds int
	fmt.Print("Enter amount in seconds: ")
	fmt.Scanf("%d\n", &seconds)

	for i := 0; i < seconds; i++ {
		fmt.Println(seconds - i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("TIME IS UP!!")
}
