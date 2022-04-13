package connection

import (
	"context"
	"database/sql"
	"github.com/117s/mdm/internal/schema"
	"go.uber.org/zap"
)

type Connection struct {
	DB *sql.DB
	Dialector
}

// Connect connects to a database based on given dialector
func Connect(ctx context.Context, dialector Dialector, options Options) (*Connection, error) {
	db, err := dialector.Connect(ctx, options)
	if err != nil {
		return nil, err
	}
	conn := &Connection{
		DB:        db,
		Dialector: dialector,
	}
	return conn, nil
}

type Dialector interface {
	Connect(context context.Context, options Options) (*sql.DB, error)
	GetFullDataType(property schema.Property) string
}

type Options struct {
	DatabaseType     string
	ConnectionString string
	Timezone         string
	Logger           *zap.Logger
}
