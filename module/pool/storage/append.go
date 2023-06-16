package poolstorage

import (
	"context"

	"github.com/khoindq/tcbHomework/common"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
)

func (s *Store) AppendPool(ctx context.Context, data *poolmodel.PoolAppend) error {
	common.FakeDB.Lock()
	defer common.FakeDB.Unlock()

	common.FakeDB.Pools[data.PoolID] = append(common.FakeDB.Pools[data.PoolID], data.PoolValues...)

	return nil
}
