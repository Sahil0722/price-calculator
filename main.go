package main

import (
	"fmt"

	"github.com/Sahil0722/price-calculator/conversions"
	"github.com/Sahil0722/price-calculator/filemanager"
	"github.com/Sahil0722/price-calculator/prices"
)

func main() {
	fmTaxRates := filemanager.New("tax_rates.txt", "tax_rates.json")
	lines, err := fmTaxRates.ReadLines()

	if err != nil {
		println(err)
		return
	}

	taxRates, err := conversions.StringToFloat64(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
