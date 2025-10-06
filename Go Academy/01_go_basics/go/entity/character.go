package entity

import "github.com/google/uuid"

type Character struct {
	ID   uuid.UUID
	Name string
}

func NewCharacter(name string) Character {
	return Character{
		ID:   uuid.New(),
		Name: name,
	}
}
