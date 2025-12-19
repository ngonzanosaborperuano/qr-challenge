/**
 * Servicio para calcular estadísticas sobre matrices
 */

/**
 * Calcula estadísticas sobre todas las matrices recibidas
 * @param {Object} matricesData - Objeto con matrices q, r y rotated (opcional)
 * @returns {Object} Estadísticas: max, min, avg, sum, anyDiagonal
 */
function calculateStats(matricesData) {
  const { q, r, rotated } = matricesData;
  
  // Combinar todas las matrices en un array plano de valores
  const allValues = [];
  
  // Agregar valores de Q
  if (q && Array.isArray(q)) {
    q.forEach(row => {
      if (Array.isArray(row)) {
        row.forEach(val => {
          if (typeof val === 'number' && !isNaN(val)) {
            allValues.push(val);
          }
        });
      }
    });
  }
  
  // Agregar valores de R
  if (r && Array.isArray(r)) {
    r.forEach(row => {
      if (Array.isArray(row)) {
        row.forEach(val => {
          if (typeof val === 'number' && !isNaN(val)) {
            allValues.push(val);
          }
        });
      }
    });
  }
  
  // Agregar valores de rotated si existe
  if (rotated && Array.isArray(rotated)) {
    rotated.forEach(row => {
      if (Array.isArray(row)) {
        row.forEach(val => {
          if (typeof val === 'number' && !isNaN(val)) {
            allValues.push(val);
          }
        });
      }
    });
  }
  
  if (allValues.length === 0) {
    throw new Error('No se encontraron valores numéricos válidos en las matrices');
  }
  
  // Calcular estadísticas de manera eficiente
  // Usar reduce en lugar de Math.max/min para evitar problemas con matrices grandes
  let max = allValues[0];
  let min = allValues[0];
  let sum = 0;
  
  for (let i = 0; i < allValues.length; i++) {
    const val = allValues[i];
    if (val > max) max = val;
    if (val < min) min = val;
    sum += val;
  }
  
  const avg = sum / allValues.length;
  
  // Verificar si alguna matriz es diagonal
  const anyDiagonal = isDiagonal(q) || isDiagonal(r) || (rotated && isDiagonal(rotated));
  
  return {
    max,
    min,
    avg,
    sum,
    anyDiagonal
  };
}

/**
 * Verifica si una matriz es diagonal
 * Una matriz es diagonal si:
 * 1. Es cuadrada (mismo número de filas y columnas)
 * 2. Todos los elementos fuera de la diagonal principal son iguales a 0
 * @param {Array<Array<number>>} matrix - Matriz a verificar
 * @returns {boolean} true si la matriz es diagonal, false en caso contrario
 */
function isDiagonal(matrix) {
  if (!matrix || !Array.isArray(matrix) || matrix.length === 0) {
    return false;
  }
  
  const rows = matrix.length;
  const cols = matrix[0] ? matrix[0].length : 0;
  
  // Debe ser cuadrada
  if (rows !== cols) {
    return false;
  }
  
  // Verificar que todos los elementos fuera de la diagonal sean 0
  const tolerance = 1e-10; // Tolerancia para comparación de números flotantes
  for (let i = 0; i < rows; i++) {
    if (!Array.isArray(matrix[i])) {
      return false;
    }
    for (let j = 0; j < cols; j++) {
      // Si no está en la diagonal principal y el valor no es 0 (con tolerancia)
      if (i !== j && Math.abs(matrix[i][j]) > tolerance) {
        return false;
      }
    }
  }
  
  return true;
}

module.exports = {
  calculateStats,
  isDiagonal
};


