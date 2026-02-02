package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	// Подключение к Postgres
	return pgx.Connect(
		ctx,
		"postgres://mydb:mydb@localhost:5432/mydb?sslmode=disable",
	)
}

