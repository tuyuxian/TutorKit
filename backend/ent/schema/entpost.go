package schema

import "entgo.io/ent"

// EntPost holds the schema definition for the EntPost entity.
type EntPost struct {
	ent.Schema
}

// Fields of the EntPost.
func (EntPost) Fields() []ent.Field {
	return nil
}

// Edges of the EntPost.
func (EntPost) Edges() []ent.Edge {
	return nil
}
