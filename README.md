# 🚀 ETL-Electro-hogar

Este proyecto es una solución integral de **ETL (Extract, Transform, Load)** diseñada para la migración y limpieza de datos de clientes desde archivos planos (CSV) hacia una base de datos relacional PostgreSQL.

Desarrollado con un enfoque en **Arquitectura Limpia (Clean Architecture)**, el sistema garantiza escalabilidad, mantenibilidad y un código altamente organizado y profesional.

---

## 👤 Autor
**Nicolas Parada Cuervo**  
*TECNOLOGÍAS ETL - Universidad Distrital Francisco José de Caldas*

---

## 🛠️ Arquitectura y Funcionamiento
El flujo de datos sigue el estándar industrial de tres fases:
1.  **Extracción (Extract)**: Lectura robusta del archivo CSV localizado en `data/`, manejando delimitadores personalizados y errores de lectura.
2.  **Transformación (Transform)**: Limpieza profunda de datos que incluye:
    *   Eliminación de espacios en blanco redundantes.
    *   Normalización de nombres de ciudades (Capitalización estándar).
    *   Normalización de género y manejo de valores nulos o inconsistentes.
3.  **Carga (Load)**: Persistencia en la base de datos utilizando **Procedimientos Almacenados** para delegar la lógica de integridad y generación de identificadores únicos (UUID) al motor de base de datos.

---

## 📋 Requisitos Previos (Base de Datos)
Antes de ejecutar la aplicación, **es obligatorio** preparar el esquema de la base de datos ejecutando el siguiente script SQL. Este script crea el procedimiento almacenado necesario para la inserción de clientes.

> [!IMPORTANT]
> **Scripts de Base de Datos**:
> *   [Sc-crear_tablas.sql](internal/shared/infrastructure/db/sql/Sc-crear_tablas.sql)
> *   [Pr-insertar_cliente.sql](internal/shared/infrastructure/db/sql/Pr-insertar_cliente.sql)
>
> **Permisos de Ejecución**:
> Para que la aplicación pueda consumir el procedimiento, debes otorgar permisos de ejecución con el siguiente comando:
> ```sql
> GRANT EXECUTE ON PROCEDURE <DB_SCHEMA>.insertar_cliente(
>     VARCHAR, VARCHAR, VARCHAR, VARCHAR, VARCHAR, CHAR, SMALLINT
> ) TO <DB_SCHEMA>;
> ```

---

## ⚙️ Configuración (.env)
La aplicación se configura mediante un archivo de variables de entorno `.dev.env`. Asegúrate de definir los siguientes parámetros:

| Variable | Descripción | Ejemplo |
| :--- | :--- | :--- |
| `TIMEOUT` | Tiempo límite de ejecución (ms) | `20000` |
| `DB_HOST` | Dirección del servidor DB | `localhost` |
| `DB_PORT` | Puerto de conexión | `5432` |
| `DB_USER` | Usuario de base de datos | `postgres` |
| `DB_PASSWORD` | Contraseña del usuario | `******` |
| `DB_NAME` | Nombre de la base de datos | `electro_hogar` |
| `DB_SCHEMA` | Esquema del almacén de datos | `electro_hogar` |

---

## 🚀 Cómo Ejecutar

1.  **Instalar dependencias**:
    ```bash
    go mod tidy
    ```

2.  **Ejecutar la aplicación**:
    Asegúrate de que la variable de entorno `ENV_FILE` apunte a tu archivo `.dev.env`.
    ```bash
    $env:ENV_FILE=".dev.env"; go run cmd/main.go
    ```

---

## 📁 Estructura del Proyecto
*   `cmd/`: Punto de entrada de la aplicación.
*   `internal/customer/domain/`: Reglas de negocio y entidades de cliente.
*   `internal/customer/application/`: Casos de uso (Orquestación del ETL).
*   `internal/customer/infrastructure/`: Implementaciones técnicas (SQL Repo, CSV Parser).
*   `internal/shared/infrastructure/db/`: Configuración global y scripts de base de datos.
*   `data/`: Contiene los archivos fuente para el proceso de extracción.

---

> [!TIP]
> El sistema está diseñado para ser **funcionalmente escalable**. Si necesitas importar un nuevo tipo de datos (ej. proveedores), simplemente añade un nuevo módulo en `internal/` siguiendo el patrón establecido.