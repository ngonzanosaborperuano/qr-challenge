package services

// RotateMatrix90Clockwise rota una matriz 90 grados en sentido horario (derecha)
// Ejemplo: [[1,2],[3,4]] -> [[3,1],[4,2]]
func RotateMatrix90Clockwise(matrix [][]float64) [][]float64 {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	rows := len(matrix)
	cols := len(matrix[0])

	// Crear matriz rotada con dimensiones invertidas
	rotated := make([][]float64, cols)
	for i := range rotated {
		rotated[i] = make([]float64, rows)
	}

	// Rotar: el elemento en [i][j] va a [j][rows-1-i]
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = matrix[i][j]
		}
	}

	return rotated
}


