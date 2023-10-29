package services

import (
	"domain-driven-go/aggregate"
	"testing"

	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Some beer", 1.20)

	if err != nil {
		t.Error(err)
	}

	peanuts, err := aggregate.NewProduct("Peanuts", "Some Peanuts", 2.10)

	if err != nil {
		t.Error(err)
	}

	cola, err := aggregate.NewProduct("Cola", "Some cola", 2.50)

	if err != nil {
		t.Error(err)
	}

	products := []aggregate.Product{
		beer, peanuts, cola,
	}

	return products
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Error(err)
	}

	// add customer
	cust, err := aggregate.NewCustomer("John")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	// Order for 1 beer
	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)

	if err != nil {
		t.Error(err)
	}
}
