package poolstorage

import (
	"context"

	"github.com/khoindq/tcbHomework/common"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
)

func (s *Store) InsertPool(ctx context.Context, data *poolmodel.PoolInsert) error {
	// InsertPool is a method of the Store type.
	// It inserts a new pool into the store.

	// Lock the FakeDB to ensure exclusive access to the data.
	common.FakeDB.Lock()
	defer common.FakeDB.Unlock()

	// Assign the pool values to the specified pool ID in FakeDB.
	common.FakeDB.Pools[*data.PoolID] = data.PoolValues

	return nil
}
