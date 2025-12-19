/**
 * Middleware de autenticación JWT
 */

const jwt = require('jsonwebtoken');

const JWT_SECRET = process.env.JWT_SECRET || 'your-secret-key-change-in-production';

/**
 * Middleware para verificar token JWT
 */
function authenticateToken(req, res, next) {
  const authHeader = req.headers['authorization'];
  const token = authHeader && authHeader.split(' ')[1]; // Bearer TOKEN

  if (!token) {
    return res.status(401).json({
      error: 'Token de acceso requerido',
      message: 'Agrega el header: Authorization: Bearer <token>'
    });
  }

  jwt.verify(token, JWT_SECRET, (err, user) => {
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
  JWT_SECRET
};

