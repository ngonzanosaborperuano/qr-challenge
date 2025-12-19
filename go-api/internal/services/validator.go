package services

import (
	"fmt"
)

// ValidateMatrix valida que la matriz sea rectangular y contenga solo valores numéricos
// Retorna error si la matriz está vacía, no es rectangular o contiene valores no numéricos
func ValidateMatrix(matrix [][]float64) error {
	if len(matrix) == 0 {
		return fmt.Errorf("la matriz no puede estar vacía")
	}

	if len(matrix[0]) == 0 {
		return fmt.Errorf("las filas de la matriz no pueden estar vacías")
	}

	// Verificar que todas las filas tengan el mismo tamaño (matriz rectangular)
	expectedCols := len(matrix[0])
	for i, row := range matrix {
		if len(row) != expectedCols {
			return fmt.Errorf("la matriz no es rectangular: la fila %d tiene %d columnas, se esperaban %d", i, len(row), expectedCols)
		}
	}

	return nil
}


