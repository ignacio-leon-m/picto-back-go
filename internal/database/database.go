package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB es una estructura que contiene el pool de conexiones a la base de datos.
type DB struct {
	Pool *pgxpool.Pool
}

// Connect establece una conexión con la base de datos PostgreSQL.
// Implementa una lógica de reintentos para manejar el caso en que la aplicación
// inicie antes de que la base de datos esté lista para aceptar conexiones.
func Connect() (*DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL no está definida en las variables de entorno")
	}

	var pool *pgxpool.Pool
	var err error
	maxRetries := 5
	retryDelay := time.Second * 2

	for i := 0; i < maxRetries; i++ {
		// Creamos un contexto con timeout para este intento específico.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

		pool, err = pgxpool.New(ctx, dbURL)
		if err == nil {
			// Si la creación del pool fue exitosa, verificamos la conexión con un ping.
			if pingErr := pool.Ping(ctx); pingErr == nil {
				log.Println("¡Conexión a la base de datos establecida exitosamente!")
				cancel() // Cancelamos el contexto ya que no lo necesitamos más.
				return &DB{Pool: pool}, nil
			} else {
				// Si el ping falla, cerramos el pool recién creado para no dejarlo abierto.
				pool.Close()
				err = pingErr // Actualizamos el error con el del ping.
			}
		}

		cancel() // Es crucial cancelar el contexto al final de cada iteración.
		log.Printf("No se pudo conectar a la base de datos (intento %d/%d): %v. Reintentando en %v...", i+1, maxRetries, err, retryDelay)
		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("no se pudo conectar a la base de datos después de %d intentos: %w", maxRetries, err)
}

// Close cierra el pool de conexiones a la base de datos.
func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}

// Ping verifica que la conexión con la base de datos está activa.
func (db *DB) Ping(ctx context.Context) error {
	return db.Pool.Ping(ctx)
}
