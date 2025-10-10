package repository

import (
	"testing"

	"example.com/go_basics/go/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateMovieAndCharacter(t *testing.T) {
	mem := db.New()
	repo := New(mem)

	movie, err := repo.CreateMovie("Shrek", 2001)
	assert.NoError(t, err)
	character, err := repo.CreateCharacter("Shrek")
	assert.NoError(t, err)

	assert.Equal(t, "Shrek", movie.Title)
	assert.Equal(t, 2001, movie.Year)
	assert.Equal(t, "Shrek", character.Name)
	assert.NotEmpty(t, movie.ID)
	assert.NotEmpty(t, character.ID)

	_, err = repo.CreateMovie("", 2001)
	assert.Error(t, err)

	_, err = repo.CreateCharacter("")
	assert.Error(t, err)
}

func TestAddAppearanceAndGetCharactersByMovie(t *testing.T) {
	mem := db.New()
	repo := New(mem)

	movie, _ := repo.CreateMovie("Shrek 2", 2004)
	character, _ := repo.CreateCharacter("Donkey")

	err := repo.AddAppearance(movie.ID, character.ID)
	assert.NoError(t, err)

	chars, err := repo.GetCharactersByMovie(movie.ID)
	assert.NoError(t, err)
	assert.Len(t, chars, 1)
	assert.Equal(t, "Donkey", chars[0].Name)

	fakeID := uuid.New()
	err = repo.AddAppearance(fakeID, character.ID)
	assert.Error(t, err)

	_, err = repo.GetCharactersByMovie(fakeID)
	assert.Error(t, err)
}

func TestGetMoviesByCharacter(t *testing.T) {
	mem := db.New()
	repo := New(mem)

	m1, _ := repo.CreateMovie("Shrek", 2001)
	m2, _ := repo.CreateMovie("Shrek 2", 2004)
	c, _ := repo.CreateCharacter("Fiona")

	repo.AddAppearance(m1.ID, c.ID)
	repo.AddAppearance(m2.ID, c.ID)

	movies, err := repo.GetMoviesByCharacter(c.ID)
	assert.NoError(t, err)
	assert.Len(t, movies, 2)
	assert.Contains(t, []string{movies[0].Title, movies[1].Title}, "Shrek")
	assert.Contains(t, []string{movies[0].Title, movies[1].Title}, "Shrek 2")

	_, err = repo.GetMoviesByCharacter(uuid.New())
	assert.Error(t, err)
}

func TestGetCharactersByMovieTitle(t *testing.T) {
	mem := db.New()
	repo := New(mem)

	m, _ := repo.CreateMovie("The Lion King", 1994)
	c, _ := repo.CreateCharacter("Simba")
	repo.AddAppearance(m.ID, c.ID)

	chars, err := repo.GetCharactersByMovieTitle("The Lion King")
	assert.NoError(t, err)
	assert.Len(t, chars, 1)
	assert.Equal(t, "Simba", chars[0].Name)

	_, err = repo.GetCharactersByMovieTitle("Unknown")
	assert.Error(t, err)
}

func TestGetMovieTitlesByCharacterName(t *testing.T) {
	mem := db.New()
	repo := New(mem)

	m1, _ := repo.CreateMovie("Shrek", 2001)
	m2, _ := repo.CreateMovie("Shrek 2", 2004)
	c, _ := repo.CreateCharacter("Puss in Boots")

	repo.AddAppearance(m1.ID, c.ID)
	repo.AddAppearance(m2.ID, c.ID)

	titles, err := repo.GetMovieTitlesByCharacterName("Puss in Boots")
	assert.NoError(t, err)
	assert.Len(t, titles, 2)
	assert.Contains(t, titles, "Shrek")
	assert.Contains(t, titles, "Shrek 2")

	_, err = repo.GetMovieTitlesByCharacterName("Scar")
	assert.Error(t, err)
}

func TestListAllMoviesAndCharacters(t *testing.T) {
	mem := db.New()
	repo := New(mem)

	repo.CreateMovie("Shrek", 2001)
	repo.CreateMovie("Shrek 2", 2004)
	repo.CreateMovie("The Lion King", 1994)

	repo.CreateCharacter("Shrek")
	repo.CreateCharacter("Donkey")
	repo.CreateCharacter("Simba")

	movies, err := repo.ListAllMovies()
	assert.NoError(t, err)
	assert.Len(t, movies, 3)

	chars, err := repo.ListAllCharacters()
	assert.NoError(t, err)
	assert.Len(t, chars, 3)

	memEmpty := db.New()
	repoEmpty := New(memEmpty)

	_, err = repoEmpty.ListAllMovies()
	assert.Error(t, err)

	_, err = repoEmpty.ListAllCharacters()
	assert.Error(t, err)
}

func TestUpdateCharacter(t *testing.T) {
	mem := db.New()
	repo := New(mem)

	char, _ := repo.CreateCharacter("Donkey")
	err := repo.UpdateCharacter(char.ID, "Donkey the Brave")
	assert.NoError(t, err)
	assert.Equal(t, "Donkey the Brave", repo.DB.Characters[char.ID].Name)

	err = repo.UpdateCharacter(uuid.New(), "Ghost")
	assert.Error(t, err)

	err = repo.UpdateCharacter(char.ID, "")
	assert.Error(t, err)
}

func TestDeleteMovie(t *testing.T) {
	mem := db.New()
	repo := New(mem)

	movie, _ := repo.CreateMovie("Shrek Forever After", 2010)
	char, _ := repo.CreateCharacter("Rumpelstiltskin")
	repo.AddAppearance(movie.ID, char.ID)

	err := repo.DeleteMovie(movie.ID)
	assert.NoError(t, err)
	assert.NotContains(t, repo.DB.Movies, movie.ID)

	for _, a := range repo.DB.Appearances {
		assert.NotEqual(t, a.MovieID, movie.ID)
	}

	err = repo.DeleteMovie(uuid.New())
	assert.Error(t, err)
}

func TestDeleteCharacter(t *testing.T) {
	mem := db.New()
	repo := New(mem)

	movie, _ := repo.CreateMovie("The Lion King", 1994)
	char, _ := repo.CreateCharacter("Scar")
	repo.AddAppearance(movie.ID, char.ID)

	err := repo.DeleteCharacter(char.ID)
	assert.NoError(t, err)
	assert.NotContains(t, repo.DB.Characters, char.ID)

	for _, a := range repo.DB.Appearances {
		assert.NotEqual(t, a.CharacterID, char.ID)
	}

	err = repo.DeleteCharacter(uuid.New())
	assert.Error(t, err)
}
