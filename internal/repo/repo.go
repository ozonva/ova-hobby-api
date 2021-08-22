package repo

import (
	"github.com/ozonva/ova-hobby-api/pkg/models"
)

// Repo is an interface of storage for Hobby type
type Repo interface {
	AddHobbies(hobbies []models.Hobby) error
	ListHobbies(limit, offset uint64) ([]models.Hobby, error)
	DescribeHobby(hobbyID uint64) (*models.Hobby, error)
}
