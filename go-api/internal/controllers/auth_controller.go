package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-api/internal/middleware"
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
func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return []byte("your-secret-key-change-in-production")
	}
	return []byte(secret)
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTSecret())
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

