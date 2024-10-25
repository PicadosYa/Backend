package database

import (
	"context"
	"fmt"
	"time"

	"picadosYa/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
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
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("no se pudo conectar a la base de datos: %w", err)
}
