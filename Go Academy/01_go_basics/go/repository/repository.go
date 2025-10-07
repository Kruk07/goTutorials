package repository

import (
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

func (r *Repository) CreateMovie(title string, year int) entity.Movie {
	movie := entity.NewMovie(
		entity.WithTitle(title),
		entity.WithYear(year),
	)

	r.DB.Movies[movie.ID] = movie
	log.Printf("Movie added: %s (%d) [ID: %s]", title, year, movie.ID)
	return movie
}

func (r *Repository) CreateCharacter(name string) entity.Character {
	character := entity.NewCharacter(entity.WithName(name))
	r.DB.Characters[character.ID] = character
	log.Printf("Character added: %s [ID: %s]", name, character.ID)
	return character
}

func (r *Repository) AddAppearance(movieID, characterID uuid.UUID) {
	r.DB.Appearances = append(r.DB.Appearances, entity.New(entity.WithMovieId(movieID), entity.WithCharacterId(characterID)))

	movie, _ := r.DB.Movies[movieID]
	character, _ := r.DB.Characters[characterID]
	log.Printf("Linked character '%s' to movie '%s'", character.Name, movie.Title)
}

func (r *Repository) GetCharactersByMovie(movieID uuid.UUID) []entity.Character {
	var result []entity.Character
	for _, a := range r.DB.Appearances {
		if a.MovieID == movieID {
			if c, ok := r.DB.Characters[a.CharacterID]; ok {
				result = append(result, c)
			}
		}
	}
	log.Printf("Retrieved %d characters for movie ID %s", len(result), movieID)
	return result
}

func (r *Repository) GetMoviesByCharacter(characterID uuid.UUID) []entity.Movie {
	var result []entity.Movie
	for _, a := range r.DB.Appearances {
		if a.CharacterID == characterID {
			if m, ok := r.DB.Movies[a.MovieID]; ok {
				result = append(result, m)
			}
		}
	}
	log.Printf("Retrieved %d movies for character ID %s", len(result), characterID)
	return result
}

func (r *Repository) GetCharactersByMovieTitle(title string) []entity.Character {
	for _, m := range r.DB.Movies {
		if m.Title == title {
			log.Printf("Searching characters for movie title: %s", title)
			return r.GetCharactersByMovie(m.ID)
		}
	}
	log.Printf("No movie found with title: %s", title)
	return []entity.Character{}
}

func (r *Repository) GetMovieTitlesByCharacterName(name string) []string {
	for _, c := range r.DB.Characters {
		if c.Name == name {
			log.Printf("Searching movies for character name: %s", name)
			movies := r.GetMoviesByCharacter(c.ID)
			var titles []string
			for _, m := range movies {
				titles = append(titles, m.Title)
			}
			log.Printf("Found %d movies for character '%s'", len(titles), name)
			return titles
		}
	}
	log.Printf("No character found with name: %s", name)
	return []string{}
}

func (r *Repository) ListAllMovies() {
	if len(r.DB.Movies) == 0 {
		log.Println("No movies found in the database.")
		return
	}
	log.Println("All movies:")
	for _, m := range r.DB.Movies {
		log.Printf("- %s (%d) [ID: %s]", m.Title, m.Year, m.ID)
	}
}

func (r *Repository) ListAllCharacters() {
	if len(r.DB.Characters) == 0 {
		log.Println("No characters found in the database.")
		return
	}
	log.Println("All characters:")
	for _, c := range r.DB.Characters {
		log.Printf("- %s [ID: %s]", c.Name, c.ID)
	}
}

func (r *Repository) UpdateCharacter(id uuid.UUID, newName string) bool {
	if character, ok := r.DB.Characters[id]; ok {
		character.Name = newName
		r.DB.Characters[id] = character
		log.Printf("Character updated: %s [ID: %s]", newName, id)
		return true
	}
	log.Printf("Character not found for update [ID: %s]", id)
	return false
}

func (r *Repository) DeleteMovie(id uuid.UUID) bool {
	if _, ok := r.DB.Movies[id]; ok {
		delete(r.DB.Movies, id)
		var updated []entity.Appearance
		for _, a := range r.DB.Appearances {
			if a.MovieID != id {
				updated = append(updated, a)
			}
		}
		r.DB.Appearances = updated
		log.Printf("Movie deleted [ID: %s]", id)
		return true
	}
	log.Printf("Movie not found for deletion [ID: %s]", id)
	return false
}

func (r *Repository) DeleteCharacter(id uuid.UUID) bool {
	if _, ok := r.DB.Characters[id]; ok {
		delete(r.DB.Characters, id)
		var updated []entity.Appearance
		for _, a := range r.DB.Appearances {
			if a.CharacterID != id {
				updated = append(updated, a)
			}
		}
		r.DB.Appearances = updated
		log.Printf("Character deleted [ID: %s]", id)
		return true
	}
	log.Printf("Character not found for deletion [ID: %s]", id)
	return false
}
