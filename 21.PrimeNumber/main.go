package main

import "fmt"

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	if number <= 3 {
		return true
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	i := 5
	for i*i <= number {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}

func generatePrimeNums(n int) []int {
	primes := []int{}
	number := 2
	for len(primes) < n {
		if isPrime(number) {
			primes = append(primes, number)
		}
		number++
	}
	return primes
}

// Get N's Prime Number
func getNPrime(n int) int {
	if n <= 0 {
		return -1
	}
	count := 0
	number := 2
	for {
		if isPrime(number) {
			count++
			if count == n {
				return number
			}
		}
		number++
	}
}

func main() {
	fmt.Println("Is 2 Prime:", isPrime(2))
	fmt.Println("Is 3 Prime:", isPrime(3))
	fmt.Println("Is 4 Prime:", isPrime(4))
	fmt.Println()
	fmt.Println("Generate 10:", generatePrimeNums(20))
	fmt.Println("Get Prime N:", getNPrime(6))
}
