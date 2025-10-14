package entity

import "github.com/google/uuid"

type Movie struct {
	ID    uuid.UUID
	Title string `json:"title" validate:"required"`
	Year  int `json:"year" validate:"required,min=1900"`
}

func NewMovie(options ...func(*Movie)) Movie {
	mov := Movie{
		ID: uuid.New(),
	}
	for _, o := range options {
		o(&mov)
	}
	return mov
}

func WithTitle(title string) func(*Movie) {
	return func(m *Movie) {
		m.Title = title
	}
}

func WithYear(year int) func(*Movie) {
	return func(m *Movie) {
		m.Year = year
	}
}
