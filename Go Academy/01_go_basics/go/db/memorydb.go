package db

import (
	"sync"

	"example.com/go_basics/go/entity"
)

type MemoryDB struct {
	Movies      sync.Map
	Characters  sync.Map
	Appearances []entity.Appearance
	Mutex       sync.Mutex
}

func New() *MemoryDB {
	return &MemoryDB{
		Appearances: make([]entity.Appearance, 0),
	}
}
