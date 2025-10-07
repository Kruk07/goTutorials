package db

import (
	"example.com/go_basics/go/entity"
	"github.com/google/uuid"
)

type MemoryDB struct {
	Movies      map[uuid.UUID]entity.Movie
	Characters  map[uuid.UUID]entity.Character
	Appearances []entity.Appearance
}

func New() *MemoryDB {
	return &MemoryDB{
		Movies:      make(map[uuid.UUID]entity.Movie, 3),
		Characters:  make(map[uuid.UUID]entity.Character, 3),
		Appearances: []entity.Appearance{},
	}
}
