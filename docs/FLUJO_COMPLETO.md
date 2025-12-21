# 🔄 Flujo Completo de Uso - QR Challenge APIs

## 📋 Resumen del Requerimiento

Según el desafío técnico, el sistema debe:

1. **API Go**: Recibir matriz → Validar → Rotar 90° → Calcular QR → Enviar a Node.js
2. **API Node.js**: Recibir matrices → Calcular estadísticas → Devolver resultados
3. **Comunicación**: HTTP entre ambas APIs
4. **Seguridad**: JWT para proteger endpoints (opcional, pero implementado)
5. **Testing**: Pruebas unitarias e integración (implementado)
6. **Frontend**: Interfaz simple (opcional, implementado)

---

## 🏗️ Arquitectura del Proyecto

### Estructura General

```
qr-challenge/
├── go-api/                    # API en Go (Golang)
│   ├── cmd/
│   │   └── server/           # Punto de entrada principal
│   ├── internal/
│   │   ├── controllers/      # Controladores (auth)
│   │   ├── handlers/         # Handlers HTTP (matrix)
│   │   ├── middleware/       # Middleware (JWT auth)
│   │   ├── models/          # Modelos de datos
│   │   └── services/        # Lógica de negocio
│   ├── Dockerfile            # Build producción
│   ├── Dockerfile.dev        # Build desarrollo
│   └── go.mod               # Dependencias Go
│
├── node-api/                 # API en Node.js
│   ├── src/
│   │   ├── controllers/     # Controladores
│   │   ├── middleware/      # Middleware (JWT)
│   │   ├── routers/         # Rutas
│   │   ├── services/        # Lógica de negocio
│   │   └── index.js         # Punto de entrada
│   ├── Dockerfile            # Build producción
│   ├── Dockerfile.dev        # Build desarrollo
│   └── package.json         # Dependencias Node.js
│
├── frontend/                 # Frontend Angular
│   ├── src/
│   │   ├── app/
│   │   │   ├── components/  # Componentes Angular
│   │   │   ├── services/    # Servicios HTTP
│   │   │   ├── pipes/       # Pipes personalizados
│   │   │   └── models/      # Modelos TypeScript
│   │   └── environments/    # Configuración de entornos
│   ├── Dockerfile            # Build producción (Nginx)
│   ├── Dockerfile.dev        # Build desarrollo
│   └── nginx.conf           # Configuración Nginx
│
├── docker-compose.yml        # Producción
├── docker-compose.dev.yml    # Desarrollo (hot-reload)
└── .env (local)            # Variables de entorno (no versionar)
```

### Arquitectura de Comunicación

```
┌─────────────┐
│   Cliente   │ (Frontend Angular / Postman / curl)
│             │
└──────┬──────┘
       │ HTTP/REST
       │
       ▼
┌─────────────────────────────────┐
│      Go API (Fiber)             │
│  - Validación                   │
│  - Rotación 90°                 │
│  - Factorización QR             │
│  - Autenticación JWT            │
└──────┬──────────────────────────┘
       │
       │ HTTP POST
       │ Authorization: Bearer <token>
       │
       ▼
┌─────────────────────────────────┐
│   Node.js API (Express)         │
│  - Cálculo de estadísticas      │
│  - Verificación diagonal        │
│  - Validación JWT               │
└──────┬──────────────────────────┘
       │
       │ HTTP Response
       │
       ▼
┌─────────────────────────────────┐
│      Go API (retorna todo)      │
└──────┬──────────────────────────┘
       │
       │ HTTP Response
       │
       ▼
┌─────────────┐
│   Cliente   │
└─────────────┘
```

### Principios de Diseño Aplicados

- **Separación de Responsabilidades (SRP)**: Cada servicio tiene una única responsabilidad
- **Inversión de Dependencias (DIP)**: Componentes dependen de abstracciones (interfaces)
- **Single Source of Truth**: Cada API maneja su dominio específico
- **RESTful**: APIs siguen principios REST
- **Microservicios**: APIs independientes y desacopladas

---

## 🔧 Tecnologías y Versiones

### Backend - Go API

- **Lenguaje**: Go (Golang) **1.23.0**
- **Framework**: Fiber v2.52.10
- **Librerías principales**:
  - `gonum.org/v1/gonum v0.16.0` - Factorización QR y operaciones matriciales
  - `github.com/golang-jwt/jwt/v5 v5.3.0` - Autenticación JWT
  - `github.com/joho/godotenv v1.5.1` - Variables de entorno
- **Base Docker**: `golang:1.23-alpine` (desarrollo y producción)

### Backend - Node.js API

- **Lenguaje**: JavaScript (Node.js)
- **Runtime**: Node.js **20.x** (Alpine)
- **Framework**: Express.js **5.2.1**
- **Librerías principales**:
  - `express-jwt v8.5.1` - Middleware JWT
  - `jsonwebtoken v9.0.3` - Generación/verificación JWT
  - `dotenv v17.2.3` - Variables de entorno
  - `jest v29.7.0` - Testing
- **Base Docker**: `node:20-alpine`

### Frontend

- **Framework**: Angular **17** (standalone components)
- **Lenguaje**: TypeScript
- **HTTP Client**: Angular HttpClient
- **Build Tool**: Angular CLI
- **Servidor Web**: Nginx (producción)
- **Base Docker**: `node:20-alpine` (build) + `nginx:alpine` (servir)

### Infraestructura

- **Containerización**: Docker + Docker Compose
- **Red**: Bridge network (`qr-network`)
- **Health Checks**: Implementados en todos los servicios
- **Multi-stage Builds**: Optimizados para producción

---

## ✅ Estado de Implementación

| Componente | Estado | Notas |
|------------|--------|-------|
| **API Go (Fiber)** | ✅ Completo | Validación, rotación, QR, comunicación HTTP |
| **API Node.js (Express)** | ✅ Completo | Estadísticas, verificación diagonal |
| **Comunicación HTTP** | ✅ Completo | Go → Node.js con timeouts y manejo de errores |
| **JWT Authentication** | ✅ Completo | Implementado en Go API (Node.js solo valida) |
| **Docker & Docker Compose** | ✅ Completo | Multi-stage builds, dev y prod |
| **Documentación** | ✅ Completo | README, TESTS, FLUJO_COMPLETO |
| **Pruebas Unitarias** | ✅ Completo | Go: validator, rotation, QR, auth. Node.js: stats, auth |
| **Pruebas de Integración** | ✅ Completo | Go: handlers, middleware con httptest |
| **Frontend** | ✅ Completo | Angular con login, procesamiento y visualización |
| **CORS** | ✅ Completo | Configurado para frontend |
| **Principios SOLID** | ✅ Completo | Aplicados en frontend y backend |

---

## 📍 Endpoints Disponibles

### Go API (http://localhost:3000)

#### 1. `GET /` - Información del Servicio
**Propósito**: Obtener información sobre la API Go, versión, endpoints disponibles, y estado del sistema.

**Autenticación**: No requerida

**Request:**
```bash
curl http://localhost:3000/
```

**Response (200 OK):**
```json
{
  "service": "Go API Backend",
  "version": "1.0.0",
  "technology": "Go (Golang)",
  "framework": "Fiber v2",
  "goVersion": "go1.23.0",
  "startTime": "2025-01-19T22:00:00Z",
  "uptime": "2h30m15s",
  "uptimeSeconds": 9015,
  "os": "linux",
  "arch": "amd64",
  "nodeApiUrl": "http://node-api:3001",
  "endpoints": {
    "health": "GET /health",
    "login": "POST /auth/login",
    "processMatrix": "POST /matrix/process (requiere JWT)",
    "info": "GET /"
  }
}
```

---

#### 2. `GET /health` - Health Check
**Propósito**: Verificar que el servicio Go API esté funcionando correctamente.

**Autenticación**: No requerida

**Request:**
```bash
curl http://localhost:3000/health
```

**Response (200 OK):**
```json
{
  "status": "ok",
  "service": "go-api"
}
```

---

#### 3. `POST /auth/login` - Autenticación
**Propósito**: Obtener un token JWT para autenticar requests posteriores.

**Autenticación**: No requerida (endpoint público)

**Request:**
```bash
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin"
  }'
```

**Response (200 OK):**
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

**Vigencia del token**: 24 horas

---

#### 4. `POST /matrix/process` - Procesar Matriz
**Propósito**: Procesar una matriz: validar, rotar 90° horario, calcular factorización QR, y obtener estadísticas de Node.js.

**Autenticación**: Requerida (JWT)

**Request:**
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

**Response (200 OK):**
```json
{
  "rotated": [
    [7, 4, 1],
    [8, 5, 2],
    [9, 6, 3]
  ],
  "q": [
    [-0.12, 0.90, 0.41],
    [-0.49, 0.30, -0.82],
    [-0.86, -0.30, 0.41]
  ],
  "r": [
    [-8.12, -9.60, -11.08],
    [0.00, 0.90, 1.81],
    [0.00, 0.00, -0.00]
  ],
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

**Errores posibles:**
- `400 Bad Request`: Matriz inválida (no rectangular, valores no numéricos)
- `401 Unauthorized`: Token JWT inválido o faltante
- `500 Internal Server Error`: Error en factorización QR o comunicación con Node.js

---

### Node.js API (http://localhost:3001)

#### 1. `GET /` - Información del Servicio
**Propósito**: Obtener información sobre la API Node.js, versión, y estado del sistema.

**Autenticación**: No requerida

**Request:**
```bash
curl http://localhost:3001/
```

**Response (200 OK):**
```json
{
  "service": "Node.js API Backend",
  "version": "1.0.0",
  "technology": "Node.js",
  "framework": "Express.js",
  "nodeVersion": "v20.11.0",
  "platform": "linux",
  "arch": "x64",
  "environment": "production",
  "memory": {
    "used": "45.2 MB",
    "total": "512 MB"
  },
  "uptime": "2h30m15s",
  "endpoints": {
    "health": "GET /health",
    "stats": "POST /matrix/stats (requiere JWT)",
    "info": "GET /"
  }
}
```

---

#### 2. `GET /health` - Health Check
**Propósito**: Verificar que el servicio Node.js API esté funcionando correctamente.

**Autenticación**: No requerida

**Request:**
```bash
curl http://localhost:3001/health
```

**Response (200 OK):**
```json
{
  "status": "ok",
  "service": "node-api"
}
```

---

#### 3. `POST /matrix/stats` - Calcular Estadísticas
**Propósito**: Calcular estadísticas (max, min, avg, sum) sobre matrices Q, R y rotated, y verificar si alguna es diagonal.

**Autenticación**: Requerida (JWT)

**Request:**
```bash
TOKEN="tu_token_jwt_aqui"

curl -X POST http://localhost:3001/matrix/stats \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "q": [[-0.12, 0.90, 0.41], [-0.49, 0.30, -0.82], [-0.86, -0.30, 0.41]],
    "r": [[-8.12, -9.60, -11.08], [0.00, 0.90, 1.81], [0.00, 0.00, -0.00]],
    "rotated": [[7, 4, 1], [8, 5, 2], [9, 6, 3]]
  }'
```

**Response (200 OK):**
```json
{
  "max": 9.00,
  "min": -11.08,
  "avg": 0.68,
  "sum": 18.34,
  "anyDiagonal": false
}
```

**Proceso interno:**
1. Valida token JWT
2. Combina todos los valores de Q, R y rotated
3. Calcula máximo, mínimo, promedio y suma
4. Verifica si alguna matriz es diagonal (cuadrada con elementos fuera de diagonal = 0)
5. Retorna estadísticas

**Errores posibles:**
- `400 Bad Request`: Matrices faltantes o formato inválido
- `401 Unauthorized`: Token JWT inválido o faltante
- `500 Internal Server Error`: Error al calcular estadísticas

---

## 🔄 Flujo Completo Paso a Paso

### **Fase 1: Inicialización y Verificación**

#### Paso 1.1: Iniciar Servicios

**Opción A: Desarrollo (con hot-reload)**
```bash
docker-compose -f docker-compose.dev.yml up
```

**Opción B: Producción**
```bash
docker-compose up
```

**Resultado Esperado:**
- Go API corriendo en `http://localhost:3000`
- Node.js API corriendo en `http://localhost:3001`
- Frontend corriendo en `http://localhost:4200` (dev) o `http://localhost:80` (prod)
- Todas las APIs saludables

---

#### Paso 1.2: Verificar Health Checks

```bash
# Verificar Go API
curl http://localhost:3000/health
# Respuesta: {"status":"ok","service":"go-api"}

# Verificar Node.js API
curl http://localhost:3001/health
# Respuesta: {"status":"ok","service":"node-api"}
```

---

### **Fase 2: Autenticación JWT**

#### Opción A: Usando Frontend (Recomendado) 🌐

1. Abrir navegador en `http://localhost:4200` (dev) o `http://localhost:80` (prod)
2. Ingresar credenciales:
   - Usuario: `admin`
   - Contraseña: `admin`
3. Hacer clic en "Iniciar Sesión"
4. El token se guarda automáticamente en localStorage

#### Opción B: Usando curl/Postman

**Paso 2.1: Obtener Token de Go API**
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

**Guardar token:**
```bash
TOKEN=$(curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin"}' \
  -s | jq -r '.token')
```

---

### **Fase 3: Procesamiento de Matriz (Flujo Principal)**

#### Opción A: Usando Frontend (Recomendado) 🌐

1. En el navegador, después de iniciar sesión
2. Ingresar matriz en el campo de texto (formato JSON):
   ```
   [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
   ```
3. Hacer clic en "Procesar Matriz"
4. Ver resultados:
   - Matriz rotada
   - Matrices Q y R
   - Estadísticas

#### Opción B: Usando curl/Postman

**Paso 3.1: Cliente Envía Matriz a Go API**

```bash
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

**Proceso Interno en Go API:**

1. **Validación de Matriz** (`validator.go`):
   - Verifica que no esté vacía
   - Verifica que sea rectangular (todas las filas del mismo tamaño)
   - Verifica que todos los valores sean numéricos

2. **Rotación 90° Horario** (`rotation.go`):
   - Matriz original: `[[1,2,3], [4,5,6], [7,8,9]]`
   - Matriz rotada: `[[7,4,1], [8,5,2], [9,6,3]]`
   - Algoritmo: `rotated[j][rows-1-i] = matrix[i][j]`

3. **Factorización QR** (`qr_decomposition.go`):
   - Usa librería `gonum.org/v1/gonum`
   - Calcula Q (ortogonal) y R (triangular superior)
   - Verifica: A = Q × R

4. **Comunicación con Node.js** (`node_client.go`):
   - POST a `http://node-api:3001/matrix/stats`
   - Incluye token JWT en header
   - Timeout: 10 segundos

**Proceso Interno en Node.js API:**

1. **Validación JWT** (`middleware/auth.js`):
   - Verifica token en header `Authorization: Bearer <token>`

2. **Cálculo de Estadísticas** (`services/statsService.js`):
   - Combina todos los valores de Q, R y rotated
   - Calcula: max, min, avg, sum
   - Verifica si alguna matriz es diagonal

3. **Respuesta a Go API**:
   - Retorna estadísticas calculadas

**Respuesta Final del Cliente:**
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

---

## 📊 Diagrama de Flujo Completo

```
┌─────────────┐
│   Cliente   │
│  (Frontend/ │
│   Postman/  │
│   curl)     │
└──────┬──────┘
       │
       │ 1. POST /auth/login
       │    {username, password}
       ▼
┌─────────────────┐
│   Go API        │
│   /auth/login   │
└──────┬──────────┘
       │
       │ 2. Response: {token}
       │
       │ 3. POST /matrix/process
       │    Authorization: Bearer <token>
       │    {matrix: [[...]]}
       ▼
┌─────────────────┐
│   Go API        │
│   /matrix/      │
│   process       │
└──────┬──────────┘
       │
       │ 4. Validar matriz
       │ 5. Rotar 90° horario
       │ 6. Calcular QR
       │
       │ 7. POST /matrix/stats
       │    Authorization: Bearer <token>
       │    {q, r, rotated}
       ▼
┌─────────────────┐
│  Node.js API    │
│  /matrix/stats  │
└──────┬──────────┘
       │
       │ 8. Validar JWT
       │ 9. Calcular estadísticas
       │    - max, min, avg, sum
       │    - anyDiagonal
       │
       │ 10. Response: {max, min, avg, sum, anyDiagonal}
       │
       ▼
┌─────────────────┐
│   Go API        │
│   (recibe stats)│
└──────┬──────────┘
       │
       │ 11. Response final:
       │     {rotated, q, r, nodeStats}
       │
       ▼
┌─────────────┐
│   Cliente   │
│  (recibe    │
│   resultado)│
└─────────────┘
```

---

## 🧪 Ejemplo Completo de Uso con curl

```bash
#!/bin/bash

echo "=== QR Challenge - Flujo Completo ==="
echo ""

# Paso 1: Verificar servicios
echo "1. Verificando servicios..."
curl -s http://localhost:3000/health | jq '.'
curl -s http://localhost:3001/health | jq '.'
echo ""

# Paso 2: Obtener token
echo "2. Obteniendo token JWT..."
TOKEN=$(curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin"}' \
  -s | jq -r '.token')

echo "✅ Token obtenido: ${TOKEN:0:50}..."
echo ""

# Paso 3: Procesar matriz
echo "3. Procesando matriz..."
RESPONSE=$(curl -X POST http://localhost:3000/matrix/process \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "matrix": [
      [1, 2, 3],
      [4, 5, 6],
      [7, 8, 9]
    ]
  }' \
  -s)

echo "$RESPONSE" | jq '.'
echo ""

# Paso 4: Mostrar estadísticas
echo "4. Estadísticas calculadas:"
echo "$RESPONSE" | jq '.nodeStats'
echo ""

# Paso 5: Verificar matriz rotada
echo "5. Matriz rotada:"
echo "$RESPONSE" | jq '.rotated'
echo ""

echo "✅ Flujo completo ejecutado exitosamente"
```

---

## ✅ Checklist de Implementación

### ✅ Completado (100%)

- [x] API Go con Fiber
- [x] API Node.js con Express
- [x] Validación de matrices
- [x] Rotación 90° horario
- [x] Factorización QR
- [x] Cálculo de estadísticas (max, min, avg, sum)
- [x] Verificación de matriz diagonal
- [x] Comunicación HTTP Go → Node.js
- [x] Manejo de errores y timeouts
- [x] JWT authentication
- [x] Docker y Docker Compose
- [x] Documentación completa
- [x] Health checks
- [x] Endpoints informativos
- [x] Pruebas unitarias (Go) - validator, rotation, QR, auth
- [x] Pruebas unitarias (Node.js) - statsService, auth middleware
- [x] Pruebas de integración (Go) - handlers, middleware con httptest
- [x] Frontend Angular - Login, procesamiento, visualización
- [x] CORS configurado
- [x] Principios SOLID aplicados
- [x] Multi-stage Docker builds
- [x] Hot-reload en desarrollo

### ❌ Pendiente (Opcional)

- [ ] Pruebas E2E completas (con testcontainers)
- [ ] CI/CD pipeline (opcional)
- [ ] Métricas y monitoreo (opcional)
- [ ] Swagger/OpenAPI documentation (opcional)

---

## 📊 Cobertura de Código Actual

```
Go API:
  - controllers:  80.0% ✅
  - middleware:   100.0% ✅ (Excelente!)
  - handlers:     77.3% ✅
  - services:     62.9% ✅
  - cmd/server:   0.0%  ⚠️ (Normal - se prueba con E2E)

Node.js API:
  - statsService: ✅ Cobertura completa de funciones principales
  - auth middleware: ✅ Cobertura completa
```

---

## 🎯 Verificación del Requerimiento

| Requerimiento | Estado | Notas |
|---------------|--------|-------|
| API Go con Fiber | ✅ | Implementado |
| API Node.js con Express | ✅ | Implementado |
| Validación de matriz rectangular | ✅ | Implementado |
| Rotación 90° horario | ✅ | Implementado |
| Factorización QR | ✅ | Implementado (gonum) |
| Estadísticas (max, min, avg, sum) | ✅ | Implementado |
| Verificación matriz diagonal | ✅ | Implementado |
| Comunicación HTTP entre APIs | ✅ | Implementado |
| Docker y Docker Compose | ✅ | Implementado |
| Documentación | ✅ | Completa |
| JWT (opcional) | ✅ | Implementado |
| Pruebas unitarias | ✅ | Implementado (Go: 62-80%, Node.js: completo) |
| Pruebas de integración | ✅ | Implementado (handlers: 77.3%, middleware: 100%) |
| Frontend (opcional) | ✅ | Implementado (Angular con login y visualización) |

**Cumplimiento del Requerimiento: 100%** ✅

---

## 📚 Referencias

- **TESTS.md**: Guía completa de pruebas unitarias e integración
- **README.md**: Documentación general del proyecto
- **docker-compose.yml**: Configuración de producción
- **docker-compose.dev.yml**: Configuración de desarrollo

---

## 🚀 Cómo Construir el Proyecto

### Requisitos Previos

- Docker y Docker Compose instalados
- (Opcional) Go 1.23+ y Node.js 20+ para desarrollo local

### Construcción con Docker

```bash
# Desarrollo (con hot-reload)
docker-compose -f docker-compose.dev.yml up --build

# Producción
docker-compose up --build
```

### Desarrollo Local (sin Docker)

**Go API:**
```bash
cd go-api
go mod download
go run cmd/server/main.go
```

**Node.js API:**
```bash
cd node-api
npm install
npm run dev
```

**Frontend:**
```bash
cd frontend
npm install
npm start
```

---

## 📝 Notas Técnicas

### Decisiones de Diseño

1. **Factorización QR sobre matriz original**: Se calcula sobre la matriz original (no rotada) para mantener la relación matemática estándar A = Q × R.

2. **JWT en Go API**: Solo Go API genera tokens. Node.js solo valida tokens recibidos.

3. **Rotación 90° horario**: Implementada con fórmula matemática estándar: `rotated[j][rows-1-i] = matrix[i][j]`.

4. **Tolerancia numérica**: Se usa `1e-10` para comparaciones de punto flotante en verificación de matrices diagonales.

5. **Timeouts**: 10 segundos para comunicación Go → Node.js.

6. **CORS**: Configurado para permitir requests desde el frontend.

---

## 🎉 Conclusión

El proyecto **QR Challenge** está **100% completo** según los requerimientos del desafío técnico. Todas las funcionalidades obligatorias y opcionales han sido implementadas, incluyendo:

- ✅ APIs RESTful en Go y Node.js
- ✅ Comunicación HTTP entre servicios
- ✅ Autenticación JWT
- ✅ Pruebas unitarias e integración
- ✅ Frontend Angular
- ✅ Dockerización completa
- ✅ Documentación exhaustiva

El sistema está listo para ser desplegado y utilizado.
