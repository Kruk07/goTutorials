package entity

import "github.com/google/uuid"

type Movie struct {
	ID    uuid.UUID
	Title string
	Year  int
}

func NewMovie(title string, year int) Movie {
	return Movie{
		ID:    uuid.New(),
		Title: title,
		Year:  year,
	}
}
