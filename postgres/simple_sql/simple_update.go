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
