package utils

import (
	"testing"
)

// test case are  compared with np percentile python
func TestGetQuantileNearestRangeDataOnly(t *testing.T) {
	testCases := []struct {
		data             []float64
		reqPercentile    float64
		expectedQuantile float64
	}{
		// Test case 1
		{
			data:             []float64{10.0, 20.0, 30.0, 40.0, 50.0},
			reqPercentile:    50.0,
			expectedQuantile: 30.0,
		},
		// Test case 2
		{
			data:             []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			reqPercentile:    80.0,
			expectedQuantile: 4.0,
		},
		// Test case 3
		{
			data:             []float64{35.0, 15.0, 35.0, 45.0, 55.0},
			reqPercentile:    25.0,
			expectedQuantile: 35.0,
		},
		// Test case 4
		{
			data:             []float64{15.0},
			reqPercentile:    100,
			expectedQuantile: 15,
		},
		// Test case 5
		{
			data:             []float64{15.0},
			reqPercentile:    0,
			expectedQuantile: 15,
		},
		// Test case 6
		{
			data:             []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20},
			reqPercentile:    41.5,
			expectedQuantile: 8,
		},
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

	// Testing invalid request percentile
	reqPercentile := -10.0
	data := []float64{15.0}
	expectedError := "percentile must be in range 0 to 100"

	_, err := GetQuantileInNearestRanks(reqPercentile, data)
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error '%s', got: %v", expectedError, err)
	}

	// Testing invalid data length
	emptyData := []float64{}
	reqPercentile = 5.0
	expectedError = "len of data must be greater than zero"

	_, err = GetQuantileInNearestRanks(reqPercentile, emptyData)
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error '%s', got: %v", expectedError, err)
	}
}
