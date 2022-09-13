package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EntAttendance holds the schema definition for the EntAttendance entity.
type EntAttendance struct {
	ent.Schema
}

// Time fields
func (EntAttendance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the EntAttendance.
func (EntAttendance) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date"),
		field.Time("startTime"),
		field.Time("endTime"),
		field.Time("day"),
		field.String("note").
			Optional(),
		field.Float("hours").
			Positive().
			Optional(),
		field.Bool("checkedByTutor").
			Default(false),
		field.Bool("checkedByStudent").
			Default(false),
		field.Bool("checkedByParent").
			Default(false),
	}
}

// Edges of the EntAttendance.
func (EntAttendance) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "attendanceFor" between EntAttendance and EntCourse (M to 1)
		edge.From("attendanceFor", EntCourse.Type).
			Ref("attendance").
			Unique(),

		// Create an inverse-edge called "ownedBy" between EntAttendance and EntUser (M to 1)
		edge.From("ownedBy", EntUser.Type).
			Ref("attendance").
			Unique(),
	}
}
