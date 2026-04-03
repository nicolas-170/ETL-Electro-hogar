package repository

import (
	"context"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/customer/domain"
)

// CustomerRepository define el contrato para persistir datos de clientes
type CustomerRepository interface {
	Save(ctx context.Context, customer domain.Customer) error
}
