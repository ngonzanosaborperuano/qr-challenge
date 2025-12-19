package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func TestLogin(t *testing.T) {
	// Configurar JWT_SECRET para tests
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	app := fiber.New()
	app.Post("/auth/login", Login)

	tests := []struct {
		name           string
		body           LoginRequest
		expectedStatus int
		checkToken     bool
	}{
		{
			name: "login exitoso con credenciales válidas",
			body: LoginRequest{
				Username: "admin",
				Password: "admin",
			},
			expectedStatus: http.StatusOK,
			checkToken:     true,
		},
		{
			name: "login fallido con credenciales inválidas",
			body: LoginRequest{
				Username: "admin",
				Password: "wrong",
			},
			expectedStatus: http.StatusUnauthorized,
			checkToken:     false,
		},
		{
			name: "login fallido con usuario incorrecto",
			body: LoginRequest{
				Username: "user",
				Password: "admin",
			},
			expectedStatus: http.StatusUnauthorized,
			checkToken:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bodyJSON, _ := json.Marshal(tt.body)
			req := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(bodyJSON))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Error al hacer request: %v", err)
			}

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Status code = %d, want %d", resp.StatusCode, tt.expectedStatus)
			}

			if tt.checkToken {
				var response LoginResponse
				json.NewDecoder(resp.Body).Decode(&response)
				if !response.Success {
					t.Errorf("Expected success = true, got false")
				}
				if response.Token == "" {
					t.Errorf("Expected token, got empty string")
				}
				// Verificar que el token es válido
				claims := &jwt.RegisteredClaims{}
				_, err := jwt.ParseWithClaims(response.Token, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte("test-secret-key"), nil
				})
				if err != nil {
					t.Errorf("Token inválido: %v", err)
				}
			}
		})
	}
}

