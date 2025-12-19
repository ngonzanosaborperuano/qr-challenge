package models

// MatrixRequest representa la petición del cliente con una matriz
type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

// MatrixStatsRequest representa la petición que se envía a Node.js
type MatrixStatsRequest struct {
	Q        [][]float64 `json:"q"`
	R        [][]float64 `json:"r"`
	Rotated  [][]float64 `json:"rotated,omitempty"`
}


