package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Location())
	fmt.Println(time.Since(time.Time{}))
	fmt.Println(time.April)
	fmt.Println(time.Second)
	fmt.Println(time.Minute)
	fmt.Println(time.Hour)
	fmt.Println(time.Monday)
	fmt.Println(time.Date(2025, 02, 10, 13, 30, 30, 1, time.Now().Location()))
}
