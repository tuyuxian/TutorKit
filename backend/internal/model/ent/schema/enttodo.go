package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EntTodo holds the schema definition for the EntTodo entity.
type EntTodo struct {
	ent.Schema
}

// Time fields
func (EntTodo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the EntTodo.
func (EntTodo) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date"),
		field.Time("startTime"),
		field.Time("endTime"),
		field.Time("day"),
		field.String("lesson").
			Optional(),
		field.String("homework").
			Optional(),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS"),
	}
}

// Edges of the EntTodo.
func (EntTodo) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "todoFor" between EntTodo and EntCourse (M to 1)
		edge.From("todoFor", EntCourse.Type).
			Ref("todo").
			Unique(),

		// Create an inverse-edge called "ownedBy" between EntTodo and EntUser (M to 1)
		edge.From("ownedBy", EntUser.Type).
			Ref("todo").
			Unique(),
	}
}
