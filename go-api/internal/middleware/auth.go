package middleware

import (
	"errors"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// getJWTSecret obtiene el secreto JWT desde variable de entorno
func getJWTSecret() (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET no está configurado en las variables de entorno")
	}
	return secret, nil
}

// getJWTSecretBytes obtiene el secreto JWT como []byte (lee dinámicamente)
func getJWTSecretBytes() ([]byte, error) {
	secret, err := getJWTSecret()
	if err != nil {
		return nil, err
	}
	return []byte(secret), nil
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
	// Validar configuración antes de intentar verificar tokens
	if _, err := getJWTSecret(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error de configuración del servidor",
			"message": err.Error(),
		})
	}

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
		return getJWTSecretBytes()
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
