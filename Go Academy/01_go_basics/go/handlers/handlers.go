package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/go_basics/go/repository"
	"github.com/google/uuid"
)

type Handlers struct {
	Repo *repository.Repository
}

func New(repo *repository.Repository) *Handlers {
	return &Handlers{Repo: repo}
}

func (h *Handlers) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Year  int    `json:"year"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	movie := h.Repo.CreateMovie(input.Title, input.Year)
	json.NewEncoder(w).Encode(movie)
}

func (h *Handlers) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	char := h.Repo.CreateCharacter(input.Name)
	json.NewEncoder(w).Encode(char)
}

func (h *Handlers) AddAppearance(w http.ResponseWriter, r *http.Request) {
	var input struct {
		MovieID     uuid.UUID `json:"movie_id"`
		CharacterID uuid.UUID `json:"character_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	h.Repo.AddAppearance(input.MovieID, input.CharacterID)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) ListAllMovies(w http.ResponseWriter, r *http.Request) {
	movies := h.Repo.ListAllMovies()
	json.NewEncoder(w).Encode(movies)
}

func (h *Handlers) ListAllCharacters(w http.ResponseWriter, r *http.Request) {
	chars := h.Repo.ListAllCharacters()
	json.NewEncoder(w).Encode(chars)
}

func (h *Handlers) GetCharactersByMovieTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Missing title", http.StatusBadRequest)
		return
	}
	chars := h.Repo.GetCharactersByMovieTitle(title)
	json.NewEncoder(w).Encode(chars)
}

func (h *Handlers) GetMovieTitlesByCharacterName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name", http.StatusBadRequest)
		return
	}
	titles := h.Repo.GetMovieTitlesByCharacterName(name)
	json.NewEncoder(w).Encode(titles)
}

func (h *Handlers) UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	h.Repo.UpdateCharacter(input.ID, input.Name)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	h.Repo.DeleteMovie(id)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	h.Repo.DeleteCharacter(id)
	w.WriteHeader(http.StatusNoContent)
}
