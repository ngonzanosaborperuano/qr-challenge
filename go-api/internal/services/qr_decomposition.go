package services

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// QRDecomposition calcula la factorización QR de una matriz usando el método de Gram-Schmidt
// Retorna las matrices Q y R, donde Q es ortogonal y R es triangular superior
// La matriz original debe ser A = Q * R
func QRDecomposition(matrix [][]float64) ([][]float64, [][]float64, error) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil, nil, fmt.Errorf("la matriz no puede estar vacía")
	}

	rows := len(matrix)
	cols := len(matrix[0])

	// Convertir a matriz de gonum
	data := make([]float64, 0, rows*cols)
	for _, row := range matrix {
		data = append(data, row...)
	}
	A := mat.NewDense(rows, cols, data)

	// Calcular factorización QR
	var qr mat.QR
	qr.Factorize(A)

	// Extraer Q
	qDense := &mat.Dense{}
	qr.QTo(qDense)

	// Extraer R
	rDense := &mat.Dense{}
	qr.RTo(rDense)

	// Convertir Q a [][]float64
	Q := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		Q[i] = make([]float64, qDense.RawMatrix().Cols)
		for j := 0; j < qDense.RawMatrix().Cols; j++ {
			Q[i][j] = qDense.At(i, j)
		}
	}

	// Convertir R a [][]float64
	R := make([][]float64, rDense.RawMatrix().Rows)
	for i := 0; i < rDense.RawMatrix().Rows; i++ {
		R[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			R[i][j] = rDense.At(i, j)
		}
	}

	return Q, R, nil
}


