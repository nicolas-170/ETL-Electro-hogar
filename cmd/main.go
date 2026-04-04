package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	customerInfra "github.com/nicolas-170/ETL-Electro-hogar/internal/customer/infrastructure"
	timeInfra "github.com/nicolas-170/ETL-Electro-hogar/internal/time/infrastructure"
	database "github.com/nicolas-170/ETL-Electro-hogar/internal/shared/infrastructure/db"
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

	// 1. Ejecutar proceso ETL del modulo Cliente
	log.Println(">>> Ejecutando ETL de Clientes...")
	if err := customerInfra.RunCustomerProcess(ctx, db, &cfg.Database); err != nil {
		log.Fatalf("Falla crítica en el proceso ETL de Clientes: %v", err)
	}

	// 2. Ejecutar proceso ETL del modulo Tiempo
	log.Println(">>> Ejecutando ETL de Tiempo...")
	if err := timeInfra.RunTimeProcess(ctx, db, &cfg.Database); err != nil {
		log.Fatalf("Falla crítica en el proceso ETL de Tiempo: %v", err)
	}

	printInfoETL()
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
