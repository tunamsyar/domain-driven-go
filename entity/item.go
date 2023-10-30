package entity

import (
	"github.com/google/uuid"
)

// Entity is a struct with an identifier and is mutable
// Item represents Item in all subdomains
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
