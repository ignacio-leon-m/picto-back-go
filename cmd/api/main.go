package main // Declara que este es un paquete ejecutable (el punto de entrada de la aplicación).

import (
	"fmt"      // Para formateo de salida (ej. imprimir mensajes en consola)
	"log"      // Para registrar mensajes de error o información
	"os"       // Para interactuar con el sistema operativo (ej. obtener variables de entorno)

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

	// Crea una nueva instancia de la aplicación Fiber.
	app := fiber.New()

	// Define una ruta GET en el endpoint raíz ("/").
	// Cuando se acceda a http://localhost:8080/, este manejador se ejecutará.
	app.Get("/", func(c *fiber.Ctx) error {
		// Retorna una respuesta de texto "Hello, Fiber!".
		return c.SendString("Hello, Fiber!")
	})

	// Obtiene la URL de la base de datos de las variables de entorno.
	// Esto es para demostrar cómo acceder a ellas; no se usa en este "Hello World".
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("Advertencia: La variable de entorno DATABASE_URL no está definida.")
	} else {
		fmt.Printf("DATABASE_URL cargada: %s\n", dbURL)
	}

	// Inicia el servidor Fiber en el puerto 8080.
	// log.Fatal hará que el programa termine si hay un error al iniciar el servidor.
	log.Fatal(app.Listen(":8080"))
}