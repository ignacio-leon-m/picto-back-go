package models

import "time"

// Pictogram representa la estructura de datos para un pictograma en nuestra aplicación.
// Este modelo es una versión simplificada y adaptada del modelo de ARASAAC,
// conteniendo solo los campos relevantes para nuestra lógica de negocio.
type Pictogram struct {
	ID         int       `json:"id" db:"id"`                 // Mapeado desde _id de ARASAAC
	Name       string    `json:"name" db:"name"`             // El keyword principal
	Categories []string  `json:"categories" db:"categories"` // Lista de categorías para filtrado
	Tags       []string  `json:"tags" db:"tags"`             // Lista de tags para búsqueda
	ImageURL   string    `json:"image_url" db:"image_url"`   // URL a la imagen en el CDN de ARASAAC
	Schematic  bool      `json:"schematic" db:"schematic"`   // Flag booleano
	Violence   bool      `json:"violence" db:"violence"`     // Flag booleano para filtros de contenido
	Sex        bool      `json:"sex" db:"sex"`               // Flag booleano para filtros de contenido
	CreatedAt  time.Time `json:"created_at" db:"created_at"` // Mapeado desde created
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"` // Mapeado desde lastUpdated
}
