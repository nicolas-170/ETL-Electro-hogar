package repository

import (
	"context"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/customer/domain"
)

// CustomerRepository defines the contract for persisting customer data
type CustomerRepository interface {
	Save(ctx context.Context, customer domain.Customer) error
}
