package main

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/fx"

	"example.com/go_basics/go/db"
	"example.com/go_basics/go/handlers"
	"example.com/go_basics/go/repository"
	"example.com/go_basics/go/routes"
	"example.com/go_basics/go/testdata"

	"github.com/labstack/echo/v4"
)

func main() {
	app := fx.New(
		fx.Provide(
			db.New,
			repository.New,
			handlers.New,
			routes.NewEchoRouter,
		),
		fx.Invoke(
			StartEchoServer,
			testdata.LoadTestData,
		),
	)
	app.Run()
}

func StartEchoServer(lc fx.Lifecycle, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			server := &http.Server{
				Addr:    ":8080",
				Handler: e,
			}
			fmt.Println("Starting Echo server at :8080")
			go func() {
				if err := e.StartServer(server); err != nil && err != http.ErrServerClosed {
					e.Logger.Error("Echo server stopped with error:", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
