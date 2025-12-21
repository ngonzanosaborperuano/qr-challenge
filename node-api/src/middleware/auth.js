/**
 * Middleware de autenticación JWT
 */

const jwt = require('jsonwebtoken');

function getJwtSecret() {
  return process.env.JWT_SECRET;
}

/**
 * Middleware para verificar token JWT
 */
function authenticateToken(req, res, next) {
  const authHeader = req.headers['authorization'];
  if (!authHeader) {
    return res.status(401).json({
      error: 'Token de acceso requerido',
      message: 'Agrega el header: Authorization: Bearer <token>'
    });
  }

  // Validar formato: "Bearer <token>"
  const parts = authHeader.split(' ');
  if (parts.length !== 2 || parts[0] !== 'Bearer' || !parts[1]) {
    return res.status(401).json({
      error: 'Token de acceso requerido',
      message: 'Agrega el header: Authorization: Bearer <token>'
    });
  }

  const token = parts[1];

  if (!token) {
    return res.status(401).json({
      error: 'Token de acceso requerido',
      message: 'Agrega el header: Authorization: Bearer <token>'
    });
  }

  const jwtSecret = getJwtSecret();
  if (!jwtSecret) {
    return res.status(500).json({
      error: 'Error de configuración del servidor',
      message: 'JWT_SECRET no está configurado en las variables de entorno'
    });
  }

  jwt.verify(token, jwtSecret, (err, user) => {
    if (err) {
      return res.status(403).json({
        error: 'Token inválido o expirado',
        message: err.message
      });
    }

    req.user = user;
    next();
  });
}

module.exports = {
  authenticateToken,
  getJwtSecret
};

