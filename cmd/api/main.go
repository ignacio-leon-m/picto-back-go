package main

import (
	"context"
	"fmt" // Para formateo de salida en consola
	"log" // registros
	"os"  // ej. obtener variables de entorno
	"time"

	"pictogram-project/internal/database"

	"github.com/gofiber/fiber/v2" // Importa el framework Fiber (versión 2)
	"github.com/joho/godotenv"    // Importa una librería para cargar variables de entorno desde un archivo .env (útil localmente)
)

func main() { // Función principal, el punto de inicio de la ejecución de la aplicación Go.
	// Carga variables de entorno desde un archivo .env si existe.
	// Esto es útil para el desarrollo local fuera de Docker, pero en Docker Compose las variables
	// se leen directamente del `environment` en docker-compose.yml.
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Advertencia: No se pudo cargar .env o el archivo no existe: %v", err)
	}

	// --- Conexión a la Base de Datos ---
	// Usamos nuestra función Connect del paquete database.
	// Esta función ya incluye la lógica de reintentos, crucial para Docker Compose.
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Fallo crítico: No se pudo conectar a la base de datos: %v", err)
	}
	// La conexión se cerrará cuando la aplicación termine. Para un cierre más elegante
	// (graceful shutdown), se implementarán manejadores de señales en fases posteriores.
	log.Println("Conexión a la base de datos verificada y lista.")

	// Crea una nueva instancia de la aplicación Fiber.
	app := fiber.New()

	// Define una ruta GET en el endpoint raíz ("/").
	// Cuando se acceda a http://localhost:8080/, este manejador se ejecutará.
	app.Get("/", func(c *fiber.Ctx) error {
		// Retorna una respuesta de texto simple.
		return c.SendString("¡Hola! Soy el backend de Pictogram en Go.")
	})

	// --- Endpoint de Health Check ---
	// Una ruta esencial para verificar que el servicio y sus dependencias (como la BD) funcionan.
	app.Get("/health", func(c *fiber.Ctx) error {
		// Creamos un contexto con un timeout corto para el ping.
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		// Usamos el método Ping de nuestra conexión para verificar el estado de la BD.
		if err := db.Ping(ctx); err != nil {
			// Si hay un error, devolvemos un estado de "servicio no disponible".
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"status":  "error",
				"message": "La base de datos no está disponible.",
			})
		}

		// Si todo está bien, devolvemos un OK.
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "Servicio funcionando y base de datos conectada.",
		})
	})

	// Inicia el servidor Fiber en el puerto 8080.
	// log.Fatal hará que el programa termine si hay un error al iniciar el servidor.
	fmt.Println("Servidor escuchando en el puerto 8080. Visita http://localhost:8080")
	fmt.Println("Health check disponible en http://localhost:8080/health")
	log.Fatal(app.Listen(":8080"))
}
