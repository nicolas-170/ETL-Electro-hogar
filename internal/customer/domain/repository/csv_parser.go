package repository

import "github.com/nicolas-170/ETL-Electro-hogar/internal/customer/domain"

type CSVParser interface {
	Parse() ([]domain.Customer, error)
}
