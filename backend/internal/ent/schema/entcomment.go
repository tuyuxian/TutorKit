package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EntComment holds the schema definition for the EntComment entity.
type EntComment struct {
	ent.Schema
}

// Time fields
func (EntComment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the EntComment.
func (EntComment) Fields() []ent.Field {
	return []ent.Field{
		field.Time("timestamp").Default(time.Now()),
		field.String("content").Optional(),
		field.Enum("share").Values("public", "private").Optional(),
	}
}

// Edges of the EntComment.
func (EntComment) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "belongsTo" between EntComment and EntPost (M to 1)
		edge.From("belongsTo", EntPost.Type).
			Ref("comment").
			Unique(),

		// Create an inverse-edge called "ownedBy" between EntComment and EntUser (M to 1)
		edge.From("ownedBy", EntUser.Type).
			Ref("comment").
			Unique(),
	}
}
