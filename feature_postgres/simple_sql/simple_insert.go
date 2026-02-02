package simple_sql

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn, 
	task TaskModel,
) error {
	query := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ($1, $2, FALSE, NOW())
	ON CONFLICT (title) DO NOTHING;
`
	_, err := conn.Exec(
		ctx, 
		query,
		task.Title, 
		task.Description,
		task.Completed,
		task.CreatedAt,
	) 
	return err
}
