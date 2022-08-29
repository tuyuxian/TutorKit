package schema

import "entgo.io/ent"

// EntTodo holds the schema definition for the EntTodo entity.
type EntTodo struct {
	ent.Schema
}

// Fields of the EntTodo.
func (EntTodo) Fields() []ent.Field {
	return nil
}

// Edges of the EntTodo.
func (EntTodo) Edges() []ent.Edge {
	return nil
}
