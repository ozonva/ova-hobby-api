package repo

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-hobby-api/pkg/models"
)

// Repo is an interface of storage for Hobby type
type Repo interface {
	AddHobby(hobby models.Hobby) error
	AddHobbies(hobbies []models.Hobby) error
	ListHobbies(limit, offset uint64) ([]models.Hobby, error)
	DescribeHobby(hobbyID uuid.UUID) (*models.Hobby, error)
	RemoveHobby(hobbyID uuid.UUID) error
}

func NewRepo(dbConnect *sqlx.DB) Repo {
	return &repo{db: dbConnect}
}

type repo struct {
	db *sqlx.DB
}

func (r *repo) AddHobby(hobby models.Hobby) error {
	_, err := r.db.Query(
		`INSERT INTO hobby_db.public.hobby (id, name, kind, user_id) VALUES (:id, :name, :kind, :user_id)`,
		map[string]interface{}{
			"id":      hobby.ID,
			"name":    hobby.Name,
			"kind":    hobby.Kind,
			"user_id": hobby.UserID,
		})
	return err
}

func (r *repo) AddHobbies(hobbies []models.Hobby) error {
	panic("Not implemented")
}

func (r *repo) DescribeHobby(hobbyID uuid.UUID) (*models.Hobby, error) {
	var hobby *models.Hobby
	err := r.db.Get(
		hobby,
		"SELECT id, name, kind, user_id FROM hobby_db.public.hobby WHERE id=$1", hobbyID,
	)
	return hobby, err
}

func (r *repo) ListHobbies(limit, offset uint64) ([]models.Hobby, error) {
	var hobbies []models.Hobby
	err := r.db.Select(
		&hobbies,
		"SELECT id, name, kind, user_id FROM hobby_db.public.hobby OFFSET $1 LIMIT $2",
		offset,
		limit,
	)
	return hobbies, err
}

func (r *repo) RemoveHobby(hobbyID uuid.UUID) error {
	_, err := r.db.Query("DELETE FROM hobby_db.public.hobby WHERE id = $1", hobbyID)
	return err
}
