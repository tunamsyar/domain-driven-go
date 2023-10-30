package value_object

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a value object with no identifier
// Immutable
// Maybe for this there should be an ID for transactions?
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
