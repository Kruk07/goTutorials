package repository

import (
	"errors"
	"fmt"
	"log"

	"example.com/go_basics/go/db"
	"example.com/go_basics/go/entity"
	"github.com/google/uuid"
)

type Repository struct {
	DB *db.MemoryDB
}

func New(db *db.MemoryDB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateMovie(title string, year int) (entity.Movie, error) {
	if title == "" {
		return entity.Movie{}, errors.New("movie title cannot be empty")
	}
	movie := entity.NewMovie(entity.WithTitle(title), entity.WithYear(year))
	r.DB.Movies[movie.ID] = movie
	log.Printf("Movie added: %s (%d) [ID: %s]", title, year, movie.ID)
	return movie, nil
}

func (r *Repository) CreateCharacter(name string) (entity.Character, error) {
	if name == "" {
		return entity.Character{}, errors.New("character name cannot be empty")
	}
	character := entity.NewCharacter(entity.WithName(name))
	r.DB.Characters[character.ID] = character
	log.Printf("Character added: %s [ID: %s]", name, character.ID)
	return character, nil
}

func (r *Repository) AddAppearance(movieID, characterID uuid.UUID) error {
	movie, ok := r.DB.Movies[movieID]
	if !ok {
		return fmt.Errorf("movie not found [ID: %s]", movieID)
	}
	character, ok := r.DB.Characters[characterID]
	if !ok {
		return fmt.Errorf("character not found [ID: %s]", characterID)
	}
	r.DB.Appearances = append(r.DB.Appearances, entity.New(
		entity.WithMovieId(movieID),
		entity.WithCharacterId(characterID),
	))
	log.Printf("Linked character '%s' to movie '%s'", character.Name, movie.Title)
	return nil
}

func (r *Repository) GetCharactersByMovie(movieID uuid.UUID) ([]entity.Character, error) {
	if _, ok := r.DB.Movies[movieID]; !ok {
		return nil, fmt.Errorf("movie not found [ID: %s]", movieID)
	}
	var result []entity.Character
	for _, a := range r.DB.Appearances {
		if a.MovieID == movieID {
			if c, ok := r.DB.Characters[a.CharacterID]; ok {
				result = append(result, c)
			}
		}
	}
	log.Printf("Retrieved %d characters for movie ID %s", len(result), movieID)
	return result, nil
}

func (r *Repository) GetMoviesByCharacter(characterID uuid.UUID) ([]entity.Movie, error) {
	if _, ok := r.DB.Characters[characterID]; !ok {
		return nil, fmt.Errorf("character not found [ID: %s]", characterID)
	}
	var result []entity.Movie
	for _, a := range r.DB.Appearances {
		if a.CharacterID == characterID {
			if m, ok := r.DB.Movies[a.MovieID]; ok {
				result = append(result, m)
			}
		}
	}
	log.Printf("Retrieved %d movies for character ID %s", len(result), characterID)
	return result, nil
}

func (r *Repository) GetCharactersByMovieTitle(title string) ([]entity.Character, error) {
	for _, m := range r.DB.Movies {
		if m.Title == title {
			log.Printf("Searching characters for movie title: %s", title)
			return r.GetCharactersByMovie(m.ID)
		}
	}
	log.Printf("No movie found with title: %s", title)
	return nil, fmt.Errorf("movie not found with title: %s", title)
}

func (r *Repository) GetMovieTitlesByCharacterName(name string) ([]string, error) {
	for _, c := range r.DB.Characters {
		if c.Name == name {
			log.Printf("Searching movies for character name: %s", name)
			movies, err := r.GetMoviesByCharacter(c.ID)
			if err != nil {
				return nil, err
			}
			var titles []string
			for _, m := range movies {
				titles = append(titles, m.Title)
			}
			log.Printf("Found %d movies for character '%s'", len(titles), name)
			return titles, nil
		}
	}
	log.Printf("No character found with name: %s", name)
	return nil, fmt.Errorf("character not found with name: %s", name)
}

func (r *Repository) ListAllMovies() (map[uuid.UUID]entity.Movie, error) {
	if len(r.DB.Movies) == 0 {
		log.Println("No movies found in the database.")
		return nil, errors.New("no movies available")
	}
	log.Println("All movies:")
	for _, m := range r.DB.Movies {
		log.Printf("- %s (%d) [ID: %s]", m.Title, m.Year, m.ID)
	}
	return r.DB.Movies, nil
}

func (r *Repository) ListAllCharacters() (map[uuid.UUID]entity.Character, error) {
	if len(r.DB.Characters) == 0 {
		log.Println("No characters found in the database.")
		return nil, errors.New("no characters available")
	}
	log.Println("All characters:")
	for _, c := range r.DB.Characters {
		log.Printf("- %s [ID: %s]", c.Name, c.ID)
	}
	return r.DB.Characters, nil
}

func (r *Repository) UpdateCharacter(id uuid.UUID, newName string) error {
	if newName == "" {
		return errors.New("new character name cannot be empty")
	}
	character, ok := r.DB.Characters[id]
	if !ok {
		return fmt.Errorf("character not found [ID: %s]", id)
	}
	character.Name = newName
	r.DB.Characters[id] = character
	log.Printf("Character updated: %s [ID: %s]", newName, id)
	return nil
}

func (r *Repository) DeleteMovie(id uuid.UUID) error {
	if _, ok := r.DB.Movies[id]; !ok {
		return fmt.Errorf("movie not found [ID: %s]", id)
	}
	delete(r.DB.Movies, id)
	var updated []entity.Appearance
	for _, a := range r.DB.Appearances {
		if a.MovieID != id {
			updated = append(updated, a)
		}
	}
	r.DB.Appearances = updated
	log.Printf("Movie deleted [ID: %s]", id)
	return nil
}

func (r *Repository) DeleteCharacter(id uuid.UUID) error {
	if _, ok := r.DB.Characters[id]; !ok {
		return fmt.Errorf("character not found [ID: %s]", id)
	}
	delete(r.DB.Characters, id)
	var updated []entity.Appearance
	for _, a := range r.DB.Appearances {
		if a.CharacterID != id {
			updated = append(updated, a)
		}
	}
	r.DB.Appearances = updated
	log.Printf("Character deleted [ID: %s]", id)
	return nil
}
