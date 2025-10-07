package entity

import "github.com/google/uuid"

type Movie struct {
	ID    uuid.UUID
	Title string
	Year  int
}

// func NewMovie(title string, year int) Movie {
// 	return Movie{
// 		ID:    uuid.New(),
// 		Title: title,
// 		Year:  year,
// 	}
// }

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
