package poolbiz

import (
	"context"

	"github.com/khoindq/tcbHomework/common"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
)

type PoolStore interface {
	InsertPool(ctx context.Context, data *poolmodel.PoolInsert) error
	AppendPool(ctx context.Context, data *poolmodel.PoolAppend) error
	FindPool(ctx context.Context, id int64) (result *poolmodel.Pool, found bool)
}

type insertAppendBiz struct {
	poolStore PoolStore
}

func NewInsertAppendPoolBiz(poolStore PoolStore) *insertAppendBiz {
	// NewInsertAppendPoolBiz is a constructor function for the insertAppendBiz type.
	// It creates and returns a new instance of insertAppendBiz with the provided poolStore.
	return &insertAppendBiz{poolStore: poolStore}
}

func (biz *insertAppendBiz) InsertAppendPool(ctx context.Context, data *poolmodel.Pool) (*poolmodel.PoolStatus, error) {
	// InsertAppendPool is a method of the insertAppendBiz type.
	// It inserts or appends a pool based on the given data.

	// First, check if the pool with the specified ID exists in the poolStore.
	_, found := biz.poolStore.FindPool(ctx, data.PoolID)

	// If the pool is not found, insert the pool.
	if !found {
		poolInsertData := poolmodel.PoolInsert{
			Pool: *data,
		}
		if err := poolInsertData.Validate(); err != nil {
			return nil, err
		}

		if err := biz.poolStore.InsertPool(ctx, &poolInsertData); err != nil {
			return nil, err
		}
		return poolmodel.PoolStatusInserted.ToPointer(), nil
	} else { // If the pool is found, append the pool
		poolAppendData := poolmodel.PoolAppend{
			Pool: *data,
		}
		if err := poolAppendData.Validate(); err != nil {
			return nil, err
		}

		if err := biz.poolStore.AppendPool(ctx, &poolAppendData); err != nil {
			return nil, common.ErrCannotUpdateEntity(poolmodel.EntityName, err)
		}
		return poolmodel.PoolStatusAppended.ToPointer(), nil
	}
}
