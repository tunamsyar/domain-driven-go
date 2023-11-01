package services

import (
	"context"
	"domain-driven-go/aggregate"
	"domain-driven-go/domain/customer"
	memory "domain-driven-go/domain/customer/memory"
	"domain-driven-go/domain/customer/mongo"
	"domain-driven-go/domain/product"
	prodmemory "domain-driven-go/domain/product/memory"
	"log"

	"github.com/google/uuid"
)

// OrderConfiguration is an alias for a function that will take OrderService as a pointer
type OrderConfiguration func(os *OrderService) error

// OrderService is an implementation of OrderService
type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// Create OrderService
	os := &OrderService{}

	// Apply configs
	for _, cfg := range cfgs {
		// Pass the service into config function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

// WithCustomerRepository applies a given customer repo to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// returns a function that matches the OrderConfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// WithMemoryCustomerRepository applies a memory Customer Repository to the OrderService
func WithMemoryCustomerRepository() OrderConfiguration {

	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

// WithMemoryProductRepository applies a memory Product Repository to the OrderService
func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmemory.New()

		for _, p := range products {
			err := pr.Add(p)

			if err != nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}

// CreateOrder is called by the OrderService to create Order
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Get Customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []aggregate.Product
	var price float64

	// Get products and sum the price
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)

		if err != nil {
			return 0, err
		}

		products = append(products, p)
		price += p.GetPrice()
	}

	log.Printf("Customer: %s has orderd %d products", c.GetID(), len(products))

	return price, nil
}
