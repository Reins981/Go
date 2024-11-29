package main

import "fmt"

// Anonymous functions and Closures

func main() {
	numbers := []int{1, 2, 3}

	/*transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(transformed)
	*/

	double := createTransformer(2)
	triple := createTransformer(3)

	doubled := transformNumbers(&numbers, double)
	trippled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(trippled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor // This is the Closure, access to the outer scope function
	}
}
