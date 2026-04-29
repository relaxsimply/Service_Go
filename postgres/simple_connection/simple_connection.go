package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateConnetcion(ctx context.Context) (*pgx.Conn, error) {

	return pgx.Connect(ctx, "postgres://postgres:2007@localhost:5432/postgres")

}
