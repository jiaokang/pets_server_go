package services

import (
	"jk.com/pets/internal/models"
	"jk.com/pets/internal/respositories"
)

func GetAllOwners() ([]models.Owner, error) {
	owners, err := respositories.GetAllOwners()
	if err != nil {
		return nil, err
	}
	return owners, nil
}
