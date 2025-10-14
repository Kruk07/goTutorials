package entity

import "github.com/google/uuid"

type Character struct {
	ID   uuid.UUID
	Name string `json:"name" validate:"required"`
}

func NewCharacter(options ...func(*Character)) Character {
	char := Character{
		ID: uuid.New(),
	}
	for _, o := range options {
		o(&char)
	}
	return char
}

func WithName(name string) func(*Character) {
	return func(c *Character) {
		c.Name = name
	}
}
