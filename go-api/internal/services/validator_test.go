package services

import (
	"testing"
)

func TestValidateMatrix(t *testing.T) {
	tests := []struct {
		name    string
		matrix  [][]float64
		wantErr bool
		errMsg  string
	}{
		{
			name:    "matriz válida 2x2",
			matrix:  [][]float64{{1, 2}, {3, 4}},
			wantErr: false,
		},
		{
			name:    "matriz válida 3x3",
			matrix:  [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			wantErr: false,
		},
		{
			name:    "matriz válida rectangular 2x3",
			matrix:  [][]float64{{1, 2, 3}, {4, 5, 6}},
			wantErr: false,
		},
		{
			name:    "matriz vacía",
			matrix:  [][]float64{},
			wantErr: true,
			errMsg:  "la matriz no puede estar vacía",
		},
		{
			name:    "fila vacía",
			matrix:  [][]float64{{}},
			wantErr: true,
			errMsg:  "las filas de la matriz no pueden estar vacías",
		},
		{
			name:    "matriz no rectangular - fila 1 más corta",
			matrix:  [][]float64{{1, 2, 3}, {4, 5}},
			wantErr: true,
			errMsg:  "la matriz no es rectangular",
		},
		{
			name:    "matriz no rectangular - fila 1 más larga",
			matrix:  [][]float64{{1, 2}, {3, 4, 5}},
			wantErr: true,
			errMsg:  "la matriz no es rectangular",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMatrix(tt.matrix)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMatrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errMsg != "" {
				if err == nil || err.Error() == "" {
					t.Errorf("ValidateMatrix() expected error message containing '%s', got: %v", tt.errMsg, err)
				}
			}
		})
	}
}

