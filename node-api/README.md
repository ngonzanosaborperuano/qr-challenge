# Node.js API - QR Challenge

API RESTful desarrollada con Node.js y Express.js para calcular estadísticas sobre matrices recibidas desde la API de Go.

## 🚀 Características

- ✅ Cálculo de estadísticas (max, min, avg, sum)
- ✅ Verificación de matrices diagonales
- ✅ Validación JWT
- ✅ Health checks
- ✅ Manejo robusto de errores
- ✅ Tests unitarios con Jest
- ✅ CORS configurado para frontend

## 📦 Instalación Local

### Requisitos

- Node.js 20.x o superior
- npm 9.x o superior
- Variables de entorno configuradas (archivo `.env` local **no versionado** o variables exportadas)

### Pasos

```bash
cd node-api

# Instalar dependencias
npm install

# Configurar variables de entorno
# Crea un archivo `.env` (no versionar) con tus valores

# Ejecutar en modo desarrollo (con nodemon)
npm run dev

# O ejecutar en modo producción
npm start
```

La API estará disponible en `http://localhost:3001`

## 🐳 Docker

### Desarrollo (con hot-reload)

```bash
docker-compose -f docker-compose.dev.yml up node-api
```

### Producción

```bash
docker-compose up node-api
```

## 📍 Endpoints

### `GET /` - Información del Servicio
Obtiene información sobre la API, versión, y estado del sistema.

**Autenticación:** No requerida

**Ejemplo:**
```bash
curl http://localhost:3001/
```

**Response:**
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

### `GET /health` - Health Check
Verifica que el servicio esté funcionando correctamente.

**Autenticación:** No requerida

**Ejemplo:**
```bash
curl http://localhost:3001/health
```

**Response:**
```json
{
  "status": "ok",
  "service": "node-api"
}
```

---

### `POST /matrix/stats` - Calcular Estadísticas
Calcula estadísticas (max, min, avg, sum) sobre matrices Q, R y rotated, y verifica si alguna es diagonal.

**Autenticación:** Requerida (JWT)

**Request:**
```json
{
  "q": [[-0.12, 0.90, 0.41], [-0.49, 0.30, -0.82], [-0.86, -0.30, 0.41]],
  "r": [[-8.12, -9.60, -11.08], [0.00, 0.90, 1.81], [0.00, 0.00, -0.00]],
  "rotated": [[7, 4, 1], [8, 5, 2], [9, 6, 3]]
}
```

**Ejemplo:**
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

**Response:**
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

**Definición de matriz diagonal:**
- Debe ser cuadrada (mismo número de filas y columnas)
- Todos los elementos fuera de la diagonal principal deben ser 0 (con tolerancia numérica de 1e-10)

---

## 🔧 Configuración

### Variables de Entorno

Crea un archivo `.env` en el directorio `node-api/`:

```env
PORT=3001
NODE_ENV=development
JWT_SECRET=REPLACE_ME_WITH_A_LONG_RANDOM_STRING
```

**Variables disponibles:**
- `PORT`: Puerto donde escucha el servidor (default: 3001)
- `NODE_ENV`: Entorno de ejecución (development/production)
- `JWT_SECRET`: Secreto para verificar tokens JWT (debe coincidir con Go API)

### Estructura del Proyecto

```
node-api/
├── src/
│   ├── controllers/          # Controladores
│   │   └── matrixController.js
│   ├── middleware/           # Middleware (JWT)
│   │   ├── auth.js
│   │   └── auth.test.js
│   ├── routers/              # Rutas
│   │   └── matrixRouter.js
│   ├── services/             # Lógica de negocio
│   │   ├── statsService.js
│   │   └── statsService.test.js
│   └── index.js              # Punto de entrada
├── Dockerfile                # Build producción
├── Dockerfile.dev            # Build desarrollo
├── package.json              # Dependencias
├── jest.config.js            # Configuración Jest
└── README.md                 # Este archivo
```

---

## 🧪 Testing

### Ejecutar Tests

```bash
# Todas las pruebas
npm test

# Modo watch (re-ejecuta al cambiar archivos)
npm run test:watch

# Con cobertura
npm run test:coverage
```

### Cobertura Actual

- `statsService`: ✅ Cobertura completa de funciones principales
- `auth middleware`: ✅ Cobertura completa

---

## 📚 Dependencias Principales

### Producción

- **express v5.2.1**: Framework web minimalista y flexible
- **express-jwt v8.5.1**: Middleware para validar tokens JWT
- **jsonwebtoken v9.0.3**: Librería para trabajar con JWT
- **dotenv v17.2.3**: Carga de variables de entorno

### Desarrollo

- **jest v29.7.0**: Framework de testing
- **nodemon v3.1.11**: Auto-reload en desarrollo

---

## 🔍 Decisiones Técnicas

1. **Cálculo eficiente de estadísticas**: Usa loop único para calcular max, min y sum en una sola pasada, evitando problemas de stack overflow con matrices grandes.

2. **Tolerancia numérica**: Usa `1e-10` para comparaciones de punto flotante en verificación de matrices diagonales.

3. **Validación JWT**: Solo valida tokens recibidos. No genera tokens (eso lo hace Go API).

4. **CORS**: Configurado para permitir requests desde el frontend.

---

## 🐛 Solución de Problemas

### Error: "Cannot find module"

```bash
npm install
```

### Error: "port already in use"

Cambia el puerto en `.env` o detén el proceso que está usando el puerto.

### Error: "JWT_SECRET must be provided"

Verifica que el archivo `.env` exista y contenga `JWT_SECRET`.

### Error: "package-lock.json desincronizado"

```bash
rm package-lock.json
npm install
```

---

## 🎯 Versión

- **Node.js**: 20.x (Alpine)
- **Express**: v5.2.1
- **Jest**: v29.7.0

---

## 📝 Scripts Disponibles

```bash
npm start          # Ejecutar en producción
npm run dev        # Ejecutar en desarrollo (con nodemon)
npm test           # Ejecutar tests
npm run test:watch # Ejecutar tests en modo watch
npm run test:coverage # Ejecutar tests con cobertura
```

---

¡Listo para usar! 🚀

