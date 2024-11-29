package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	// sum := sumup(numbers)
	sum := sumup(1, 10, 15)
	fmt.Println(sum)
	anotherSum := sumup(numbers...) // unpack numbers into standalone values
	fmt.Println(anotherSum)
}

func sumup(numbers ...int) int { // ... collect the standalone values and create a slice behind the scenes (collect all parameter)
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
