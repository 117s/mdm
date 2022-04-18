package parser

import (
	"encoding/json"
	"errors"
	"github.com/117s/mdm/internal/global"
	"github.com/117s/mdm/internal/schema"
)

// ValidateJSON checks if the input json string is a valid DataModel
func ValidateJSON(target string) (*schema.DataModel, error) {
	global.Log.Sugar().Debugf("validate target: %s", target)

	var val schema.DataModel
	err := json.Unmarshal([]byte(target), &val)
	if err != nil {
		global.Log.Sugar().Warnf("failed to marshal json, err: %s", err.Error())
		return nil, err
	}

	err = global.Validator.Struct(val)
	if err != nil {
		global.Log.Sugar().Warnf("input json is not valid, err: %s", err.Error())
		return nil, err
	}
	global.Log.Debug("pass schema validation")

	err = validateProperties(&val)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func validateProperties(dm *schema.DataModel) error {
	if len(dm.Properties) == 0 {
		return errors.New("missing required DataModel properties")
	}

	m := make(map[string]schema.Property, len(dm.Properties))
	for _, property := range dm.Properties {
		if _, ok := m[property.ID]; ok {
			return errors.New("duplicated property: " + property.ID)
		}
		m[property.ID] = property
		err := property.Validate()
		if err != nil {
			return err
		}
	}

	// check primary keys
	for _, key := range dm.PrimaryKeys {
		pk, ok := m[key]
		if !ok {
			return errors.New("missing required primary key property")
		}
		if pk.Required == false {
			return errors.New("PrimaryKey must be a required property")
		}
	}
	return nil
}
