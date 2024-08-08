package conversions

import (
	"errors"
	"strconv"
)

func StringToFloat64(lines []string) ([]float64, error) {
	floatValues := make([]float64, len(lines))

	for index, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return nil, errors.New("Failed to convert string to float.")
		}

		floatValues[index] = floatPrice
	}

	return floatValues, nil
}
