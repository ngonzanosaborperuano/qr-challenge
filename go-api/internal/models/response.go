package models

// MatrixStatsResponse representa la respuesta de Node.js con estadísticas
type MatrixStatsResponse struct {
	Max         float64 `json:"max"`
	Min         float64 `json:"min"`
	Avg         float64 `json:"avg"`
	Sum         float64 `json:"sum"`
	AnyDiagonal bool    `json:"anyDiagonal"`
}

// MatrixProcessResponse representa la respuesta final al cliente
type MatrixProcessResponse struct {
	Rotated   [][]float64        `json:"rotated"`
	Q         [][]float64        `json:"q"`
	R         [][]float64        `json:"r"`
	NodeStats *MatrixStatsResponse `json:"nodeStats,omitempty"`
	Error     string             `json:"error,omitempty"`
}


