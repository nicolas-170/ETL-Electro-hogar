package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	customerInfra "github.com/nicolas-170/ETL-Electro-hogar/internal/customer/infrastructure"
	timeInfra "github.com/nicolas-170/ETL-Electro-hogar/internal/time/infrastructure"
	productInfra "github.com/nicolas-170/ETL-Electro-hogar/internal/product/infrastructure"
	storeInfra "github.com/nicolas-170/ETL-Electro-hogar/internal/store/infrastructure"
	saleInfra "github.com/nicolas-170/ETL-Electro-hogar/internal/sale/infrastructure"
	database "github.com/nicolas-170/ETL-Electro-hogar/internal/shared/infrastructure/db"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/shared/infrastructure/etl"
)

func main() {
	printInfoInit()

	// Carga de configuración
	printInfoConfig()
	cfg := config.LoadConfig()

	// Conexión a la Base de Datos
	db, err := database.NewDB(cfg.Database)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Contexto global con tiempo límite
	ctx, cancel := context.WithTimeout(context.Background(), *cfg.App.Timeout)
	defer cancel()

	// Ejecutar orquestación de procesos ETL
	if err := runETLProcess(ctx, db, &cfg.Database); err != nil {
		log.Fatalf("Error crítico en la ejecución del proceso ETL: %v", err)
	}

	printInfoETL()
}

// runETLProcess centraliza la ejecución de todos los módulos ETL registrados
func runETLProcess(ctx context.Context, db *sql.DB, cfg *config.Database) error {
	orchestrator := etl.NewOrchestrator()

	// Registro de procesos modulares
	orchestrator.Register(&customerInfra.CustomerTask{})
	orchestrator.Register(&timeInfra.TimeTask{})
	orchestrator.Register(&productInfra.ProductTask{})
	orchestrator.Register(&storeInfra.StoreTask{})
	orchestrator.Register(&saleInfra.SaleTask{})

	// Ejecución de todas las tareas
	return orchestrator.RunAll(ctx, db, cfg)
}

func printInfoInit() {
	print("Iniciando ETL...")
	print("Proyecto ETL-Electro-hogar, creado por: Nicolas Parada Cuervo")
}

func printInfoConfig() {
	print("Cargando configuración...")
}

func printInfoETL() {
	print("Proceso finalizado exitosamente.")
}

func print(str string) {
	fmt.Println(str)
}
