// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/entcourse"
	"backend/ent/enttodo"
	"backend/ent/entuser"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntTodoCreate is the builder for creating a EntTodo entity.
type EntTodoCreate struct {
	config
	mutation *EntTodoMutation
	hooks    []Hook
}

// SetDate sets the "date" field.
func (etc *EntTodoCreate) SetDate(t time.Time) *EntTodoCreate {
	etc.mutation.SetDate(t)
	return etc
}

// SetStartTime sets the "startTime" field.
func (etc *EntTodoCreate) SetStartTime(t time.Time) *EntTodoCreate {
	etc.mutation.SetStartTime(t)
	return etc
}

// SetNillableStartTime sets the "startTime" field if the given value is not nil.
func (etc *EntTodoCreate) SetNillableStartTime(t *time.Time) *EntTodoCreate {
	if t != nil {
		etc.SetStartTime(*t)
	}
	return etc
}

// SetEndTime sets the "endTime" field.
func (etc *EntTodoCreate) SetEndTime(t time.Time) *EntTodoCreate {
	etc.mutation.SetEndTime(t)
	return etc
}

// SetNillableEndTime sets the "endTime" field if the given value is not nil.
func (etc *EntTodoCreate) SetNillableEndTime(t *time.Time) *EntTodoCreate {
	if t != nil {
		etc.SetEndTime(*t)
	}
	return etc
}

// SetDay sets the "day" field.
func (etc *EntTodoCreate) SetDay(t time.Time) *EntTodoCreate {
	etc.mutation.SetDay(t)
	return etc
}

// SetNillableDay sets the "day" field if the given value is not nil.
func (etc *EntTodoCreate) SetNillableDay(t *time.Time) *EntTodoCreate {
	if t != nil {
		etc.SetDay(*t)
	}
	return etc
}

// SetLesson sets the "lesson" field.
func (etc *EntTodoCreate) SetLesson(s string) *EntTodoCreate {
	etc.mutation.SetLesson(s)
	return etc
}

// SetNillableLesson sets the "lesson" field if the given value is not nil.
func (etc *EntTodoCreate) SetNillableLesson(s *string) *EntTodoCreate {
	if s != nil {
		etc.SetLesson(*s)
	}
	return etc
}

// SetHomework sets the "homework" field.
func (etc *EntTodoCreate) SetHomework(s string) *EntTodoCreate {
	etc.mutation.SetHomework(s)
	return etc
}

// SetNillableHomework sets the "homework" field if the given value is not nil.
func (etc *EntTodoCreate) SetNillableHomework(s *string) *EntTodoCreate {
	if s != nil {
		etc.SetHomework(*s)
	}
	return etc
}

// SetStatus sets the "status" field.
func (etc *EntTodoCreate) SetStatus(e enttodo.Status) *EntTodoCreate {
	etc.mutation.SetStatus(e)
	return etc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (etc *EntTodoCreate) SetNillableStatus(e *enttodo.Status) *EntTodoCreate {
	if e != nil {
		etc.SetStatus(*e)
	}
	return etc
}

// SetTodoForID sets the "todoFor" edge to the EntCourse entity by ID.
func (etc *EntTodoCreate) SetTodoForID(id int) *EntTodoCreate {
	etc.mutation.SetTodoForID(id)
	return etc
}

// SetNillableTodoForID sets the "todoFor" edge to the EntCourse entity by ID if the given value is not nil.
func (etc *EntTodoCreate) SetNillableTodoForID(id *int) *EntTodoCreate {
	if id != nil {
		etc = etc.SetTodoForID(*id)
	}
	return etc
}

// SetTodoFor sets the "todoFor" edge to the EntCourse entity.
func (etc *EntTodoCreate) SetTodoFor(e *EntCourse) *EntTodoCreate {
	return etc.SetTodoForID(e.ID)
}

// SetOwnedByID sets the "ownedBy" edge to the EntUser entity by ID.
func (etc *EntTodoCreate) SetOwnedByID(id int) *EntTodoCreate {
	etc.mutation.SetOwnedByID(id)
	return etc
}

// SetNillableOwnedByID sets the "ownedBy" edge to the EntUser entity by ID if the given value is not nil.
func (etc *EntTodoCreate) SetNillableOwnedByID(id *int) *EntTodoCreate {
	if id != nil {
		etc = etc.SetOwnedByID(*id)
	}
	return etc
}

// SetOwnedBy sets the "ownedBy" edge to the EntUser entity.
func (etc *EntTodoCreate) SetOwnedBy(e *EntUser) *EntTodoCreate {
	return etc.SetOwnedByID(e.ID)
}

// Mutation returns the EntTodoMutation object of the builder.
func (etc *EntTodoCreate) Mutation() *EntTodoMutation {
	return etc.mutation
}

// Save creates the EntTodo in the database.
func (etc *EntTodoCreate) Save(ctx context.Context) (*EntTodo, error) {
	var (
		err  error
		node *EntTodo
	)
	etc.defaults()
	if len(etc.hooks) == 0 {
		if err = etc.check(); err != nil {
			return nil, err
		}
		node, err = etc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntTodoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = etc.check(); err != nil {
				return nil, err
			}
			etc.mutation = mutation
			if node, err = etc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(etc.hooks) - 1; i >= 0; i-- {
			if etc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = etc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, etc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EntTodo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EntTodoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (etc *EntTodoCreate) SaveX(ctx context.Context) *EntTodo {
	v, err := etc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (etc *EntTodoCreate) Exec(ctx context.Context) error {
	_, err := etc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etc *EntTodoCreate) ExecX(ctx context.Context) {
	if err := etc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etc *EntTodoCreate) defaults() {
	if _, ok := etc.mutation.Status(); !ok {
		v := enttodo.DefaultStatus
		etc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (etc *EntTodoCreate) check() error {
	if _, ok := etc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "EntTodo.date"`)}
	}
	if _, ok := etc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "EntTodo.status"`)}
	}
	if v, ok := etc.mutation.Status(); ok {
		if err := enttodo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "EntTodo.status": %w`, err)}
		}
	}
	return nil
}

func (etc *EntTodoCreate) sqlSave(ctx context.Context) (*EntTodo, error) {
	_node, _spec := etc.createSpec()
	if err := sqlgraph.CreateNode(ctx, etc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (etc *EntTodoCreate) createSpec() (*EntTodo, *sqlgraph.CreateSpec) {
	var (
		_node = &EntTodo{config: etc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: enttodo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: enttodo.FieldID,
			},
		}
	)
	if value, ok := etc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: enttodo.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := etc.mutation.StartTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: enttodo.FieldStartTime,
		})
		_node.StartTime = value
	}
	if value, ok := etc.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: enttodo.FieldEndTime,
		})
		_node.EndTime = value
	}
	if value, ok := etc.mutation.Day(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: enttodo.FieldDay,
		})
		_node.Day = value
	}
	if value, ok := etc.mutation.Lesson(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: enttodo.FieldLesson,
		})
		_node.Lesson = value
	}
	if value, ok := etc.mutation.Homework(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: enttodo.FieldHomework,
		})
		_node.Homework = value
	}
	if value, ok := etc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: enttodo.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := etc.mutation.TodoForIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   enttodo.TodoForTable,
			Columns: []string{enttodo.TodoForColumn},
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
		_node.ent_course_todo = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.OwnedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   enttodo.OwnedByTable,
			Columns: []string{enttodo.OwnedByColumn},
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
		_node.ent_user_todo = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EntTodoCreateBulk is the builder for creating many EntTodo entities in bulk.
type EntTodoCreateBulk struct {
	config
	builders []*EntTodoCreate
}

// Save creates the EntTodo entities in the database.
func (etcb *EntTodoCreateBulk) Save(ctx context.Context) ([]*EntTodo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(etcb.builders))
	nodes := make([]*EntTodo, len(etcb.builders))
	mutators := make([]Mutator, len(etcb.builders))
	for i := range etcb.builders {
		func(i int, root context.Context) {
			builder := etcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntTodoMutation)
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
					_, err = mutators[i+1].Mutate(root, etcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, etcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, etcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (etcb *EntTodoCreateBulk) SaveX(ctx context.Context) []*EntTodo {
	v, err := etcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (etcb *EntTodoCreateBulk) Exec(ctx context.Context) error {
	_, err := etcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etcb *EntTodoCreateBulk) ExecX(ctx context.Context) {
	if err := etcb.Exec(ctx); err != nil {
		panic(err)
	}
}
