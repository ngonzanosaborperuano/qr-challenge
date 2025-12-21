package handlers

import (
	"log"

	"go-api/internal/models"
	"go-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

// MatrixHandler maneja las peticiones relacionadas con matrices
type MatrixHandler struct {
	NodeClient *services.NodeClient
}

// NewMatrixHandler crea un nuevo handler de matrices
func NewMatrixHandler(nodeClient *services.NodeClient) *MatrixHandler {
	return &MatrixHandler{
		NodeClient: nodeClient,
	}
}

// ProcessMatrix procesa una matriz: valida, rota, calcula QR y obtiene estadísticas de Node
func (h *MatrixHandler) ProcessMatrix(c *fiber.Ctx) error {
	var req models.MatrixRequest

	// Parsear request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "formato JSON inválido: " + err.Error(),
		})
	}

	// Validar matriz
	if err := services.ValidateMatrix(req.Matrix); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Rotar matriz 90° en sentido horario
	rotated := services.RotateMatrix90Clockwise(req.Matrix)

	// Calcular factorización QR de la matriz original (no rotada)
	// Decisión técnica: calculamos QR de la matriz original para mantener
	// la relación matemática estándar A = Q * R
	Q, R, err := services.QRDecomposition(req.Matrix)
	if err != nil {
		log.Printf("Error en factorización QR: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error al calcular factorización QR: " + err.Error(),
		})
	}

	// Obtener token JWT del header Authorization
	// El middleware ya validó que existe, así que podemos extraerlo directamente
	authHeader := c.Get("Authorization")
	tokenJWT := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenJWT = authHeader[7:]
	}

	// Obtener estadísticas de Node.js
	var nodeStats *models.MatrixStatsResponse
	nodeStats, err = h.NodeClient.GetMatrixStats(Q, R, rotated, tokenJWT)
	if err != nil {
		log.Printf("Error al obtener estadísticas de Node.js: %v", err)
		// Continuamos sin las estadísticas de Node, pero las incluimos en el error
		return c.Status(fiber.StatusOK).JSON(models.MatrixProcessResponse{
			Rotated: rotated,
			Q:       Q,
			R:       R,
			Error:   "No se pudieron obtener estadísticas de Node.js: " + err.Error(),
		})
	}

	// Respuesta exitosa
	response := models.MatrixProcessResponse{
		Rotated:   rotated,
		Q:         Q,
		R:         R,
		NodeStats: nodeStats,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
