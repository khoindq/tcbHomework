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
	return &insertAppendBiz{poolStore: poolStore}
}

func (biz *insertAppendBiz) InsertAppendPool(ctx context.Context, data *poolmodel.Pool) (*poolmodel.PoolStatus, error) {
	//find the pool with the id first
	_, found := biz.poolStore.FindPool(ctx, data.PoolID)

	//not found case, we will insert to the pooll
	if !found {
		poolInsertData := poolmodel.PoolInsert{
			Pool: *data,
		}
		if err := poolInsertData.Validate(); err != nil {
			return nil, err
		}

		if err := biz.poolStore.InsertPool(ctx, &poolInsertData); err != nil {
			return nil, common.ErrCannotCreateEntity(poolmodel.EntityName, err)
		}
		return poolmodel.PoolStatusInserted.ToPointer(), nil
	} else { // found case, we will append
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
