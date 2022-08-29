package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const emailRegex string = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

// EntUser holds the schema definition for the EntUser entity.
type EntUser struct {
	ent.Schema
}

// Fields of the User.
func (EntUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(), // TODO: name length limit
		field.String("email").
			NotEmpty().
			Match(regexp.MustCompile(emailRegex)),
		field.String("password").NotEmpty(),
		field.String("phone").Optional(), // TODO: phone regex
		field.String("profilePictureUrl").Optional(),
		field.Bool("isTutor"),
		field.Bool("isStudent"),
		field.Bool("isParent"),
	}
}

// Edges of the User.
func (EntUser) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an edge called "course" to EntCourse
		edge.To("course", EntCourse.Type),

		// Below are the edge between EntUser

		// Create an bidirectional edge called "parent/children"
		edge.To("parent", EntUser.Type).
			From("children"),
		// Create an bidirectional edge called "tutor/student"
		edge.To("tutor", EntUser.Type).
			From("student"),
		// Create an bidirectional edge called "STutor/SParent"
		edge.To("STutor", EntUser.Type).
			From("SParent"),
	}
}
