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
	r.DB.Movies.Store(movie.ID, movie)
	log.Printf("Movie added: %s (%d) [ID: %s]", title, year, movie.ID)
	return movie, nil
}

func (r *Repository) CreateCharacter(name string) (entity.Character, error) {
	if name == "" {
		return entity.Character{}, errors.New("character name cannot be empty")
	}
	character := entity.NewCharacter(entity.WithName(name))
	r.DB.Characters.Store(character.ID, character)
	log.Printf("Character added: %s [ID: %s]", name, character.ID)
	return character, nil
}

func (r *Repository) AddAppearance(movieID, characterID uuid.UUID) error {
	mRaw, ok := r.DB.Movies.Load(movieID)
	if !ok {
		return fmt.Errorf("movie not found [ID: %s]", movieID)
	}
	cRaw, ok := r.DB.Characters.Load(characterID)
	if !ok {
		return fmt.Errorf("character not found [ID: %s]", characterID)
	}
	movie := mRaw.(entity.Movie)
	character := cRaw.(entity.Character)

	r.DB.Mutex.Lock()
	r.DB.Appearances = append(r.DB.Appearances, entity.New(
		entity.WithMovieId(movieID),
		entity.WithCharacterId(characterID),
	))
	r.DB.Mutex.Unlock()

	log.Printf("Linked character '%s' to movie '%s'", character.Name, movie.Title)
	return nil
}

func (r *Repository) GetCharactersByMovie(movieID uuid.UUID) ([]entity.Character, error) {
	if _, ok := r.DB.Movies.Load(movieID); !ok {
		return nil, fmt.Errorf("movie not found [ID: %s]", movieID)
	}
	var result []entity.Character
	r.DB.Mutex.Lock()
	for _, a := range r.DB.Appearances {
		if a.MovieID == movieID {
			if cRaw, ok := r.DB.Characters.Load(a.CharacterID); ok {
				if c, ok := cRaw.(entity.Character); ok {
					result = append(result, c)
				}
			}
		}
	}
	r.DB.Mutex.Unlock()
	log.Printf("Retrieved %d characters for movie ID %s", len(result), movieID)
	return result, nil
}

func (r *Repository) GetMoviesByCharacter(characterID uuid.UUID) ([]entity.Movie, error) {
	if _, ok := r.DB.Characters.Load(characterID); !ok {
		return nil, fmt.Errorf("character not found [ID: %s]", characterID)
	}
	var result []entity.Movie
	r.DB.Mutex.Lock()
	for _, a := range r.DB.Appearances {
		if a.CharacterID == characterID {
			if mRaw, ok := r.DB.Movies.Load(a.MovieID); ok {
				if m, ok := mRaw.(entity.Movie); ok {
					result = append(result, m)
				}
			}
		}
	}
	r.DB.Mutex.Unlock()
	log.Printf("Retrieved %d movies for character ID %s", len(result), characterID)
	return result, nil
}

func (r *Repository) GetCharactersByMovieTitle(title string) ([]entity.Character, error) {
	var movieID uuid.UUID
	found := false
	r.DB.Movies.Range(func(_, value any) bool {
		m := value.(entity.Movie)
		if m.Title == title {
			movieID = m.ID
			found = true
			return false
		}
		return true
	})
	if !found {
		log.Printf("No movie found with title: %s", title)
		return nil, fmt.Errorf("movie not found with title: %s", title)
	}
	log.Printf("Searching characters for movie title: %s", title)
	return r.GetCharactersByMovie(movieID)
}

func (r *Repository) GetMovieTitlesByCharacterName(name string) ([]string, error) {
	var characterID uuid.UUID
	found := false
	r.DB.Characters.Range(func(_, value any) bool {
		c := value.(entity.Character)
		if c.Name == name {
			characterID = c.ID
			found = true
			return false
		}
		return true
	})
	if !found {
		log.Printf("No character found with name: %s", name)
		return nil, fmt.Errorf("character not found with name: %s", name)
	}
	movies, err := r.GetMoviesByCharacter(characterID)
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

func (r *Repository) ListAllMovies() (map[uuid.UUID]entity.Movie, error) {
	result := make(map[uuid.UUID]entity.Movie)
	r.DB.Movies.Range(func(key, value any) bool {
		id := key.(uuid.UUID)
		m := value.(entity.Movie)
		result[id] = m
		log.Printf("- %s (%d) [ID: %s]", m.Title, m.Year, m.ID)
		return true
	})
	if len(result) == 0 {
		log.Println("No movies found in the database.")
		return nil, errors.New("no movies available")
	}
	log.Println("All movies listed.")
	return result, nil
}

func (r *Repository) ListAllCharacters() (map[uuid.UUID]entity.Character, error) {
	result := make(map[uuid.UUID]entity.Character)
	r.DB.Characters.Range(func(key, value any) bool {
		id := key.(uuid.UUID)
		c := value.(entity.Character)
		result[id] = c
		log.Printf("- %s [ID: %s]", c.Name, c.ID)
		return true
	})
	if len(result) == 0 {
		log.Println("No characters found in the database.")
		return nil, errors.New("no characters available")
	}
	log.Println("All characters listed.")
	return result, nil
}

func (r *Repository) UpdateCharacter(id uuid.UUID, newName string) error {
	if newName == "" {
		return errors.New("new character name cannot be empty")
	}
	cRaw, ok := r.DB.Characters.Load(id)
	if !ok {
		return fmt.Errorf("character not found [ID: %s]", id)
	}
	character := cRaw.(entity.Character)
	character.Name = newName
	r.DB.Characters.Store(id, character)
	log.Printf("Character updated: %s [ID: %s]", newName, id)
	return nil
}

func (r *Repository) DeleteMovie(id uuid.UUID) error {
	if _, ok := r.DB.Movies.Load(id); !ok {
		return fmt.Errorf("movie not found [ID: %s]", id)
	}
	r.DB.Movies.Delete(id)
	r.DB.Mutex.Lock()
	var updated []entity.Appearance
	for _, a := range r.DB.Appearances {
		if a.MovieID != id {
			updated = append(updated, a)
		}
	}
	r.DB.Appearances = updated
	r.DB.Mutex.Unlock()
	log.Printf("Movie deleted [ID: %s]", id)
	return nil
}

func (r *Repository) DeleteCharacter(id uuid.UUID) error {
	if _, ok := r.DB.Characters.Load(id); !ok {
		return fmt.Errorf("character not found [ID: %s]", id)
	}
	r.DB.Characters.Delete(id)
	r.DB.Mutex.Lock()
	var updated []entity.Appearance
	for _, a := range r.DB.Appearances {
		if a.CharacterID != id {
			updated = append(updated, a)
		}
	}
	r.DB.Appearances = updated
	r.DB.Mutex.Unlock()
	log.Printf("Character deleted [ID: %s]", id)
	return nil
}
