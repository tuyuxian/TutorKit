// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/internal/ent/entattendance"
	"backend/internal/ent/entcourse"
	"backend/internal/ent/entuser"
	"backend/internal/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntAttendanceUpdate is the builder for updating EntAttendance entities.
type EntAttendanceUpdate struct {
	config
	hooks    []Hook
	mutation *EntAttendanceMutation
}

// Where appends a list predicates to the EntAttendanceUpdate builder.
func (eau *EntAttendanceUpdate) Where(ps ...predicate.EntAttendance) *EntAttendanceUpdate {
	eau.mutation.Where(ps...)
	return eau
}

// SetUpdatedAt sets the "updatedAt" field.
func (eau *EntAttendanceUpdate) SetUpdatedAt(t time.Time) *EntAttendanceUpdate {
	eau.mutation.SetUpdatedAt(t)
	return eau
}

// SetDate sets the "date" field.
func (eau *EntAttendanceUpdate) SetDate(t time.Time) *EntAttendanceUpdate {
	eau.mutation.SetDate(t)
	return eau
}

// SetStartTime sets the "startTime" field.
func (eau *EntAttendanceUpdate) SetStartTime(t time.Time) *EntAttendanceUpdate {
	eau.mutation.SetStartTime(t)
	return eau
}

// SetNillableStartTime sets the "startTime" field if the given value is not nil.
func (eau *EntAttendanceUpdate) SetNillableStartTime(t *time.Time) *EntAttendanceUpdate {
	if t != nil {
		eau.SetStartTime(*t)
	}
	return eau
}

// ClearStartTime clears the value of the "startTime" field.
func (eau *EntAttendanceUpdate) ClearStartTime() *EntAttendanceUpdate {
	eau.mutation.ClearStartTime()
	return eau
}

// SetEndTime sets the "endTime" field.
func (eau *EntAttendanceUpdate) SetEndTime(t time.Time) *EntAttendanceUpdate {
	eau.mutation.SetEndTime(t)
	return eau
}

// SetNillableEndTime sets the "endTime" field if the given value is not nil.
func (eau *EntAttendanceUpdate) SetNillableEndTime(t *time.Time) *EntAttendanceUpdate {
	if t != nil {
		eau.SetEndTime(*t)
	}
	return eau
}

// ClearEndTime clears the value of the "endTime" field.
func (eau *EntAttendanceUpdate) ClearEndTime() *EntAttendanceUpdate {
	eau.mutation.ClearEndTime()
	return eau
}

// SetDay sets the "day" field.
func (eau *EntAttendanceUpdate) SetDay(t time.Time) *EntAttendanceUpdate {
	eau.mutation.SetDay(t)
	return eau
}

// SetNillableDay sets the "day" field if the given value is not nil.
func (eau *EntAttendanceUpdate) SetNillableDay(t *time.Time) *EntAttendanceUpdate {
	if t != nil {
		eau.SetDay(*t)
	}
	return eau
}

// ClearDay clears the value of the "day" field.
func (eau *EntAttendanceUpdate) ClearDay() *EntAttendanceUpdate {
	eau.mutation.ClearDay()
	return eau
}

// SetNote sets the "note" field.
func (eau *EntAttendanceUpdate) SetNote(s string) *EntAttendanceUpdate {
	eau.mutation.SetNote(s)
	return eau
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (eau *EntAttendanceUpdate) SetNillableNote(s *string) *EntAttendanceUpdate {
	if s != nil {
		eau.SetNote(*s)
	}
	return eau
}

// ClearNote clears the value of the "note" field.
func (eau *EntAttendanceUpdate) ClearNote() *EntAttendanceUpdate {
	eau.mutation.ClearNote()
	return eau
}

// SetHours sets the "hours" field.
func (eau *EntAttendanceUpdate) SetHours(f float64) *EntAttendanceUpdate {
	eau.mutation.ResetHours()
	eau.mutation.SetHours(f)
	return eau
}

// SetNillableHours sets the "hours" field if the given value is not nil.
func (eau *EntAttendanceUpdate) SetNillableHours(f *float64) *EntAttendanceUpdate {
	if f != nil {
		eau.SetHours(*f)
	}
	return eau
}

// AddHours adds f to the "hours" field.
func (eau *EntAttendanceUpdate) AddHours(f float64) *EntAttendanceUpdate {
	eau.mutation.AddHours(f)
	return eau
}

// ClearHours clears the value of the "hours" field.
func (eau *EntAttendanceUpdate) ClearHours() *EntAttendanceUpdate {
	eau.mutation.ClearHours()
	return eau
}

// SetCheckedByTutor sets the "checkedByTutor" field.
func (eau *EntAttendanceUpdate) SetCheckedByTutor(b bool) *EntAttendanceUpdate {
	eau.mutation.SetCheckedByTutor(b)
	return eau
}

// SetNillableCheckedByTutor sets the "checkedByTutor" field if the given value is not nil.
func (eau *EntAttendanceUpdate) SetNillableCheckedByTutor(b *bool) *EntAttendanceUpdate {
	if b != nil {
		eau.SetCheckedByTutor(*b)
	}
	return eau
}

// SetCheckedByStudent sets the "checkedByStudent" field.
func (eau *EntAttendanceUpdate) SetCheckedByStudent(b bool) *EntAttendanceUpdate {
	eau.mutation.SetCheckedByStudent(b)
	return eau
}

// SetNillableCheckedByStudent sets the "checkedByStudent" field if the given value is not nil.
func (eau *EntAttendanceUpdate) SetNillableCheckedByStudent(b *bool) *EntAttendanceUpdate {
	if b != nil {
		eau.SetCheckedByStudent(*b)
	}
	return eau
}

// SetCheckedByParent sets the "checkedByParent" field.
func (eau *EntAttendanceUpdate) SetCheckedByParent(b bool) *EntAttendanceUpdate {
	eau.mutation.SetCheckedByParent(b)
	return eau
}

// SetNillableCheckedByParent sets the "checkedByParent" field if the given value is not nil.
func (eau *EntAttendanceUpdate) SetNillableCheckedByParent(b *bool) *EntAttendanceUpdate {
	if b != nil {
		eau.SetCheckedByParent(*b)
	}
	return eau
}

// SetAttendanceForID sets the "attendanceFor" edge to the EntCourse entity by ID.
func (eau *EntAttendanceUpdate) SetAttendanceForID(id int) *EntAttendanceUpdate {
	eau.mutation.SetAttendanceForID(id)
	return eau
}

// SetNillableAttendanceForID sets the "attendanceFor" edge to the EntCourse entity by ID if the given value is not nil.
func (eau *EntAttendanceUpdate) SetNillableAttendanceForID(id *int) *EntAttendanceUpdate {
	if id != nil {
		eau = eau.SetAttendanceForID(*id)
	}
	return eau
}

// SetAttendanceFor sets the "attendanceFor" edge to the EntCourse entity.
func (eau *EntAttendanceUpdate) SetAttendanceFor(e *EntCourse) *EntAttendanceUpdate {
	return eau.SetAttendanceForID(e.ID)
}

// SetOwnedByID sets the "ownedBy" edge to the EntUser entity by ID.
func (eau *EntAttendanceUpdate) SetOwnedByID(id int) *EntAttendanceUpdate {
	eau.mutation.SetOwnedByID(id)
	return eau
}

// SetNillableOwnedByID sets the "ownedBy" edge to the EntUser entity by ID if the given value is not nil.
func (eau *EntAttendanceUpdate) SetNillableOwnedByID(id *int) *EntAttendanceUpdate {
	if id != nil {
		eau = eau.SetOwnedByID(*id)
	}
	return eau
}

// SetOwnedBy sets the "ownedBy" edge to the EntUser entity.
func (eau *EntAttendanceUpdate) SetOwnedBy(e *EntUser) *EntAttendanceUpdate {
	return eau.SetOwnedByID(e.ID)
}

// Mutation returns the EntAttendanceMutation object of the builder.
func (eau *EntAttendanceUpdate) Mutation() *EntAttendanceMutation {
	return eau.mutation
}

// ClearAttendanceFor clears the "attendanceFor" edge to the EntCourse entity.
func (eau *EntAttendanceUpdate) ClearAttendanceFor() *EntAttendanceUpdate {
	eau.mutation.ClearAttendanceFor()
	return eau
}

// ClearOwnedBy clears the "ownedBy" edge to the EntUser entity.
func (eau *EntAttendanceUpdate) ClearOwnedBy() *EntAttendanceUpdate {
	eau.mutation.ClearOwnedBy()
	return eau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eau *EntAttendanceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	eau.defaults()
	if len(eau.hooks) == 0 {
		affected, err = eau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntAttendanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eau.mutation = mutation
			affected, err = eau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eau.hooks) - 1; i >= 0; i-- {
			if eau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eau *EntAttendanceUpdate) SaveX(ctx context.Context) int {
	affected, err := eau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eau *EntAttendanceUpdate) Exec(ctx context.Context) error {
	_, err := eau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eau *EntAttendanceUpdate) ExecX(ctx context.Context) {
	if err := eau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eau *EntAttendanceUpdate) defaults() {
	if _, ok := eau.mutation.UpdatedAt(); !ok {
		v := entattendance.UpdateDefaultUpdatedAt()
		eau.mutation.SetUpdatedAt(v)
	}
}

func (eau *EntAttendanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entattendance.Table,
			Columns: entattendance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entattendance.FieldID,
			},
		},
	}
	if ps := eau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldUpdatedAt,
		})
	}
	if value, ok := eau.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldDate,
		})
	}
	if value, ok := eau.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldStartTime,
		})
	}
	if eau.mutation.StartTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: entattendance.FieldStartTime,
		})
	}
	if value, ok := eau.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldEndTime,
		})
	}
	if eau.mutation.EndTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: entattendance.FieldEndTime,
		})
	}
	if value, ok := eau.mutation.Day(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldDay,
		})
	}
	if eau.mutation.DayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: entattendance.FieldDay,
		})
	}
	if value, ok := eau.mutation.Note(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entattendance.FieldNote,
		})
	}
	if eau.mutation.NoteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: entattendance.FieldNote,
		})
	}
	if value, ok := eau.mutation.Hours(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entattendance.FieldHours,
		})
	}
	if value, ok := eau.mutation.AddedHours(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entattendance.FieldHours,
		})
	}
	if eau.mutation.HoursCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: entattendance.FieldHours,
		})
	}
	if value, ok := eau.mutation.CheckedByTutor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entattendance.FieldCheckedByTutor,
		})
	}
	if value, ok := eau.mutation.CheckedByStudent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entattendance.FieldCheckedByStudent,
		})
	}
	if value, ok := eau.mutation.CheckedByParent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entattendance.FieldCheckedByParent,
		})
	}
	if eau.mutation.AttendanceForCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entattendance.AttendanceForTable,
			Columns: []string{entattendance.AttendanceForColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entcourse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eau.mutation.AttendanceForIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entattendance.AttendanceForTable,
			Columns: []string{entattendance.AttendanceForColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entcourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eau.mutation.OwnedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entattendance.OwnedByTable,
			Columns: []string{entattendance.OwnedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eau.mutation.OwnedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entattendance.OwnedByTable,
			Columns: []string{entattendance.OwnedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entattendance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EntAttendanceUpdateOne is the builder for updating a single EntAttendance entity.
type EntAttendanceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntAttendanceMutation
}

// SetUpdatedAt sets the "updatedAt" field.
func (eauo *EntAttendanceUpdateOne) SetUpdatedAt(t time.Time) *EntAttendanceUpdateOne {
	eauo.mutation.SetUpdatedAt(t)
	return eauo
}

// SetDate sets the "date" field.
func (eauo *EntAttendanceUpdateOne) SetDate(t time.Time) *EntAttendanceUpdateOne {
	eauo.mutation.SetDate(t)
	return eauo
}

// SetStartTime sets the "startTime" field.
func (eauo *EntAttendanceUpdateOne) SetStartTime(t time.Time) *EntAttendanceUpdateOne {
	eauo.mutation.SetStartTime(t)
	return eauo
}

// SetNillableStartTime sets the "startTime" field if the given value is not nil.
func (eauo *EntAttendanceUpdateOne) SetNillableStartTime(t *time.Time) *EntAttendanceUpdateOne {
	if t != nil {
		eauo.SetStartTime(*t)
	}
	return eauo
}

// ClearStartTime clears the value of the "startTime" field.
func (eauo *EntAttendanceUpdateOne) ClearStartTime() *EntAttendanceUpdateOne {
	eauo.mutation.ClearStartTime()
	return eauo
}

// SetEndTime sets the "endTime" field.
func (eauo *EntAttendanceUpdateOne) SetEndTime(t time.Time) *EntAttendanceUpdateOne {
	eauo.mutation.SetEndTime(t)
	return eauo
}

// SetNillableEndTime sets the "endTime" field if the given value is not nil.
func (eauo *EntAttendanceUpdateOne) SetNillableEndTime(t *time.Time) *EntAttendanceUpdateOne {
	if t != nil {
		eauo.SetEndTime(*t)
	}
	return eauo
}

// ClearEndTime clears the value of the "endTime" field.
func (eauo *EntAttendanceUpdateOne) ClearEndTime() *EntAttendanceUpdateOne {
	eauo.mutation.ClearEndTime()
	return eauo
}

// SetDay sets the "day" field.
func (eauo *EntAttendanceUpdateOne) SetDay(t time.Time) *EntAttendanceUpdateOne {
	eauo.mutation.SetDay(t)
	return eauo
}

// SetNillableDay sets the "day" field if the given value is not nil.
func (eauo *EntAttendanceUpdateOne) SetNillableDay(t *time.Time) *EntAttendanceUpdateOne {
	if t != nil {
		eauo.SetDay(*t)
	}
	return eauo
}

// ClearDay clears the value of the "day" field.
func (eauo *EntAttendanceUpdateOne) ClearDay() *EntAttendanceUpdateOne {
	eauo.mutation.ClearDay()
	return eauo
}

// SetNote sets the "note" field.
func (eauo *EntAttendanceUpdateOne) SetNote(s string) *EntAttendanceUpdateOne {
	eauo.mutation.SetNote(s)
	return eauo
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (eauo *EntAttendanceUpdateOne) SetNillableNote(s *string) *EntAttendanceUpdateOne {
	if s != nil {
		eauo.SetNote(*s)
	}
	return eauo
}

// ClearNote clears the value of the "note" field.
func (eauo *EntAttendanceUpdateOne) ClearNote() *EntAttendanceUpdateOne {
	eauo.mutation.ClearNote()
	return eauo
}

// SetHours sets the "hours" field.
func (eauo *EntAttendanceUpdateOne) SetHours(f float64) *EntAttendanceUpdateOne {
	eauo.mutation.ResetHours()
	eauo.mutation.SetHours(f)
	return eauo
}

// SetNillableHours sets the "hours" field if the given value is not nil.
func (eauo *EntAttendanceUpdateOne) SetNillableHours(f *float64) *EntAttendanceUpdateOne {
	if f != nil {
		eauo.SetHours(*f)
	}
	return eauo
}

// AddHours adds f to the "hours" field.
func (eauo *EntAttendanceUpdateOne) AddHours(f float64) *EntAttendanceUpdateOne {
	eauo.mutation.AddHours(f)
	return eauo
}

// ClearHours clears the value of the "hours" field.
func (eauo *EntAttendanceUpdateOne) ClearHours() *EntAttendanceUpdateOne {
	eauo.mutation.ClearHours()
	return eauo
}

// SetCheckedByTutor sets the "checkedByTutor" field.
func (eauo *EntAttendanceUpdateOne) SetCheckedByTutor(b bool) *EntAttendanceUpdateOne {
	eauo.mutation.SetCheckedByTutor(b)
	return eauo
}

// SetNillableCheckedByTutor sets the "checkedByTutor" field if the given value is not nil.
func (eauo *EntAttendanceUpdateOne) SetNillableCheckedByTutor(b *bool) *EntAttendanceUpdateOne {
	if b != nil {
		eauo.SetCheckedByTutor(*b)
	}
	return eauo
}

// SetCheckedByStudent sets the "checkedByStudent" field.
func (eauo *EntAttendanceUpdateOne) SetCheckedByStudent(b bool) *EntAttendanceUpdateOne {
	eauo.mutation.SetCheckedByStudent(b)
	return eauo
}

// SetNillableCheckedByStudent sets the "checkedByStudent" field if the given value is not nil.
func (eauo *EntAttendanceUpdateOne) SetNillableCheckedByStudent(b *bool) *EntAttendanceUpdateOne {
	if b != nil {
		eauo.SetCheckedByStudent(*b)
	}
	return eauo
}

// SetCheckedByParent sets the "checkedByParent" field.
func (eauo *EntAttendanceUpdateOne) SetCheckedByParent(b bool) *EntAttendanceUpdateOne {
	eauo.mutation.SetCheckedByParent(b)
	return eauo
}

// SetNillableCheckedByParent sets the "checkedByParent" field if the given value is not nil.
func (eauo *EntAttendanceUpdateOne) SetNillableCheckedByParent(b *bool) *EntAttendanceUpdateOne {
	if b != nil {
		eauo.SetCheckedByParent(*b)
	}
	return eauo
}

// SetAttendanceForID sets the "attendanceFor" edge to the EntCourse entity by ID.
func (eauo *EntAttendanceUpdateOne) SetAttendanceForID(id int) *EntAttendanceUpdateOne {
	eauo.mutation.SetAttendanceForID(id)
	return eauo
}

// SetNillableAttendanceForID sets the "attendanceFor" edge to the EntCourse entity by ID if the given value is not nil.
func (eauo *EntAttendanceUpdateOne) SetNillableAttendanceForID(id *int) *EntAttendanceUpdateOne {
	if id != nil {
		eauo = eauo.SetAttendanceForID(*id)
	}
	return eauo
}

// SetAttendanceFor sets the "attendanceFor" edge to the EntCourse entity.
func (eauo *EntAttendanceUpdateOne) SetAttendanceFor(e *EntCourse) *EntAttendanceUpdateOne {
	return eauo.SetAttendanceForID(e.ID)
}

// SetOwnedByID sets the "ownedBy" edge to the EntUser entity by ID.
func (eauo *EntAttendanceUpdateOne) SetOwnedByID(id int) *EntAttendanceUpdateOne {
	eauo.mutation.SetOwnedByID(id)
	return eauo
}

// SetNillableOwnedByID sets the "ownedBy" edge to the EntUser entity by ID if the given value is not nil.
func (eauo *EntAttendanceUpdateOne) SetNillableOwnedByID(id *int) *EntAttendanceUpdateOne {
	if id != nil {
		eauo = eauo.SetOwnedByID(*id)
	}
	return eauo
}

// SetOwnedBy sets the "ownedBy" edge to the EntUser entity.
func (eauo *EntAttendanceUpdateOne) SetOwnedBy(e *EntUser) *EntAttendanceUpdateOne {
	return eauo.SetOwnedByID(e.ID)
}

// Mutation returns the EntAttendanceMutation object of the builder.
func (eauo *EntAttendanceUpdateOne) Mutation() *EntAttendanceMutation {
	return eauo.mutation
}

// ClearAttendanceFor clears the "attendanceFor" edge to the EntCourse entity.
func (eauo *EntAttendanceUpdateOne) ClearAttendanceFor() *EntAttendanceUpdateOne {
	eauo.mutation.ClearAttendanceFor()
	return eauo
}

// ClearOwnedBy clears the "ownedBy" edge to the EntUser entity.
func (eauo *EntAttendanceUpdateOne) ClearOwnedBy() *EntAttendanceUpdateOne {
	eauo.mutation.ClearOwnedBy()
	return eauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eauo *EntAttendanceUpdateOne) Select(field string, fields ...string) *EntAttendanceUpdateOne {
	eauo.fields = append([]string{field}, fields...)
	return eauo
}

// Save executes the query and returns the updated EntAttendance entity.
func (eauo *EntAttendanceUpdateOne) Save(ctx context.Context) (*EntAttendance, error) {
	var (
		err  error
		node *EntAttendance
	)
	eauo.defaults()
	if len(eauo.hooks) == 0 {
		node, err = eauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntAttendanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eauo.mutation = mutation
			node, err = eauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(eauo.hooks) - 1; i >= 0; i-- {
			if eauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, eauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EntAttendance)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EntAttendanceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (eauo *EntAttendanceUpdateOne) SaveX(ctx context.Context) *EntAttendance {
	node, err := eauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eauo *EntAttendanceUpdateOne) Exec(ctx context.Context) error {
	_, err := eauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eauo *EntAttendanceUpdateOne) ExecX(ctx context.Context) {
	if err := eauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eauo *EntAttendanceUpdateOne) defaults() {
	if _, ok := eauo.mutation.UpdatedAt(); !ok {
		v := entattendance.UpdateDefaultUpdatedAt()
		eauo.mutation.SetUpdatedAt(v)
	}
}

func (eauo *EntAttendanceUpdateOne) sqlSave(ctx context.Context) (_node *EntAttendance, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entattendance.Table,
			Columns: entattendance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entattendance.FieldID,
			},
		},
	}
	id, ok := eauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EntAttendance.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entattendance.FieldID)
		for _, f := range fields {
			if !entattendance.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != entattendance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldUpdatedAt,
		})
	}
	if value, ok := eauo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldDate,
		})
	}
	if value, ok := eauo.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldStartTime,
		})
	}
	if eauo.mutation.StartTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: entattendance.FieldStartTime,
		})
	}
	if value, ok := eauo.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldEndTime,
		})
	}
	if eauo.mutation.EndTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: entattendance.FieldEndTime,
		})
	}
	if value, ok := eauo.mutation.Day(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldDay,
		})
	}
	if eauo.mutation.DayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: entattendance.FieldDay,
		})
	}
	if value, ok := eauo.mutation.Note(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entattendance.FieldNote,
		})
	}
	if eauo.mutation.NoteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: entattendance.FieldNote,
		})
	}
	if value, ok := eauo.mutation.Hours(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entattendance.FieldHours,
		})
	}
	if value, ok := eauo.mutation.AddedHours(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entattendance.FieldHours,
		})
	}
	if eauo.mutation.HoursCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: entattendance.FieldHours,
		})
	}
	if value, ok := eauo.mutation.CheckedByTutor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entattendance.FieldCheckedByTutor,
		})
	}
	if value, ok := eauo.mutation.CheckedByStudent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entattendance.FieldCheckedByStudent,
		})
	}
	if value, ok := eauo.mutation.CheckedByParent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entattendance.FieldCheckedByParent,
		})
	}
	if eauo.mutation.AttendanceForCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entattendance.AttendanceForTable,
			Columns: []string{entattendance.AttendanceForColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entcourse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eauo.mutation.AttendanceForIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entattendance.AttendanceForTable,
			Columns: []string{entattendance.AttendanceForColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entcourse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eauo.mutation.OwnedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entattendance.OwnedByTable,
			Columns: []string{entattendance.OwnedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eauo.mutation.OwnedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entattendance.OwnedByTable,
			Columns: []string{entattendance.OwnedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: entuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EntAttendance{config: eauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entattendance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
