package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/store/domain"
)

type StoreCSVParser struct {
	filePath string
}

func NewCSVParser(filePath string) *StoreCSVParser {
	return &StoreCSVParser{filePath: filePath}
}

// Parse extrae los datos del CSV de Tiendas
func (p *StoreCSVParser) Parse() ([]domain.Store, error) {
	file, err := os.Open(p.filePath)
	if err != nil {
		return nil, fmt.Errorf("error abriendo archivo CSV de tiendas: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';' // Delimitador detectado en inspección previa

	// Leer cabecera
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("error leyendo cabecera de tiendas: %w", err)
	}

	var stores []domain.Store
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error leyendo registro de tienda: %w", err)
		}

		// IdCiudad;Ciudad;Departamento;Region
		if len(record) < 4 {
			continue
		}

		store := domain.Store{
			ID:           strings.TrimSpace(record[0]),
			Nombre:       strings.TrimSpace(record[1]),
			Departamento: strings.TrimSpace(record[2]),
			Region:       strings.TrimSpace(record[3]),
		}

		stores = append(stores, store)
	}

	return stores, nil
}
