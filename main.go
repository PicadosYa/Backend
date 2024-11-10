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

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	_ "picadosYa/docs" // Importa el paquete generado por swag init

	echoSwagger "github.com/swaggo/echo-swagger" // Importa echo-swagger
)

// @title PicadosYa API
// @version 1.0
// @descrition API para la administracion de canchas de futbol 5
// @host localhost:8080
// @BasePath /api
// @schema http
func main() {

	if err := godotenv.Load("./.env"); err != nil {
		panic(err)
	}

	app := fx.New(
		fx.Provide(
			context.Background,
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
func setLifeCycle(lc fx.Lifecycle, a *api.API, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT"))

			// Agrega la ruta de Swagger
			e.GET("/swagger/*", echoSwagger.WrapHandler)

			go a.Start(e, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
