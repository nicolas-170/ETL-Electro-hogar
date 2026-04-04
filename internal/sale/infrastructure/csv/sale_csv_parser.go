package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/sale/domain"
)

type SaleCSVParser struct {
	filePath string
}

func NewCSVParser(filePath string) *SaleCSVParser {
	return &SaleCSVParser{filePath: filePath}
}

// Parse extrae los datos del CSV de Ventas
func (p *SaleCSVParser) Parse() ([]domain.Sale, error) {
	file, err := os.Open(p.filePath)
	if err != nil {
		return nil, fmt.Errorf("error abriendo archivo CSV de ventas: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';' // Delimitador detectado

	// Leer cabecera
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("error leyendo cabecera de ventas: %w", err)
	}

	var sales []domain.Sale
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error leyendo registro de venta: %w", err)
		}

		// IdVenta;IdTiempo;IdCliente;IdProducto;IdTienda;Cantidad;PrecioUnitario;TotalVenta;CostoUnitario;MargenUnitario;MargenTotal
		if len(record) < 11 {
			continue
		}

		cantidad, _ := strconv.Atoi(strings.TrimSpace(record[5]))
		precioUnitario, _ := strconv.ParseFloat(strings.TrimSpace(record[6]), 64)
		totalVenta, _ := strconv.ParseFloat(strings.TrimSpace(record[7]), 64)
		costoUnitario, _ := strconv.ParseFloat(strings.TrimSpace(record[8]), 64)
		margenUnitario, _ := strconv.ParseFloat(strings.TrimSpace(record[9]), 64)
		margenTotal, _ := strconv.ParseFloat(strings.TrimSpace(record[10]), 64)

		sale := domain.Sale{
			ID:             strings.TrimSpace(record[0]),
			IDTiempo:       strings.TrimSpace(record[1]),
			IDCliente:      strings.TrimSpace(record[2]),
			IDProducto:     strings.TrimSpace(record[3]),
			IDTienda:       strings.TrimSpace(record[4]),
			Cantidad:       cantidad,
			PrecioUnitario: precioUnitario,
			TotalVenta:     totalVenta,
			CostoUnitario:  costoUnitario,
			MargenUnitario: margenUnitario,
			MargenTotal:    margenTotal,
		}

		sales = append(sales, sale)
	}

	return sales, nil
}
