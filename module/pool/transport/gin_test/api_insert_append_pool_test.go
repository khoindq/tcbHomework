package gintest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
	"github.com/khoindq/tcbHomework/common"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
	poolgin "github.com/khoindq/tcbHomework/module/pool/transport/gin"
)

func TestAPIInsertOrAppendPoolHandlerReturnSuccess(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Define the route for the API handler
	router.POST("/api/handler", poolgin.NewPoolController().InsertOrAppendPoolHandler())

	// Send 1,000 requests to the API handler
	var i int64
	for i = 0; i <= 1000; i++ {
		// Create a test request body
		pool := poolmodel.Pool{
			PoolID:     aws.Int64(i),
			PoolValues: []float64{1, 2, 4},
		}
		requestBody, err := json.Marshal(pool)

		if err != nil {
			panic(err)
		}

		// Create a new test request
		reqBody := bytes.NewReader(requestBody)

		// Create a new test request
		req, err := http.NewRequest("POST", "/api/handler", reqBody)
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

		//check inserted status
		type data struct {
			Status string `json:"status"`
		}

		var response struct {
			Data data `json:"data"`
		}

		err = json.Unmarshal(res.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("Failed to unmarshal response body: %v", err)
		}

		// Check the "status" field value
		if response.Data.Status != "inserted" {
			t.Errorf("Request %d: Expected 'status' to be 'inserted', got '%s'", i+1, response.Data.Status)
		}
		t.Cleanup(func() {
			common.FakeDB.Clean()
		})

	}
}
