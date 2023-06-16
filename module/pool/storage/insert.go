package poolstorage

import (
	"context"

	"github.com/khoindq/tcbHomework/common"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
)

func (s *Store) InsertPool(ctx context.Context, data *poolmodel.PoolInsert) error {
	common.FakeDB.Lock()
	defer common.FakeDB.Unlock()
	common.FakeDB.Pools[data.PoolID] = data.PoolValues
	return nil
}
