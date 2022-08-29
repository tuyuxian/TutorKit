package schema

import "entgo.io/ent"

// EntComment holds the schema definition for the EntComment entity.
type EntComment struct {
	ent.Schema
}

// Fields of the EntComment.
func (EntComment) Fields() []ent.Field {
	return nil
}

// Edges of the EntComment.
func (EntComment) Edges() []ent.Edge {
	return nil
}
