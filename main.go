package main

import (
	"context"
	"fmt"

	"picadosYa/database"
	"picadosYa/internal/api"
	"picadosYa/internal/repository"
	"picadosYa/internal/service"
	"picadosYa/settings"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			repository.NewFieldRepository,
			service.New,
			service.NewFieldService,
			api.New,
			echo.New,
		),
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

// la app se mantiene indefinidamente
func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%d", s.Port)
			go a.Start(e, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
