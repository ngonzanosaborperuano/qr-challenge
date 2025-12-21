# 🚀 Inicio Rápido - QR Challenge

Guía completa para levantar el proyecto completo (Go API, Node.js API y Frontend Angular).

---

## 📋 Requisitos Previos

### Opción A: Con Docker (Recomendado) 🐳

- **Docker** versión 20.10 o superior
- **Docker Compose** versión 2.0 o superior

Verificar instalación:
```bash
docker --version
docker-compose --version
```

### Opción B: Desarrollo Local (Sin Docker)

- **Go** 1.23.0 o superior
- **Node.js** 20.x o superior
- **npm** 9.x o superior
- **Angular CLI** 17.x (se instala con npm)

Verificar instalación:
```bash
go version      # Debe ser go1.23.0 o superior
node --version  # Debe ser v20.x o superior
npm --version   # Debe ser 9.x o superior
```

---

## 🐳 Opción 1: Levantar con Docker (Recomendado)

### Paso 1: Configurar Variables de Entorno

Configura variables de entorno (recomendado con un archivo `.env` local **no versionado** o exportando variables en tu shell).

Crea/edita `.env` en la raíz del repo:
```bash
# Puertos
GO_API_PORT=3000
NODE_API_PORT=3001
FRONTEND_PORT=4200

# JWT Secret (OBLIGATORIO: no hay default en el código ni en docker-compose)
JWT_SECRET=REPLACE_ME_WITH_A_LONG_RANDOM_STRING

# Entorno
NODE_ENV=development
```

### Paso 2: Levantar Servicios

#### Desarrollo (con hot-reload) 🔥

```bash
docker-compose -f docker-compose.dev.yml up -d --build
```

**Características:**
- ✅ Hot-reload automático en Go, Node.js y Angular
- ✅ Volúmenes montados para desarrollo
- ✅ Logs en tiempo real
- ✅ Frontend en puerto 4200

**Ver logs:**
```bash
docker-compose -f docker-compose.dev.yml logs -f
```

**Detener servicios:**
```bash
docker-compose -f docker-compose.dev.yml down
```

#### Producción 🚀

```bash
docker-compose up -d --build
```

**Características:**
- ✅ Imágenes optimizadas (multi-stage builds)
- ✅ Sin hot-reload (más rápido)
- ✅ Frontend servido por Nginx en puerto 80

**Ver logs:**
```bash
docker-compose logs -f
```

**Detener servicios:**
```bash
docker-compose down
```

### Paso 3: Verificar que Todo Esté Funcionando

```bash
# Verificar Go API
curl http://localhost:3000/health
# Respuesta esperada: {"status":"ok","service":"go-api"}

# Verificar Node.js API
curl http://localhost:3001/health
# Respuesta esperada: {"status":"ok","service":"node-api"}

# Verificar Frontend (desarrollo)
Abrir navegador en: http://localhost:4200

# Verificar Frontend (producción)
# Abrir navegador en: http://localhost:80
```

### Paso 4: Acceder a la Aplicación

1. **Frontend (Desarrollo):** http://localhost:4200
2. **Frontend (Producción):** http://localhost:80
3. **Go API:** http://localhost:3000
4. **Node.js API:** http://localhost:3001

---

## 💻 Opción 2: Desarrollo Local (Sin Docker)

### Paso 1: Configurar Variables de Entorno

#### Go API

```bash
cd go-api
```

Edita `go-api/.env` (no versionar):
```env
PORT=3000
NODE_API_URL=http://localhost:3001
JWT_SECRET=REPLACE_ME_WITH_A_LONG_RANDOM_STRING
```

#### Node.js API

```bash
cd node-api
```

Edita `node-api/.env` (no versionar):
```env
PORT=3001
NODE_ENV=development
JWT_SECRET=REPLACE_ME_WITH_A_LONG_RANDOM_STRING
```

#### Frontend

No requiere `.env` (usa `src/environments/environment.ts`)

---

### Paso 2: Instalar Dependencias

#### Go API

```bash
cd go-api
go mod download
```

#### Node.js API

```bash
cd node-api
npm install
```

#### Frontend

```bash
cd frontend
npm install
```

---

### Paso 3: Levantar Servicios

**⚠️ IMPORTANTE:** Debes levantar los servicios en este orden:

#### Terminal 1: Node.js API

```bash
cd node-api
npm run dev
```

**Resultado esperado:**
```
Servidor Node.js iniciado en puerto 3001
```

#### Terminal 2: Go API

```bash
cd go-api
go run cmd/server/main.go
```

**Resultado esperado:**
```
Servidor Go iniciado en puerto 3000
Conectado a Node.js API en: http://localhost:3001
```

#### Terminal 3: Frontend Angular

```bash
cd frontend
npm start
```

**Resultado esperado:**
```
** Angular Live Development Server is listening on localhost:4200 **
```

---

### Paso 4: Verificar que Todo Esté Funcionando

```bash
# Verificar Go API
curl http://localhost:3000/health

# Verificar Node.js API
curl http://localhost:3001/health

# Abrir navegador en: http://localhost:4200
```

---

## 🔧 Comandos Útiles

### Docker

```bash
# Ver estado de contenedores
docker-compose ps

# Ver logs de un servicio específico
docker-compose logs -f go-api
docker-compose logs -f node-api
docker-compose logs -f frontend

# Reiniciar un servicio específico
docker-compose restart go-api

# Reconstruir un servicio específico
docker-compose build go-api

# Detener y eliminar contenedores
docker-compose down

# Detener y eliminar contenedores + volúmenes
docker-compose down -v
```

### Desarrollo Local

```bash
# Go API - Ejecutar tests
cd go-api
go test ./...

# Go API - Ejecutar tests con cobertura
go test -cover ./...

# Node.js API - Ejecutar tests
cd node-api
npm test

# Node.js API - Ejecutar tests con cobertura
npm run test:coverage

# Frontend - Build de producción
cd frontend
npm run build
```

---

## 🐛 Solución de Problemas

### Error: Puerto ya en uso

**Solución:** Cambia los puertos en `.env` o detén el proceso que está usando el puerto.

```bash
# Ver qué proceso usa el puerto 3000
lsof -i :3000

# Matar proceso (macOS/Linux)
kill -9 <PID>
```

### Error: Docker no puede construir imágenes

**Solución:** Verifica que Docker esté corriendo y que tengas espacio en disco.

```bash
docker system df
docker system prune  # Limpiar espacio (cuidado: elimina imágenes no usadas)
```

### Error: npm ci falla (package-lock.json desincronizado)

**Solución:** Regenera el package-lock.json

```bash
cd node-api
rm package-lock.json
npm install
```

### Error: Go no encuentra módulos

**Solución:** Verifica que estés en el directorio correcto y descarga dependencias.

```bash
cd go-api
go mod download
go mod tidy
```

### Error: Frontend no se conecta a las APIs

**Solución:** Verifica que las APIs estén corriendo y que las URLs en `environment.ts` sean correctas.

```bash
# Verificar que las APIs respondan
curl http://localhost:3000/health
curl http://localhost:3001/health
```

---

## 📊 Estructura de Puertos

| Servicio | Puerto Desarrollo | Puerto Producción | URL |
|----------|------------------|-------------------|-----|
| Go API | 3000 | 3000 | http://localhost:3000 |
| Node.js API | 3001 | 3001 | http://localhost:3001 |
| Frontend (Dev) | 4200 | - | http://localhost:4200 |
| Frontend (Prod) | - | 80 | http://localhost:80 |

---

## 🔐 Credenciales por Defecto

Para acceder a la aplicación:

- **Usuario:** `admin`
- **Contraseña:** `admin`

⚠️ **IMPORTANTE:** Cambia estas credenciales en producción.

---

## 📝 Próximos Pasos

Una vez que los servicios estén corriendo:

1. **Abrir Frontend:** http://localhost:4200 (dev) o http://localhost:80 (prod)
2. **Iniciar Sesión:** Usa las credenciales por defecto
3. **Procesar Matriz:** Ingresa una matriz en formato JSON
4. **Ver Resultados:** Matriz rotada, QR, y estadísticas

Para más detalles sobre el uso de los endpoints, consulta:
- **FLUJO_COMPLETO.md**: Flujo detallado de uso
- **TESTS.md**: Cómo ejecutar pruebas

---

## 🎯 Resumen Rápido

### Docker (Desarrollo)
```bash
# Asegúrate de tener un `.env` en la raíz con JWT_SECRET (obligatorio)
docker-compose -f docker-compose.dev.yml up -d --build
# Abrir: http://localhost:4200
```

### Docker (Producción)
```bash
# Asegúrate de tener un `.env` en la raíz con JWT_SECRET (obligatorio)
docker-compose up -d --build
# Abrir: http://localhost:80
```

### Desarrollo Local
```bash
# Terminal 1
cd node-api && npm install && npm run dev

# Terminal 2
cd go-api && go mod download && go run cmd/server/main.go

# Terminal 3
cd frontend && npm install && npm start
# Abrir: http://localhost:4200
```

---

## ✅ Checklist de Verificación

- [ ] Docker instalado y corriendo (si usas Docker)
- [ ] Variables de entorno configuradas (`.env`)
- [ ] Servicios levantados correctamente
- [ ] Health checks responden (3000 y 3001)
- [ ] Frontend accesible (4200 o 80)
- [ ] Puedo hacer login con `admin/admin`
- [ ] Puedo procesar una matriz

---

## 📚 Documentación Adicional

- **FLUJO_COMPLETO.md**: Flujo completo de uso y arquitectura
- **TESTS.md**: Guía de pruebas
- **README.md**: Documentación general del proyecto

---

## 🆘 Soporte

Si encuentras problemas:

1. Verifica los logs: `docker-compose logs -f`
2. Verifica que los puertos no estén ocupados
3. Verifica que las variables de entorno estén correctas
4. Consulta la sección "Solución de Problemas" arriba

---

¡Listo! 🎉 Ahora puedes comenzar a usar el proyecto.

