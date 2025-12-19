const { calculateStats, isDiagonal } = require('./statsService');

describe('statsService', () => {
  describe('calculateStats', () => {
    test('debe calcular estadísticas correctamente para matrices válidas', () => {
      const matricesData = {
        q: [[1, 0], [0, 1]],
        r: [[2, 0], [0, 3]],
        rotated: [[4, 1], [3, 2]]
      };

      const result = calculateStats(matricesData);

      expect(result).toHaveProperty('max');
      expect(result).toHaveProperty('min');
      expect(result).toHaveProperty('avg');
      expect(result).toHaveProperty('sum');
      expect(result).toHaveProperty('anyDiagonal');

      expect(result.max).toBe(4);
      expect(result.min).toBe(0);
      expect(result.sum).toBe(14);
      expect(result.avg).toBeCloseTo(1.75, 2);
    });

    test('debe calcular estadísticas solo con Q y R', () => {
      const matricesData = {
        q: [[1, 2], [3, 4]],
        r: [[5, 6], [7, 8]]
      };

      const result = calculateStats(matricesData);

      expect(result.max).toBe(8);
      expect(result.min).toBe(1);
      expect(result.sum).toBe(36);
    });

    test('debe lanzar error si no hay valores numéricos válidos', () => {
      const matricesData = {
        q: [],
        r: []
      };

      expect(() => calculateStats(matricesData)).toThrow('No se encontraron valores numéricos válidos');
    });

    test('debe manejar matrices con valores negativos', () => {
      const matricesData = {
        q: [[-1, -2], [-3, -4]],
        r: [[1, 2], [3, 4]]
      };

      const result = calculateStats(matricesData);

      expect(result.max).toBe(4);
      expect(result.min).toBe(-4);
      expect(result.sum).toBe(0);
    });

    test('debe manejar valores decimales', () => {
      const matricesData = {
        q: [[1.5, 2.5], [3.5, 4.5]],
        r: [[0.5, 1.5], [2.5, 3.5]]
      };

      const result = calculateStats(matricesData);

      expect(result.max).toBe(4.5);
      expect(result.min).toBe(0.5);
      expect(result.avg).toBeCloseTo(2.5, 1);
    });
  });

  describe('isDiagonal', () => {
    test('debe retornar true para matriz diagonal válida', () => {
      const matrix = [
        [2, 0, 0],
        [0, 3, 0],
        [0, 0, 4]
      ];

      expect(isDiagonal(matrix)).toBe(true);
    });

    test('debe retornar false para matriz no diagonal', () => {
      const matrix = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
      ];

      expect(isDiagonal(matrix)).toBe(false);
    });

    test('debe retornar false para matriz no cuadrada', () => {
      const matrix = [
        [1, 2, 3],
        [4, 5, 6]
      ];

      expect(isDiagonal(matrix)).toBe(false);
    });

    test('debe retornar true para matriz identidad', () => {
      const matrix = [
        [1, 0, 0],
        [0, 1, 0],
        [0, 0, 1]
      ];

      expect(isDiagonal(matrix)).toBe(true);
    });

    test('debe retornar false para matriz vacía', () => {
      expect(isDiagonal([])).toBe(false);
      expect(isDiagonal(null)).toBe(false);
      expect(isDiagonal(undefined)).toBe(false);
    });

    test('debe manejar valores muy pequeños (tolerancia)', () => {
      const matrix = [
        [2, 0.0000000001, 0],
        [0, 3, 0],
        [0, 0, 4]
      ];

      // Con tolerancia, debería ser false porque hay un valor fuera de la diagonal
      expect(isDiagonal(matrix)).toBe(false);
    });

    test('debe retornar true para matriz 1x1', () => {
      const matrix = [[5]];
      expect(isDiagonal(matrix)).toBe(true);
    });
  });

  describe('calculateStats con anyDiagonal', () => {
    test('debe detectar matriz diagonal en Q', () => {
      const matricesData = {
        q: [[2, 0], [0, 3]],
        r: [[1, 2], [3, 4]]
      };

      const result = calculateStats(matricesData);
      expect(result.anyDiagonal).toBe(true);
    });

    test('debe detectar matriz diagonal en R', () => {
      const matricesData = {
        q: [[1, 2], [3, 4]],
        r: [[5, 0], [0, 6]]
      };

      const result = calculateStats(matricesData);
      expect(result.anyDiagonal).toBe(true);
    });

    test('debe detectar matriz diagonal en rotated', () => {
      const matricesData = {
        q: [[1, 2], [3, 4]],
        r: [[5, 6], [7, 8]],
        rotated: [[2, 0], [0, 3]]
      };

      const result = calculateStats(matricesData);
      expect(result.anyDiagonal).toBe(true);
    });

    test('debe retornar false si ninguna matriz es diagonal', () => {
      const matricesData = {
        q: [[1, 2], [3, 4]],
        r: [[5, 6], [7, 8]]
      };

      const result = calculateStats(matricesData);
      expect(result.anyDiagonal).toBe(false);
    });
  });
});

