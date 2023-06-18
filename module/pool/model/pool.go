package poolmodel

import (
	"errors"
)

const EntityName = "Pool"

var FakePoolDB map[int64]Pool

// Pool represents a pool with an ID and a collection of values.
type Pool struct {
	PoolID     *int64    `json:"poolId" binding:"required"`     // ID of the pool
	PoolValues []float64 `json:"poolValues" binding:"required"` // ID of the pool
}

// PoolInsert is used for inserting a pool.
type PoolInsert struct {
	Pool // Embedding the Pool struct for reuse
}

// Validate validates the PoolInsert struct.
func (pool *PoolInsert) Validate() error {
	if len(pool.PoolValues) == 0 {
		return errors.New("pool.poolValues is empty")
	}
	return nil
}

// PoolAppend is used for appending to a pool.
type PoolAppend struct {
	Pool // Embedding the Pool struct for reuse
}

// Validate validates the PoolAppend struct.
func (pool *PoolAppend) Validate() error {
	//nothing to validate here
	return nil
}

// PoolQuantileGet represents a request to get the quantile of a pool.
type PoolQuantileGet struct {
	PoolID     *int64  `json:"poolId" binding:"required"`     // ID of the pool
	Percentile float64 `json:"percentile" binding:"required"` // Desired percentile
}

// PoolQuantileResp represents the response containing the quantile information.
type PoolQuantileResp struct {
	TotalCount         int64   `json:"totalCount"`         // Total count of values in the pool
	CalculatedQuantile float64 `json:"calculatedQuantile"` // Calculated quantile value
}

// Validate validates the PoolQuantileGet struct.
func (pool *PoolQuantileGet) Validate() error {
	if pool.Percentile < 0 || pool.Percentile > 100 {
		return errors.New("pool.Percentile must be between 0 and 100")
	}
	return nil
}

// PoolStatus represents the status of inserted or appended pool
type PoolStatus string

// PoolStatus constants for different pool statuses.
const PoolStatusInserted = PoolStatus("inserted")
const PoolStatusAppended = PoolStatus("appended")

// ToPointer converts a PoolStatus value to a pointer.
func (p PoolStatus) ToPointer() *PoolStatus {
	return &p
}
