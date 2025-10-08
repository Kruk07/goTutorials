package testdata

import (
	"log"

	"example.com/go_basics/go/repository"
)

func LoadTestData(repo *repository.Repository) {
	shrek := repo.CreateMovie("Shrek", 2001)
	shrek2 := repo.CreateMovie("Shrek 2", 2004)
	lionKing := repo.CreateMovie("The Lion King", 1994)

	shrekChar := repo.CreateCharacter("Shrek")
	donkey := repo.CreateCharacter("Donkey")
	fiona := repo.CreateCharacter("Fiona")
	simba := repo.CreateCharacter("Simba")
	pumbaa := repo.CreateCharacter("Pumbaa")

	repo.AddAppearance(shrek.ID, shrekChar.ID)
	repo.AddAppearance(shrek.ID, donkey.ID)
	repo.AddAppearance(shrek.ID, fiona.ID)

	repo.AddAppearance(shrek2.ID, shrekChar.ID)
	repo.AddAppearance(shrek2.ID, donkey.ID)
	repo.AddAppearance(shrek2.ID, fiona.ID)

	repo.AddAppearance(lionKing.ID, simba.ID)
	repo.AddAppearance(lionKing.ID, pumbaa.ID)

	log.Println("Test data loaded successfully")
}
