package routes

import (
	"log"

	"example.com/go_basics/go/api"      // wygenerowany kod z openapi.yaml
	"example.com/go_basics/go/handlers" // Twoje handlery

	"github.com/labstack/echo/v4"
)

func NewEchoRouter(h *handlers.Handlers) *echo.Echo {
	e := echo.New()

	// Załaduj specyfikację OpenAPI
	_, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("Failed to load OpenAPI spec: %v", err)
	}

	// Rejestracja handlerów zgodnych z interfejsem ServerInterface
	api.RegisterHandlers(e, h)

	return e
}
