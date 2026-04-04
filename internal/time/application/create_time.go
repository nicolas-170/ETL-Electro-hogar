package application

import (
	"context"
	"fmt"
	"log"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/time/domain"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/time/domain/repository"
)

// CreateTime orquesta el proceso ETL para la dimensión Tiempo
type CreateTime struct {
	repo      repository.TimeRepository
	csvParser repository.CSVParser
}

func NewCreateTime(repo repository.TimeRepository, parser repository.CSVParser) *CreateTime {
	return &CreateTime{
		repo:      repo,
		csvParser: parser,
	}
}

// Execute ejecuta el proceso ETL de Tiempo
func (u *CreateTime) Execute(ctx context.Context) error {
	// No persistir si ya existen registros
	count, err := u.repo.Count(ctx)
	if err != nil {
		return fmt.Errorf("error en consulta previa de tiempo: %w", err)
	}

	if count > 0 {
		log.Printf("[!] El proceso ETL de Tiempo se saltará: La tabla ya contiene %d registros.\n", count)
		return nil
	}

	log.Println("Extrayendo datos del CSV de Tiempo...")
	times, err := u.csvParser.Parse()
	if err != nil {
		return fmt.Errorf("fallo la extraccion: %w", err)
	}

	log.Printf("Iniciando carga de %d registros de tiempo...\n", len(times))
	successCount := 0
	errorCount := 0

	for _, t := range times {
		// En Tiempo la transformación es mínima ya que el CSV está bien estructurado
		// No obstante, se puede añadir lógica de normalización si fuera necesario.

		// --- Carga (L de ETL) ---
		err := u.Load(ctx, t)
		if err != nil {
			log.Printf("[!] Error insertando registro de tiempo %s: %v\n", t.ID, err)
			errorCount++
			continue
		}

		successCount++
	}

	log.Printf("ETL Tiempo Finalizado: %d cargados, %d errores.\n", successCount, errorCount)
	return nil
}

func (u *CreateTime) Load(ctx context.Context, t domain.Time) error {
	return u.repo.Save(ctx, t)
}
