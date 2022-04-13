package schema

import (
	"encoding/json"
	"errors"
)

// Property defines all necessary information of one field
type Property struct {
	// ID is used to identify this property, which will be used as the name of this column
	// should follow pattern ^[a-z_][a-z0-9_]*$
	ID string `json:"id" validate:"required"`
	// Name is used for business observers
	Name string `json:"name" validate:"required"`
	// Type can be any of "string", "text", "number", "boolean", "date"
	// string must have a max length, while text does not require this
	// string => VARCHAR
	// text => TEXT
	Type     string                 `json:"type" validate:"required"`
	Options  map[string]interface{} `json:"options"`
	Comments string                 `json:"comments"`
	Required bool                   `json:"required" validate:"required"`
}

type VarCharOptions struct {
	MaxLength int `json:"maxLength"`
}

var validPropertyType map[string]int8

const (
	TextField     = "text"
	VarcharField  = "string"
	IntField      = "int"
	FloatField    = "float"
	DoubleField   = "double"
	BooleanField  = "boolean"
	DatetimeField = "datetime"
)

func init() {
	validPropertyType = make(map[string]int8)
	validPropertyType[TextField] = 1
	validPropertyType[VarcharField] = 1
	validPropertyType[IntField] = 1
	validPropertyType[FloatField] = 1
	validPropertyType[DoubleField] = 1
	validPropertyType[BooleanField] = 1
	validPropertyType[DatetimeField] = 1
}

func (p *Property) validate() error {
	if _, ok := validPropertyType[p.Type]; !ok {
		return errors.New("unsupported property type: " + p.Type)
	}
	if p.Type == VarcharField {
		_, err := p.GetOptionsForVarCharField()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Property) GetOptionsForVarCharField() (*VarCharOptions, error) {
	// FIXME: need a better way to convert interface to Options
	optStr, err := json.Marshal(p.Options)
	if err != nil {
		return nil, err
	}
	var opt VarCharOptions
	err = json.Unmarshal(optStr, &opt)
	if err != nil {
		return nil, err
	}

	if opt.MaxLength <= 0 {
		return nil, errors.New("varchar field should have options.length >= 0")
	}
	return &opt, nil
}
