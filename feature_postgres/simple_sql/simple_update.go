package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn, task Task) error {
	query := `
	UPDATE tasks 	
	SET title=$1, description=$2, completed=$3, completed_at=$4
	WHERE id=$5;
`
	_, err := conn.Exec(ctx, query, 
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		task.ID,
	)

	return err
}

func UpdateTask(
	ctx context.Context,
	conn *pgx.Conn,
	task TaskModel,
) error {
	query := `
	UPDATE tasks 
	SET title=$1, description=$2, completed=$3, completed_at=$4
	WHERE id=$5
	`

	_, err := conn.Exec(ctx, query,
		task.Title,       // $1
		task.Description, // $2
		task.Completed,   // $3
		task.CompletedAt, // $4
		task.ID,          // $5
	)
	return err

	return err
}