package prices

import (
	"fmt"
	"math"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"` // tag struct
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"` // Input price converted to a string maps to a result price in float64
	IOManager         iomanager.IOManager `json:"-"`                   // tag structs (metadata), ignore the TaxRate key and value in the output
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.ConvertStringToFloat64s(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) { // The receiver argument (job TaxIncludedPriceJob) turns this function into a method
	result := make(map[string]string)

	err := job.LoadData()

	if err != nil {
		errorChan <- err
		return // empty return statements work in go routines
	}

	for _, price := range job.InputPrices {
		priceText := fmt.Sprintf("%.2f", price)
		taxIncludedPrice := math.Round(price*(1+job.TaxRate)*100) / 100
		result[priceText] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	doneChan <- true

}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: iom,
	}
}
