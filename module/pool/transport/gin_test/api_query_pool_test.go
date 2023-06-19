package gintest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
	"github.com/khoindq/tcbHomework/common"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
	poolgin "github.com/khoindq/tcbHomework/module/pool/transport/gin"
)

func TestAPIQuantileSuccess(t *testing.T) {
	// Create a new Gin router
	getEndpoint := "/api/quantile/get"
	insertOrAppendEndpoint := "/api/pool/insertappend"
	router := gin.Default()

	// Define the route for the API handler
	router.POST(getEndpoint, poolgin.NewPoolController().GetQuantileHandler())
	router.POST(insertOrAppendEndpoint, poolgin.NewPoolController().InsertOrAppendPoolHandler())

	// Send 1,000 requests to the API handler
	for i := 0; i <= 1000; i++ {
		// Create an insert or append to the pool
		pool := poolmodel.Pool{
			PoolID:     aws.Int64(int64(i)),
			PoolValues: []float64{1, 2, 4, 5, 6, 7, 8, 9},
		}
		requestBody, err := json.Marshal(pool)
		if err != nil {
			panic(err)
		}

		// Create a new test request
		reqBody := bytes.NewReader(requestBody)
		req, err := http.NewRequest("POST", insertOrAppendEndpoint, reqBody)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Create a new test response recorder
		res := httptest.NewRecorder()

		// Serve the request using the test router
		router.ServeHTTP(res, req)

		// Check the response status code
		if res.Code != http.StatusOK {
			t.Errorf("Request %d failed with status code %d", i+1, res.Code)
		}
	}

	// Send the get query
	for i := 0; i <= 1000; i++ {
		// Create an resquest data object
		poolGet := poolmodel.PoolQuantileGet{
			PoolID:     aws.Int64(int64(i)),
			Percentile: aws.Float64(50),
		}
		requestBody, err := json.Marshal(poolGet)
		if err != nil {
			panic(err)
		}

		// Create a new test request
		reqBody := bytes.NewReader(requestBody)
		req, err := http.NewRequest("POST", getEndpoint, reqBody)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Create a new test response recorder
		res := httptest.NewRecorder()

		// Serve the request using the test router
		router.ServeHTTP(res, req)

		// Check the response status code
		if res.Code != http.StatusOK {
			t.Errorf("Request %d failed with status code %d", i+1, res.Code)
		}

		// Test the get quantile and total count result
		var response struct {
			Data poolmodel.PoolQuantileResp `json:"data"`
		}
		err = json.Unmarshal(res.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("Failed to unmarshal response body: %v", err)
		}
		fmt.Println(response)

		if response.Data.CalculatedQuantile != 6 {
			t.Errorf("Request %d: Expected calculated quantile is 6, got '%f'", i+1, response.Data.CalculatedQuantile)
		}
		if response.Data.TotalCount != 8 {
			t.Errorf("Request %d: Expected totalCount is 8, got '%d'", i+1, response.Data.TotalCount)
		}
	}
	t.Cleanup(func() {
		common.FakeDB.Clean()
	})
}
