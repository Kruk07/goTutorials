package routes

import (
	"net/http"

	"example.com/go_basics/go/handlers"
)

func NewServeMux(h *handlers.Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateMovie(w, r)
		case http.MethodGet:
			h.ListAllMovies(w, r)
		case http.MethodDelete:
			h.DeleteMovie(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/characters", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateCharacter(w, r)
		case http.MethodGet:
			h.ListAllCharacters(w, r)
		case http.MethodPut:
			h.UpdateCharacter(w, r)
		case http.MethodDelete:
			h.DeleteCharacter(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/appearances", h.AddAppearance)
	mux.HandleFunc("/characters/by-movie", h.GetCharactersByMovieTitle)
	mux.HandleFunc("/movies/by-character", h.GetMovieTitlesByCharacterName)

	return mux
}
