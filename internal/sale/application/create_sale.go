package application

import (
	"context"
	"fmt"
	"log"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/sale/domain"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/sale/domain/repository"
)

type CreateSale struct {
	repo      repository.SaleRepository
	csvParser repository.CSVParser
}

func NewCreateSale(repo repository.SaleRepository, parser repository.CSVParser) *CreateSale {
	return &CreateSale{
		repo:      repo,
		csvParser: parser,
	}
}

func (u *CreateSale) Execute(ctx context.Context) error {
	// Requerimiento: No persistir si ya existen registros
	count, err := u.repo.Count(ctx)
	if err != nil {
		return fmt.Errorf("error en consulta previa de ventas: %w", err)
	}

	if count > 0 {
		log.Printf("[!] El proceso ETL de Ventas se saltará: La tabla ya contiene %d registros.\n", count)
		return nil
	}

	log.Println("Extrayendo datos del CSV de Ventas...")
	sales, err := u.csvParser.Parse()
	if err != nil {
		return fmt.Errorf("fallo la extraccion: %w", err)
	}

	log.Printf("Iniciando carga de %d ventas...\n", len(sales))
	successCount := 0
	errorCount := 0

	for _, s := range sales {
		if err := u.repo.Save(ctx, s); err != nil {
			log.Printf("[!] Error insertando venta %s: %v\n", s.ID, err)
			errorCount++
			continue
		}
		successCount++
	}

	log.Printf("ETL Ventas Finalizado: %d cargadas, %d errores.\n", successCount, errorCount)
	return nil
}

func (u *CreateSale) Load(ctx context.Context, s domain.Sale) error {
	return u.repo.Save(ctx, s)
}
