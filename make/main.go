package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// this is a slice, not an array but of course it points to an array underlying.
	// That means co creates a new array whenever a new entry is inserted.
	//userNames := []string{}

	// Create a preallocated slice with make that has 2 empty slots (Good for memory management)
	userNames := make([]string, 2, 5) // 5 is the capacity (the maximum number of array items). Go will reserve space for 5 items in memory
	// The next 2 elements will be appended but will not occupy the 2 slots
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	// occupy first 2 slots
	userNames[0] = "Julie"
	userNames[1] = "Mark"
	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9
	courseRatings["node"] = 5.0

	courseRatings.output()

	// fmt.Println(courseRatings)

	for index, value := range userNames {
		// ...
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRatings {
		// ...
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}

}
