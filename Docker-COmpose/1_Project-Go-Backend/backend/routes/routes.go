package routes

import (
	"employees/controller"
	"github.com/gofiber/fiber/v2"
)

// RegisterRoute registers all application routes
func RegisterRoute(app *fiber.App, employeeController *controller.Employee) {
	// Health check endpoints for Kubernetes and ALB probes
	app.Get("/healthz", HealthCheck)
	app.Get("/health", HealthCheck)

	// Employee routes - support both root and /api prefixed routes
	if employeeController != nil {
		app.Post("/employees", employeeController.Create)
		app.Get("/employees", employeeController.GetAll)
		app.Post("/api/employees", employeeController.Create)
		app.Get("/api/employees", employeeController.GetAll)
	}
}

// HealthCheck handles liveness/readiness probe requests
func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "healthy",
	})
}
