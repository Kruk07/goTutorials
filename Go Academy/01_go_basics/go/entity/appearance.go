package entity

import "github.com/google/uuid"

type Appearance struct {
	MovieID     uuid.UUID
	CharacterID uuid.UUID
}

func New(movieID, characterID uuid.UUID) Appearance {
	return Appearance{
		MovieID:     movieID,
		CharacterID: characterID,
	}
}