# --- Etapa de Compilación (Build Stage) ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copia los archivos de módulos primero para aprovechar el caché de Docker.
# La capa de 'go mod download' solo se reconstruirá si go.mod o go.sum cambian.
COPY go.mod go.sum ./

# Descarga todas las dependencias del módulo Go y asegura que go.mod/go.sum estén correctos.
# go mod download: Descarga las dependencias.
# go mod tidy: Limpia y añade dependencias según las importaciones en el código.
RUN go mod download
RUN go mod tidy

# Ahora copia el resto del código fuente de la aplicación.
COPY . .

# Compila tu aplicación Go.
# CGO_ENABLED=0: Deshabilita CGO para compilaciones estáticas.
# GOOS=linux: Especifica el sistema operativo objetivo como Linux.
# -a -installsuffix nocgo: Asegura una compilación estática.
# -ldflags "-s -w": Reduce el tamaño del binario eliminando símbolos de depuración.
# -o /app/backend: Especifica que el binario compilado se llamará 'backend' y se guardará en /app/.
# ./cmd/api/main.go: Ruta al archivo principal de tu aplicación Go.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -ldflags "-s -w" -o /app/backend ./cmd/api


# --- Etapa Final (Run Stage) ---
FROM alpine:latest

WORKDIR /app

# Copia solo el binario compilado de la etapa de compilación.
COPY --from=builder /app/backend .

# Expone el puerto en el que escucha tu aplicación.
EXPOSE 8080

# Comando para ejecutar la aplicación cuando el contenedor se inicie.
CMD ["./backend"]