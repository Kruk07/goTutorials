package routes

import (
	"log"

	"example.com/go_basics/go/api"
	"example.com/go_basics/go/handlers"

	"github.com/labstack/echo/v4"
)

func NewEchoRouter(h *handlers.Handlers) *echo.Echo {
	e := echo.New()

	_, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("Failed to load OpenAPI spec: %v", err)
	}

	api.RegisterHandlers(e, h)

	return e
}
