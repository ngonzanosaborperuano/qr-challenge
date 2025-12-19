/**
 * Controller para manejar peticiones relacionadas con matrices
 */

const statsService = require('../services/statsService');

/**
 * Calcula estadísticas sobre las matrices recibidas
 * POST /matrix/stats
 * Body: { q: [[...]], r: [[...]], rotated: [[...]] (opcional) }
 */
function getMatrixStats(req, res) {
  try {
    const { q, r, rotated, matrices } = req.body;
    
    // Validar que al menos q y r estén presentes
    if (!q || !r) {
      // Intentar formato alternativo con array 'matrices'
      if (matrices && Array.isArray(matrices) && matrices.length >= 2) {
        const matricesData = {
          q: matrices[0],
          r: matrices[1],
          rotated: matrices[2] || rotated
        };
        const stats = statsService.calculateStats(matricesData);
        return res.status(200).json(stats);
      }
      
      return res.status(400).json({
        error: 'Se requieren las matrices q y r en el body'
      });
    }
    
    // Validar que sean arrays
    if (!Array.isArray(q) || !Array.isArray(r)) {
      return res.status(400).json({
        error: 'q y r deben ser arrays'
      });
    }
    
    // Calcular estadísticas
    const stats = statsService.calculateStats({ q, r, rotated });
    
    res.status(200).json(stats);
  } catch (error) {
    console.error('Error al calcular estadísticas:', error);
    res.status(500).json({
      error: 'Error interno del servidor: ' + error.message
    });
  }
}

module.exports = {
  getMatrixStats
};


