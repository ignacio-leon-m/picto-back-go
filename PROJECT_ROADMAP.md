# Project Roadmap: Pictogram Backend (Go & Docker)

Este documento es nuestra hoja de ruta para el desarrollo del backend del proyecto Pictogram. Lo usaremos para definir metas, seguir el progreso y asegurarnos de que construimos el proyecto de manera lógica y ordenada.

**Stack Tecnológico:**
*   **Lenguaje/Framework:** Go (con Fiber)
*   **Base de Datos:** PostgreSQL
*   **Entorno:** Contenedores Docker y Docker Compose

---

## ✅ Fase 0: Configuración del Entorno de Desarrollo y "Hello, World!"

*Estado: **Completado***

Hemos establecido una base sólida para el desarrollo, asegurando que el entorno sea reproducible, eficiente y esté listo para la codificación.

### Hitos Alcanzados:
-   [x] **Estructura del Proyecto:** Creada la estructura de directorios inicial (`cmd/api`).
-   [x] **Dockerfile de Producción:** Definido un `Dockerfile` multi-etapa para compilaciones optimizadas.
-   [x] **Dockerfile de Desarrollo:** Creado un `Dockerfile.dev` para un entorno de desarrollo con todas las herramientas necesarias.
-   [x] **Docker Compose:** Configurado `docker-compose.yml` para orquestar el servicio de backend.
-   [x] **Live Reloading:** Integrado `air` para la recarga automática del servidor al detectar cambios en el código, configurado a través de `.air.toml`.
-   [x] **Endpoint Básico:** La aplicación levanta un servidor web y responde en la ruta raíz (`/`).

---

## ⬜ Fase 1: Integración de la Base de Datos y Cimientos de la API

*Estado: **En Progreso***

El objetivo de esta fase es conectar nuestra aplicación Go con una base de datos PostgreSQL, todo orquestado por Docker Compose. Esto nos permitirá persistir y consultar datos.

### Tareas:
-   [x] **Añadir PostgreSQL a Docker Compose:** Actualizar `docker-compose.yml` para incluir un servicio de base de datos PostgreSQL.
-   [x] **Persistencia de Datos:** Configurar un volumen de Docker para que los datos de PostgreSQL no se pierdan al reiniciar los contenedores.
-   [x] **Conexión a la Base de Datos:** Escribir el código en Go para establecer y verificar una conexión con la base de datos PostgreSQL desde la aplicación Fiber.
-   [ ] **Estructura del Modelo:** Definir el `struct` de Go que representará un `Pictogram` en nuestro código.
-   [x] **Endpoint de "Health Check":** Crear una nueva ruta en la API (ej. `/health`) que verifique el estado de la conexión a la base de datos y devuelva un `OK`.

---

## ⬜ Fase 2: API Core - Gestión de Pictogramas (CRUD)

*Estado: **Pendiente***

Con la base de datos conectada, construiremos los endpoints fundamentales para gestionar los pictogramas.

### Tareas:
-   [ ] **Patrón Repositorio:** Implementar el patrón de diseño "Repository" para abstraer la lógica de acceso a datos.
-   [ ] **Crear Pictograma:** Implementar el endpoint `POST /pictograms`.
-   [ ] **Obtener un Pictograma:** Implementar el endpoint `GET /pictograms/:id`.
-   [ ] **Listar Pictogramas:** Implementar el endpoint `GET /pictograms`.
-   [ ] **Actualizar Pictograma:** Implementar el endpoint `PUT /pictograms/:id`.
-   [ ] **Eliminar Pictograma:** Implementar el endpoint `DELETE /pictograms/:id`.

---

## ⬜ Fase 3: Funcionalidades Avanzadas de la API

*Estado: **Pendiente***

Refinaremos el endpoint de listado para incluir las funcionalidades de filtrado y ordenamiento requeridas.

### Tareas:
-   [ ] **Paginación:** Añadir parámetros de consulta (`page`, `pageSize`) al endpoint `GET /pictograms`.
-   [ ] **Filtrado:** Permitir filtrar los resultados por nombre, categoría, etc.
-   [ ] **Ordenamiento:** Permitir ordenar los resultados por diferentes campos.

---

## ⬜ Fase 4: Generación de Documentos PDF

*Estado: **Pendiente***

Implementaremos una de las funcionalidades clave del proyecto: la creación de PDFs a partir de una selección de pictogramas.

### Tareas:
-   [ ] **Investigar Librería de PDF:** Seleccionar e integrar una librería de Go para la creación de PDFs (ej. `gofpdf`).
-   [ ] **Endpoint de Descarga:** Crear un endpoint (ej. `POST /pictograms/download`) que reciba una lista de IDs de pictogramas.
-   [ ] **Lógica de Generación:** Implementar la lógica para obtener los pictogramas de la base de datos y dibujarlos en un documento PDF con el formato especificado (dos por página, con imagen y nombre).
-   [ ] **Devolver el Archivo:** Configurar el endpoint para que devuelva el archivo PDF generado para su descarga en el navegador.