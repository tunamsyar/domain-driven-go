package services

import (
	"domain-driven-go/aggregate"
	"testing"

	"github.com/google/uuid"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Error(err)
	}

	// Setup
	tavern, err := NewTavern(WithOrderService(os))

	if err != nil {
		t.Error(err)
	}

	customer, err := aggregate.NewCustomer("Some Customer")

	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(customer)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	// Execute Tavern Order
	err = tavern.Order(customer.GetID(), order)

	if err != nil {
		t.Error(err)
	}
}
