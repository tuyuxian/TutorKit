// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/entcourse"
	"backend/ent/entuser"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntCourseCreate is the builder for creating a EntCourse entity.
type EntCourseCreate struct {
	config
	mutation *EntCourseMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ecc *EntCourseCreate) SetName(s string) *EntCourseCreate {
	ecc.mutation.SetName(s)
	return ecc
}

// SetCourseUrl sets the "courseUrl" field.
func (ecc *EntCourseCreate) SetCourseUrl(s string) *EntCourseCreate {
	ecc.mutation.SetCourseUrl(s)
	return ecc
}

// SetNillableCourseUrl sets the "courseUrl" field if the given value is not nil.
func (ecc *EntCourseCreate) SetNillableCourseUrl(s *string) *EntCourseCreate {
	if s != nil {
		ecc.SetCourseUrl(*s)
	}
	return ecc
}

// SetPaymentMethod sets the "paymentMethod" field.
func (ecc *EntCourseCreate) SetPaymentMethod(em entcourse.PaymentMethod) *EntCourseCreate {
	ecc.mutation.SetPaymentMethod(em)
	return ecc
}

// SetNillablePaymentMethod sets the "paymentMethod" field if the given value is not nil.
func (ecc *EntCourseCreate) SetNillablePaymentMethod(em *entcourse.PaymentMethod) *EntCourseCreate {
	if em != nil {
		ecc.SetPaymentMethod(*em)
	}
	return ecc
}

// SetPaymentAmount sets the "paymentAmount" field.
func (ecc *EntCourseCreate) SetPaymentAmount(f float64) *EntCourseCreate {
	ecc.mutation.SetPaymentAmount(f)
	return ecc
}

// SetNillablePaymentAmount sets the "paymentAmount" field if the given value is not nil.
func (ecc *EntCourseCreate) SetNillablePaymentAmount(f *float64) *EntCourseCreate {
	if f != nil {
		ecc.SetPaymentAmount(*f)
	}
	return ecc
}

// SetStartDate sets the "startDate" field.
func (ecc *EntCourseCreate) SetStartDate(t time.Time) *EntCourseCreate {
	ecc.mutation.SetStartDate(t)
	return ecc
}

// SetNillableStartDate sets the "startDate" field if the given value is not nil.
func (ecc *EntCourseCreate) SetNillableStartDate(t *time.Time) *EntCourseCreate {
	if t != nil {
		ecc.SetStartDate(*t)
	}
	return ecc
}

// SetEndDate sets the "endDate" field.
func (ecc *EntCourseCreate) SetEndDate(t time.Time) *EntCourseCreate {
	ecc.mutation.SetEndDate(t)
	return ecc
}

// SetNillableEndDate sets the "endDate" field if the given value is not nil.
func (ecc *EntCourseCreate) SetNillableEndDate(t *time.Time) *EntCourseCreate {
	if t != nil {
		ecc.SetEndDate(*t)
	}
	return ecc
}

// SetMonday sets the "monday" field.
func (ecc *EntCourseCreate) SetMonday(b bool) *EntCourseCreate {
	ecc.mutation.SetMonday(b)
	return ecc
}

// SetNillableMonday sets the "monday" field if the given value is not nil.
func (ecc *EntCourseCreate) SetNillableMonday(b *bool) *EntCourseCreate {
	if b != nil {
		ecc.SetMonday(*b)
	}
	return ecc
}

// SetTuesday sets the "tuesday" field.
func (ecc *EntCourseCreate) SetTuesday(b bool) *EntCourseCreate {
	ecc.mutation.SetTuesday(b)
	return ecc
}

// SetNillableTuesday sets the "tuesday" field if the given value is not nil.
func (ecc *EntCourseCreate) SetNillableTuesday(b *bool) *EntCourseCreate {
	if b != nil {
		ecc.SetTuesday(*b)
	}
	return ecc
}

// SetWednesday sets the "wednesday" field.
func (ecc *EntCourseCreate) SetWednesday(b bool) *EntCourseCreate {
	ecc.mutation.SetWednesday(b)
	return ecc
}

// SetNillableWednesday sets the "wednesday" field if the given value is not nil.
func (ecc *EntCourseCreate) SetNillableWednesday(b *bool) *EntCourseCreate {
	if b != nil {
		ecc.SetWednesday(*b)
	}
	return ecc
}

// SetThursday sets the "thursday" field.
func (ecc *EntCourseCreate) SetThursday(b bool) *EntCourseCreate {
	ecc.mutation.SetThursday(b)
	return ecc
}

// SetNillableThursday sets the "thursday" field if the given value is not nil.
func (ecc *EntCourseCreate) SetNillableThursday(b *bool) *EntCourseCreate {
	if b != nil {
		ecc.SetThursday(*b)
	}
	return ecc
}

// SetFriday sets the "friday" field.
func (ecc *EntCourseCreate) SetFriday(b bool) *EntCourseCreate {
	ecc.mutation.SetFriday(b)
	return ecc
}

// SetNillableFriday sets the "friday" field if the given value is not nil.
func (ecc *EntCourseCreate) SetNillableFriday(b *bool) *EntCourseCreate {
	if b != nil {
		ecc.SetFriday(*b)
	}
	return ecc
}

// SetSaturday sets the "saturday" field.
func (ecc *EntCourseCreate) SetSaturday(b bool) *EntCourseCreate {
	ecc.mutation.SetSaturday(b)
	return ecc
}

// SetNillableSaturday sets the "saturday" field if the given value is not nil.
func (ecc *EntCourseCreate) SetNillableSaturday(b *bool) *EntCourseCreate {
	if b != nil {
		ecc.SetSaturday(*b)
	}
	return ecc
}

// SetSunday sets the "sunday" field.
func (ecc *EntCourseCreate) SetSunday(b bool) *EntCourseCreate {
	ecc.mutation.SetSunday(b)
	return ecc
}

// SetNillableSunday sets the "sunday" field if the given value is not nil.
func (ecc *EntCourseCreate) SetNillableSunday(b *bool) *EntCourseCreate {
	if b != nil {
		ecc.SetSunday(*b)
	}
	return ecc
}

// AddCourseOwnerIDs adds the "courseOwner" edge to the EntUser entity by IDs.
func (ecc *EntCourseCreate) AddCourseOwnerIDs(ids ...int) *EntCourseCreate {
	ecc.mutation.AddCourseOwnerIDs(ids...)
	return ecc
}

// AddCourseOwner adds the "courseOwner" edges to the EntUser entity.
func (ecc *EntCourseCreate) AddCourseOwner(e ...*EntUser) *EntCourseCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecc.AddCourseOwnerIDs(ids...)
}

// Mutation returns the EntCourseMutation object of the builder.
func (ecc *EntCourseCreate) Mutation() *EntCourseMutation {
	return ecc.mutation
}

// Save creates the EntCourse in the database.
func (ecc *EntCourseCreate) Save(ctx context.Context) (*EntCourse, error) {
	var (
		err  error
		node *EntCourse
	)
	if len(ecc.hooks) == 0 {
		if err = ecc.check(); err != nil {
			return nil, err
		}
		node, err = ecc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntCourseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ecc.check(); err != nil {
				return nil, err
			}
			ecc.mutation = mutation
			if node, err = ecc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ecc.hooks) - 1; i >= 0; i-- {
			if ecc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ecc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ecc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EntCourse)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EntCourseMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ecc *EntCourseCreate) SaveX(ctx context.Context) *EntCourse {
	v, err := ecc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecc *EntCourseCreate) Exec(ctx context.Context) error {
	_, err := ecc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecc *EntCourseCreate) ExecX(ctx context.Context) {
	if err := ecc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecc *EntCourseCreate) check() error {
	if _, ok := ecc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "EntCourse.name"`)}
	}
	if v, ok := ecc.mutation.Name(); ok {
		if err := entcourse.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "EntCourse.name": %w`, err)}
		}
	}
	if v, ok := ecc.mutation.PaymentMethod(); ok {
		if err := entcourse.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "paymentMethod", err: fmt.Errorf(`ent: validator failed for field "EntCourse.paymentMethod": %w`, err)}
		}
	}
	return nil
}

func (ecc *EntCourseCreate) sqlSave(ctx context.Context) (*EntCourse, error) {
	_node, _spec := ecc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ecc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ecc *EntCourseCreate) createSpec() (*EntCourse, *sqlgraph.CreateSpec) {
	var (
		_node = &EntCourse{config: ecc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: entcourse.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entcourse.FieldID,
			},
		}
	)
	if value, ok := ecc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entcourse.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ecc.mutation.CourseUrl(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entcourse.FieldCourseUrl,
		})
		_node.CourseUrl = value
	}
	if value, ok := ecc.mutation.PaymentMethod(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: entcourse.FieldPaymentMethod,
		})
		_node.PaymentMethod = value
	}
	if value, ok := ecc.mutation.PaymentAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entcourse.FieldPaymentAmount,
		})
		_node.PaymentAmount = value
	}
	if value, ok := ecc.mutation.StartDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entcourse.FieldStartDate,
		})
		_node.StartDate = value
	}
	if value, ok := ecc.mutation.EndDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entcourse.FieldEndDate,
		})
		_node.EndDate = value
	}
	if value, ok := ecc.mutation.Monday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldMonday,
		})
		_node.Monday = value
	}
	if value, ok := ecc.mutation.Tuesday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldTuesday,
		})
		_node.Tuesday = value
	}
	if value, ok := ecc.mutation.Wednesday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldWednesday,
		})
		_node.Wednesday = value
	}
	if value, ok := ecc.mutation.Thursday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldThursday,
		})
		_node.Thursday = value
	}
	if value, ok := ecc.mutation.Friday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldFriday,
		})
		_node.Friday = value
	}
	if value, ok := ecc.mutation.Saturday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldSaturday,
		})
		_node.Saturday = value
	}
	if value, ok := ecc.mutation.Sunday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldSunday,
		})
		_node.Sunday = value
	}
	if nodes := ecc.mutation.CourseOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   entcourse.CourseOwnerTable,
			Columns: entcourse.CourseOwnerPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EntCourseCreateBulk is the builder for creating many EntCourse entities in bulk.
type EntCourseCreateBulk struct {
	config
	builders []*EntCourseCreate
}

// Save creates the EntCourse entities in the database.
func (eccb *EntCourseCreateBulk) Save(ctx context.Context) ([]*EntCourse, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eccb.builders))
	nodes := make([]*EntCourse, len(eccb.builders))
	mutators := make([]Mutator, len(eccb.builders))
	for i := range eccb.builders {
		func(i int, root context.Context) {
			builder := eccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntCourseMutation)
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
					_, err = mutators[i+1].Mutate(root, eccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eccb *EntCourseCreateBulk) SaveX(ctx context.Context) []*EntCourse {
	v, err := eccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eccb *EntCourseCreateBulk) Exec(ctx context.Context) error {
	_, err := eccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eccb *EntCourseCreateBulk) ExecX(ctx context.Context) {
	if err := eccb.Exec(ctx); err != nil {
		panic(err)
	}
}
