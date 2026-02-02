package simple_sql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// SelectRows возвращает все задачи
func SelectRows(ctx context.Context, conn *pgx.Conn) ([]Task, error) {
	query := `
		SELECT id, title, description, completed, created_at, completed_at
		FROM tasks
		ORDER BY id ASC;
	`

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var t Task
		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Completed,
			&t.CreatedAt,
			&t.CompletedAt,
		)
		if err != nil {
			return nil, err
		}

		// Печать одной задачи
		// printTask(t)

		tasks = append(tasks, t)
	}

	return tasks, nil
}

// printTask выводит одну задачу красиво
func printTask(t Task) {
	fmt.Println("------------------------------------")
	fmt.Println("id:", t.ID)
	fmt.Println("title:", t.Title)
	fmt.Println("description:", t.Description)
	fmt.Println("completed:", t.Completed)
	fmt.Println("createdAt:", t.CreatedAt)
	if t.CompletedAt != nil {
		fmt.Println("completedAt:", t.CompletedAt.Format(time.RFC3339))
	} else {
		fmt.Println("completedAt: not completed yet")
	}
}
