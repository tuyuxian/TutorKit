package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EntPost holds the schema definition for the EntPost entity.
type EntPost struct {
	ent.Schema
}

// Time fields
func (EntPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the EntPost.
func (EntPost) Fields() []ent.Field {
	return []ent.Field{
		field.Time("timestamp").Default(time.Now()),
		field.String("content").Optional(),
		field.Enum("share").Values("public", "private").Optional(),
	}
}

// Edges of the EntPost.
func (EntPost) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an edge called "comment" between EntPost and EntComment (1 to M)
		edge.To("comment", EntComment.Type),

		// Create an edge called "shareWith" between EntPost and EntUser (M to M)
		edge.To("shareWith", EntUser.Type),

		// Create an inverse-edge called "belongsTo" between EntPost and EntCourse (M to 1)
		edge.From("belongsTo", EntCourse.Type).
			Ref("post").
			Unique(),

		// Create an inverse-edge called "ownedBy" between EntPost and EntUser (M to 1)
		edge.From("ownedBy", EntUser.Type).
			Ref("post").
			Unique(),
	}
}
