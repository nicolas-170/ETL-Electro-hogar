package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
)

// NewDB crea y retorna el pool de conexiones a la base de datos
func NewDB(cfg config.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	// Abrir la conexión usando el driver correspondiente ("postgres", "mysql", "sqlserver", ...)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la base de datos: %w", err)
	}

	// Comprobar que realmente podemos alcanzar la base de datos
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	return db, nil
}
