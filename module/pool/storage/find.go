package poolstorage

import (
	"context"

	"github.com/khoindq/tcbHomework/common"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
)

func (s *Store) FindPool(ctx context.Context, id int64) (result *poolmodel.Pool, found bool) {
	common.FakeDB.RLock()
	defer common.FakeDB.RUnlock()
	value, ok := common.FakeDB.Pools[id]
	if !ok {
		return nil, false // not found
	}

	return &poolmodel.Pool{
		PoolID:     id,
		PoolValues: value,
	}, true // found
}
