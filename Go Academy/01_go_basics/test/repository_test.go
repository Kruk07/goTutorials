package test

import (
	"testing"

	"example.com/go_basics/go/db"
	"example.com/go_basics/go/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateMovieAndCharacter(t *testing.T) {
	mem := db.New()
	repo := repository.New(mem)

	movie := repo.CreateMovie("Shrek", 2001)
	character := repo.CreateCharacter("Shrek")

	assert.Equal(t, "Shrek", movie.Title)
	assert.Equal(t, 2001, movie.Year)
	assert.Equal(t, "Shrek", character.Name)
	assert.NotEmpty(t, movie.ID)
	assert.NotEmpty(t, character.ID)
}

func TestAddAppearanceAndGetCharactersByMovie(t *testing.T) {
	mem := db.New()
	repo := repository.New(mem)

	movie := repo.CreateMovie("Shrek 2", 2004)
	character := repo.CreateCharacter("Donkey")

	repo.AddAppearance(movie.ID, character.ID)

	chars := repo.GetCharactersByMovie(movie.ID)
	assert.Len(t, chars, 1)
	assert.Equal(t, "Donkey", chars[0].Name)
}

func TestGetMoviesByCharacter(t *testing.T) {
	mem := db.New()
	repo := repository.New(mem)

	m1 := repo.CreateMovie("Shrek", 2001)
	m2 := repo.CreateMovie("Shrek 2", 2004)
	c := repo.CreateCharacter("Fiona")

	repo.AddAppearance(m1.ID, c.ID)
	repo.AddAppearance(m2.ID, c.ID)

	movies := repo.GetMoviesByCharacter(c.ID)
	assert.Len(t, movies, 2)
	assert.Contains(t, []string{movies[0].Title, movies[1].Title}, "Shrek")
	assert.Contains(t, []string{movies[0].Title, movies[1].Title}, "Shrek 2")
}

func TestGetCharactersByMovieTitle(t *testing.T) {
	mem := db.New()
	repo := repository.New(mem)

	m := repo.CreateMovie("The Lion King", 1994)
	c := repo.CreateCharacter("Simba")

	repo.AddAppearance(m.ID, c.ID)

	chars := repo.GetCharactersByMovieTitle("The Lion King")
	assert.Len(t, chars, 1)
	assert.Equal(t, "Simba", chars[0].Name)

	none := repo.GetCharactersByMovieTitle("Unknown")
	assert.Len(t, none, 0)
}

func TestGetMovieTitlesByCharacterName(t *testing.T) {
	mem := db.New()
	repo := repository.New(mem)

	m1 := repo.CreateMovie("Shrek", 2001)
	m2 := repo.CreateMovie("Shrek 2", 2004)
	c := repo.CreateCharacter("Puss in Boots")

	repo.AddAppearance(m1.ID, c.ID)
	repo.AddAppearance(m2.ID, c.ID)

	titles := repo.GetMovieTitlesByCharacterName("Puss in Boots")
	assert.Len(t, titles, 2)
	assert.Contains(t, titles, "Shrek")
	assert.Contains(t, titles, "Shrek 2")

	none := repo.GetMovieTitlesByCharacterName("Scar")
	assert.Len(t, none, 0)
}

func TestListAllMoviesAndCharacters(t *testing.T) {
	mem := db.New()
	repo := repository.New(mem)

	repo.CreateMovie("Shrek", 2001)
	repo.CreateMovie("Shrek 2", 2004)
	repo.CreateMovie("The Lion King", 1994)

	repo.CreateCharacter("Shrek")
	repo.CreateCharacter("Donkey")
	repo.CreateCharacter("Simba")

	assert.Len(t, repo.DB.Movies, 3)
	assert.Len(t, repo.DB.Characters, 3)

	repo.ListAllMovies()
	repo.ListAllCharacters()
}

func TestUpdateCharacter(t *testing.T) {
	mem := db.New()
	repo := repository.New(mem)

	char := repo.CreateCharacter("Donkey")
	updated := repo.UpdateCharacter(char.ID, "Donkey the Brave")

	assert.True(t, updated)
	assert.Equal(t, "Donkey the Brave", repo.DB.Characters[char.ID].Name)

	fakeID := uuid.New()
	notUpdated := repo.UpdateCharacter(fakeID, "Ghost")
	assert.False(t, notUpdated)
}

func TestDeleteMovie(t *testing.T) {
	mem := db.New()
	repo := repository.New(mem)

	movie := repo.CreateMovie("Shrek Forever After", 2010)
	char := repo.CreateCharacter("Rumpelstiltskin")
	repo.AddAppearance(movie.ID, char.ID)

	deleted := repo.DeleteMovie(movie.ID)
	assert.True(t, deleted)
	assert.NotContains(t, repo.DB.Movies, movie.ID)

	for _, a := range repo.DB.Appearances {
		assert.NotEqual(t, a.MovieID, movie.ID)
	}

	fakeID := uuid.New()
	notDeleted := repo.DeleteMovie(fakeID)
	assert.False(t, notDeleted)
}

func TestDeleteCharacter(t *testing.T) {
	mem := db.New()
	repo := repository.New(mem)

	movie := repo.CreateMovie("The Lion King", 1994)
	char := repo.CreateCharacter("Scar")
	repo.AddAppearance(movie.ID, char.ID)

	deleted := repo.DeleteCharacter(char.ID)
	assert.True(t, deleted)
	assert.NotContains(t, repo.DB.Characters, char.ID)

	for _, a := range repo.DB.Appearances {
		assert.NotEqual(t, a.CharacterID, char.ID)
	}

	fakeID := uuid.New()
	notDeleted := repo.DeleteCharacter(fakeID)
	assert.False(t, notDeleted)
}
