package saver

import (
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/ozonva/ova-hobby-api/internal/flusher"
	"github.com/ozonva/ova-hobby-api/pkg/models"
)

var sugar *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	sugar = logger.Sugar()
}

// Saver is intended to save a models.Hobby
type Saver interface {
	Save(entity models.Hobby) error
	Close() ([]models.Hobby, error)
}

// NewSaver returns Saver with periodic saving using some buffer and initialize it
func NewSaver(capacity uint, flusher flusher.Flusher, saveFrequency time.Duration) (Saver, error) {
	if capacity == 0 {
		return nil, errors.New("capacity must be positive number")
	}
	if flusher == nil {
		return nil, errors.New("flusher must be initialized")
	}

	newSaver := &saver{
		buffer:        make([]models.Hobby, 0, capacity),
		flusher:       flusher,
		saveFrequency: saveFrequency,
		quit:          make(chan struct{}),
	}
	newSaver.init()
	return newSaver, nil

}

// Save adds a hobby to a buffer. When the buffer is full, it tries to flush them by itself
func (s *saver) Save(entity models.Hobby) error {
	if s.isClosed {
		return errors.New("saver is already closed")
	}
	if !s.hasEnoughCapacity() {
		s.flush()
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.buffer = append(s.buffer, entity)
	return nil
}

// Close stops and waits for the goroutine that handles hobbies saving
func (s *saver) Close() ([]models.Hobby, error) {
	if s.isClosed {
		return nil, errors.New("saver is already closed")
	}

	close(s.quit)
	s.closeWaitGroup.Wait()
	s.isClosed = true

	s.flush()
	if len(s.buffer) != 0 {
		return s.buffer, errors.New("some of the hobbies remain unsaved")
	}
	return nil, nil
}

type saver struct {
	mu             sync.RWMutex
	flusher        flusher.Flusher
	buffer         []models.Hobby
	saveFrequency  time.Duration
	quit           chan struct{}
	closeWaitGroup sync.WaitGroup
	isClosed       bool
}

// init runs a goroutine for flushing hobbies with saveFrequency
func (s *saver) init() {
	var initWaitGroup sync.WaitGroup
	initWaitGroup.Add(1)

	s.closeWaitGroup.Add(1)
	go func() {
		initWaitGroup.Done()
		defer s.closeWaitGroup.Done()

		flushTicker := time.NewTicker(s.saveFrequency)
		defer flushTicker.Stop()

		for {
			select {
			case <-flushTicker.C:
				s.flush()
			case <-s.quit:
				return
			}
		}
	}()
	initWaitGroup.Wait()
}

func (s *saver) flush() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.buffer) == 0 {
		return
	}
	unsavedHobbies := s.flusher.Flush(s.buffer)

	s.buffer = s.buffer[:0]

	if unsavedHobbies != nil {
		unableToFlushHobbies := s.flusher.Flush(unsavedHobbies)
		if unableToFlushHobbies != nil {
			sugar.Errorf("unable to flush hobbies: %v", unableToFlushHobbies)
		}
	}
}

func (s *saver) hasEnoughCapacity() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.buffer) != cap(s.buffer)
}
