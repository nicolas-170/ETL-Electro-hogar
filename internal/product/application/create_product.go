package application

import (
	"context"
	"fmt"
	"log"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/product/domain"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/product/domain/repository"
)

type CreateProduct struct {
	repo      repository.ProductRepository
	csvParser repository.CSVParser
}

func NewCreateProduct(repo repository.ProductRepository, parser repository.CSVParser) *CreateProduct {
	return &CreateProduct{
		repo:      repo,
		csvParser: parser,
	}
}

func (u *CreateProduct) Execute(ctx context.Context) error {
	// Requerimiento: No persistir si ya existen registros
	count, err := u.repo.Count(ctx)
	if err != nil {
		return fmt.Errorf("error en consulta previa de productos: %w", err)
	}

	if count > 0 {
		log.Printf("[!] El proceso ETL de Productos se saltará: La tabla ya contiene %d registros.\n", count)
		return nil
	}

	log.Println("Extrayendo datos del CSV de Productos...")
	products, err := u.csvParser.Parse()
	if err != nil {
		return fmt.Errorf("fallo la extraccion: %w", err)
	}

	log.Printf("Iniciando carga de %d productos...\n", len(products))
	successCount := 0
	errorCount := 0

	for _, p := range products {
		if err := u.repo.Save(ctx, p); err != nil {
			log.Printf("[!] Error insertando producto %s: %v\n", p.ID, err)
			errorCount++
			continue
		}
		successCount++
	}

	log.Printf("ETL Productos Finalizado: %d cargados, %d errores.\n", successCount, errorCount)
	return nil
}

func (u *CreateProduct) Load(ctx context.Context, p domain.Product) error {
	return u.repo.Save(ctx, p)
}
