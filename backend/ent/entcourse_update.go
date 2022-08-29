// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/entcourse"
	"backend/ent/entuser"
	"backend/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntCourseUpdate is the builder for updating EntCourse entities.
type EntCourseUpdate struct {
	config
	hooks    []Hook
	mutation *EntCourseMutation
}

// Where appends a list predicates to the EntCourseUpdate builder.
func (ecu *EntCourseUpdate) Where(ps ...predicate.EntCourse) *EntCourseUpdate {
	ecu.mutation.Where(ps...)
	return ecu
}

// SetName sets the "name" field.
func (ecu *EntCourseUpdate) SetName(s string) *EntCourseUpdate {
	ecu.mutation.SetName(s)
	return ecu
}

// SetCourseUrl sets the "courseUrl" field.
func (ecu *EntCourseUpdate) SetCourseUrl(s string) *EntCourseUpdate {
	ecu.mutation.SetCourseUrl(s)
	return ecu
}

// SetNillableCourseUrl sets the "courseUrl" field if the given value is not nil.
func (ecu *EntCourseUpdate) SetNillableCourseUrl(s *string) *EntCourseUpdate {
	if s != nil {
		ecu.SetCourseUrl(*s)
	}
	return ecu
}

// ClearCourseUrl clears the value of the "courseUrl" field.
func (ecu *EntCourseUpdate) ClearCourseUrl() *EntCourseUpdate {
	ecu.mutation.ClearCourseUrl()
	return ecu
}

// SetPaymentMethod sets the "paymentMethod" field.
func (ecu *EntCourseUpdate) SetPaymentMethod(em entcourse.PaymentMethod) *EntCourseUpdate {
	ecu.mutation.SetPaymentMethod(em)
	return ecu
}

// SetNillablePaymentMethod sets the "paymentMethod" field if the given value is not nil.
func (ecu *EntCourseUpdate) SetNillablePaymentMethod(em *entcourse.PaymentMethod) *EntCourseUpdate {
	if em != nil {
		ecu.SetPaymentMethod(*em)
	}
	return ecu
}

// ClearPaymentMethod clears the value of the "paymentMethod" field.
func (ecu *EntCourseUpdate) ClearPaymentMethod() *EntCourseUpdate {
	ecu.mutation.ClearPaymentMethod()
	return ecu
}

// SetPaymentAmount sets the "paymentAmount" field.
func (ecu *EntCourseUpdate) SetPaymentAmount(f float64) *EntCourseUpdate {
	ecu.mutation.ResetPaymentAmount()
	ecu.mutation.SetPaymentAmount(f)
	return ecu
}

// SetNillablePaymentAmount sets the "paymentAmount" field if the given value is not nil.
func (ecu *EntCourseUpdate) SetNillablePaymentAmount(f *float64) *EntCourseUpdate {
	if f != nil {
		ecu.SetPaymentAmount(*f)
	}
	return ecu
}

// AddPaymentAmount adds f to the "paymentAmount" field.
func (ecu *EntCourseUpdate) AddPaymentAmount(f float64) *EntCourseUpdate {
	ecu.mutation.AddPaymentAmount(f)
	return ecu
}

// ClearPaymentAmount clears the value of the "paymentAmount" field.
func (ecu *EntCourseUpdate) ClearPaymentAmount() *EntCourseUpdate {
	ecu.mutation.ClearPaymentAmount()
	return ecu
}

// SetStartDate sets the "startDate" field.
func (ecu *EntCourseUpdate) SetStartDate(t time.Time) *EntCourseUpdate {
	ecu.mutation.SetStartDate(t)
	return ecu
}

// SetNillableStartDate sets the "startDate" field if the given value is not nil.
func (ecu *EntCourseUpdate) SetNillableStartDate(t *time.Time) *EntCourseUpdate {
	if t != nil {
		ecu.SetStartDate(*t)
	}
	return ecu
}

// ClearStartDate clears the value of the "startDate" field.
func (ecu *EntCourseUpdate) ClearStartDate() *EntCourseUpdate {
	ecu.mutation.ClearStartDate()
	return ecu
}

// SetEndDate sets the "endDate" field.
func (ecu *EntCourseUpdate) SetEndDate(t time.Time) *EntCourseUpdate {
	ecu.mutation.SetEndDate(t)
	return ecu
}

// SetNillableEndDate sets the "endDate" field if the given value is not nil.
func (ecu *EntCourseUpdate) SetNillableEndDate(t *time.Time) *EntCourseUpdate {
	if t != nil {
		ecu.SetEndDate(*t)
	}
	return ecu
}

// ClearEndDate clears the value of the "endDate" field.
func (ecu *EntCourseUpdate) ClearEndDate() *EntCourseUpdate {
	ecu.mutation.ClearEndDate()
	return ecu
}

// SetMonday sets the "monday" field.
func (ecu *EntCourseUpdate) SetMonday(b bool) *EntCourseUpdate {
	ecu.mutation.SetMonday(b)
	return ecu
}

// SetNillableMonday sets the "monday" field if the given value is not nil.
func (ecu *EntCourseUpdate) SetNillableMonday(b *bool) *EntCourseUpdate {
	if b != nil {
		ecu.SetMonday(*b)
	}
	return ecu
}

// ClearMonday clears the value of the "monday" field.
func (ecu *EntCourseUpdate) ClearMonday() *EntCourseUpdate {
	ecu.mutation.ClearMonday()
	return ecu
}

// SetTuesday sets the "tuesday" field.
func (ecu *EntCourseUpdate) SetTuesday(b bool) *EntCourseUpdate {
	ecu.mutation.SetTuesday(b)
	return ecu
}

// SetNillableTuesday sets the "tuesday" field if the given value is not nil.
func (ecu *EntCourseUpdate) SetNillableTuesday(b *bool) *EntCourseUpdate {
	if b != nil {
		ecu.SetTuesday(*b)
	}
	return ecu
}

// ClearTuesday clears the value of the "tuesday" field.
func (ecu *EntCourseUpdate) ClearTuesday() *EntCourseUpdate {
	ecu.mutation.ClearTuesday()
	return ecu
}

// SetWednesday sets the "wednesday" field.
func (ecu *EntCourseUpdate) SetWednesday(b bool) *EntCourseUpdate {
	ecu.mutation.SetWednesday(b)
	return ecu
}

// SetNillableWednesday sets the "wednesday" field if the given value is not nil.
func (ecu *EntCourseUpdate) SetNillableWednesday(b *bool) *EntCourseUpdate {
	if b != nil {
		ecu.SetWednesday(*b)
	}
	return ecu
}

// ClearWednesday clears the value of the "wednesday" field.
func (ecu *EntCourseUpdate) ClearWednesday() *EntCourseUpdate {
	ecu.mutation.ClearWednesday()
	return ecu
}

// SetThursday sets the "thursday" field.
func (ecu *EntCourseUpdate) SetThursday(b bool) *EntCourseUpdate {
	ecu.mutation.SetThursday(b)
	return ecu
}

// SetNillableThursday sets the "thursday" field if the given value is not nil.
func (ecu *EntCourseUpdate) SetNillableThursday(b *bool) *EntCourseUpdate {
	if b != nil {
		ecu.SetThursday(*b)
	}
	return ecu
}

// ClearThursday clears the value of the "thursday" field.
func (ecu *EntCourseUpdate) ClearThursday() *EntCourseUpdate {
	ecu.mutation.ClearThursday()
	return ecu
}

// SetFriday sets the "friday" field.
func (ecu *EntCourseUpdate) SetFriday(b bool) *EntCourseUpdate {
	ecu.mutation.SetFriday(b)
	return ecu
}

// SetNillableFriday sets the "friday" field if the given value is not nil.
func (ecu *EntCourseUpdate) SetNillableFriday(b *bool) *EntCourseUpdate {
	if b != nil {
		ecu.SetFriday(*b)
	}
	return ecu
}

// ClearFriday clears the value of the "friday" field.
func (ecu *EntCourseUpdate) ClearFriday() *EntCourseUpdate {
	ecu.mutation.ClearFriday()
	return ecu
}

// SetSaturday sets the "saturday" field.
func (ecu *EntCourseUpdate) SetSaturday(b bool) *EntCourseUpdate {
	ecu.mutation.SetSaturday(b)
	return ecu
}

// SetNillableSaturday sets the "saturday" field if the given value is not nil.
func (ecu *EntCourseUpdate) SetNillableSaturday(b *bool) *EntCourseUpdate {
	if b != nil {
		ecu.SetSaturday(*b)
	}
	return ecu
}

// ClearSaturday clears the value of the "saturday" field.
func (ecu *EntCourseUpdate) ClearSaturday() *EntCourseUpdate {
	ecu.mutation.ClearSaturday()
	return ecu
}

// SetSunday sets the "sunday" field.
func (ecu *EntCourseUpdate) SetSunday(b bool) *EntCourseUpdate {
	ecu.mutation.SetSunday(b)
	return ecu
}

// SetNillableSunday sets the "sunday" field if the given value is not nil.
func (ecu *EntCourseUpdate) SetNillableSunday(b *bool) *EntCourseUpdate {
	if b != nil {
		ecu.SetSunday(*b)
	}
	return ecu
}

// ClearSunday clears the value of the "sunday" field.
func (ecu *EntCourseUpdate) ClearSunday() *EntCourseUpdate {
	ecu.mutation.ClearSunday()
	return ecu
}

// AddCourseOwnerIDs adds the "courseOwner" edge to the EntUser entity by IDs.
func (ecu *EntCourseUpdate) AddCourseOwnerIDs(ids ...int) *EntCourseUpdate {
	ecu.mutation.AddCourseOwnerIDs(ids...)
	return ecu
}

// AddCourseOwner adds the "courseOwner" edges to the EntUser entity.
func (ecu *EntCourseUpdate) AddCourseOwner(e ...*EntUser) *EntCourseUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecu.AddCourseOwnerIDs(ids...)
}

// Mutation returns the EntCourseMutation object of the builder.
func (ecu *EntCourseUpdate) Mutation() *EntCourseMutation {
	return ecu.mutation
}

// ClearCourseOwner clears all "courseOwner" edges to the EntUser entity.
func (ecu *EntCourseUpdate) ClearCourseOwner() *EntCourseUpdate {
	ecu.mutation.ClearCourseOwner()
	return ecu
}

// RemoveCourseOwnerIDs removes the "courseOwner" edge to EntUser entities by IDs.
func (ecu *EntCourseUpdate) RemoveCourseOwnerIDs(ids ...int) *EntCourseUpdate {
	ecu.mutation.RemoveCourseOwnerIDs(ids...)
	return ecu
}

// RemoveCourseOwner removes "courseOwner" edges to EntUser entities.
func (ecu *EntCourseUpdate) RemoveCourseOwner(e ...*EntUser) *EntCourseUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecu.RemoveCourseOwnerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ecu *EntCourseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ecu.hooks) == 0 {
		if err = ecu.check(); err != nil {
			return 0, err
		}
		affected, err = ecu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntCourseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ecu.check(); err != nil {
				return 0, err
			}
			ecu.mutation = mutation
			affected, err = ecu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ecu.hooks) - 1; i >= 0; i-- {
			if ecu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ecu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ecu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ecu *EntCourseUpdate) SaveX(ctx context.Context) int {
	affected, err := ecu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ecu *EntCourseUpdate) Exec(ctx context.Context) error {
	_, err := ecu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecu *EntCourseUpdate) ExecX(ctx context.Context) {
	if err := ecu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecu *EntCourseUpdate) check() error {
	if v, ok := ecu.mutation.Name(); ok {
		if err := entcourse.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "EntCourse.name": %w`, err)}
		}
	}
	if v, ok := ecu.mutation.PaymentMethod(); ok {
		if err := entcourse.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "paymentMethod", err: fmt.Errorf(`ent: validator failed for field "EntCourse.paymentMethod": %w`, err)}
		}
	}
	return nil
}

func (ecu *EntCourseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entcourse.Table,
			Columns: entcourse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entcourse.FieldID,
			},
		},
	}
	if ps := ecu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entcourse.FieldName,
		})
	}
	if value, ok := ecu.mutation.CourseUrl(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entcourse.FieldCourseUrl,
		})
	}
	if ecu.mutation.CourseUrlCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: entcourse.FieldCourseUrl,
		})
	}
	if value, ok := ecu.mutation.PaymentMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: entcourse.FieldPaymentMethod,
		})
	}
	if ecu.mutation.PaymentMethodCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: entcourse.FieldPaymentMethod,
		})
	}
	if value, ok := ecu.mutation.PaymentAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entcourse.FieldPaymentAmount,
		})
	}
	if value, ok := ecu.mutation.AddedPaymentAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entcourse.FieldPaymentAmount,
		})
	}
	if ecu.mutation.PaymentAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: entcourse.FieldPaymentAmount,
		})
	}
	if value, ok := ecu.mutation.StartDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entcourse.FieldStartDate,
		})
	}
	if ecu.mutation.StartDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: entcourse.FieldStartDate,
		})
	}
	if value, ok := ecu.mutation.EndDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entcourse.FieldEndDate,
		})
	}
	if ecu.mutation.EndDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: entcourse.FieldEndDate,
		})
	}
	if value, ok := ecu.mutation.Monday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldMonday,
		})
	}
	if ecu.mutation.MondayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldMonday,
		})
	}
	if value, ok := ecu.mutation.Tuesday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldTuesday,
		})
	}
	if ecu.mutation.TuesdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldTuesday,
		})
	}
	if value, ok := ecu.mutation.Wednesday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldWednesday,
		})
	}
	if ecu.mutation.WednesdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldWednesday,
		})
	}
	if value, ok := ecu.mutation.Thursday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldThursday,
		})
	}
	if ecu.mutation.ThursdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldThursday,
		})
	}
	if value, ok := ecu.mutation.Friday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldFriday,
		})
	}
	if ecu.mutation.FridayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldFriday,
		})
	}
	if value, ok := ecu.mutation.Saturday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldSaturday,
		})
	}
	if ecu.mutation.SaturdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldSaturday,
		})
	}
	if value, ok := ecu.mutation.Sunday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldSunday,
		})
	}
	if ecu.mutation.SundayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldSunday,
		})
	}
	if ecu.mutation.CourseOwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.RemovedCourseOwnerIDs(); len(nodes) > 0 && !ecu.mutation.CourseOwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.CourseOwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ecu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entcourse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EntCourseUpdateOne is the builder for updating a single EntCourse entity.
type EntCourseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntCourseMutation
}

// SetName sets the "name" field.
func (ecuo *EntCourseUpdateOne) SetName(s string) *EntCourseUpdateOne {
	ecuo.mutation.SetName(s)
	return ecuo
}

// SetCourseUrl sets the "courseUrl" field.
func (ecuo *EntCourseUpdateOne) SetCourseUrl(s string) *EntCourseUpdateOne {
	ecuo.mutation.SetCourseUrl(s)
	return ecuo
}

// SetNillableCourseUrl sets the "courseUrl" field if the given value is not nil.
func (ecuo *EntCourseUpdateOne) SetNillableCourseUrl(s *string) *EntCourseUpdateOne {
	if s != nil {
		ecuo.SetCourseUrl(*s)
	}
	return ecuo
}

// ClearCourseUrl clears the value of the "courseUrl" field.
func (ecuo *EntCourseUpdateOne) ClearCourseUrl() *EntCourseUpdateOne {
	ecuo.mutation.ClearCourseUrl()
	return ecuo
}

// SetPaymentMethod sets the "paymentMethod" field.
func (ecuo *EntCourseUpdateOne) SetPaymentMethod(em entcourse.PaymentMethod) *EntCourseUpdateOne {
	ecuo.mutation.SetPaymentMethod(em)
	return ecuo
}

// SetNillablePaymentMethod sets the "paymentMethod" field if the given value is not nil.
func (ecuo *EntCourseUpdateOne) SetNillablePaymentMethod(em *entcourse.PaymentMethod) *EntCourseUpdateOne {
	if em != nil {
		ecuo.SetPaymentMethod(*em)
	}
	return ecuo
}

// ClearPaymentMethod clears the value of the "paymentMethod" field.
func (ecuo *EntCourseUpdateOne) ClearPaymentMethod() *EntCourseUpdateOne {
	ecuo.mutation.ClearPaymentMethod()
	return ecuo
}

// SetPaymentAmount sets the "paymentAmount" field.
func (ecuo *EntCourseUpdateOne) SetPaymentAmount(f float64) *EntCourseUpdateOne {
	ecuo.mutation.ResetPaymentAmount()
	ecuo.mutation.SetPaymentAmount(f)
	return ecuo
}

// SetNillablePaymentAmount sets the "paymentAmount" field if the given value is not nil.
func (ecuo *EntCourseUpdateOne) SetNillablePaymentAmount(f *float64) *EntCourseUpdateOne {
	if f != nil {
		ecuo.SetPaymentAmount(*f)
	}
	return ecuo
}

// AddPaymentAmount adds f to the "paymentAmount" field.
func (ecuo *EntCourseUpdateOne) AddPaymentAmount(f float64) *EntCourseUpdateOne {
	ecuo.mutation.AddPaymentAmount(f)
	return ecuo
}

// ClearPaymentAmount clears the value of the "paymentAmount" field.
func (ecuo *EntCourseUpdateOne) ClearPaymentAmount() *EntCourseUpdateOne {
	ecuo.mutation.ClearPaymentAmount()
	return ecuo
}

// SetStartDate sets the "startDate" field.
func (ecuo *EntCourseUpdateOne) SetStartDate(t time.Time) *EntCourseUpdateOne {
	ecuo.mutation.SetStartDate(t)
	return ecuo
}

// SetNillableStartDate sets the "startDate" field if the given value is not nil.
func (ecuo *EntCourseUpdateOne) SetNillableStartDate(t *time.Time) *EntCourseUpdateOne {
	if t != nil {
		ecuo.SetStartDate(*t)
	}
	return ecuo
}

// ClearStartDate clears the value of the "startDate" field.
func (ecuo *EntCourseUpdateOne) ClearStartDate() *EntCourseUpdateOne {
	ecuo.mutation.ClearStartDate()
	return ecuo
}

// SetEndDate sets the "endDate" field.
func (ecuo *EntCourseUpdateOne) SetEndDate(t time.Time) *EntCourseUpdateOne {
	ecuo.mutation.SetEndDate(t)
	return ecuo
}

// SetNillableEndDate sets the "endDate" field if the given value is not nil.
func (ecuo *EntCourseUpdateOne) SetNillableEndDate(t *time.Time) *EntCourseUpdateOne {
	if t != nil {
		ecuo.SetEndDate(*t)
	}
	return ecuo
}

// ClearEndDate clears the value of the "endDate" field.
func (ecuo *EntCourseUpdateOne) ClearEndDate() *EntCourseUpdateOne {
	ecuo.mutation.ClearEndDate()
	return ecuo
}

// SetMonday sets the "monday" field.
func (ecuo *EntCourseUpdateOne) SetMonday(b bool) *EntCourseUpdateOne {
	ecuo.mutation.SetMonday(b)
	return ecuo
}

// SetNillableMonday sets the "monday" field if the given value is not nil.
func (ecuo *EntCourseUpdateOne) SetNillableMonday(b *bool) *EntCourseUpdateOne {
	if b != nil {
		ecuo.SetMonday(*b)
	}
	return ecuo
}

// ClearMonday clears the value of the "monday" field.
func (ecuo *EntCourseUpdateOne) ClearMonday() *EntCourseUpdateOne {
	ecuo.mutation.ClearMonday()
	return ecuo
}

// SetTuesday sets the "tuesday" field.
func (ecuo *EntCourseUpdateOne) SetTuesday(b bool) *EntCourseUpdateOne {
	ecuo.mutation.SetTuesday(b)
	return ecuo
}

// SetNillableTuesday sets the "tuesday" field if the given value is not nil.
func (ecuo *EntCourseUpdateOne) SetNillableTuesday(b *bool) *EntCourseUpdateOne {
	if b != nil {
		ecuo.SetTuesday(*b)
	}
	return ecuo
}

// ClearTuesday clears the value of the "tuesday" field.
func (ecuo *EntCourseUpdateOne) ClearTuesday() *EntCourseUpdateOne {
	ecuo.mutation.ClearTuesday()
	return ecuo
}

// SetWednesday sets the "wednesday" field.
func (ecuo *EntCourseUpdateOne) SetWednesday(b bool) *EntCourseUpdateOne {
	ecuo.mutation.SetWednesday(b)
	return ecuo
}

// SetNillableWednesday sets the "wednesday" field if the given value is not nil.
func (ecuo *EntCourseUpdateOne) SetNillableWednesday(b *bool) *EntCourseUpdateOne {
	if b != nil {
		ecuo.SetWednesday(*b)
	}
	return ecuo
}

// ClearWednesday clears the value of the "wednesday" field.
func (ecuo *EntCourseUpdateOne) ClearWednesday() *EntCourseUpdateOne {
	ecuo.mutation.ClearWednesday()
	return ecuo
}

// SetThursday sets the "thursday" field.
func (ecuo *EntCourseUpdateOne) SetThursday(b bool) *EntCourseUpdateOne {
	ecuo.mutation.SetThursday(b)
	return ecuo
}

// SetNillableThursday sets the "thursday" field if the given value is not nil.
func (ecuo *EntCourseUpdateOne) SetNillableThursday(b *bool) *EntCourseUpdateOne {
	if b != nil {
		ecuo.SetThursday(*b)
	}
	return ecuo
}

// ClearThursday clears the value of the "thursday" field.
func (ecuo *EntCourseUpdateOne) ClearThursday() *EntCourseUpdateOne {
	ecuo.mutation.ClearThursday()
	return ecuo
}

// SetFriday sets the "friday" field.
func (ecuo *EntCourseUpdateOne) SetFriday(b bool) *EntCourseUpdateOne {
	ecuo.mutation.SetFriday(b)
	return ecuo
}

// SetNillableFriday sets the "friday" field if the given value is not nil.
func (ecuo *EntCourseUpdateOne) SetNillableFriday(b *bool) *EntCourseUpdateOne {
	if b != nil {
		ecuo.SetFriday(*b)
	}
	return ecuo
}

// ClearFriday clears the value of the "friday" field.
func (ecuo *EntCourseUpdateOne) ClearFriday() *EntCourseUpdateOne {
	ecuo.mutation.ClearFriday()
	return ecuo
}

// SetSaturday sets the "saturday" field.
func (ecuo *EntCourseUpdateOne) SetSaturday(b bool) *EntCourseUpdateOne {
	ecuo.mutation.SetSaturday(b)
	return ecuo
}

// SetNillableSaturday sets the "saturday" field if the given value is not nil.
func (ecuo *EntCourseUpdateOne) SetNillableSaturday(b *bool) *EntCourseUpdateOne {
	if b != nil {
		ecuo.SetSaturday(*b)
	}
	return ecuo
}

// ClearSaturday clears the value of the "saturday" field.
func (ecuo *EntCourseUpdateOne) ClearSaturday() *EntCourseUpdateOne {
	ecuo.mutation.ClearSaturday()
	return ecuo
}

// SetSunday sets the "sunday" field.
func (ecuo *EntCourseUpdateOne) SetSunday(b bool) *EntCourseUpdateOne {
	ecuo.mutation.SetSunday(b)
	return ecuo
}

// SetNillableSunday sets the "sunday" field if the given value is not nil.
func (ecuo *EntCourseUpdateOne) SetNillableSunday(b *bool) *EntCourseUpdateOne {
	if b != nil {
		ecuo.SetSunday(*b)
	}
	return ecuo
}

// ClearSunday clears the value of the "sunday" field.
func (ecuo *EntCourseUpdateOne) ClearSunday() *EntCourseUpdateOne {
	ecuo.mutation.ClearSunday()
	return ecuo
}

// AddCourseOwnerIDs adds the "courseOwner" edge to the EntUser entity by IDs.
func (ecuo *EntCourseUpdateOne) AddCourseOwnerIDs(ids ...int) *EntCourseUpdateOne {
	ecuo.mutation.AddCourseOwnerIDs(ids...)
	return ecuo
}

// AddCourseOwner adds the "courseOwner" edges to the EntUser entity.
func (ecuo *EntCourseUpdateOne) AddCourseOwner(e ...*EntUser) *EntCourseUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecuo.AddCourseOwnerIDs(ids...)
}

// Mutation returns the EntCourseMutation object of the builder.
func (ecuo *EntCourseUpdateOne) Mutation() *EntCourseMutation {
	return ecuo.mutation
}

// ClearCourseOwner clears all "courseOwner" edges to the EntUser entity.
func (ecuo *EntCourseUpdateOne) ClearCourseOwner() *EntCourseUpdateOne {
	ecuo.mutation.ClearCourseOwner()
	return ecuo
}

// RemoveCourseOwnerIDs removes the "courseOwner" edge to EntUser entities by IDs.
func (ecuo *EntCourseUpdateOne) RemoveCourseOwnerIDs(ids ...int) *EntCourseUpdateOne {
	ecuo.mutation.RemoveCourseOwnerIDs(ids...)
	return ecuo
}

// RemoveCourseOwner removes "courseOwner" edges to EntUser entities.
func (ecuo *EntCourseUpdateOne) RemoveCourseOwner(e ...*EntUser) *EntCourseUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecuo.RemoveCourseOwnerIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ecuo *EntCourseUpdateOne) Select(field string, fields ...string) *EntCourseUpdateOne {
	ecuo.fields = append([]string{field}, fields...)
	return ecuo
}

// Save executes the query and returns the updated EntCourse entity.
func (ecuo *EntCourseUpdateOne) Save(ctx context.Context) (*EntCourse, error) {
	var (
		err  error
		node *EntCourse
	)
	if len(ecuo.hooks) == 0 {
		if err = ecuo.check(); err != nil {
			return nil, err
		}
		node, err = ecuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntCourseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ecuo.check(); err != nil {
				return nil, err
			}
			ecuo.mutation = mutation
			node, err = ecuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ecuo.hooks) - 1; i >= 0; i-- {
			if ecuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ecuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ecuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (ecuo *EntCourseUpdateOne) SaveX(ctx context.Context) *EntCourse {
	node, err := ecuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ecuo *EntCourseUpdateOne) Exec(ctx context.Context) error {
	_, err := ecuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecuo *EntCourseUpdateOne) ExecX(ctx context.Context) {
	if err := ecuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecuo *EntCourseUpdateOne) check() error {
	if v, ok := ecuo.mutation.Name(); ok {
		if err := entcourse.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "EntCourse.name": %w`, err)}
		}
	}
	if v, ok := ecuo.mutation.PaymentMethod(); ok {
		if err := entcourse.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "paymentMethod", err: fmt.Errorf(`ent: validator failed for field "EntCourse.paymentMethod": %w`, err)}
		}
	}
	return nil
}

func (ecuo *EntCourseUpdateOne) sqlSave(ctx context.Context) (_node *EntCourse, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entcourse.Table,
			Columns: entcourse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entcourse.FieldID,
			},
		},
	}
	id, ok := ecuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EntCourse.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ecuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entcourse.FieldID)
		for _, f := range fields {
			if !entcourse.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != entcourse.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ecuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entcourse.FieldName,
		})
	}
	if value, ok := ecuo.mutation.CourseUrl(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entcourse.FieldCourseUrl,
		})
	}
	if ecuo.mutation.CourseUrlCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: entcourse.FieldCourseUrl,
		})
	}
	if value, ok := ecuo.mutation.PaymentMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: entcourse.FieldPaymentMethod,
		})
	}
	if ecuo.mutation.PaymentMethodCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: entcourse.FieldPaymentMethod,
		})
	}
	if value, ok := ecuo.mutation.PaymentAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entcourse.FieldPaymentAmount,
		})
	}
	if value, ok := ecuo.mutation.AddedPaymentAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entcourse.FieldPaymentAmount,
		})
	}
	if ecuo.mutation.PaymentAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: entcourse.FieldPaymentAmount,
		})
	}
	if value, ok := ecuo.mutation.StartDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entcourse.FieldStartDate,
		})
	}
	if ecuo.mutation.StartDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: entcourse.FieldStartDate,
		})
	}
	if value, ok := ecuo.mutation.EndDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entcourse.FieldEndDate,
		})
	}
	if ecuo.mutation.EndDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: entcourse.FieldEndDate,
		})
	}
	if value, ok := ecuo.mutation.Monday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldMonday,
		})
	}
	if ecuo.mutation.MondayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldMonday,
		})
	}
	if value, ok := ecuo.mutation.Tuesday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldTuesday,
		})
	}
	if ecuo.mutation.TuesdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldTuesday,
		})
	}
	if value, ok := ecuo.mutation.Wednesday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldWednesday,
		})
	}
	if ecuo.mutation.WednesdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldWednesday,
		})
	}
	if value, ok := ecuo.mutation.Thursday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldThursday,
		})
	}
	if ecuo.mutation.ThursdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldThursday,
		})
	}
	if value, ok := ecuo.mutation.Friday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldFriday,
		})
	}
	if ecuo.mutation.FridayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldFriday,
		})
	}
	if value, ok := ecuo.mutation.Saturday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldSaturday,
		})
	}
	if ecuo.mutation.SaturdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldSaturday,
		})
	}
	if value, ok := ecuo.mutation.Sunday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entcourse.FieldSunday,
		})
	}
	if ecuo.mutation.SundayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entcourse.FieldSunday,
		})
	}
	if ecuo.mutation.CourseOwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.RemovedCourseOwnerIDs(); len(nodes) > 0 && !ecuo.mutation.CourseOwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.CourseOwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EntCourse{config: ecuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ecuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entcourse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
