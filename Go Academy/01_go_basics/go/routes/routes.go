package routes

import (
	"example.com/go_basics/go/handlers"
	"github.com/labstack/echo/v4"
)

func NewEchoRouter(h *handlers.Handlers) *echo.Echo {
	e := echo.New()

	// Movies
	e.GET("/movies", h.ListAllMovies)
	e.POST("/movies", h.CreateMovie)
	e.DELETE("/movies", h.DeleteMovie)

	// Characters
	e.GET("/characters", h.ListAllCharacters)
	e.POST("/characters", h.CreateCharacter)
	e.PUT("/characters", h.UpdateCharacter)
	e.DELETE("/characters", h.DeleteCharacter)

	// Appearances
	e.POST("/appearances", h.AddAppearance)

	// Queries
	e.GET("/characters/by-movie", h.GetCharactersByMovieTitle)
	e.GET("/movies/by-character", h.GetMovieTitlesByCharacterName)

	return e
}
