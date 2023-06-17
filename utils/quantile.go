package utils

import (
	"errors"
	"sort"

	"github.com/khoindq/tcbHomework/common"
)

func GetQuantileInNearestRanks(reqPercentile float64, data []float64) (calculatedQuantile *float64, err error) {
	// Check if the request percentile is within the valid range of 0 to 100
	if reqPercentile < 0 || reqPercentile > 100 {
		return nil, common.ErrInvalidRequest(errors.New("percentile must be in range 0 to 100"))
	}

	// Get the length of the data
	lengthData := len(data)

	// Check if the data length is valid (greater than zero)
	if lengthData < 1 {
		return nil, common.ErrInvalidRequest(errors.New("len of data must be greater than zero"))
	}

	// Sort the data in ascending order
	sort.Float64s(data)

	// Calculate the index based on the request percentile
	index := int64((reqPercentile / 100.0) * float64(lengthData-1))

	// Get the value at the calculated index (nearest rank)
	value := data[index]

	return &value, nil
}
