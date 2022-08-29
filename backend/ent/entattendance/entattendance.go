// Code generated by ent, DO NOT EDIT.

package entattendance

const (
	// Label holds the string label denoting the entattendance type in the database.
	Label = "ent_attendance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the entattendance in the database.
	Table = "ent_attendances"
)

// Columns holds all SQL columns for entattendance fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
