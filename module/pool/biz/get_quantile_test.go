package poolbiz

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
	"github.com/stretchr/testify/assert"
)

func TestGetQuantile(t *testing.T) {
	ctx := context.Background()

	// Create a mock poolStore
	poolStore := NewMockPoolStore(t)

	// Create the getQuantileBiz instance
	biz := &getQuantileBiz{
		poolStore: poolStore,
	}

	t.Run("Returns quantile value when pool and request are valid", func(t *testing.T) {
		// Set up test data
		req := &poolmodel.PoolQuantileGet{
			PoolID:     aws.Int64(1),
			Percentile: 50,
		}
		poolValues := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

		// Mock the FindPool method to return pool found
		poolStore.On("FindPool", ctx, req.PoolID).Return(&poolmodel.Pool{
			PoolID:     req.PoolID,
			PoolValues: poolValues,
			// Set other necessary fields
		}, true)

		// Mock the GetQuantileInNearestRanks method to succeed
		expectedQuantile := 3.0 // Mocked quantile value

		// Call the method
		resp, err := biz.GetQuantile(ctx, req)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, int64(len(poolValues)), resp.TotalCount)
		assert.Equal(t, expectedQuantile, resp.CalculatedQuantile)

		// Verify the method calls
		poolStore.AssertCalled(t, "FindPool", ctx, req.PoolID)
	})

	t.Run("Returns an error when request validation fails", func(t *testing.T) {
		// Set up test data
		req := &poolmodel.PoolQuantileGet{
			PoolID:     aws.Int64(1),
			Percentile: 150, // Invalid percentile value
		}

		// Call the method
		resp, err := biz.GetQuantile(ctx, req)

		// Assertions
		assert.EqualError(t, err, "pool.Percentile must be between 0 and 100")
		assert.Nil(t, resp)
	})

	t.Run("Returns an error when pool is not found", func(t *testing.T) {
		// Set up test data
		req := &poolmodel.PoolQuantileGet{
			PoolID:     aws.Int64(6),
			Percentile: 0.7,
		}

		// Mock the FindPool method to return pool not found
		poolStore.On("FindPool", ctx, req.PoolID).Return(nil, false)

		// Call the method
		resp, err := biz.GetQuantile(ctx, req)

		// Assertions
		assert.EqualError(t, err, "pool not found")
		assert.Nil(t, resp)

		// Verify the method calls
		poolStore.AssertCalled(t, "FindPool", ctx, req.PoolID)
	})
}
