package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"picadosYa/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	fmt.Println("Connection string:", connectionString)

	var db *sqlx.DB
	var err error
	for i := 0; i < 10; i++ { // Intenta 10 veces
		db, err = sqlx.ConnectContext(ctx, "mysql", connectionString)
		if err == nil {
			return db, nil
		}
		fmt.Println("Error conectando a la base de datos, reintentando en 5 segundos...")
		fmt.Println("Error:", err)
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("no se pudo conectar a la base de datos: %w", err)
}
