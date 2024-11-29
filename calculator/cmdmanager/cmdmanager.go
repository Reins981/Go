package cmdmanager

import "fmt"

type CMDMangager struct{}

func (m CMDMangager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER, 0 for EXIT")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil

}

func (m CMDMangager) WriteResult(data interface{}) error { // interface{} or any -> Accept any value
	fmt.Println(data)
	return nil
}

func New() CMDMangager {
	return CMDMangager{}
}
