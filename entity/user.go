package entity

import (
	"github.com/google/uuid"
)

// User represents User throughout domains
type User struct {
	ID   uuid.UUID
	Name string
	Age  int
}
