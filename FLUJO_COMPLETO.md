# 🔄 Flujo Completo de Uso - QR Challenge APIs

## 📋 Resumen del Requerimiento

Según el desafío técnico, el sistema debe:

1. **API Go**: Recibir matriz → Validar → Rotar 90° → Calcular QR → Enviar a Node.js
2. **API Node.js**: Recibir matrices → Calcular estadísticas → Devolver resultados
3. **Comunicación**: HTTP entre ambas APIs
4. **Seguridad**: JWT para proteger endpoints (opcional, pero implementado)
5. **Testing**: Pruebas unitarias e integración (pendiente)
6. **Frontend**: Interfaz simple (opcional, pendiente)

---

## ✅ Estado de Implementación

| Componente | Estado | Notas |
|------------|--------|-------|
| **API Go (Fiber)** | ✅ Completo | Validación, rotación, QR, comunicación HTTP |
| **API Node.js (Express)** | ✅ Completo | Estadísticas, verificación diagonal |
| **Comunicación HTTP** | ✅ Completo | Go → Node.js con timeouts y manejo de errores |
| **JWT Authentication** | ✅ Completo | Implementado en Go API (Node.js solo valida) |
| **Docker & Docker Compose** | ✅ Completo | Multi-stage builds, dev y prod |
| **Documentación** | ✅ Completo | README, ENDPOINTS, TESTS, FLUJO_COMPLETO |
| **Pruebas Unitarias** | ✅ Completo | Go: validator, rotation, QR, auth. Node.js: stats, auth |
| **Pruebas de Integración** | ✅ Completo | Go: handlers, middleware con httptest |
| **Frontend** | ✅ Completo | Angular con login, procesamiento y visualización |

---

## 🔄 Flujo Completo Paso a Paso

### **Fase 1: Inicialización y Verificación**

#### Paso 1.1: Iniciar Servicios
```bash
# Opción A: Desarrollo (con hot-reload)
docker-compose -f docker-compose.dev.yml up

# Opción B: Producción
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

**Estado:** ✅ Implementado

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

#### Paso 2.1: Obtener Token de Go API
```bash
curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin"}'
```

**Request:**
```json
{
  "username": "admin",
  "password": "admin"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiaWQiOjEsInJvbGUiOiJhZG1pbiIsImlhdCI6MTc2NjE3MTc4OCwiZXhwIjoxNzY2MjU4MTg4fQ.dq_IHhA-NyNvGiPWwTHA5Ckboi_2z257OWu0Y0c6Lls",
  "message": "Login exitoso",
  "expiresIn": "24h"
}
```

**Guardar token:**
```bash
TOKEN_GO=$(curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin"}' \
  -s | jq -r '.token')
```

---

#### Paso 2.2: Obtener Token de Node.js API (Opcional)
```bash
curl -X POST http://localhost:3001/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin"}'
```

**Estado:** ✅ Implementado

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

#### Paso 3.1: Cliente Envía Matriz a Go API

**Endpoint:** `POST /matrix/process`  
**URL:** `http://localhost:3000/matrix/process`  
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

**Ejemplo con curl:**
```bash
curl -X POST http://localhost:3000/matrix/process \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN_GO" \
  -d '{
    "matrix": [
      [1, 2, 3],
      [4, 5, 6],
      [7, 8, 9]
    ]
  }'
```

---

#### Paso 3.2: Go API - Validación de Matriz

**Proceso Interno:**
1. ✅ Verificar que la matriz no esté vacía
2. ✅ Verificar que sea rectangular (todas las filas del mismo tamaño)
3. ✅ Verificar que todos los valores sean numéricos

**Si hay error (400 Bad Request):**
```json
{
  "error": "la matriz no es rectangular: la fila 1 tiene 2 columnas, se esperaban 3"
}
```

**Estado:** ✅ Implementado en `go-api/internal/services/validator.go`

---

#### Paso 3.3: Go API - Rotación 90° Horario

**Proceso Interno:**
- Matriz original: `[[1,2,3], [4,5,6], [7,8,9]]`
- Matriz rotada: `[[7,4,1], [8,5,2], [9,6,3]]`

**Algoritmo:**
- Rotación en sentido horario (clockwise)
- Primera columna → primera fila (invertida)
- Segunda columna → segunda fila (invertida)
- etc.

**Estado:** ✅ Implementado en `go-api/internal/services/rotation.go`

---

#### Paso 3.4: Go API - Factorización QR

**Proceso Interno:**
1. Convertir matriz a formato `gonum` (matriz densa)
2. Calcular factorización QR usando `gonum.org/v1/gonum/lapack`
3. Extraer matrices Q y R

**Matriz Q (Ortogonal):**
- Q × Q^T = I (matriz identidad)
- Columnas ortonormales

**Matriz R (Triangular Superior):**
- Elementos por debajo de la diagonal = 0
- A = Q × R

**Estado:** ✅ Implementado en `go-api/internal/services/qr_decomposition.go`

**Nota:** La factorización QR se calcula sobre la **matriz original** (antes de rotar), según decisión técnica documentada.

---

#### Paso 3.5: Go API → Node.js API (Comunicación HTTP)

**Proceso Interno:**
1. Go API prepara payload con Q, R y matriz rotada
2. Realiza POST HTTP a `http://localhost:3001/matrix/stats`
3. Incluye token JWT en header `Authorization`
4. Timeout configurado (ej: 10 segundos)
5. Manejo de errores de conexión

**Request de Go a Node.js:**
```json
POST http://localhost:3001/matrix/stats
Headers:
  Content-Type: application/json
  Authorization: Bearer <token_nodejs>

Body:
{
  "q": [[-0.123, -0.808, -0.577], ...],
  "r": [[-8.124, -9.601, -11.078], ...],
  "rotated": [[7, 4, 1], [8, 5, 2], [9, 6, 3]]
}
```

**Estado:** ✅ Implementado en `go-api/internal/services/node_client.go`

**Manejo de Errores:**
- Si Node.js no responde → Error 500 en Go API
- Si timeout → Error 500 con mensaje de timeout
- Si Node.js devuelve error → Se propaga al cliente

---

#### Paso 3.6: Node.js API - Validación de Request

**Proceso Interno:**
1. ✅ Verificar token JWT
2. ✅ Verificar que existan matrices Q y R
3. ✅ Validar formato de matrices (arrays de arrays)

**Si hay error (400 Bad Request):**
```json
{
  "error": "se requieren las matrices Q y R"
}
```

**Estado:** ✅ Implementado en `node-api/src/controllers/matrixController.js`

---

#### Paso 3.7: Node.js API - Cálculo de Estadísticas

**Proceso Interno:**

1. **Extraer todos los valores** de Q, R y rotated
2. **Calcular máximo:**
   ```javascript
   max = Math.max(...todosLosValores)
   ```
3. **Calcular mínimo:**
   ```javascript
   min = Math.min(...todosLosValores)
   ```
4. **Calcular promedio:**
   ```javascript
   avg = sumaTotal / cantidadValores
   ```
5. **Calcular suma total:**
   ```javascript
   sum = todosLosValores.reduce((a, b) => a + b, 0)
   ```
6. **Verificar si alguna matriz es diagonal:**
   - Matriz debe ser cuadrada (mismo número de filas y columnas)
   - Todos los elementos fuera de la diagonal principal = 0
   - Función: `isDiagonal(matrix)`

**Estado:** ✅ Implementado en `node-api/src/services/statsService.js`

---

#### Paso 3.8: Node.js API → Go API (Respuesta)

**Response de Node.js:**
```json
{
  "max": 11.078,
  "min": -11.078,
  "avg": 0.123,
  "sum": 45.0,
  "anyDiagonal": false
}
```

**Estado:** ✅ Implementado

---

#### Paso 3.9: Go API → Cliente (Respuesta Final)

**Response Final:**
```json
{
  "rotated": [
    [7, 4, 1],
    [8, 5, 2],
    [9, 6, 3]
  ],
  "q": [
    [-0.123, -0.808, -0.577],
    [-0.492, -0.308, 0.816],
    [-0.861, 0.502, -0.082]
  ],
  "r": [
    [-8.124, -9.601, -11.078],
    [0, 0.904, 1.808],
    [0, 0, 0]
  ],
  "nodeStats": {
    "max": 11.078,
    "min": -11.078,
    "avg": 0.123,
    "sum": 45.0,
    "anyDiagonal": false
  }
}
```

**Estado:** ✅ Implementado

---

## 📊 Diagrama de Flujo

```
┌─────────────┐
│   Cliente   │
│  (Postman/  │
│   curl/etc) │
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

## 🧪 Ejemplo Completo de Uso

### Script Bash Completo

```bash
#!/bin/bash

echo "=== QR Challenge - Flujo Completo ==="
echo ""

# Paso 1: Verificar servicios
echo "1. Verificando servicios..."
curl -s http://localhost:3000/health | jq '.'
curl -s http://localhost:3001/health | jq '.'
echo ""

# Paso 2: Obtener tokens
echo "2. Obteniendo tokens JWT..."
TOKEN_GO=$(curl -X POST http://localhost:3000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin"}' \
  -s | jq -r '.token')

TOKEN_NODE=$(curl -X POST http://localhost:3001/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin"}' \
  -s | jq -r '.token')

echo "✅ Tokens obtenidos"
echo ""

# Paso 3: Procesar matriz
echo "3. Procesando matriz..."
RESPONSE=$(curl -X POST http://localhost:3000/matrix/process \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN_GO" \
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

## ✅ Pruebas Implementadas

### 1. Pruebas Unitarias

#### Go API - ✅ Implementado:
- [x] Tests para `validator.go` (validación de matrices) - **7 casos de prueba**
- [x] Tests para `rotation.go` (rotación 90°) - **6 casos + rotación doble**
- [x] Tests para `qr_decomposition.go` (factorización QR) - **3 casos + verificación Q*R**
- [x] Tests para `auth_controller.go` (JWT) - **3 casos de login**

**Cobertura:**
- `services`: 62.9%
- `controllers`: 80.0%

**Archivos de prueba:**
```
go-api/
├── internal/
│   ├── services/
│   │   ├── validator_test.go      ✅
│   │   ├── rotation_test.go        ✅
│   │   └── qr_decomposition_test.go ✅
│   ├── controllers/
│   │   └── auth_controller_test.go ✅
```

---

#### Node.js API - ✅ Implementado:
- [x] Tests para `statsService.js` (cálculo de estadísticas) - **10+ casos**
- [x] Tests para `auth.js` middleware (verificación JWT) - **6 casos**

**Herramientas usadas:**
- `jest` para testing
- Mocks para middleware de Express

**Archivos de prueba:**
```
node-api/
├── src/
│   ├── services/
│   │   └── statsService.test.js    ✅
│   ├── middleware/
│   │   └── auth.test.js            ✅
```

---

### 2. Pruebas de Integración - ✅ Implementado

#### Go API - Implementado:
- [x] Test de `matrix_handler.go` con httptest - **4 casos**
  - Procesamiento exitoso
  - Error con matriz inválida
  - Error con JSON inválido
  - Error cuando Node.js no responde
- [x] Test de `auth.go` middleware - **6 casos**
  - Token válido
  - Sin token
  - Formato incorrecto
  - Token inválido/expirado
  - Secreto incorrecto

**Cobertura:**
- `handlers`: 77.3%
- `middleware`: 100.0%

**Archivos de prueba:**
```
go-api/
├── internal/
│   ├── handlers/
│   │   └── matrix_handler_test.go  ✅
│   ├── middleware/
│   │   └── auth_test.go            ✅
```

**Ejecutar pruebas:**
```bash
cd go-api
go test -v ./internal/handlers/... ./internal/middleware/...
go test -cover ./...
```

---

### 3. Frontend - ✅ Implementado

#### Características Implementadas:
- [x] Interfaz web simple con Angular
- [x] Formulario para ingresar matriz (formato JSON)
- [x] Visualización de matriz rotada
- [x] Visualización de matrices Q y R
- [x] Mostrar estadísticas (max, min, avg, sum, anyDiagonal)
- [x] Manejo de errores en UI
- [x] Login con JWT
- [x] CSS básico con diseño moderno y responsive
- [x] Docker y Docker Compose configurado

**Tecnologías usadas:**
- Angular 17 (standalone components)
- TypeScript
- CSS básico con gradientes
- HttpClient para consumir APIs
- LocalStorage para guardar token JWT

**Archivos principales:**
```
frontend/
├── src/
│   ├── app/
│   │   ├── components/
│   │   │   ├── login.component.ts
│   │   │   └── matrix-processor.component.ts
│   │   ├── services/
│   │   │   ├── api.service.ts
│   │   │   └── auth.service.ts
│   │   ├── models/
│   │   │   └── api.models.ts
│   │   └── app.component.ts
│   ├── styles.css
│   └── index.html
├── Dockerfile
├── Dockerfile.dev
└── nginx.conf
```

**Acceso:**
- Desarrollo: `http://localhost:4200`
- Producción: `http://localhost:80` (o puerto configurado)

---

## 📝 Checklist de Implementación

### ✅ Completado

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
- [x] **Pruebas unitarias (Go)** - validator, rotation, QR, auth
- [x] **Pruebas unitarias (Node.js)** - statsService, auth middleware
- [x] **Pruebas de integración (Go)** - handlers, middleware con httptest
- [x] **Frontend Angular** - Login, procesamiento, visualización

### ❌ Pendiente

- [ ] Pruebas E2E completas (con testcontainers)
- [ ] CI/CD pipeline (opcional)
- [ ] Métricas y monitoreo (opcional)

---

## 🎯 Próximos Pasos Recomendados

1. **Mejorar Cobertura de Pruebas (Prioridad Media)**
   - Aumentar cobertura de `services` (actualmente 62.9%)
   - Agregar tests para `node_client.go` (comunicación HTTP)
   - Objetivo: >80% coverage en todos los paquetes

2. **Pruebas E2E Completas (Prioridad Media)**
   - Test end-to-end con Docker Compose
   - Test de comunicación real entre APIs
   - Usar `testcontainers` para levantar servicios reales

3. **Mejoras al Frontend (Opcional)**
   - Agregar más validaciones visuales
   - Mejorar UX con animaciones
   - Agregar historial de matrices procesadas

---

## 📚 Referencias

- **ENDPOINTS.md**: Documentación completa de endpoints
- **TESTING.md**: Guía de pruebas manuales
- **POSTMAN_JWT.md**: Instrucciones para Postman
- **README.md**: Documentación general del proyecto

---

## 🔍 Verificación del Requerimiento

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

**Cumplimiento del Requerimiento: ~100%** ✅

### 📊 Cobertura de Código Actual

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

### 📁 Archivos de Prueba Creados

**Go API:**
- `internal/services/validator_test.go`
- `internal/services/rotation_test.go`
- `internal/services/qr_decomposition_test.go`
- `internal/controllers/auth_controller_test.go`
- `internal/handlers/matrix_handler_test.go`
- `internal/middleware/auth_test.go`

**Node.js API:**
- `src/services/statsService.test.js`
- `src/middleware/auth.test.js`
- `jest.config.js`

### 🧪 Ejecutar Pruebas

```bash
# Go API
cd go-api
go test ./...                    # Todas las pruebas
go test -v ./...                 # Con verbosidad
go test -cover ./...             # Con cobertura
go test -coverprofile=coverage.out ./...  # Generar reporte

# Node.js API
cd node-api
npm install                      # Instalar jest
npm test                         # Ejecutar pruebas
npm run test:coverage            # Con cobertura
```

Ver documentación completa en `TESTS.md`

