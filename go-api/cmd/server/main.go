package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"go-api/internal/controllers"
	"go-api/internal/handlers"
	"go-api/internal/middleware"
	"go-api/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

var (
	startTime = time.Now()
	version   = "1.0.0"
)

func main() {
	// Cargar variables de entorno desde .env (solo en desarrollo local)
	// En producción/Docker, las variables vienen del sistema
	if os.Getenv("NODE_ENV") != "production" {
		godotenv.Load()
	}

	// Obtener URL de Node.js desde variable de entorno o usar default
	nodeURL := os.Getenv("NODE_API_URL")
	if nodeURL == "" {
		nodeURL = "http://localhost:3001"
	}

	// Crear cliente para Node.js
	nodeClient := services.NewNodeClient(nodeURL)

	// Crear handler
	matrixHandler := handlers.NewMatrixHandler(nodeClient)

	// Crear app Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middlewares
	app.Use(recover.New())
	app.Use(logger.New())
	
	// CORS para permitir requests desde el frontend
	// IMPORTANTE: Debe estar ANTES de las rutas para manejar OPTIONS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, HEAD",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length",
		MaxAge:           86400, // 24 horas
	}))

	// Handler manual para OPTIONS (preflight requests)
	app.Options("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "go-api",
		})
	})

	// Rutas públicas (sin autenticación)
	app.Post("/auth/login", controllers.Login)

	// Endpoint informativo del backend (público)
	app.Get("/", func(c *fiber.Ctx) error {
		uptime := time.Since(startTime)
		return c.JSON(fiber.Map{
			"service":       "Go API Backend",
			"version":       version,
			"technology":    "Go (Golang)",
			"framework":     "Fiber v2",
			"goVersion":     runtime.Version(),
			"startTime":     startTime.Format(time.RFC3339),
			"uptime":        uptime.String(),
			"uptimeSeconds": int(uptime.Seconds()),
			"os":            runtime.GOOS,
			"arch":          runtime.GOARCH,
			"nodeApiUrl":    nodeURL,
			"endpoints": fiber.Map{
				"health":        "GET /health",
				"login":          "POST /auth/login",
				"processMatrix": "POST /matrix/process (requiere JWT)",
				"info":          "GET /",
			},
		})
	})

	// Rutas protegidas (requieren JWT)
	app.Post("/matrix/process", middleware.AuthenticateToken, matrixHandler.ProcessMatrix)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Servidor Go iniciado en puerto %s", port)
	log.Printf("Conectado a Node.js API en: %s", nodeURL)

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error al iniciar servidor: %v", err)
	}
}
