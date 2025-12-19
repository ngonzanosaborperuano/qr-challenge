# Dockerfile para API Go
# Etapa 1: Builder - Compilar la aplicación
FROM golang:1.23-alpine AS builder

# Instalar dependencias del sistema necesarias para compilación
RUN apk add --no-cache git

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos de dependencias primero (mejora cache de Docker)
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fuente
COPY . .

# Compilar la aplicación como binario estático
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main ./cmd/server

# Etapa 2: Producción - Imagen final mínima
FROM alpine:latest

# Instalar ca-certificates para HTTPS y wget para healthchecks
RUN apk --no-cache add ca-certificates wget

# Crear usuario no-root para mayor seguridad
RUN addgroup -g 1001 -S appuser && \
    adduser -S appuser -u 1001

WORKDIR /app

# Copiar el binario compilado desde builder
COPY --from=builder /app/main .

# Cambiar ownership al usuario no-root
RUN chown -R appuser:appuser /app

# Cambiar a usuario no-root
USER appuser

# Exponer puerto
EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:3000/health || exit 1

# Comando para ejecutar la aplicación
CMD ["./main"]

