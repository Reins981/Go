package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))   // create channel slice with empty slots
	errorChans := make([]chan error, len(taxRates)) // create channel slice with empty slots

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index]) // go routines do not return values, instead we need to use channels

		/* if err != nil {
			fmt.Println("could not process Job!")
			fmt.Println(err)
		} */
	}

	// Wait loop until channels have emitted values
	for index := range taxRates {
		select {
		// Select which channel emits data (Either the error or done channel). // Only the channel which emits data earlier will be executed.
		// So if a value from a channel is received, select will move on and will not care about the other case (other channel)
		// That means, wait for either an error or the success case
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]: // we dont care about the boolean value
			fmt.Println("done!")
		}
	}

}
