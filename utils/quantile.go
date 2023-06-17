package utils

import (
	"errors"
	"math"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/khoindq/tcbHomework/common"
)

func GetQuantileInNearestRanks(reqPercentile float64, data []float64) (calculatedQuantile *float64, err error) {
	// Check if the request percentile is within the valid range of 0 to 100
	if reqPercentile < 0 || reqPercentile > 100 {
		return nil, common.ErrInvalidRequest(errors.New("percentile must be in range 0 to 100"))
	}

	// Get the length of the data
	lengthData := len(data)

	// Sort the data in ascending order
	sort.Float64s(data)

	if len(data) == 0 {
		return aws.Float64(math.NaN()), nil
	}

	if reqPercentile == 0 {
		return aws.Float64(data[0]), nil
	}
	if reqPercentile == 100 {
		return aws.Float64(data[lengthData-1]), nil
	}
	// Find the index of the quantile
	i := int(math.Round(reqPercentile / 100.0 * float64(lengthData-1)))
	return aws.Float64(data[i]), nil
}
