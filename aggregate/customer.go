package aggregate

import (
	"domain-driven-go/entity"
	"domain-driven-go/value_object"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidUser = errors.New("invalid user")
)

type Customer struct {
	// user is Root Entity
	user *entity.User
	// customer has many products
	products []*entity.Item
	// customer has many transactions
	transactions []*value_object.Transaction
}

// NewCustomer factory pattern. Creates a Customer Aggregate
// Move to factory package?
func NewCustomer(name string) (Customer, error) {

	// Validate name presence
	if name == "" {
		return Customer{}, ErrInvalidUser
	}

	// Create new user and generate ID
	user := &entity.User{
		Name: name,
		ID:   uuid.New(),
	}

	// Returns customer object with initialized data to avoid nil pointers
	return Customer{
		user:         user,
		products:     make([]*entity.Item, 0),
		transactions: make([]*value_object.Transaction, 0),
	}, nil
}

// GetID returns customer root user id
func (c *Customer) GetID() uuid.UUID {
	return c.user.ID
}

// SetID sets the root user id
func (c *Customer) SetID(id uuid.UUID) {
	if c.user == nil {
		c.user = &entity.User{}
	}
	c.user.ID = id
}

// SetName sets the name of the root user
func (c *Customer) SetName(name string) {
	if c.user == nil {
		c.user = &entity.User{}
	}
	c.user.Name = name
}

// GetName returns the root user name
func (c *Customer) GetName() string {
	return c.user.Name
}
