package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func TestAuthenticateToken(t *testing.T) {
	// Configurar JWT_SECRET para tests
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	app := fiber.New()

	// Ruta de prueba protegida
	app.Get("/protected", AuthenticateToken, func(c *fiber.Ctx) error {
		user := c.Locals("user")
		if user == nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "user not found in context",
			})
		}
		return c.JSON(fiber.Map{
			"message": "acceso permitido",
			"user":    user,
		})
	})

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
		shouldHaveUser bool
	}{
		{
			name:           "acceso permitido con token válido",
			authHeader:     "Bearer " + createValidToken(t),
			expectedStatus: http.StatusOK,
			shouldHaveUser: true,
		},
		{
			name:           "acceso denegado sin token",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
			shouldHaveUser: false,
		},
		{
			name:           "acceso denegado con formato incorrecto",
			authHeader:     "InvalidFormat token",
			expectedStatus: http.StatusUnauthorized,
			shouldHaveUser: false,
		},
		{
			name:           "acceso denegado sin Bearer",
			authHeader:     "token-sin-bearer",
			expectedStatus: http.StatusUnauthorized,
			shouldHaveUser: false,
		},
		{
			name:           "acceso denegado con token inválido",
			authHeader:     "Bearer invalid-token-12345",
			expectedStatus: http.StatusForbidden,
			shouldHaveUser: false,
		},
		{
			name:           "acceso denegado con token con secreto incorrecto",
			authHeader:     "Bearer " + createTokenWithWrongSecret(t),
			expectedStatus: http.StatusForbidden,
			shouldHaveUser: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/protected", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Error al hacer request: %v", err)
			}

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Status code = %d, want %d", resp.StatusCode, tt.expectedStatus)
			}

			// Verificar que el usuario está en el contexto si debería estar
			if tt.shouldHaveUser && resp.StatusCode == http.StatusOK {
				// El usuario debería estar en el contexto (verificado en el handler de prueba)
			}
		})
	}
}

// createValidToken crea un token JWT válido para pruebas
func createValidToken(t *testing.T) string {
	claims := &Claims{
		Username: "admin",
		ID:       1,
		Role:     "admin",
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("test-secret-key"))
	if err != nil {
		t.Fatalf("Error al crear token: %v", err)
	}
	return tokenString
}

// createTokenWithWrongSecret crea un token con un secreto incorrecto
func createTokenWithWrongSecret(t *testing.T) string {
	claims := &Claims{
		Username: "admin",
		ID:       1,
		Role:     "admin",
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("wrong-secret-key"))
	if err != nil {
		t.Fatalf("Error al crear token: %v", err)
	}
	return tokenString
}

