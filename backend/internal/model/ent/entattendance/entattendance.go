// Code generated by ent, DO NOT EDIT.

package entattendance

import (
	"time"
)

const (
	// Label holds the string label denoting the entattendance type in the database.
	Label = "ent_attendance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldStartTime holds the string denoting the starttime field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the endtime field in the database.
	FieldEndTime = "end_time"
	// FieldDay holds the string denoting the day field in the database.
	FieldDay = "day"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldHours holds the string denoting the hours field in the database.
	FieldHours = "hours"
	// FieldCheckedByTutor holds the string denoting the checkedbytutor field in the database.
	FieldCheckedByTutor = "checked_by_tutor"
	// FieldCheckedByStudent holds the string denoting the checkedbystudent field in the database.
	FieldCheckedByStudent = "checked_by_student"
	// FieldCheckedByParent holds the string denoting the checkedbyparent field in the database.
	FieldCheckedByParent = "checked_by_parent"
	// EdgeAttendanceFor holds the string denoting the attendancefor edge name in mutations.
	EdgeAttendanceFor = "attendanceFor"
	// EdgeOwnedBy holds the string denoting the ownedby edge name in mutations.
	EdgeOwnedBy = "ownedBy"
	// Table holds the table name of the entattendance in the database.
	Table = "ent_attendances"
	// AttendanceForTable is the table that holds the attendanceFor relation/edge.
	AttendanceForTable = "ent_attendances"
	// AttendanceForInverseTable is the table name for the EntCourse entity.
	// It exists in this package in order to avoid circular dependency with the "entcourse" package.
	AttendanceForInverseTable = "ent_courses"
	// AttendanceForColumn is the table column denoting the attendanceFor relation/edge.
	AttendanceForColumn = "ent_course_attendance"
	// OwnedByTable is the table that holds the ownedBy relation/edge.
	OwnedByTable = "ent_attendances"
	// OwnedByInverseTable is the table name for the EntUser entity.
	// It exists in this package in order to avoid circular dependency with the "entuser" package.
	OwnedByInverseTable = "ent_users"
	// OwnedByColumn is the table column denoting the ownedBy relation/edge.
	OwnedByColumn = "ent_user_attendance"
)

// Columns holds all SQL columns for entattendance fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDate,
	FieldStartTime,
	FieldEndTime,
	FieldDay,
	FieldNote,
	FieldHours,
	FieldCheckedByTutor,
	FieldCheckedByStudent,
	FieldCheckedByParent,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ent_attendances"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"ent_course_attendance",
	"ent_user_attendance",
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
	// HoursValidator is a validator for the "hours" field. It is called by the builders before save.
	HoursValidator func(float64) error
	// DefaultCheckedByTutor holds the default value on creation for the "checkedByTutor" field.
	DefaultCheckedByTutor bool
	// DefaultCheckedByStudent holds the default value on creation for the "checkedByStudent" field.
	DefaultCheckedByStudent bool
	// DefaultCheckedByParent holds the default value on creation for the "checkedByParent" field.
	DefaultCheckedByParent bool
)
