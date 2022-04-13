package migrator

import (
	"context"
	"github.com/117s/mdm/internal/connection"
	"github.com/117s/mdm/internal/connection/conn"
	"github.com/117s/mdm/internal/global"
	"github.com/117s/mdm/internal/schema"
	"github.com/117s/mdm/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMigrator_CreateTable(t *testing.T) {
	utils.PreTest()

	t.Run("success", func(t *testing.T) {
		dm := schema.DataModel{
			ID:        "user",
			Name:      "User",
			TableName: "users",
			Properties: []schema.Property{
				{
					ID:       "id",
					Name:     "User ID",
					Type:     "string",
					Required: true,
					Options: map[string]interface{}{
						"maxLength": 16,
					},
				},
			},
			PrimaryKeys: []string{"id"},
		}

		opt := connection.Options{
			DatabaseType:     "postgresql",
			ConnectionString: "host=127.0.0.1 user=postgres password=auth dbname=postgres port=5432",
			Timezone:         "Asia/Shanghai",
			Logger:           global.Log,
		}

		c, err := conn.NewConnection(context.Background(), opt)
		assert.Nil(t, err)

		m := &Migrator{
			Conn:   c,
			Logger: global.Log,
		}
		err = m.CreateTable(context.Background(), dm)
		assert.Nil(t, err)
	})
}
