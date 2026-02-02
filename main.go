package main

import (
	"context"
	"fmt"
	"time"

	"study/feature_postgres/simple_connection"
	"study/feature_postgres/simple_sql"
)

func main() {
	ctx := context.Background()

	// Создаём соединение
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	fmt.Println("succeed!")

	// SELECT всех задач
	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n--- Задачи до обновления ---")
	for _, task := range tasks {
		fmt.Printf("[%d] %s — %s (completed=%v)\n",
			task.ID, task.Title, task.Description, task.Completed)
	}

	// Обновляем задачу с ID=5
	for i := range tasks {
		if tasks[i].ID == 5 {
			tasks[i].Title = "Покормить кошку"
			tasks[i].Description = "Отсыпать кошке 30 грамм корма"
			tasks[i].Completed = true
			now := time.Now()
			tasks[i].CompletedAt = &now

			// Сохраняем изменения в базе
			if err := simple_sql.UpdateRow(ctx, conn, tasks[i]); err != nil {
				panic(err)
			}
		}
	}

	// SELECT после обновления
	tasks, err = simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n--- Задачи после обновления ---")
	for _, task := range tasks {
		fmt.Printf("[%d] %s — %s (completed=%v)\n",
			task.ID, task.Title, task.Description, task.Completed)
		if task.CompletedAt != nil {
			fmt.Println("completedAt:", task.CompletedAt.Format(time.RFC3339))
		}
	}
}
