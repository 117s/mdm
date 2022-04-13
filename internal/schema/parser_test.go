package schema

import (
	"encoding/json"
	"github.com/117s/mdm/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateJSON(t *testing.T) {
	utils.PreTest()

	t.Run("id is required for data type", func(t *testing.T) {
		schemaJson := map[string]interface{}{
			"name": "User",
			"properties": []interface{}{
				map[string]interface{}{
					"id":       "cellphone",
					"type":     "string",
					"required": true,
				},
			},
		}
		schemaBytes, _ := json.Marshal(schemaJson)
		schemaStr := string(schemaBytes)

		_, err := ValidateJSON(schemaStr)
		assert.NotNil(t, err, "123")
	})

	t.Run("validate without invalid json", func(t *testing.T) {

	})

	t.Run("success", func(t *testing.T) {
		schemaJson := map[string]interface{}{
			"id":   "user",
			"name": "User",
			"properties": []interface{}{
				map[string]interface{}{
					"id":       "cellphone",
					"type":     "string",
					"required": true,
				},
			},
		}
		schemaBytes, _ := json.Marshal(schemaJson)
		schemaStr := string(schemaBytes)

		ValidateJSON(schemaStr)
	})
}

func TestValidateProperties(t *testing.T) {
	utils.PreTest()

	t.Run("all properties should have unique id", func(t *testing.T) {
		dm := DataModel{
			ID:        "user",
			Name:      "User",
			TableName: "users",
			Properties: []Property{
				{
					ID:       "id",
					Name:     "User ID",
					Type:     "string",
					Required: true,
				},
				{
					ID:       "id",
					Name:     "User ID",
					Type:     "string",
					Required: true,
				},
			},
			PrimaryKeys: []string{"id"},
		}
		err := validateProperties(&dm)
		assert.NotNil(t, err)
	})

	t.Run("pk must be required", func(t *testing.T) {
		dm := DataModel{
			ID:        "user",
			Name:      "User",
			TableName: "users",
			Properties: []Property{
				{
					ID:       "id",
					Name:     "User ID",
					Type:     "string",
					Required: false,
				},
			},
			PrimaryKeys: []string{"id"},
		}
		err := validateProperties(&dm)
		assert.NotNil(t, err)
		assert.Equal(t, "PrimaryKey must be a required property", err.Error())
	})

	t.Run("unsupported property type should fail", func(t *testing.T) {
		dm := DataModel{
			ID:        "user",
			Name:      "User",
			TableName: "users",
			Properties: []Property{
				{
					ID:       "id",
					Name:     "User ID",
					Type:     "Unknown",
					Required: false,
				},
			},
			PrimaryKeys: []string{"id"},
		}
		err := validateProperties(&dm)
		assert.NotNil(t, err)
		assert.Equal(t, "unsupported property type: Unknown", err.Error())
	})

	t.Run("success", func(t *testing.T) {
		dm := DataModel{
			ID:        "user",
			Name:      "User",
			TableName: "users",
			Properties: []Property{
				{
					ID:       "id",
					Name:     "User ID",
					Type:     "string",
					Required: true,
				},
			},
			PrimaryKeys: []string{"id"},
		}
		err := validateProperties(&dm)
		assert.Nil(t, err)
	})
}
