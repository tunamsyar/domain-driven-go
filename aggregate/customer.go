package aggregate

import (
	"domain-driven-go/entity"
	"domain-driven-go/value_object"
)

type Customer struct {
	user         *entity.User
	products     []*entity.Item
	transactions []*value_object.Transaction
}
