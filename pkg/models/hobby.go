package models

import (
	"fmt"

	"github.com/google/uuid"
)

// HobbyKind is a component of Hobby struct
type HobbyKind uint8

const (
	_ HobbyKind = iota

	// Enrichment is like foreign language study, reading, writing/blogging, music/ musical instruments, etc.
	Enrichment

	// Sports is like jogging, horseback riding, yoga and team sports such as volley ball, bowling, soccer, etc.
	Sports

	// Social is like card games, dinner or movie club ballroom dancing, etc.
	Social

	// Creative is like scrapbooking, needle arts, jewelry making, drawing, painting, photography, etc.
	Creative

	// Collecting is like collecting antiques, décor, postcards, genealogy, etc.
	Collecting

	// Outdoors is like hiking/letterboxing, geocaching, bird-watching, hunting, fishing, etc.
	Outdoors

	// Domestic is like cooking, baking, knitting, quilting, etc.
	Domestic

	// Uncategorised is for everything that cannot be classified
	Uncategorised
)

// Hobby represents a single hobby
type Hobby struct {
	ID     uuid.UUID
	UserID uint64
	Name   string
	Kind   HobbyKind
}

// NewHobby is a constructor for Hobby
func NewHobby(name string, userID uint64, kind HobbyKind) Hobby {
	return Hobby{Name: name, UserID: userID, Kind: kind, ID: uuid.New()}
}

func (h Hobby) String() string {
	return fmt.Sprintf("Hobby{Name: %v, ID: %v, UserID: %v, Kind: %v}", h.Name, h.ID, h.UserID, h.Kind)
}
