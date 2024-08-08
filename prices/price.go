package prices

import (
	"fmt"

	"github.com/Sahil0722/price-calculator/conversions"
	"github.com/Sahil0722/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOmanager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	Inputprices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOmanager: fm,
		TaxRate:   taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.Inputprices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	err = job.IOmanager.WriteJsonToFile(job)

	if err != nil {
		return err
	}

	return nil
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOmanager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversions.StringToFloat64(lines)

	if err != nil {
		fmt.Println(err)
		return err
	}

	job.Inputprices = prices
	return nil
}
