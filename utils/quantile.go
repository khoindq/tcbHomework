package utils

import (
	"errors"
	"sort"

	"github.com/khoindq/tcbHomework/common"
)

func GetQuantileInNearestRanks(reqPercentile float64, data []float64) (calculatedQuantile *float64, err error) {
	if reqPercentile < 0 || reqPercentile > 100 {
		return nil, common.ErrInvalidRequest(errors.New("percentile must be in range 0 to 100"))
	}

	lengthData := len(data)
	if lengthData < 1 {
		return nil, common.ErrInvalidRequest(errors.New("len of data must be greater than zero"))
	}

	sort.Float64s(data)

	index := int64((reqPercentile / 100.0) * float64(lengthData-1))
	value := data[index]

	return &value, nil

}
