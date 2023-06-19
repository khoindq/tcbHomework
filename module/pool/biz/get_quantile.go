package poolbiz

import (
	"context"

	"github.com/khoindq/tcbHomework/common"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
	"github.com/khoindq/tcbHomework/utils"
)

type getQuantileBiz struct {
	poolStore PoolStore
}

func NewGetQuantileBiz(poolStore PoolStore) *getQuantileBiz {
	// NewGetQuantileBiz is a constructor function for the getQuantileBiz type.
	// It creates and returns a new instance of getQuantileBiz with the provided poolStore.
	return &getQuantileBiz{poolStore: poolStore}
}

func (biz *getQuantileBiz) GetQuantile(ctx context.Context, req *poolmodel.PoolQuantileGet) (*poolmodel.PoolQuantileResp, error) {
	// GetQuantile is a method of the getQuantileBiz type.
	// It retrieves the quantile value from a pool based on the given request.

	// First, validate the request object.
	if err := req.Validate(); err != nil {
		return nil, common.ErrValidateFailed(err)
	}

	// Find the pool with the specified ID using the poolStore.
	// If the pool is not found, return an error indicating the entity was not found.
	foundPool, found := biz.poolStore.FindPool(ctx, req.PoolID)
	if !found {
		return nil, common.ErrEntityNotFound(poolmodel.EntityName, nil)
	}

	// Calculate the quantile value using the utils.GetQuantileInNearestRanks function.
	calculatedQuantile, err := utils.GetQuantileInNearestRanks(*req.Percentile, foundPool.PoolValues)
	if err != nil {
		return nil, err
	}

	// Return the pool quantile response object with the total count of pool values
	// and the calculated quantile value.
	return &poolmodel.PoolQuantileResp{
		TotalCount:         int64(len(foundPool.PoolValues)),
		CalculatedQuantile: *calculatedQuantile,
	}, nil

}
