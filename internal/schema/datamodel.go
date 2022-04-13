/*
Package schema implements the DataModel management.
*/
package schema

// DataModel is a notation for describing data or information. Including
//
// Structure of the data
//
// Operations on the data. Users are generally allowed to perform a limited set of actions on the data, e.g., a
// property is readonly, while another property (think about the password field of DataModel User) is write-only
//
// Constraints on the data. Describe limitations on what the data can be.
//
type DataModel struct {
	ID          string     `validate:"required" json:"id"`
	Name        string     `validate:"required" json:"name"`
	TableName   string     `json:"tableName"`
	Properties  []Property `validate:"required" json:"properties"`
	PrimaryKeys []string   `validate:"required" json:"-"`
}
