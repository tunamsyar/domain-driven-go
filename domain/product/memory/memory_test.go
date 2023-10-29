package memory

import (
	"domain-driven-go/aggregate"
	"domain-driven-go/domain/product"
	"testing"

	"github.com/google/uuid"
)

func TestMemoryProductRepository_Add(t *testing.T) {
	repo := New()
	product, err := aggregate.NewProduct("Beer", "Gitcher goin", 1.20)

	if err != nil {
		t.Error(err)
	}

	repo.Add(product)

	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}
}

func TestMemoryProductRepository_Get(t *testing.T) {

	repo := New()
	existingprod, err := aggregate.NewProduct("Beer", "Gitcher goin", 1.20)

	if err != nil {
		t.Error(err)
	}

	repo.Add(existingprod)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "ExistingProd",
			id:          existingprod.GetID(),
			expectedErr: nil,
		},
		{
			name:        "Non existent prod",
			id:          uuid.New(),
			expectedErr: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)

			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemoryProductRepository_Delete(t *testing.T) {
	repo := New()
	existingprod, err := aggregate.NewProduct("Beer", "Gitcher goin", 1.20)

	if err != nil {
		t.Error(err)
	}

	repo.Add(existingprod)
	if len(repo.products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.products))
	}

	err = repo.Delete(existingprod.GetID())

	if err != nil {
		t.Error(err)
	}

	if len(repo.products) != 0 {
		t.Errorf("Expected 0 products, got %d", len(repo.products))
	}
}
