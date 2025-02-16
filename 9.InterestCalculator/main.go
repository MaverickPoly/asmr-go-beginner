package main

import (
	"fmt"
	"reflect"
)

func CalculateSimpleInterest(principal, rate, time float64) float64 {
	return (principal * rate * time) / 100
}

func main() {
	var principal, rate, time float64

	fmt.Print("Enter the principal amount: ")
	fmt.Scanln(&principal)

	fmt.Print("Enter the interest rate (in %): ")
	fmt.Scanln(&rate)

	fmt.Print("Enter the time (in years): ")
	fmt.Scanln(&time)

	// Get type of variables
	fmt.Println("Principal: ", reflect.TypeOf(principal))
	fmt.Println("Rate: ", reflect.TypeOf(rate))
	fmt.Println("Time: ", reflect.TypeOf(time))

	interestAmount := CalculateSimpleInterest(principal, rate, time)
	fmt.Printf("Simple Interest: %.2f\n", interestAmount)

}
