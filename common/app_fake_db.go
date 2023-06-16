package common

import "sync"

type DB struct {
	sync.RWMutex
	Pools map[int64][]float64
}

var FakeDB = DB{
	Pools: make(map[int64][]float64),
}
