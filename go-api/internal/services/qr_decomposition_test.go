package services

import (
	"math"
	"testing"
)

func TestQRDecomposition(t *testing.T) {
	tests := []struct {
		name    string
		matrix  [][]float64
		wantErr bool
	}{
		{
			name:    "matriz 2x2 válida",
			matrix:  [][]float64{{1, 2}, {3, 4}},
			wantErr: false,
		},
		{
			name:    "matriz 3x3 válida",
			matrix:  [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			wantErr: false,
		},
		{
			name:    "matriz identidad 3x3",
			matrix:  [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Q, R, err := QRDecomposition(tt.matrix)
			if (err != nil) != tt.wantErr {
				t.Errorf("QRDecomposition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Verificar dimensiones
				if len(Q) != len(tt.matrix) {
					t.Errorf("Q tiene %d filas, esperaba %d", len(Q), len(tt.matrix))
				}
				if len(R) != len(tt.matrix) {
					t.Errorf("R tiene %d filas, esperaba %d", len(R), len(tt.matrix))
				}
				// Verificar que R es triangular superior (elementos debajo de la diagonal son 0 o muy pequeños)
				for i := 0; i < len(R); i++ {
					for j := 0; j < i; j++ {
						if math.Abs(R[i][j]) > 1e-10 {
							t.Errorf("R[%d][%d] = %f, debería ser 0 (R debe ser triangular superior)", i, j, R[i][j])
						}
					}
				}
			}
		})
	}
}

func TestQRDecomposition_Verification(t *testing.T) {
	// Test que verifica que Q * R ≈ A (matriz original)
	matrix := [][]float64{{1, 2}, {3, 4}}
	Q, R, err := QRDecomposition(matrix)
	if err != nil {
		t.Fatalf("QRDecomposition() error = %v", err)
	}

	// Calcular Q * R
	rows := len(Q)
	cols := len(R[0])
	QR := make([][]float64, rows)
	for i := range QR {
		QR[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			for k := 0; k < len(R); k++ {
				QR[i][j] += Q[i][k] * R[k][j]
			}
		}
	}

	// Verificar que Q * R ≈ matrix (con tolerancia para errores de punto flotante)
	tolerance := 1e-6
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			diff := math.Abs(QR[i][j] - matrix[i][j])
			if diff > tolerance {
				t.Errorf("QR[%d][%d] = %f, matrix[%d][%d] = %f, diferencia = %f (mayor que tolerancia %f)",
					i, j, QR[i][j], i, j, matrix[i][j], diff, tolerance)
			}
		}
	}
}

