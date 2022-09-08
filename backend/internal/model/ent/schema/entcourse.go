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
		field.String("name").NotEmpty(),
		field.String("courseUrl").Optional(),
		field.Enum("paymentMethod").Values("times", "hours").Optional(),
		field.Float("paymentAmount").Optional(),
		field.Time("startDate").Optional(),
		field.Time("endDate").Optional(),
		field.Bool("monday").Optional(),
		field.Bool("tuesday").Optional(),
		field.Bool("wednesday").Optional(),
		field.Bool("thursday").Optional(),
		field.Bool("friday").Optional(),
		field.Bool("saturday").Optional(),
		field.Bool("sunday").Optional(),
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
