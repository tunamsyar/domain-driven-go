package customer

import (
	"domain-driven-go/aggregate"
	"errors"

	"github.com/google/uuid"
)

// Error messages
var (
	ErrCustomerNotFound = errors.New("customer not found")

	ErrFailedToAddCustomer = errors.New("failed to add customer")

	ErrUpdateCustomer = errors.New("failed to update customer")
)

// Customer Repository interface
// Method logic are in memory.go
type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
