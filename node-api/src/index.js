/**
 * API Node.js para procesamiento de matrices
 * Endpoint: POST /matrix/stats
 */

// Cargar variables de entorno desde .env (solo en desarrollo local)
if (process.env.NODE_ENV !== 'production') {
  require('dotenv').config();
}

const express = require('express');
const matrixRouter = require('./routers/matrixRouter');
const { authenticateToken } = require('./middleware/auth');
const packageJson = require('../package.json');

const app = express();
const PORT = process.env.PORT || 3001;
const startTime = new Date();

// Middlewares
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// CORS para permitir requests desde el frontend
app.use((req, res, next) => {
  res.header('Access-Control-Allow-Origin', '*');
  res.header('Access-Control-Allow-Methods', 'GET, POST, OPTIONS');
  res.header('Access-Control-Allow-Headers', 'Content-Type, Authorization');
  if (req.method === 'OPTIONS') {
    return res.sendStatus(200);
  }
  next();
});

// Logging middleware simple
app.use((req, res, next) => {
  console.log(`${new Date().toISOString()} - ${req.method} ${req.path}`);
  next();
});

// Health check
app.get('/health', (req, res) => {
  res.json({
    status: 'ok',
    service: 'node-api'
  });
});

// Endpoint informativo del backend
app.get('/', (req, res) => {
  const uptime = Date.now() - startTime.getTime();
  const uptimeSeconds = Math.floor(uptime / 1000);
  
  res.json({
    service: 'Node.js API Backend',
    version: packageJson.version,
    technology: 'Node.js',
    framework: `Express ${packageJson.dependencies.express}`,
    nodeVersion: process.version,
    startTime: startTime.toISOString(),
    uptime: `${uptimeSeconds}s`,
    uptimeSeconds: uptimeSeconds,
    platform: process.platform,
    arch: process.arch,
    env: process.env.NODE_ENV || 'development',
    endpoints: {
      health: 'GET /health',
      matrixStats: 'POST /matrix/stats (requiere JWT)',
      info: 'GET /',
      note: 'Login disponible en Go API: POST http://localhost:3000/auth/login'
    },
    memory: {
      used: Math.round(process.memoryUsage().heapUsed / 1024 / 1024) + ' MB',
      total: Math.round(process.memoryUsage().heapTotal / 1024 / 1024) + ' MB',
      rss: Math.round(process.memoryUsage().rss / 1024 / 1024) + ' MB'
    }
  });
});

// Rutas protegidas (requieren JWT)
// Nota: El login se realiza en Go API (http://localhost:3000/auth/login)
app.use('/matrix', authenticateToken, matrixRouter);

// Manejo de errores
app.use((err, req, res, next) => {
  console.error('Error no manejado:', err);
  res.status(500).json({
    error: 'Error interno del servidor',
    message: err.message
  });
});

// Manejo de rutas no encontradas
app.use((req, res) => {
  res.status(404).json({
    error: 'Ruta no encontrada'
  });
});

// Iniciar servidor
app.listen(PORT, () => {
  console.log(`Servidor Node.js iniciado en puerto ${PORT}`);
});

module.exports = app;


