package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/product/domain"
)

type ProductCSVParser struct {
	filePath string
}

func NewCSVParser(filePath string) *ProductCSVParser {
	return &ProductCSVParser{filePath: filePath}
}

// Parse extrae los datos del CSV de Productos
func (p *ProductCSVParser) Parse() ([]domain.Product, error) {
	file, err := os.Open(p.filePath)
	if err != nil {
		return nil, fmt.Errorf("error abriendo archivo CSV de productos: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';' // Delimitador detectado

	// Leer cabecera
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("error leyendo cabecera de productos: %w", err)
	}

	var products []domain.Product
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error leyendo registro de producto: %w", err)
		}

		// IdProducto;ProductoNombre;Marca;Modelo;Categoria;Subcategoria;Precio
		if len(record) < 7 {
			continue
		}

		precio, _ := strconv.ParseFloat(strings.TrimSpace(record[6]), 64)

		product := domain.Product{
			ID:           strings.TrimSpace(record[0]),
			Nombre:       strings.TrimSpace(record[1]),
			Marca:        strings.TrimSpace(record[2]),
			Modelo:       strings.TrimSpace(record[3]),
			Categoria:    strings.TrimSpace(record[4]),
			Subcategoria: strings.TrimSpace(record[5]),
			Precio:       precio,
		}

		products = append(products, product)
	}

	return products, nil
}
