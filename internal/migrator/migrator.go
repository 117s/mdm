package migrator

import (
	"context"
	"fmt"
	"github.com/117s/mdm/internal/connection"
	"github.com/117s/mdm/internal/schema"
	"go.uber.org/zap"
	"strings"
)

type Migrator struct {
	Conn   *connection.Connection
	Logger *zap.Logger
}

func (m *Migrator) CreateTable(ctx context.Context, dm schema.DataModel) error {
	var (
		createTableSQL = "CREATE TABLE " + dm.TableName + " ("
		values         []interface{}
	)

	for _, property := range dm.Properties {
		// pName := fmt.Sprintf("name%d", idx)
		// pType := fmt.Sprintf("type%d", idx)
		createTableSQL += fmt.Sprintf("%s %s", property.ID, m.Conn.GetFullDataType(property))
		// values = append(values, m.Conn.GetFullDataType(property))
	}

	createTableSQL = strings.TrimSuffix(createTableSQL, ",")

	createTableSQL += ")"

	m.Logger.Sugar().Debugf("create table sql: %s, values: %v", createTableSQL, values)
	_, err := m.Conn.DB.ExecContext(ctx, createTableSQL)
	if err != nil {
		m.Logger.Sugar().Warnf("create table error: %s", err.Error())
		return err
	}

	return nil
}
