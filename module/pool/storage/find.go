package poolstorage

import (
	"context"

	"github.com/khoindq/tcbHomework/common"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
)

func (s *Store) FindPool(ctx context.Context, id int64) (result *poolmodel.Pool, found bool) {
	// FindPool is a method of the Store type.
	// It searches for a pool with the specified ID in the store.

	// Read lock the FakeDB to allow concurrent access to the data.
	common.FakeDB.RLock()
	defer common.FakeDB.RUnlock()

	// Retrieve the pool values associated with the given ID from FakeDB.
	value, ok := common.FakeDB.Pools[id]

	// If the pool is not found, return nil and false.
	if !ok {
		return nil, false
	}

	// If the pool is found, create a new poolmodel.Pool instance with the ID and values.
	return &poolmodel.Pool{
		PoolID:     id,
		PoolValues: value,
	}, true
}
