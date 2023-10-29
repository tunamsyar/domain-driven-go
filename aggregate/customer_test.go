package aggregate_test

import (
	"domain-driven-go/aggregate"
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T) {

	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	// test cases
	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidUser,
		},
		{
			test:        "Valid name",
			name:        "Some name",
			expectedErr: nil,
		},
	}

	// actual test run
	for _, tc := range testCases {
		// run test case for each
		t.Run(tc.test, func(t *testing.T) {

			// create customer
			_, err := aggregate.NewCustomer(tc.name)

			// check error matches
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
