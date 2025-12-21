# 🧪 Guía de Pruebas Unitarias

## 📋 Resumen

Se han creado pruebas unitarias básicas para ambos proyectos (Go y Node.js) que cubren las funcionalidades principales.

---

## 🔵 Go API - Pruebas Unitarias

### Archivos de Prueba

1. **`validator_test.go`** - Pruebas de validación de matrices
2. **`rotation_test.go`** - Pruebas de rotación 90° horario
3. **`qr_decomposition_test.go`** - Pruebas de factorización QR
4. **`auth_controller_test.go`** - Pruebas de autenticación JWT

### Ejecutar Pruebas

```bash
# Desde la raíz del proyecto
cd go-api

# Ejecutar todas las pruebas
go test ./...

# Ejecutar pruebas con verbosidad
go test -v ./...

# Ejecutar pruebas de un paquete específico
go test ./internal/services/...

# Ejecutar pruebas con cobertura
go test -cover ./...

# Ejecutar pruebas con cobertura detallada
# Nota: si usas `./...` se incluye `cmd/server` (sin tests) y verás 0% en `main.go`.
# Para un reporte más útil, excluye `cmd/`:
go test -coverprofile=coverage.out $(go list ./... | grep -v '/cmd/')
go tool cover -html=coverage.out
```

### Cobertura de Pruebas

#### `validator_test.go`
- ✅ Matriz válida (2x2, 3x3, rectangular)
- ✅ Matriz vacía
- ✅ Fila vacía
- ✅ Matriz no rectangular

#### `rotation_test.go`
- ✅ Rotación 2x2
- ✅ Rotación 3x3
- ✅ Rotación rectangular (2x3, 3x2)
- ✅ Matriz vacía
- ✅ Matriz 1x1
- ✅ Rotación doble (4 rotaciones = original)

#### `qr_decomposition_test.go`
- ✅ Factorización QR 2x2
- ✅ Factorización QR 3x3
- ✅ Matriz identidad
- ✅ Verificación que R es triangular superior
- ✅ Verificación que Q * R ≈ A (matriz original)

#### `auth_controller_test.go`
- ✅ Login exitoso con credenciales válidas
- ✅ Login fallido con credenciales inválidas
- ✅ Validación de token JWT generado

---

## 🟢 Node.js API - Pruebas Unitarias

### Archivos de Prueba

1. **`statsService.test.js`** - Pruebas de cálculo de estadísticas
2. **`auth.test.js`** - Pruebas de middleware de autenticación

### Instalación de Dependencias

```bash
cd node-api
npm install
```

Esto instalará `jest` como dependencia de desarrollo.

### Ejecutar Pruebas

```bash
# Desde node-api/
npm test

# Modo watch (re-ejecuta al cambiar archivos)
npm run test:watch

# Con cobertura
npm run test:coverage
```

### Cobertura de Pruebas

#### `statsService.test.js`
- ✅ Cálculo de estadísticas (max, min, avg, sum)
- ✅ Estadísticas solo con Q y R
- ✅ Manejo de valores negativos
- ✅ Manejo de valores decimales
- ✅ Error cuando no hay valores válidos
- ✅ Detección de matriz diagonal (Q, R, rotated)
- ✅ `anyDiagonal` = false cuando ninguna es diagonal

#### `isDiagonal` (dentro de statsService.test.js)
- ✅ Matriz diagonal válida
- ✅ Matriz no diagonal
- ✅ Matriz no cuadrada
- ✅ Matriz identidad
- ✅ Matriz vacía/null/undefined
- ✅ Tolerancia para valores pequeños
- ✅ Matriz 1x1

#### `auth.test.js`
- ✅ Acceso permitido con token válido
- ✅ Acceso rechazado sin token
- ✅ Acceso rechazado con token inválido
- ✅ Acceso rechazado con secreto incorrecto
- ✅ Acceso rechazado con token expirado
- ✅ Rechazo de formato de header incorrecto

---

## 📊 Ejemplo de Salida

### Go API

```bash
$ cd go-api && go test -v ./...

=== RUN   TestValidateMatrix
=== RUN   TestValidateMatrix/matriz_válida_2x2
=== RUN   TestValidateMatrix/matriz_válida_3x3
...
--- PASS: TestValidateMatrix (0.00s)
    --- PASS: TestValidateMatrix/matriz_válida_2x2 (0.00s)
    --- PASS: TestValidateMatrix/matriz_válida_3x3 (0.00s)
    ...
PASS
ok      go-api/internal/services    0.123s
```

### Node.js API

```bash
$ cd node-api && npm test

> api@1.0.0 test
> jest

 PASS  src/services/statsService.test.js
 PASS  src/middleware/auth.test.js

Test Suites: 2 passed, 2 total
Tests:       20 passed, 20 total
Snapshots:   0 total
Time:        1.234 s
```

---

## 🎯 Próximos Pasos (Opcional)

### Mejoras Futuras

1. **Pruebas de Integración**
   - Test end-to-end completo
   - Test de comunicación HTTP entre APIs
   - Test con Docker Compose

2. **Más Cobertura**
   - Tests para `node_client.go` (comunicación HTTP)
   - Tests para `matrix_handler.go` (orquestación)
   - Tests para controllers de Node.js

3. **CI/CD**
   - Integración con GitHub Actions
   - Ejecutar tests automáticamente en PRs
   - Reportes de cobertura

---

## 📝 Notas

- Las pruebas usan valores de ejemplo simples para facilitar la comprensión
- Los tests de QR usan tolerancia para errores de punto flotante
- Los tests de autenticación usan un secreto de prueba (`test-secret-key`)
- En producción, usar secretos reales y más seguros

---

## 🔧 Troubleshooting

### Go: "package not found"
```bash
# Asegúrate de estar en el directorio correcto
cd go-api
go mod tidy
go test ./...
```

### Node.js: "jest not found"
```bash
cd node-api
npm install
```

### Node.js: Tests fallan por JWT_SECRET
Los tests configuran `JWT_SECRET` automáticamente, pero si fallan:
```bash
# Verificar que jest.config.js existe
# Verificar que los mocks están correctos
```

---

## ✅ Checklist de Pruebas

- [x] Validación de matrices (Go)
- [x] Rotación de matrices (Go)
- [x] Factorización QR (Go)
- [x] Autenticación JWT (Go)
- [x] Cálculo de estadísticas (Node.js)
- [x] Verificación diagonal (Node.js)
- [x] Middleware de autenticación (Node.js)
- [ ] Pruebas de integración (Pendiente)
- [ ] Pruebas E2E (Pendiente)

