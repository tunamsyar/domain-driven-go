package services

import (
	"log"

	"github.com/google/uuid"
)

// TavernConfiguration is an alias that takes a pointer and modifies the Tavern
type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	OrderService *OrderService

	// BillingService interface{}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	// Create the Tavern
	t := &Tavern{}

	// Apply the configurations
	for _, cfg := range cfgs {
		err := cfg(t)

		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

// WithOrderService applies a given OrderService to the Tavern
func WithOrderService(os *OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)

	if err != nil {
		return err
	}

	log.Printf("BILL CUSTOMER %0.0f", price)
	// err = t.BillingService.Bill(customer, price)
	return nil
}
