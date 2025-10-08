package routes

import (
	"net/http"

	"example.com/go_basics/go/handlers"
)

func NewServeMux(h *handlers.Handlers) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/movies", h.CreateMovie)
	mux.HandleFunc("/characters", h.CreateCharacter)
	mux.HandleFunc("/appearance", h.AddAppearance)
	mux.HandleFunc("/movies/list", h.ListAllMovies)
	mux.HandleFunc("/characters/list", h.ListAllCharacters)
	mux.HandleFunc("/characters/by-movie", h.GetCharactersByMovieTitle)
	mux.HandleFunc("/movies/by-character", h.GetMovieTitlesByCharacterName)
	mux.HandleFunc("/characters/update", h.UpdateCharacter)
	mux.HandleFunc("/movies/delete", h.DeleteMovie)
	mux.HandleFunc("/characters/delete", h.DeleteCharacter)
	return mux
}
