package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-api/internal/middleware"
	"go-api/internal/services"
)

// Nota: Para simplificar, estas pruebas usan el cliente real de Node.js
// pero con una URL que no existe, lo que simula un error de conexión.
// En una implementación más completa, se usaría una interfaz para poder mockear.

func TestProcessMatrix(t *testing.T) {
	// Configurar JWT_SECRET para tests ANTES de importar middleware
	// (el middleware inicializa jwtSecret como variable global)
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	// Reinicializar el secreto en el middleware para que use el nuevo valor
	// Nota: Esto requiere acceso al paquete middleware, pero como jwtSecret es privado,
	// necesitamos asegurarnos de que el entorno esté configurado antes
	// de que cualquier código del middleware se ejecute.

	// Crear token JWT válido para las pruebas
	token := createTestToken(t, "test-secret-key")

	tests := []struct {
		name           string
		body           map[string]interface{}
		authToken      string
		nodeURL        string // URL del servidor Node.js (o inexistente para simular error)
		expectedStatus int
		checkResponse  func(*testing.T, *http.Response)
	}{
		{
			name: "procesamiento exitoso con matriz válida",
			body: map[string]interface{}{
				"matrix": [][]float64{{1, 2}, {3, 4}},
			},
			authToken:      token,
			nodeURL:        "http://localhost:3001", // URL real (puede fallar si Node.js no está corriendo)
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, resp *http.Response) {
				var result map[string]interface{}
				json.NewDecoder(resp.Body).Decode(&result)
				if result["rotated"] == nil {
					t.Error("Expected 'rotated' in response")
				}
				if result["q"] == nil {
					t.Error("Expected 'q' in response")
				}
				if result["r"] == nil {
					t.Error("Expected 'r' in response")
				}
				// Nota: nodeStats puede ser nil si Node.js no está disponible en la prueba
			},
		},
		{
			name: "error con matriz inválida (no rectangular)",
			body: map[string]interface{}{
				"matrix": [][]float64{{1, 2}, {3}},
			},
			authToken:      token,
			nodeURL:        "http://localhost:3001",
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, resp *http.Response) {
				var result map[string]interface{}
				json.NewDecoder(resp.Body).Decode(&result)
				if result["error"] == nil {
					t.Error("Expected 'error' in response")
				}
			},
		},
		{
			name: "error con JSON inválido",
			body: map[string]interface{}{
				"invalid": "data",
			},
			authToken:      token,
			nodeURL:        "http://localhost:3001",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "error cuando Node.js no responde",
			body: map[string]interface{}{
				"matrix": [][]float64{{1, 2}, {3, 4}},
			},
			authToken:      token,
			nodeURL:        "http://localhost:9999", // URL que no existe para simular error
			expectedStatus: http.StatusOK,            // El handler devuelve 200 pero con error en el campo Error
			checkResponse: func(t *testing.T, resp *http.Response) {
				var result map[string]interface{}
				json.NewDecoder(resp.Body).Decode(&result)
				if result["error"] == nil {
					t.Error("Expected 'error' field when Node.js fails")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Asegurar que JWT_SECRET esté configurado
			os.Setenv("JWT_SECRET", "test-secret-key")

			// Crear app Fiber para pruebas
			app := fiber.New()

			// Crear cliente Node.js con la URL especificada en el test
			nodeClient := services.NewNodeClient(tt.nodeURL)
			handler := NewMatrixHandler(nodeClient)

			// Configurar ruta con middleware de autenticación
			// El middleware leerá JWT_SECRET del entorno
			app.Post("/matrix/process", func(c *fiber.Ctx) error {
				// Llamar al middleware manualmente para que use el entorno actualizado
				return middleware.AuthenticateToken(c)
			}, handler.ProcessMatrix)

			// Preparar request
			bodyJSON, _ := json.Marshal(tt.body)
			req := httptest.NewRequest("POST", "/matrix/process", bytes.NewBuffer(bodyJSON))
			req.Header.Set("Content-Type", "application/json")
			if tt.authToken != "" {
				req.Header.Set("Authorization", "Bearer "+tt.authToken)
			}

			// Ejecutar request
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Error al hacer request: %v", err)
			}

			// Verificar status code
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Status code = %d, want %d", resp.StatusCode, tt.expectedStatus)
			}

			// Verificar respuesta si hay función de verificación
			if tt.checkResponse != nil {
				tt.checkResponse(t, resp)
			}
		})
	}
}

// createTestToken crea un token JWT válido para pruebas
func createTestToken(t *testing.T, secret string) string {
	claims := &middleware.Claims{
		Username: "admin",
		ID:       1,
		Role:     "admin",
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Error al crear token de prueba: %v", err)
	}
	return tokenString
}

