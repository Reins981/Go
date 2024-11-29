package conversion

import (
	"errors"
	"strconv"
)

func ConvertStringToFloat64s(inputSlice []string) ([]float64, error) {

	floats := make([]float64, len(inputSlice))

	for sliceIndex, sliceVal := range inputSlice {
		floatPrice, err := strconv.ParseFloat(sliceVal, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		floats[sliceIndex] = floatPrice
	}

	return floats, nil
}
