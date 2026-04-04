package application

import (
	"context"
	"fmt"
	"log"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/store/domain"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/store/domain/repository"
)

type CreateStore struct {
	repo      repository.StoreRepository
	csvParser repository.CSVParser
}

func NewCreateStore(repo repository.StoreRepository, parser repository.CSVParser) *CreateStore {
	return &CreateStore{
		repo:      repo,
		csvParser: parser,
	}
}

func (u *CreateStore) Execute(ctx context.Context) error {
	// Requerimiento: No persistir si ya existen registros
	count, err := u.repo.Count(ctx)
	if err != nil {
		return fmt.Errorf("error en consulta previa de tiendas: %w", err)
	}

	if count > 0 {
		log.Printf("[!] El proceso ETL de Tiendas se saltará: La tabla ya contiene %d registros.\n", count)
		return nil
	}

	log.Println("Extrayendo datos del CSV de Tiendas...")
	stores, err := u.csvParser.Parse()
	if err != nil {
		return fmt.Errorf("fallo la extraccion: %w", err)
	}

	log.Printf("Iniciando carga de %d tiendas...\n", len(stores))
	successCount := 0
	errorCount := 0

	for _, s := range stores {
		if err := u.repo.Save(ctx, s); err != nil {
			log.Printf("[!] Error insertando tienda %s: %v\n", s.ID, err)
			errorCount++
			continue
		}
		successCount++
	}

	log.Printf("ETL Tiendas Finalizado: %d cargadas, %d errores.\n", successCount, errorCount)
	return nil
}

func (u *CreateStore) Load(ctx context.Context, s domain.Store) error {
	return u.repo.Save(ctx, s)
}
