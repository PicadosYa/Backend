package main


// Comentario de prueba
import (
	"context"
	"fmt"
	"os"
	"picadosYa/database"
	"picadosYa/internal/api"
	"picadosYa/internal/repository"
	"picadosYa/internal/service"
	"picadosYa/settings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}

	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			repository.NewFieldRepository,
			repository.NewReservationRepository,
			service.New,
			service.NewFieldService,
			service.NewReservationService,
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
			address := fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT"))
			go a.Start(e, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
