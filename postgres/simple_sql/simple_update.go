package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	UPDATE tasks
	SET description = ':)'
	WHERE completed = FALSE;
	`
	_, err := conn.Exec(ctx, sqlQuery)
	return err

}

func UpdateTitle(ctx context.Context,
	conn *pgx.Conn,
	id int,
	newTitle string,
) error {

	sqlQuery := `
	UPDATE tasks
	SET title=$1
	WHERE id=$2
	`

	_, err := conn.Exec(ctx, sqlQuery, newTitle, id)

	return err
}

func UpdateDescription(
	ctx context.Context,
	conn *pgx.Conn,

	id int,
	newDescription string,

) error {

	sqlQuery := `
	UPDATE tasks
	SET description=$1
	WHERE id=$2
	`
	_, err := conn.Exec(ctx, sqlQuery, newDescription, id)
	return err
}

func UpdateTask(
	ctx context.Context,
	conn *pgx.Conn,
	task TaskModel,
) error {

	sqlQuery := `
	UPDATE tasks
	SET title=$1, description=$2, completed=$3, created_at=$4, completed_at=$5
	WHERE id=$6;
	`

	_, err := conn.Exec(
		ctx,
		sqlQuery,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompletedAt,
		task.ID,
	)

	return err

}
