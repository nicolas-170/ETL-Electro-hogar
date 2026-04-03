package infrastructure

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/customer/domain"
)

// CSVParser handles extraction of customer data from CSV files
type CSVParser struct {
	filePath string
}

func NewCSVParser(filePath string) *CSVParser {
	return &CSVParser{filePath: filePath}
}

// Parse reads the CSV and returns a slice of Customers
func (p *CSVParser) Parse() ([]domain.Customer, error) {
	file, err := os.Open(p.filePath)
	if err != nil {
		return nil, fmt.Errorf("error al abrir archivo CSV: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';' // Delimitador especificado en el archivo
	reader.LazyQuotes = true

	// Leer encabezados
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error al leer encabezados: %w", err)
	}

	var customers []domain.Customer
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error al leer registro: %w", err)
		}

		// idCliente;Nombre;Email;Telefono;Segmento;Ciudad;Sexo;Edad
		if len(record) < 8 {
			continue // Omitir líneas incompletas
		}

		edad, _ := strconv.Atoi(strings.TrimSpace(record[7]))

		customer := domain.Customer{
			ID:       strings.TrimSpace(record[0]),
			Nombre:   strings.TrimSpace(record[1]),
			Email:    strings.TrimSpace(record[2]),
			Telefono: strings.TrimSpace(record[3]),
			Segmento: strings.TrimSpace(record[4]),
			Ciudad:   strings.TrimSpace(record[5]),
			Sexo:     strings.TrimSpace(record[6]),
			Edad:     edad,
		}

		customers = append(customers, customer)
	}

	return customers, nil
}
