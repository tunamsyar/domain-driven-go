package entity

import (
	"github.com/google/uuid"
)

// Item represents Item in all subdomains
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
