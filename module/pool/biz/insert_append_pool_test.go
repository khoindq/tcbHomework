package poolbiz

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestInsertAppendPool_Insert(t *testing.T) {
	poolStore := new(MockPoolStore)
	poolBiz := NewInsertAppendPoolBiz(poolStore)

	pool := &poolmodel.Pool{
		PoolID:     aws.Int64(1),
		PoolValues: []float64{1, 2, 3},

		// Set other pool fields as needed
	}

	poolStore.On("FindPool", mock.Anything, pool.PoolID).Return(nil, false)
	poolStore.On("InsertPool", mock.Anything, mock.AnythingOfType("*poolmodel.PoolInsert")).Return(nil)

	poolStatus, err := poolBiz.InsertAppendPool(context.Background(), pool)
	assert.NoError(t, err)
	assert.Equal(t, poolmodel.PoolStatusInserted, *poolStatus)

	poolStore.AssertExpectations(t)
}

func TestInsertAppendPool_Append(t *testing.T) {
	poolStore := new(MockPoolStore)
	poolBiz := NewInsertAppendPoolBiz(poolStore)

	pool := &poolmodel.Pool{
		PoolID:     aws.Int64(1),
		PoolValues: []float64{1, 2, 3},
		// Set other pool fields as needed
	}

	existingPool := &poolmodel.Pool{
		PoolID: aws.Int64(1),
		// Set other existing pool fields as needed
	}

	poolStore.On("FindPool", mock.Anything, pool.PoolID).Return(existingPool, true)
	poolStore.On("AppendPool", mock.Anything, mock.AnythingOfType("*poolmodel.PoolAppend")).Return(nil)

	poolStatus, err := poolBiz.InsertAppendPool(context.Background(), pool)
	assert.NoError(t, err)
	assert.Equal(t, poolmodel.PoolStatusAppended, *poolStatus)

	poolStore.AssertExpectations(t)
}
