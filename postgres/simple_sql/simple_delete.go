package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	DELETE FROM tasks
	WHERE id = 4;
	`

	_, err := conn.Exec(ctx, sqlQuery)
	return err

}
