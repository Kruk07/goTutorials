package handlers

import (
	"net/http"

	"example.com/go_basics/go/api"
	"example.com/go_basics/go/repository"
	"example.com/go_basics/go/swapi"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Repo      *repository.Repository
	Validator *validator.Validate
	SWAPI     *swapi.Client
}

func New(repo *repository.Repository) *Handlers {
	return &Handlers{
		Repo:      repo,
		Validator: validator.New(),
		SWAPI:     swapi.New(),
	}
}

func (h *Handlers) PostMovies(c echo.Context) error {
	var input api.Movie
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := h.Validator.Struct(input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Validation failed", "details": err.Error()})
	}
	movie, err := h.Repo.CreateMovie(input.Title, input.ReleaseYear)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, movie)
}

func (h *Handlers) PostCharacters(c echo.Context) error {
	var input api.Character
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := h.Validator.Struct(input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Validation failed", "details": err.Error()})
	}

	if input.Movie != nil && *input.Movie == "Star Wars" {
		exists, err := h.SWAPI.CharacterExists(input.Name)
		if err != nil {
			return c.JSON(http.StatusBadGateway, echo.Map{"error": "SWAPI lookup failed", "details": err.Error()})
		}
		if !exists {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Character not found in Star Wars universe"})
		}
	}

	char, err := h.Repo.CreateCharacter(input.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, char)
}

func (h *Handlers) PostAppearances(c echo.Context) error {
	var input api.Appearance
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := h.Validator.Struct(input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Validation failed", "details": err.Error()})
	}
	movieID, err := uuid.Parse(input.MovieId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid movie_id"})
	}
	charID, err := uuid.Parse(input.CharacterId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid character_id"})
	}
	if err := h.Repo.AddAppearance(movieID, charID); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handlers) GetMovies(c echo.Context) error {
	movies, err := h.Repo.ListAllMovies()
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, movies)
}

func (h *Handlers) GetCharacters(c echo.Context) error {
	chars, err := h.Repo.ListAllCharacters()
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, chars)
}

func (h *Handlers) GetCharactersByMovie(c echo.Context, params api.GetCharactersByMovieParams) error {
	if params.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing title"})
	}
	chars, err := h.Repo.GetCharactersByMovieTitle(params.Title)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, chars)
}

func (h *Handlers) GetMoviesByCharacter(c echo.Context, params api.GetMoviesByCharacterParams) error {
	if params.Name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing name"})
	}
	titles, err := h.Repo.GetMovieTitlesByCharacterName(params.Name)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, titles)
}

func (h *Handlers) PutCharacters(c echo.Context) error {
	var input api.Character
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := h.Validator.Struct(input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Validation failed", "details": err.Error()})
	}
	idStr := c.QueryParam("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing id"})
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid UUID format"})
	}
	if err := h.Repo.UpdateCharacter(id, input.Name); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handlers) DeleteMovies(c echo.Context, params api.DeleteMoviesParams) error {
	id, err := uuid.Parse(params.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid UUID format"})
	}
	if err := h.Repo.DeleteMovie(id); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handlers) DeleteCharactersId(c echo.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid UUID format"})
	}
	if err := h.Repo.DeleteCharacter(uid); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
