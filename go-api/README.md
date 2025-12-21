# Go API - QR Challenge

API RESTful desarrollada con Go (Golang) y Fiber para procesamiento de matrices: validación, rotación 90° horario y factorización QR.

## 🚀 Características

- ✅ Validación de matrices rectangulares
- ✅ Rotación de matrices 90° en sentido horario
- ✅ Factorización QR usando librería `gonum`
- ✅ Comunicación HTTP con Node.js API
- ✅ Autenticación JWT
- ✅ Health checks
- ✅ Manejo robusto de errores y timeouts
- ✅ Tests unitarios e integración
- ✅ CORS configurado para frontend

## 📦 Instalación Local

### Requisitos

- Go 1.23.0 o superior
- Variables de entorno configuradas (archivo `.env` local **no versionado** o variables exportadas)

### Pasos

```bash
cd go-api

# Instalar dependencias
go mod download

# Configurar variables de entorno
# Crea un archivo `.env` (no versionar) con tus valores

# Ejecutar servidor
go run cmd/server/main.go
```

La API estará disponible en `http://localhost:3000`

## 🐳 Docker

### Desarrollo (con hot-reload)

```bash
docker-compose -f docker-compose.dev.yml up go-api
```

### Producción

```bash
docker-compose up go-api
```

## 📍 Endpoints

### `GET /` - Información del Servicio
Obtiene información sobre la API, versión, endpoints disponibles y estado del sistema.

**Autenticación:** No requerida

**Ejemplo:**
```bash
curl http://localhost:3000/
```

---

### `GET /health` - Health Check
Verifica que el servicio esté funcionando correctamente.

**Autenticación:** No requerida

**Ejemplo:**
```bash
curl http://localhost:3000/health
```

**Response:**
```json
{
  "status": "ok",
  "service": "go-api"
}
```

---

### `POST /auth/login` - Autenticación
Obtiene un token JWT para autenticar requests posteriores.

**Autenticación:** No requerida (endpoint público)

**Request:**
```json
{
  "username": "admin",
  "password": "admin"
}
```

**Ejemplo:**
```bash
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin"}'
```

**Response:**
```json
{
  "success": true,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "message": "Login exitoso",
  "expiresIn": "24h"
}
```

**Credenciales por defecto:**
- Usuario: `admin`
- Contraseña: `admin`

---

### `POST /matrix/process` - Procesar Matriz
Procesa una matriz: valida, rota 90° horario, calcula factorización QR y obtiene estadísticas de Node.js.

**Autenticación:** Requerida (JWT)

**Request:**
```json
{
  "matrix": [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
  ]
}
```

**Ejemplo:**
```bash
TOKEN="tu_token_jwt_aqui"

curl -X POST http://localhost:3000/matrix/process \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "matrix": [
      [1, 2, 3],
      [4, 5, 6],
      [7, 8, 9]
    ]
  }'
```

**Response:**
```json
{
  "rotated": [[7, 4, 1], [8, 5, 2], [9, 6, 3]],
  "q": [[-0.12, 0.90, 0.41], ...],
  "r": [[-8.12, -9.60, -11.08], ...],
  "nodeStats": {
    "max": 9.00,
    "min": -11.08,
    "avg": 0.68,
    "sum": 18.34,
    "anyDiagonal": false
  }
}
```

**Proceso interno:**
1. Valida que la matriz sea rectangular y numérica
2. Rota la matriz 90° en sentido horario
3. Calcula factorización QR de la matriz original
4. Envía Q, R y matriz rotada a Node.js API
5. Recibe estadísticas de Node.js
6. Retorna todo al cliente

---

## 🔧 Configuración

### Variables de Entorno

Crea un archivo `.env` en el directorio `go-api/`:

```env
PORT=3000
NODE_API_URL=http://localhost:3001
JWT_SECRET=REPLACE_ME_WITH_A_LONG_RANDOM_STRING
```

**Variables disponibles:**
- `PORT`: Puerto donde escucha el servidor (default: 3000)
- `NODE_API_URL`: URL de la API de Node.js (**obligatoria**)
- `JWT_SECRET`: Secreto para firmar tokens JWT

### Estructura del Proyecto

```
go-api/
├── cmd/
│   └── server/
│       └── main.go          # Punto de entrada
├── internal/
│   ├── controllers/          # Controladores (auth)
│   ├── handlers/             # Handlers HTTP (matrix)
│   ├── middleware/           # Middleware (JWT auth)
│   ├── models/               # Modelos de datos
│   └── services/             # Lógica de negocio
│       ├── validator.go      # Validación de matrices
│       ├── rotation.go       # Rotación 90° horario
│       ├── qr_decomposition.go  # Factorización QR
│       └── node_client.go    # Cliente HTTP para Node.js
├── Dockerfile                # Build producción
├── Dockerfile.dev            # Build desarrollo
├── go.mod                    # Dependencias
└── README.md                 # Este archivo
```

---

## 🧪 Testing

### Ejecutar Tests

```bash
# Todas las pruebas
go test ./...

# Con verbosidad
go test -v ./...

# Con cobertura
go test -cover ./...

# Generar reporte de cobertura
go test -coverprofile=coverage.out $(go list ./... | grep -v '/cmd/')
go tool cover -html=coverage.out
```

### Cobertura Actual

- `controllers`: 80.0%
- `middleware`: 100.0%
- `handlers`: 77.3%
- `services`: 62.9%

---

## 📚 Dependencias Principales

- **Fiber v2.52.10**: Framework web rápido y expresivo
- **gonum v0.16.0**: Librería para operaciones matriciales y factorización QR
- **golang-jwt/jwt/v5 v5.3.0**: Autenticación JWT
- **godotenv v1.5.1**: Carga de variables de entorno

---

## 🔍 Decisiones Técnicas

1. **Factorización QR sobre matriz original**: Se calcula sobre la matriz original (no rotada) para mantener la relación matemática estándar A = Q × R.

2. **Rotación 90° horario**: Implementada con fórmula matemática estándar: `rotated[j][rows-1-i] = matrix[i][j]`.

3. **Comunicación con Node.js**: Timeout de 10 segundos, manejo robusto de errores.

4. **JWT**: Solo Go API genera tokens. Node.js solo valida tokens recibidos.

---

## 🐛 Solución de Problemas

### Error: "cannot find package"

```bash
go mod download
go mod tidy
```

### Error: "port already in use"

Cambia el puerto en `.env` o detén el proceso que está usando el puerto.

### Error: "connection refused" al conectar con Node.js

Verifica que Node.js API esté corriendo en el puerto configurado en `NODE_API_URL`.

---

## 🎯 Versión

- **Go**: 1.23.0
- **Fiber**: v2.52.10
- **gonum**: v0.16.0

---

¡Listo para usar! 🚀

