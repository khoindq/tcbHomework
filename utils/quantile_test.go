package utils

import (
	"testing"
)

// test case are compared (rechecked with) with np percentile python
func TestGetQuantileNearestRangeDataOnly(t *testing.T) {
	testCases := []struct {
		data             []float64
		reqPercentile    float64
		expectedQuantile float64
	}{
		{[]float64{10.0, 20.0, 30.0, 40.0, 50.0}, 50.0, 30.0},
		{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}, 80.0, 4.0},
		{[]float64{35.0, 15.0, 35.0, 45.0, 55.0}, 25.0, 35.0},
		{[]float64{15.0}, 100, 15},
		{[]float64{15.0}, 0, 15},
		{[]float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20}, 41.5, 8},
		{[]float64{2, 4, 6, 8, 10}, 10, 2.0},
		{[]float64{9, 3, 7, 2, 5}, 70.0, 7.0},
		{[]float64{6, 10, 2, 4, 8}, 20.0, 4.0},
		{[]float64{50, 30, 10, 40, 20}, 90.0, 50.0},
		{[]float64{15, 5, 25, 20, 35}, 60.0, 20.0},
		{[]float64{15, 5, 25, 20, 35}, 0, 5.0},
	}

	for i, testCase := range testCases {
		calculatedQuantile, err := GetQuantileInNearestRanks(testCase.reqPercentile, testCase.data)
		if err != nil {
			t.Errorf("Test case %d: Expected no error, got: %v", i+1, err)
		}
		if *calculatedQuantile != testCase.expectedQuantile {
			t.Errorf("Test case %d: Expected quantile %f, got %f", i+1, testCase.expectedQuantile, *calculatedQuantile)
		}
	}
}

func TestGetQuantileInPercentileSpecialCase(t *testing.T) {
	// TestGetQuantileInPercentileSpecialCase is a test function for GetQuantileInNearestRanks.
	// It tests special cases where the request percentile is invalid or the data length is invalid.

	// Testing invalid request percentile
	reqPercentile := -10.0
	data := []float64{15.0}
	expectedError := "percentile must be in range 0 to 100"

	_, err := GetQuantileInNearestRanks(reqPercentile, data)
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error '%s', got: %v", expectedError, err)
	}
}
