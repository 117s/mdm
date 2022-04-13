package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/117s/mdm/internal/connection"
	"github.com/117s/mdm/internal/schema"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

type PostgresqlDialector struct {
	connection.Options

	DB *sql.DB
}

func (c *PostgresqlDialector) Connect(ctx context.Context, options connection.Options) (*sql.DB, error) {
	config, err := pgx.ParseConfig(options.ConnectionString)
	if err != nil {
		return nil, err
	}
	c.ConnectionString = options.ConnectionString

	config.PreferSimpleProtocol = false

	if options.Timezone != "" {
		config.RuntimeParams["timezone"] = options.Timezone
		c.Timezone = options.Timezone
	}

	if options.Logger != nil {
		config.Logger = NewLogger(options.Logger)
		c.Logger = options.Logger
	}

	c.DB = stdlib.OpenDB(*config)
	return c.DB, nil
}

// GetFullDataType returns the database specific data type of Property
// https://www.postgresql.org/docs/current/datatype.html
func (c *PostgresqlDialector) GetFullDataType(prop schema.Property) string {
	switch prop.Type {
	case schema.TextField:
		return "text"
	case schema.VarcharField:
		opt, _ := prop.GetOptionsForVarCharField()
		// varchar(n)
		return fmt.Sprintf("varchar(%d)", opt.MaxLength)
	case schema.BooleanField:
		return "boolean"
	case schema.DatetimeField:
		return "timestamp"
	case schema.IntField:
		return "int"
	case schema.FloatField:
		return "real"
	case schema.DoubleField:
		return "double"
	default:
		return ""
	}
}
