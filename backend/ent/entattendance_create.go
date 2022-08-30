// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/entattendance"
	"backend/ent/entcourse"
	"backend/ent/entuser"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntAttendanceCreate is the builder for creating a EntAttendance entity.
type EntAttendanceCreate struct {
	config
	mutation *EntAttendanceMutation
	hooks    []Hook
}

// SetDate sets the "date" field.
func (eac *EntAttendanceCreate) SetDate(t time.Time) *EntAttendanceCreate {
	eac.mutation.SetDate(t)
	return eac
}

// SetStartTime sets the "startTime" field.
func (eac *EntAttendanceCreate) SetStartTime(t time.Time) *EntAttendanceCreate {
	eac.mutation.SetStartTime(t)
	return eac
}

// SetNillableStartTime sets the "startTime" field if the given value is not nil.
func (eac *EntAttendanceCreate) SetNillableStartTime(t *time.Time) *EntAttendanceCreate {
	if t != nil {
		eac.SetStartTime(*t)
	}
	return eac
}

// SetEndTime sets the "endTime" field.
func (eac *EntAttendanceCreate) SetEndTime(t time.Time) *EntAttendanceCreate {
	eac.mutation.SetEndTime(t)
	return eac
}

// SetNillableEndTime sets the "endTime" field if the given value is not nil.
func (eac *EntAttendanceCreate) SetNillableEndTime(t *time.Time) *EntAttendanceCreate {
	if t != nil {
		eac.SetEndTime(*t)
	}
	return eac
}

// SetDay sets the "day" field.
func (eac *EntAttendanceCreate) SetDay(t time.Time) *EntAttendanceCreate {
	eac.mutation.SetDay(t)
	return eac
}

// SetNillableDay sets the "day" field if the given value is not nil.
func (eac *EntAttendanceCreate) SetNillableDay(t *time.Time) *EntAttendanceCreate {
	if t != nil {
		eac.SetDay(*t)
	}
	return eac
}

// SetNote sets the "note" field.
func (eac *EntAttendanceCreate) SetNote(s string) *EntAttendanceCreate {
	eac.mutation.SetNote(s)
	return eac
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (eac *EntAttendanceCreate) SetNillableNote(s *string) *EntAttendanceCreate {
	if s != nil {
		eac.SetNote(*s)
	}
	return eac
}

// SetHours sets the "hours" field.
func (eac *EntAttendanceCreate) SetHours(f float64) *EntAttendanceCreate {
	eac.mutation.SetHours(f)
	return eac
}

// SetNillableHours sets the "hours" field if the given value is not nil.
func (eac *EntAttendanceCreate) SetNillableHours(f *float64) *EntAttendanceCreate {
	if f != nil {
		eac.SetHours(*f)
	}
	return eac
}

// SetCheckedByTutor sets the "checkedByTutor" field.
func (eac *EntAttendanceCreate) SetCheckedByTutor(b bool) *EntAttendanceCreate {
	eac.mutation.SetCheckedByTutor(b)
	return eac
}

// SetNillableCheckedByTutor sets the "checkedByTutor" field if the given value is not nil.
func (eac *EntAttendanceCreate) SetNillableCheckedByTutor(b *bool) *EntAttendanceCreate {
	if b != nil {
		eac.SetCheckedByTutor(*b)
	}
	return eac
}

// SetCheckedByStudent sets the "checkedByStudent" field.
func (eac *EntAttendanceCreate) SetCheckedByStudent(b bool) *EntAttendanceCreate {
	eac.mutation.SetCheckedByStudent(b)
	return eac
}

// SetNillableCheckedByStudent sets the "checkedByStudent" field if the given value is not nil.
func (eac *EntAttendanceCreate) SetNillableCheckedByStudent(b *bool) *EntAttendanceCreate {
	if b != nil {
		eac.SetCheckedByStudent(*b)
	}
	return eac
}

// SetCheckedByParent sets the "checkedByParent" field.
func (eac *EntAttendanceCreate) SetCheckedByParent(b bool) *EntAttendanceCreate {
	eac.mutation.SetCheckedByParent(b)
	return eac
}

// SetNillableCheckedByParent sets the "checkedByParent" field if the given value is not nil.
func (eac *EntAttendanceCreate) SetNillableCheckedByParent(b *bool) *EntAttendanceCreate {
	if b != nil {
		eac.SetCheckedByParent(*b)
	}
	return eac
}

// SetAttendanceForID sets the "attendanceFor" edge to the EntCourse entity by ID.
func (eac *EntAttendanceCreate) SetAttendanceForID(id int) *EntAttendanceCreate {
	eac.mutation.SetAttendanceForID(id)
	return eac
}

// SetNillableAttendanceForID sets the "attendanceFor" edge to the EntCourse entity by ID if the given value is not nil.
func (eac *EntAttendanceCreate) SetNillableAttendanceForID(id *int) *EntAttendanceCreate {
	if id != nil {
		eac = eac.SetAttendanceForID(*id)
	}
	return eac
}

// SetAttendanceFor sets the "attendanceFor" edge to the EntCourse entity.
func (eac *EntAttendanceCreate) SetAttendanceFor(e *EntCourse) *EntAttendanceCreate {
	return eac.SetAttendanceForID(e.ID)
}

// SetOwnedByID sets the "ownedBy" edge to the EntUser entity by ID.
func (eac *EntAttendanceCreate) SetOwnedByID(id int) *EntAttendanceCreate {
	eac.mutation.SetOwnedByID(id)
	return eac
}

// SetNillableOwnedByID sets the "ownedBy" edge to the EntUser entity by ID if the given value is not nil.
func (eac *EntAttendanceCreate) SetNillableOwnedByID(id *int) *EntAttendanceCreate {
	if id != nil {
		eac = eac.SetOwnedByID(*id)
	}
	return eac
}

// SetOwnedBy sets the "ownedBy" edge to the EntUser entity.
func (eac *EntAttendanceCreate) SetOwnedBy(e *EntUser) *EntAttendanceCreate {
	return eac.SetOwnedByID(e.ID)
}

// Mutation returns the EntAttendanceMutation object of the builder.
func (eac *EntAttendanceCreate) Mutation() *EntAttendanceMutation {
	return eac.mutation
}

// Save creates the EntAttendance in the database.
func (eac *EntAttendanceCreate) Save(ctx context.Context) (*EntAttendance, error) {
	var (
		err  error
		node *EntAttendance
	)
	eac.defaults()
	if len(eac.hooks) == 0 {
		if err = eac.check(); err != nil {
			return nil, err
		}
		node, err = eac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntAttendanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eac.check(); err != nil {
				return nil, err
			}
			eac.mutation = mutation
			if node, err = eac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(eac.hooks) - 1; i >= 0; i-- {
			if eac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, eac.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (eac *EntAttendanceCreate) SaveX(ctx context.Context) *EntAttendance {
	v, err := eac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eac *EntAttendanceCreate) Exec(ctx context.Context) error {
	_, err := eac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eac *EntAttendanceCreate) ExecX(ctx context.Context) {
	if err := eac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eac *EntAttendanceCreate) defaults() {
	if _, ok := eac.mutation.CheckedByTutor(); !ok {
		v := entattendance.DefaultCheckedByTutor
		eac.mutation.SetCheckedByTutor(v)
	}
	if _, ok := eac.mutation.CheckedByStudent(); !ok {
		v := entattendance.DefaultCheckedByStudent
		eac.mutation.SetCheckedByStudent(v)
	}
	if _, ok := eac.mutation.CheckedByParent(); !ok {
		v := entattendance.DefaultCheckedByParent
		eac.mutation.SetCheckedByParent(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eac *EntAttendanceCreate) check() error {
	if _, ok := eac.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "EntAttendance.date"`)}
	}
	if _, ok := eac.mutation.CheckedByTutor(); !ok {
		return &ValidationError{Name: "checkedByTutor", err: errors.New(`ent: missing required field "EntAttendance.checkedByTutor"`)}
	}
	if _, ok := eac.mutation.CheckedByStudent(); !ok {
		return &ValidationError{Name: "checkedByStudent", err: errors.New(`ent: missing required field "EntAttendance.checkedByStudent"`)}
	}
	if _, ok := eac.mutation.CheckedByParent(); !ok {
		return &ValidationError{Name: "checkedByParent", err: errors.New(`ent: missing required field "EntAttendance.checkedByParent"`)}
	}
	return nil
}

func (eac *EntAttendanceCreate) sqlSave(ctx context.Context) (*EntAttendance, error) {
	_node, _spec := eac.createSpec()
	if err := sqlgraph.CreateNode(ctx, eac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (eac *EntAttendanceCreate) createSpec() (*EntAttendance, *sqlgraph.CreateSpec) {
	var (
		_node = &EntAttendance{config: eac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: entattendance.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entattendance.FieldID,
			},
		}
	)
	if value, ok := eac.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := eac.mutation.StartTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldStartTime,
		})
		_node.StartTime = value
	}
	if value, ok := eac.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldEndTime,
		})
		_node.EndTime = value
	}
	if value, ok := eac.mutation.Day(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entattendance.FieldDay,
		})
		_node.Day = value
	}
	if value, ok := eac.mutation.Note(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entattendance.FieldNote,
		})
		_node.Note = value
	}
	if value, ok := eac.mutation.Hours(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entattendance.FieldHours,
		})
		_node.Hours = value
	}
	if value, ok := eac.mutation.CheckedByTutor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entattendance.FieldCheckedByTutor,
		})
		_node.CheckedByTutor = value
	}
	if value, ok := eac.mutation.CheckedByStudent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entattendance.FieldCheckedByStudent,
		})
		_node.CheckedByStudent = value
	}
	if value, ok := eac.mutation.CheckedByParent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entattendance.FieldCheckedByParent,
		})
		_node.CheckedByParent = value
	}
	if nodes := eac.mutation.AttendanceForIDs(); len(nodes) > 0 {
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
		_node.ent_course_attendance = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := eac.mutation.OwnedByIDs(); len(nodes) > 0 {
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
		_node.ent_user_attendance = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EntAttendanceCreateBulk is the builder for creating many EntAttendance entities in bulk.
type EntAttendanceCreateBulk struct {
	config
	builders []*EntAttendanceCreate
}

// Save creates the EntAttendance entities in the database.
func (eacb *EntAttendanceCreateBulk) Save(ctx context.Context) ([]*EntAttendance, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eacb.builders))
	nodes := make([]*EntAttendance, len(eacb.builders))
	mutators := make([]Mutator, len(eacb.builders))
	for i := range eacb.builders {
		func(i int, root context.Context) {
			builder := eacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntAttendanceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, eacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, eacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eacb *EntAttendanceCreateBulk) SaveX(ctx context.Context) []*EntAttendance {
	v, err := eacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eacb *EntAttendanceCreateBulk) Exec(ctx context.Context) error {
	_, err := eacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eacb *EntAttendanceCreateBulk) ExecX(ctx context.Context) {
	if err := eacb.Exec(ctx); err != nil {
		panic(err)
	}
}
