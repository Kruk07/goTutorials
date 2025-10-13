package handlers

import (
	"net/http"

	"example.com/go_basics/go/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Repo *repository.Repository
}

func New(repo *repository.Repository) *Handlers {
	return &Handlers{Repo: repo}
}

func (h *Handlers) CreateMovie(c echo.Context) error {
	var input struct {
		Title string `json:"title"`
		Year  int    `json:"year"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	movie, err := h.Repo.CreateMovie(input.Title, input.Year)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, movie)
}

func (h *Handlers) CreateCharacter(c echo.Context) error {
	var input struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	char, err := h.Repo.CreateCharacter(input.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, char)
}

func (h *Handlers) AddAppearance(c echo.Context) error {
	var input struct {
		MovieID     uuid.UUID `json:"movie_id"`
		CharacterID uuid.UUID `json:"character_id"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := h.Repo.AddAppearance(input.MovieID, input.CharacterID); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handlers) ListAllMovies(c echo.Context) error {
	movies, err := h.Repo.ListAllMovies()
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, movies)
}

func (h *Handlers) ListAllCharacters(c echo.Context) error {
	chars, err := h.Repo.ListAllCharacters()
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, chars)
}

func (h *Handlers) GetCharactersByMovieTitle(c echo.Context) error {
	title := c.QueryParam("title")
	if title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing title"})
	}
	chars, err := h.Repo.GetCharactersByMovieTitle(title)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, chars)
}

func (h *Handlers) GetMovieTitlesByCharacterName(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing name"})
	}
	titles, err := h.Repo.GetMovieTitlesByCharacterName(name)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, titles)
}

func (h *Handlers) UpdateCharacter(c echo.Context) error {
	var input struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := h.Repo.UpdateCharacter(input.ID, input.Name); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handlers) DeleteMovie(c echo.Context) error {
	idStr := c.QueryParam("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing id"})
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid UUID format"})
	}
	if err := h.Repo.DeleteMovie(id); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handlers) DeleteCharacter(c echo.Context) error {
	idStr := c.QueryParam("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing id"})
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid UUID format"})
	}
	if err := h.Repo.DeleteCharacter(id); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
