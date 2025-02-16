package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Start timer: ")
	fmt.Scanln()
	start := time.Now()
	fmt.Print("End Timer: ")
	fmt.Scanln()
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("Duration in seconds:", duration.Seconds())
}
