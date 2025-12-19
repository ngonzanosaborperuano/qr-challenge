package services

import (
	"reflect"
	"testing"
)

func TestRotateMatrix90Clockwise(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]float64
		expected [][]float64
	}{
		{
			name:     "matriz 2x2",
			matrix:   [][]float64{{1, 2}, {3, 4}},
			expected: [][]float64{{3, 1}, {4, 2}},
		},
		{
			name:     "matriz 3x3",
			matrix:   [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: [][]float64{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			name:     "matriz rectangular 2x3",
			matrix:   [][]float64{{1, 2, 3}, {4, 5, 6}},
			expected: [][]float64{{4, 1}, {5, 2}, {6, 3}},
		},
		{
			name:     "matriz rectangular 3x2",
			matrix:   [][]float64{{1, 2}, {3, 4}, {5, 6}},
			expected: [][]float64{{5, 3, 1}, {6, 4, 2}},
		},
		{
			name:     "matriz vacía",
			matrix:   [][]float64{},
			expected: [][]float64{},
		},
		{
			name:     "matriz 1x1",
			matrix:   [][]float64{{42}},
			expected: [][]float64{{42}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RotateMatrix90Clockwise(tt.matrix)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RotateMatrix90Clockwise() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Test de rotación doble (debe volver a la original)
func TestRotateMatrix90Clockwise_DoubleRotation(t *testing.T) {
	original := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotatedOnce := RotateMatrix90Clockwise(original)
	rotatedTwice := RotateMatrix90Clockwise(rotatedOnce)
	rotatedThrice := RotateMatrix90Clockwise(rotatedTwice)
	rotatedFour := RotateMatrix90Clockwise(rotatedThrice)

	// Después de 4 rotaciones de 90°, debe volver a la original
	if !reflect.DeepEqual(rotatedFour, original) {
		t.Errorf("Double rotation failed: after 4 rotations got %v, want %v", rotatedFour, original)
	}
}

