package main

import "fmt"

func fibonacciN(n int) int64 {
	if n <= 2 {
		return 1
	}
	return fibonacciN(n-1) + fibonacciN(n-2)
}

func generateSequence(n int) []int64 {
	res := []int64{}
	for i := 1; i <= n; i++ {
		res = append(res, fibonacciN(i))
	}
	return res
}

func main() {
	fmt.Println("Fibonacci Sequence")
	fmt.Println("=================================")
	fmt.Println(fibonacciN(10))
	fmt.Println(fibonacciN(2))
	fmt.Println(fibonacciN(30))

	fmt.Println("Generating sequence:", generateSequence(10))
}
