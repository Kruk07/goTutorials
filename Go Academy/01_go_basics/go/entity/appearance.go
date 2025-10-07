package entity

import "github.com/google/uuid"

type Appearance struct {
	MovieID     uuid.UUID
	CharacterID uuid.UUID
}

// func New(movieID, characterID uuid.UUID) Appearance {
// 	return Appearance{
// 		MovieID:     movieID,
// 		CharacterID: characterID,
// 	}
// }

func New(options ...func(*Appearance)) Appearance {
	app := Appearance{}
	for _, o := range options {
		o(&app)
	}
	return app
}

func WithMovieId(movieID uuid.UUID) func(*Appearance) {
	return func(s *Appearance) {
		s.MovieID = movieID
	}
}

func WithCharacterId(characterID uuid.UUID) func(*Appearance) {
	return func(s *Appearance) {
		s.CharacterID = characterID
	}
}
