package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EntCourse holds the schema definition for the EntCourse entity.
type EntCourse struct {
	ent.Schema
}

// Time fields
func (EntCourse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the EntCourse.
func (EntCourse) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("courseUrl").
			NotEmpty(),
		field.Enum("paymentMethod").
			NamedValues(
				"times", "TIMES",
				"hours", "HOURS",
			),
		field.Float("paymentAmount").
			Positive().
			Optional(),
		field.Time("startDate"),
		field.Time("endDate"),
		field.Bool("monday").
			Default(false),
		field.Bool("tuesday").
			Default(false),
		field.Bool("wednesday").
			Default(false),
		field.Bool("thursday").
			Default(false),
		field.Bool("friday").
			Default(false),
		field.Bool("saturday").
			Default(false),
		field.Bool("sunday").
			Default(false),
		field.Time("mondayStartTime").
			Optional(),
		field.Time("mondayEndTime").
			Optional(),
		field.Time("tuesdayStartTime").
			Optional(),
		field.Time("tuesdayEndTime").
			Optional(),
		field.Time("wednesdayStartTime").
			Optional(),
		field.Time("wednesdayEndTime").
			Optional(),
		field.Time("thursdayStartTime").
			Optional(),
		field.Time("thursdayEndTime").
			Optional(),
		field.Time("fridayStartTime").
			Optional(),
		field.Time("fridayEndTime").
			Optional(),
		field.Time("saturdayStartTime").
			Optional(),
		field.Time("saturdayEndTime").
			Optional(),
		field.Time("sundayStartTime").
			Optional(),
		field.Time("sundayEndTime").
			Optional(),
	}
}

// Edges of the EntCourse.
func (EntCourse) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an edge called "todo" between EntCourse and EntTodo (1 to M)
		edge.To("todo", EntTodo.Type),

		// Create an edge called "attendance" between EntCourse and EntAttendance (1 to M)
		edge.To("attendance", EntAttendance.Type),

		// Create an edge called "post" between EntCourse and EntPost (1 to M)
		edge.To("post", EntPost.Type),

		// Create an inverse-edge called "ownedBy" between EntCourse and EntUser (M to M)
		edge.From("ownedBy", EntUser.Type).
			Ref("course"),

		// Create an inverse-edge called "joinedBy" between EntCourse and EntUser (M to M)
		edge.From("joinedBy", EntUser.Type).
			Ref("join"),
	}
}
