package controllers

import (
	"errors"
	"os"
	"time"

	"go-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// LoginRequest estructura para el request de login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse estructura para la respuesta de login
type LoginResponse struct {
	Success   bool   `json:"success"`
	Token     string `json:"token"`
	Message   string `json:"message"`
	ExpiresIn string `json:"expiresIn"`
}

// getJWTSecret obtiene el secreto JWT desde variable de entorno
func getJWTSecret() ([]byte, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET no está configurado en las variables de entorno")
	}
	return []byte(secret), nil
}

// Login genera un token JWT
// POST /auth/login
func Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Formato JSON inválido",
			"message": err.Error(),
		})
	}

	// Credenciales simples (en producción usar hash y base de datos)
	if req.Username != "admin" || req.Password != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Credenciales inválidas",
			"message": "Usuario o contraseña incorrectos",
		})
	}

	// Crear claims
	claims := &middleware.Claims{
		Username: "admin",
		ID:       1,
		Role:     "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Generar token
	secret, err := getJWTSecret()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error de configuración del servidor",
			"message": err.Error(),
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error al generar token",
			"message": err.Error(),
		})
	}

	return c.JSON(LoginResponse{
		Success:   true,
		Token:     tokenString,
		Message:   "Login exitoso",
		ExpiresIn: "24h",
	})
}
