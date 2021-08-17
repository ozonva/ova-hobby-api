package hobby

import (
	"fmt"

	"github.com/google/uuid"
)

type HobbyKind uint8

const (
	_             HobbyKind = iota
	Enrichment              // foreign language study, reading, writing/blogging, music/ musical instruments, etc.
	Sports                  // jogging, horseback riding, yoga and team sports such as volley ball, bowling, soccer, etc.
	Social                  // card games, dinner or movie club ballroom dancing, etc.
	Creative                // scrapbooking, needle arts, jewelry making, drawing, painting, photography, etc.
	Collecting              // collecting antiques, décor, postcards, genealogy, etc.
	Outdoors                // hiking/letterboxing, geocaching, bird-watching, hunting, fishing, etc.
	Domestic                // cooking, baking, knitting, quilting, etc.
	Uncategorised           // everything that cannot be classified
)

type Hobby struct {
	ID     uuid.UUID
	UserID uint64
	Name   string
	Kind   HobbyKind
}

func NewHobby(name string, userID uint64, kind HobbyKind) Hobby {
	return Hobby{Name: name, UserID: userID, Kind: kind, ID: uuid.New()}
}

func (h Hobby) String() string {
	return fmt.Sprintf("Hobby{Name: %v, ID: %v, UserID: %v, Kind: %v}", h.Name, h.ID, h.UserID, h.Kind)
}
