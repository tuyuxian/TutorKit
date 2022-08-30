// Code generated by ent, DO NOT EDIT.

package entpost

import (
	"backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTimestamp), v))
	})
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.EntPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTimestamp), v...))
	})
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.EntPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTimestamp), v...))
	})
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTimestamp), v))
	})
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTimestamp), v))
	})
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTimestamp), v))
	})
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTimestamp), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.EntPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.EntPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentIsNil applies the IsNil predicate on the "content" field.
func ContentIsNil() predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldContent)))
	})
}

// ContentNotNil applies the NotNil predicate on the "content" field.
func ContentNotNil() predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldContent)))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// ShareEQ applies the EQ predicate on the "share" field.
func ShareEQ(v Share) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShare), v))
	})
}

// ShareNEQ applies the NEQ predicate on the "share" field.
func ShareNEQ(v Share) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShare), v))
	})
}

// ShareIn applies the In predicate on the "share" field.
func ShareIn(vs ...Share) predicate.EntPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldShare), v...))
	})
}

// ShareNotIn applies the NotIn predicate on the "share" field.
func ShareNotIn(vs ...Share) predicate.EntPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldShare), v...))
	})
}

// ShareIsNil applies the IsNil predicate on the "share" field.
func ShareIsNil() predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldShare)))
	})
}

// ShareNotNil applies the NotNil predicate on the "share" field.
func ShareNotNil() predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldShare)))
	})
}

// HasComment applies the HasEdge predicate on the "comment" edge.
func HasComment() predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentTable, CommentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentWith applies the HasEdge predicate on the "comment" edge with a given conditions (other predicates).
func HasCommentWith(preds ...predicate.EntComment) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentTable, CommentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBelongsTo applies the HasEdge predicate on the "belongsTo" edge.
func HasBelongsTo() predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BelongsToTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BelongsToTable, BelongsToColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBelongsToWith applies the HasEdge predicate on the "belongsTo" edge with a given conditions (other predicates).
func HasBelongsToWith(preds ...predicate.EntCourse) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BelongsToInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BelongsToTable, BelongsToColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwnedBy applies the HasEdge predicate on the "ownedBy" edge.
func HasOwnedBy() predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnedByTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnedByTable, OwnedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnedByWith applies the HasEdge predicate on the "ownedBy" edge with a given conditions (other predicates).
func HasOwnedByWith(preds ...predicate.EntUser) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnedByInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnedByTable, OwnedByColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EntPost) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EntPost) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EntPost) predicate.EntPost {
	return predicate.EntPost(func(s *sql.Selector) {
		p(s.Not())
	})
}
