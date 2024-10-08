package main

import "fmt"

// Bonus
type Product struct {
	title string
	id    int
	price float64
}

func main() {

	// 1
	hobbies := [3]string{"Tennis", "Football", "Golf"}
	fmt.Println(hobbies)
	// 2
	fmt.Println(hobbies[0])
	hobbiesSlice := hobbies[1:]
	fmt.Println(hobbiesSlice)
	// 3
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5
	goals := []string{"Goal1", "Goal2"}
	fmt.Println(goals)

	// 6
	goals[1] = "NewGoal"
	goals = append(goals, "Goal3")
	fmt.Println(goals)

	productList := []Product{
		{"Main", 1, 22.0},
		{"Sub", 2, 345.7},
	}
	newProduct := Product{
		title: "Test",
		id:    3,
		price: 100.5,
	}
	productList = append(productList, newProduct)
	fmt.Println(productList)

	prices := []float64{1.2, 2.3}
	prices = append(prices, 3.5)
	discountPrices := []float64{5.9, 6.7, 4.4}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
