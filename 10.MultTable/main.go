package main

import "fmt"

func generateTable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d ", i*j)
		}
		fmt.Println()
	}
}

func main() {
	var number int
	fmt.Print("Enter the number: ")
	fmt.Scanf("%d\n", &number)
	generateTable(number)
}
