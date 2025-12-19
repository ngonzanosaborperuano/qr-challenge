package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go-api/internal/models"
)

// NodeClient maneja la comunicación con la API de Node.js
type NodeClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewNodeClient crea un nuevo cliente para la API de Node.js
func NewNodeClient(baseURL string) *NodeClient {
	return &NodeClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second, // Timeout de 10 segundos
		},
	}
}

// GetMatrixStats envía las matrices Q y R a Node.js y obtiene las estadísticas
// tokenJWT es el token JWT que se usará para autenticar la petición a Node.js
func (c *NodeClient) GetMatrixStats(q, r, rotated [][]float64, tokenJWT string) (*models.MatrixStatsResponse, error) {
	// Preparar el request
	reqBody := models.MatrixStatsRequest{
		Q:       q,
		R:       r,
		Rotated: rotated,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error al serializar request: %w", err)
	}

	// Realizar la petición HTTP
	url := fmt.Sprintf("%s/matrix/stats", c.BaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error al crear request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	// Incluir token JWT en el header Authorization
	if tokenJWT != "" {
		req.Header.Set("Authorization", "Bearer "+tokenJWT)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al realizar petición a Node.js: %w", err)
	}
	defer resp.Body.Close()

	// Leer respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer respuesta: %w", err)
	}

	// Verificar código de estado
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Node.js retornó código %d: %s", resp.StatusCode, string(body))
	}

	// Parsear respuesta
	var stats models.MatrixStatsResponse
	if err := json.Unmarshal(body, &stats); err != nil {
		return nil, fmt.Errorf("error al parsear respuesta: %w", err)
	}

	return &stats, nil
}


