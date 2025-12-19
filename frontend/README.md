# Frontend Angular - QR Challenge

Frontend simple desarrollado con Angular para consumir las APIs de Go y Node.js.

## 🚀 Características

- ✅ Login con JWT
- ✅ Formulario para ingresar matrices
- ✅ Visualización de matriz rotada
- ✅ Visualización de matrices Q y R
- ✅ Estadísticas (max, min, avg, sum, anyDiagonal)
- ✅ CSS básico con diseño moderno
- ✅ Responsive

## 📦 Instalación Local

```bash
cd frontend
npm install
npm start
```

La aplicación estará disponible en `http://localhost:4200`

## 🐳 Docker

### Desarrollo
```bash
docker-compose -f docker-compose.dev.yml up frontend
```

### Producción
```bash
docker-compose up frontend
```

## 🎨 Uso

1. **Iniciar sesión:**
   - Usuario: `admin`
   - Contraseña: `admin`

2. **Procesar matriz:**
   - Ingresa una matriz en formato JSON
   - Ejemplo: `[[1, 2, 3], [4, 5, 6], [7, 8, 9]]`
   - Haz clic en "Procesar Matriz"

3. **Ver resultados:**
   - Matriz rotada
   - Matrices Q y R
   - Estadísticas calculadas

## 🔧 Configuración

Las URLs de las APIs están configuradas en `src/app/services/api.service.ts`:
- Go API: `http://localhost:3000`
- Node.js API: `http://localhost:3001`

Para cambiar las URLs, modifica las constantes en el servicio.

