package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(getJWTSecret())

// getJWTSecret obtiene el secreto JWT desde variable de entorno
func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "your-secret-key-change-in-production"
	}
	return secret
}

// getJWTSecretBytes obtiene el secreto JWT como []byte (lee dinámicamente)
func getJWTSecretBytes() []byte {
	return []byte(getJWTSecret())
}

// Claims estructura para los claims del JWT
type Claims struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// AuthenticateToken middleware para verificar token JWT
func AuthenticateToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Token de acceso requerido",
			"message": "Agrega el header: Authorization: Bearer <token>",
		})
	}

	// Extraer token del header "Bearer TOKEN"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Formato de token inválido",
			"message": "El formato debe ser: Bearer <token>",
		})
	}

	tokenString := parts[1]

	// Parsear y validar token
	// Usar getJWTSecretBytes() para leer dinámicamente (útil en tests)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecretBytes(), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   "Token inválido o expirado",
			"message": err.Error(),
		})
	}

	// Guardar información del usuario en el contexto
	c.Locals("user", claims)

	return c.Next()
}

