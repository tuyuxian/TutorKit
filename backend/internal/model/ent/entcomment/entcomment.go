// Code generated by ent, DO NOT EDIT.

package entcomment

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the entcomment type in the database.
	Label = "ent_comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldShare holds the string denoting the share field in the database.
	FieldShare = "share"
	// EdgeBelongsTo holds the string denoting the belongsto edge name in mutations.
	EdgeBelongsTo = "belongsTo"
	// EdgeOwnedBy holds the string denoting the ownedby edge name in mutations.
	EdgeOwnedBy = "ownedBy"
	// Table holds the table name of the entcomment in the database.
	Table = "ent_comments"
	// BelongsToTable is the table that holds the belongsTo relation/edge.
	BelongsToTable = "ent_comments"
	// BelongsToInverseTable is the table name for the EntPost entity.
	// It exists in this package in order to avoid circular dependency with the "entpost" package.
	BelongsToInverseTable = "ent_posts"
	// BelongsToColumn is the table column denoting the belongsTo relation/edge.
	BelongsToColumn = "ent_post_comment"
	// OwnedByTable is the table that holds the ownedBy relation/edge.
	OwnedByTable = "ent_comments"
	// OwnedByInverseTable is the table name for the EntUser entity.
	// It exists in this package in order to avoid circular dependency with the "entuser" package.
	OwnedByInverseTable = "ent_users"
	// OwnedByColumn is the table column denoting the ownedBy relation/edge.
	OwnedByColumn = "ent_user_comment"
)

// Columns holds all SQL columns for entcomment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTimestamp,
	FieldContent,
	FieldShare,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ent_comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"ent_post_comment",
	"ent_user_comment",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultTimestamp holds the default value on creation for the "timestamp" field.
	DefaultTimestamp time.Time
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
)

// Share defines the type for the "share" enum field.
type Share string

// SharePublic is the default value of the Share enum.
const DefaultShare = SharePublic

// Share values.
const (
	SharePublic  Share = "PUBLIC"
	SharePrivate Share = "PRIVATE"
)

func (s Share) String() string {
	return string(s)
}

// ShareValidator is a validator for the "share" field enum values. It is called by the builders before save.
func ShareValidator(s Share) error {
	switch s {
	case SharePublic, SharePrivate:
		return nil
	default:
		return fmt.Errorf("entcomment: invalid enum value for share field: %q", s)
	}
}
