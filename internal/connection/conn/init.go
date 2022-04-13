package conn

import (
	"context"
	"github.com/117s/mdm/internal/connection"
	"github.com/117s/mdm/internal/connection/postgresql"
)

func NewConnection(ctx context.Context, options connection.Options) (*connection.Connection, error) {
	dialector := GetDialector(options.DatabaseType)
	conn, err := connection.Connect(ctx, dialector, options)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GetDialector(dbType string) connection.Dialector {
	switch dbType {
	case "postgresql":
		return &postgresql.PostgresqlDialector{}
	default:
		return nil
	}
}
