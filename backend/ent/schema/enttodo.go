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

// Fields of the EntTodo.
func (EntTodo) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date"),
		field.Time("startTime").Optional(),
		field.Time("endTime").Optional(),
		field.Time("day").Optional(),
		field.String("lesson").Optional(),
		field.String("homework").Optional(),
		field.Enum("status").NamedValues(
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

		// Create an inverse-edge called "todoOwnedBy" between EntTodo and EntUser (M to 1)
		edge.From("ownedBy", EntUser.Type).
			Ref("todo").
			Unique(),
	}
}
