package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const emailRegex string = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
const phoneRegex string = "^[\\+]?[(]?[0-9]{3}[)]?[-\\s\\.]?[0-9]{3}[-\\s\\.]?[0-9]{4,6}$"

// EntUser holds the schema definition for the EntUser entity.
type EntUser struct {
	ent.Schema
}

// Time fields
func (EntUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (EntUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(20),
		field.String("email").
			NotEmpty().
			Match(regexp.MustCompile(emailRegex)),
		field.String("password").
			NotEmpty(),
		field.String("country").
			Optional(),
		field.String("phone").
			Optional().
			Match(regexp.MustCompile(phoneRegex)),
		field.Time("dateOfBirth").
			Optional(),
		field.String("profilePictureUrl").
			Optional(),
		field.Bool("isTutor").
			Default(false),
		field.Bool("isStudent").
			Default(false),
		field.Bool("isParent").
			Default(false),
	}
}

// Edges of the User.
func (EntUser) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an edge called "course" between EntUser and EntCourse (M to M)
		edge.To("course", EntCourse.Type),

		// Create an edge called "todo" between EntUser and EntTodo (M to M)
		edge.To("todo", EntTodo.Type),

		// Create an edge called "attendance" between EntUser and EntAttendance (M to M)
		edge.To("attendance", EntAttendance.Type),

		// Create an edge called "post" between EntUser and EntPost (1 to M)
		edge.To("post", EntPost.Type),

		// Create an edge called "comment" between EntUser and EntPost (1 to M)
		edge.To("comment", EntComment.Type),

		// Create an edge called "join" between EntUser and EntCourse (M to M)
		edge.To("join", EntCourse.Type),

		// Below are the edge between EntUser
		// Create an bidirectional edge called "parent/children" (M to M)
		edge.To("parent", EntUser.Type).
			From("children"),

		// Create an bidirectional edge called "tutor/student" (M to M)
		edge.To("tutor", EntUser.Type).
			From("student"),

		// Create an bidirectional edge called "STutor/SParent" (M to M)
		edge.To("STutor", EntUser.Type).
			From("SParent"),

		// Create an inverse-edge called "canSee" between EntUser and EntPost (M to M)
		edge.From("canSee", EntPost.Type).
			Ref("shareWith"),
	}
}
