package poolmodel

import (
	"errors"

	"github.com/khoindq/tcbHomework/common"
)

const EntityName = "Pool"

var FakePoolDB map[int64]Pool

// Pool object
type Pool struct {
	PoolID     int64     `json:"poolId" binding:"required"`
	PoolValues []float64 `json:"poolValues" binding:"required"`
}

type PoolInsert struct {
	Pool
}

func (pool *PoolInsert) Validate() error {
	if len(pool.PoolValues) == 0 {
		return ErrPoolCreateValidateFailed(errors.New("pool.poolValues is empty"))
	}
	return nil
}

type PoolAppend struct {
	Pool
}

func (pool *PoolAppend) Validate() error {
	//nothing to validate here
	return nil
}

// PoolQuantileGet object
type PoolQuantileGet struct {
	PoolID     int64   `json:"poolId" binding:"required"`
	Percentile float64 `json:"percentile" binding:"required"`
}
type PoolQuantileResp struct {
	TotalCount         int64   `json:"totalCount"`
	CalculatedQuantile float64 `json:"calculatedQuantile"`
}

func (pool *PoolQuantileGet) Validate() error {
	if pool.Percentile < 0 || pool.Percentile > 100 {
		return ErrPoolPercentileGetValidateFailed(errors.New("pool.Percentile must be between 0 and 100"))
	}
	return nil
}

type PoolStatus string

const PoolStatusInserted = PoolStatus("inserted")
const PoolStatusAppended = PoolStatus("appended")

func (p PoolStatus) ToPointer() *PoolStatus {
	return &p
}

func ErrPoolPercentileGetValidateFailed(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot validate PoolPercentileGet",
		"ErrPoolPercentileGetValidateFailed",
	)
}

func ErrPoolCreateValidateFailed(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot validate PoolCreate",
		"ErrPoolCreatelidateFailed",
	)
}

func ErrPoolNotFound(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Pool Not Found",
		"ErrPoolNotFound",
	)
}