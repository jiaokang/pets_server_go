package respositories

import (
	"jk.com/pets/internal/database"
	"jk.com/pets/internal/models"
)

func GetAllOwners() ([]models.Owner, error) {
	var owners []models.Owner
	result := database.DB.Find(&owners)
	return owners, result.Error
}
