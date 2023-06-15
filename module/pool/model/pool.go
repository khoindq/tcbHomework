package poolmodel

import (
	"errors"

	"github.com/khoindq/tcbHomework/common"
)

const EntityName = "Pool"

type Pool struct {
	PoolID     int64     `json:"poolId"`
	PoolValues []float64 `json:"poolValues"`
}

type PoolPercentileGet struct {
	Pool       int64   `json:"poolId"`
	Percentile float64 `json:"percentile"`
}

func (pool *PoolPercentileGet) Validate() error {
	if pool.Percentile < 0 || pool.Percentile > 100 {
		return ErrPoolPercentileGetValidateFailed(errors.New("pool.Percentile must be between 0 and 100"))
	}
	return nil
}

func ErrPoolPercentileGetValidateFailed(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot validate PoolPercentileGet",
		"ErrPoolPercentileGetValidateFailed",
	)
}
