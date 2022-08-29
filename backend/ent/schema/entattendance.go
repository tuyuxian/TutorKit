package schema

import "entgo.io/ent"

// EntAttendance holds the schema definition for the EntAttendance entity.
type EntAttendance struct {
	ent.Schema
}

// Fields of the EntAttendance.
func (EntAttendance) Fields() []ent.Field {
	return nil
}

// Edges of the EntAttendance.
func (EntAttendance) Edges() []ent.Edge {
	return nil
}
