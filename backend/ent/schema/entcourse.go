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
		// Create an inverse-edge called "parent/children" between EntUser
		edge.From("courseOwner", EntUser.Type).
			Ref("course"),
	}
}
