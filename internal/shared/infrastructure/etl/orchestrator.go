package etl

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
)

// Orchestrator gestiona una lista de tareas ETL y las ejecuta secuencialmente
type Orchestrator struct {
	tasks []Task
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		tasks: make([]Task, 0),
	}
}

// Register añade una nueva tarea a la lista
func (o *Orchestrator) Register(task Task) {
	o.tasks = append(o.tasks, task)
}

// RunAll ejecuta todas las tareas registradas
func (o *Orchestrator) RunAll(ctx context.Context, db *sql.DB, cfg *config.Database) error {
	log.Printf("Iniciando orquestación de %d tareas ETL...\n", len(o.tasks))

	for i, task := range o.tasks {
		log.Printf("[%d/%d] >>> Ejecutando ETL: %s\n", i+1, len(o.tasks), task.Name())
		
		if err := task.Run(ctx, db, cfg); err != nil {
			return fmt.Errorf("error en tarea '%s': %w", task.Name(), err)
		}
	}

	log.Println("Orquestación ETL finalizada exitosamente.")
	return nil
}
