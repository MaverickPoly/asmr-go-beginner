package main

import "fmt"

func factorialN(n int) int64 {
	if n <= 1 {
		return 1
	}
	return int64(n) * factorialN(n-1)
}

func generateFactorial(n int) []int64 {
	res := []int64{}
	for i := 1; i <= n; i++ {
		res = append(res, factorialN(i))
	}
	return res
}

func main() {
	fmt.Println("Factorial Generator program:")
	fmt.Println("=======================================================")
	fmt.Println(factorialN(20))
	fmt.Println(factorialN(9))
	fmt.Println(factorialN(6))
	fmt.Println("Generated sequence:", generateFactorial(10))
	fmt.Println("Generated sequence:", generateFactorial(20))
}
