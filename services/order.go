package services

import (
	"domain-driven-go/domain/customer"
	"domain-driven-go/domain/customer/memory"
)

// OrderConfiguration is an alias for a function that will take OrderService as a pointer
type OrderConfiguration func(os *OrderService) error

// OrderService is an implementation of OrderService
type OrderService struct {
	customers customer.CustomerRepository
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
