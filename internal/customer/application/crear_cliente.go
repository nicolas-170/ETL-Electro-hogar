package application

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/customer/domain"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/customer/domain/repository"
)

// CrearCliente orquesta el proceso ETL para los clientes
type CrearCliente struct {
	repo      repository.CustomerRepository
	csvParser repository.CSVParser
}

func NewCrearCliente(repo repository.CustomerRepository, parser repository.CSVParser) *CrearCliente {
	return &CrearCliente{
		repo:      repo,
		csvParser: parser,
	}
}

// Execute ejecuta el proceso ETL
func (u *CrearCliente) Execute(ctx context.Context) error {
	log.Println("Extrayendo datos del CSV...")
	customers, err := u.csvParser.Parse()
	if err != nil {
		return fmt.Errorf("fallo la extraccion: %w", err)
	}

	log.Printf("Iniciando carga de %d clientes...\n", len(customers))
	successCount := 0
	errorCount := 0

	for _, customer := range customers {
		// --- Transformación (T de ETL) ---
		// Limpiar espacios y normalizar datos
		u.Transform(&customer)

		// Omitir si no hay nombre (datos inválidos en el CSV)
		if customer.Nombre == "" || customer.Nombre == "NULL" {
			log.Printf("[-] Saltando cliente ID %s por falta de nombre\n", customer.ID)
			errorCount++
			continue
		}

		// --- Carga (L de ETL) ---
		err := u.Load(customer, ctx)
		if err != nil {
			log.Printf("[!] Error insertando cliente %s (%s): %v\n", customer.ID, customer.Nombre, err)
			errorCount++
			continue
		}

		successCount++
	}

	log.Printf("ETL Finalizado: %d cargados, %d errores.\n", successCount, errorCount)
	return nil
}

func (u *CrearCliente) Transform(customer *domain.Customer) {
	customer.Nombre = strings.TrimSpace(customer.Nombre)
	customer.Ciudad = u.normalizeCity(customer.Ciudad)
	customer.Sexo = u.normalizeGender(customer.Sexo)
}

func (u *CrearCliente) Load(customer domain.Customer, ctx context.Context) error {
	return u.repo.Save(ctx, customer)
}

// normalizeCity maneja variaciones en el nombre de las ciudades
func (u *CrearCliente) normalizeCity(city string) string {
	city = strings.TrimSpace(city)
	if city == "NULL" || city == "" {
		return "Desconocida"
	}
	// Convertir BOGOTA -> Bogotá, MEDELLIN -> Medellín, etc.
	// Esto es una normalización básica, se puede expandir
	return strings.Title(strings.ToLower(city))
}

// normalizeGender limpia el campo sexo (e.g. "M " -> "M")
func (u *CrearCliente) normalizeGender(gender string) string {
	gender = strings.TrimSpace(gender)
	if len(gender) > 1 {
		return gender[0:1]
	}
	return gender
}
