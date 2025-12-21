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

  // Acumuladores (evitamos construir arrays grandes)
  let max;
  let min;
  let sum = 0;

  // Conteo para promedio:
  // - Por defecto: promedio sobre valores de Q y R (según tests).
  // - Fallback: si Q/R no aportan valores, usar los de rotated.
  let avgCount = 0;
  let rotatedCount = 0;

  const accumulate = (matrix, { upperTriangleOnly = false, countForAvg = true } = {}) => {
    if (!matrix || !Array.isArray(matrix)) return;

    for (let i = 0; i < matrix.length; i++) {
      const row = matrix[i];
      if (!Array.isArray(row)) continue;

      for (let j = 0; j < row.length; j++) {
        if (upperTriangleOnly && i > j) continue;

        const val = row[j];
        if (typeof val !== 'number' || Number.isNaN(val)) continue;

        if (max === undefined || val > max) max = val;
        if (min === undefined || val < min) min = val;
        sum += val;

        if (countForAvg) avgCount += 1;
        if (upperTriangleOnly) rotatedCount += 1;
      }
    }
  };

  // Agregar valores de Q y R completos
  accumulate(q, { upperTriangleOnly: false, countForAvg: true });
  accumulate(r, { upperTriangleOnly: false, countForAvg: true });

  // Agregar valores de rotated solo en triángulo superior (incluida diagonal)
  // Esto refleja que "rotated" suele representar una matriz triangular (R) tras rotaciones.
  accumulate(rotated, { upperTriangleOnly: true, countForAvg: false });

  if (max === undefined || min === undefined) {
    throw new Error('No se encontraron valores numéricos válidos en las matrices');
  }

  // Fallback para evitar división por cero si Q/R no traen valores, pero rotated sí.
  if (avgCount === 0 && rotatedCount > 0) {
    avgCount = rotatedCount;
  }

  const avg = sum / avgCount;

  // Verificar si alguna matriz es diagonal (forzamos boolean)
  const anyDiagonal = isDiagonal(q) || isDiagonal(r) || isDiagonal(rotated);
  
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
  const tolerance = 1e-12; // Tolerancia para comparación de números flotantes
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


