package poolstorage

import (
	"context"

	"github.com/khoindq/tcbHomework/common"
	poolmodel "github.com/khoindq/tcbHomework/module/pool/model"
)

func (s *Store) AppendPool(ctx context.Context, data *poolmodel.PoolAppend) error {
	// AppendPool is a method of the Store type.
	// It appends pool values to an existing pool in the store.

	// Lock the FakeDB to ensure exclusive access to the data.
	common.FakeDB.Lock()
	defer common.FakeDB.Unlock()

	// Append the pool values to the existing pool in FakeDB.
	common.FakeDB.Pools[data.PoolID] = append(common.FakeDB.Pools[data.PoolID], data.PoolValues...)

	return nil
}
