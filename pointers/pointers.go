package main

import "fmt"

func main() {
	age := 32 //regular variable

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", *agePointer)
	fmt.Println(agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult Years:", age)
	fmt.Println(agePointer)
}

func editAgeToAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18
}
