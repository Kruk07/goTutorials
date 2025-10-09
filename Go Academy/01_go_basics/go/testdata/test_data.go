package testdata

import (
	"log"

	"example.com/go_basics/go/repository"
	"github.com/google/uuid"
)

func LoadTestData(repo *repository.Repository) {
	shrek, err := repo.CreateMovie("Shrek", 2001)
	if err != nil {
		log.Printf("Error creating movie Shrek: %v", err)
	}
	shrek2, err := repo.CreateMovie("Shrek 2", 2004)
	if err != nil {
		log.Printf("Error creating movie Shrek 2: %v", err)
	}
	lionKing, err := repo.CreateMovie("The Lion King", 1994)
	if err != nil {
		log.Printf("Error creating movie The Lion King: %v", err)
	}

	shrekChar, err := repo.CreateCharacter("Shrek")
	if err != nil {
		log.Printf("Error creating character Shrek: %v", err)
	}
	donkey, err := repo.CreateCharacter("Donkey")
	if err != nil {
		log.Printf("Error creating character Donkey: %v", err)
	}
	fiona, err := repo.CreateCharacter("Fiona")
	if err != nil {
		log.Printf("Error creating character Fiona: %v", err)
	}
	simba, err := repo.CreateCharacter("Simba")
	if err != nil {
		log.Printf("Error creating character Simba: %v", err)
	}
	pumbaa, err := repo.CreateCharacter("Pumbaa")
	if err != nil {
		log.Printf("Error creating character Pumbaa: %v", err)
	}

	appearances := []struct {
		MovieID     uuid.UUID
		CharacterID uuid.UUID
	}{
		{shrek.ID, shrekChar.ID},
		{shrek.ID, donkey.ID},
		{shrek.ID, fiona.ID},
		{shrek2.ID, shrekChar.ID},
		{shrek2.ID, donkey.ID},
		{shrek2.ID, fiona.ID},
		{lionKing.ID, simba.ID},
		{lionKing.ID, pumbaa.ID},
	}

	for _, a := range appearances {
		if err := repo.AddAppearance(a.MovieID, a.CharacterID); err != nil {
			log.Printf("Error linking character %s to movie %s: %v", a.CharacterID, a.MovieID, err)
		}
	}

	log.Println("Test data loaded successfully")
}
