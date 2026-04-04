package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/time/domain"
)

type TimeCSVParser struct {
	filePath string
}

func NewCSVParser(filePath string) *TimeCSVParser {
	return &TimeCSVParser{filePath: filePath}
}

// Parse extrae los datos del CSV de Tiempo
func (p *TimeCSVParser) Parse() ([]domain.Time, error) {
	file, err := os.Open(p.filePath)
	if err != nil {
		return nil, fmt.Errorf("error abriendo archivo CSV: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';' // Delimitador detectado

	// Leer cabecera
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("error leyendo cabecera: %w", err)
	}

	var times []domain.Time
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error leyendo registro: %w", err)
		}

		// IdTiempo;Fecha;Dia;Mes;Trimestre;Año;DiaSemana;EsfinSemana
		if len(record) < 8 {
			continue
		}

		// Parsear Fecha (D/M/YYYY)
		fechaStr := strings.TrimSpace(record[1])
		fecha, err := time.Parse("2/1/2006", fechaStr)
		if err != nil {
			// Intentar con padding si falla
			fecha, err = time.Parse("02/01/2006", fechaStr)
			if err != nil {
				fmt.Printf("[!] Error parseando fecha '%s': %v\n", fechaStr, err)
				continue
			}
		}

		dia, _ := strconv.Atoi(record[2])
		mes, _ := strconv.Atoi(record[3])
		trimestre, _ := strconv.Atoi(record[4])
		anio, _ := strconv.Atoi(record[5])
		esFinSemana := record[7] == "1"

		timeEntry := domain.Time{
			ID:          strings.TrimSpace(record[0]),
			Fecha:       fecha,
			Dia:         dia,
			Mes:         mes,
			Trimestre:   trimestre,
			Anio:        anio,
			DiaSemana:   strings.TrimSpace(record[6]),
			EsFinSemana: esFinSemana,
		}

		times = append(times, timeEntry)
	}

	return times, nil
}
