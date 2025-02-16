package main

import (
	"fmt"
	"slices"
)

func reverseSlice(slice []int) {
	slices.Reverse(slice)
}

func reverseAlgo(slice []int) {
	for start, end := 0, len(slice)-1; start < end; start, end = start+1, end-1 {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6)
	fmt.Println("Max:", slices.Max(numbers))
	fmt.Println("Min:", slices.Min(numbers))
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
	fmt.Println(numbers)
	reverseSlice(numbers)
	fmt.Println(numbers)

	fmt.Println("Algorithms way:")
	numbers2 := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers2)
	reverseAlgo(numbers2)
	fmt.Println(numbers2)

}
