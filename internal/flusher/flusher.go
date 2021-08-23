package flusher

import (
	"github.com/ozonva/ova-hobby-api/internal/repo"
	"github.com/ozonva/ova-hobby-api/internal/utils"
	"github.com/ozonva/ova-hobby-api/pkg/models"
)

// Flusher is an interface to dump hobbies into a storage
type Flusher interface {
	Flush(hobbies []models.Hobby) []models.Hobby
}

// NewFlusher creates a Flusher with saving in batches
func NewFlusher(chunkSize int, hobbyRepo repo.Repo) Flusher {
	return &flusher{chunkSize: chunkSize, hobbyRepo: hobbyRepo}
}

type flusher struct {
	chunkSize int
	hobbyRepo repo.Repo
}

func (f *flusher) Flush(hobbies []models.Hobby) []models.Hobby {
	if hobbies == nil {
		return hobbies
	}
	hobbyBatches, err := utils.SliceHobbiesIntoBatches(hobbies, f.chunkSize)
	if err != nil {
		return hobbies
	}
	var unableToSaveHobbies []models.Hobby
	for _, batch := range hobbyBatches {
		err := f.hobbyRepo.AddHobbies(batch)
		if err != nil {
			unableToSaveHobbies = append(unableToSaveHobbies, batch...)
		}
	}
	return unableToSaveHobbies
}
