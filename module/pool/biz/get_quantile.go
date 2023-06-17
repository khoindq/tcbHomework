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
	return &getQuantileBiz{poolStore: poolStore}
}

func (biz *getQuantileBiz) GetQuantile(ctx context.Context, req *poolmodel.PoolQuantileGet) (*poolmodel.PoolQuantileResp, error) {
	//find the pool with the id first
	if err := req.Validate(); err != nil {
		return nil, err
	}
	foundPool, found := biz.poolStore.FindPool(ctx, req.PoolID)
	if !found {
		return nil, common.ErrEntityNotFound(poolmodel.EntityName, nil)
	}
	calculatedQuantile, err := utils.GetQuantileInNearestRanks(req.Percentile, foundPool.PoolValues)
	if err != nil {
		return nil, err
	}
	return &poolmodel.PoolQuantileResp{
		TotalCount:         int64(len(foundPool.PoolValues)),
		CalculatedQuantile: *calculatedQuantile,
	}, nil

}
