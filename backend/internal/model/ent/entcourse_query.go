// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/internal/model/ent/entattendance"
	"backend/internal/model/ent/entcourse"
	"backend/internal/model/ent/entpost"
	"backend/internal/model/ent/enttodo"
	"backend/internal/model/ent/entuser"
	"backend/internal/model/ent/predicate"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntCourseQuery is the builder for querying EntCourse entities.
type EntCourseQuery struct {
	config
	limit          *int
	offset         *int
	unique         *bool
	order          []OrderFunc
	fields         []string
	predicates     []predicate.EntCourse
	withTodo       *EntTodoQuery
	withAttendance *EntAttendanceQuery
	withPost       *EntPostQuery
	withOwnedBy    *EntUserQuery
	withJoinedBy   *EntUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EntCourseQuery builder.
func (ecq *EntCourseQuery) Where(ps ...predicate.EntCourse) *EntCourseQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit adds a limit step to the query.
func (ecq *EntCourseQuery) Limit(limit int) *EntCourseQuery {
	ecq.limit = &limit
	return ecq
}

// Offset adds an offset step to the query.
func (ecq *EntCourseQuery) Offset(offset int) *EntCourseQuery {
	ecq.offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *EntCourseQuery) Unique(unique bool) *EntCourseQuery {
	ecq.unique = &unique
	return ecq
}

// Order adds an order step to the query.
func (ecq *EntCourseQuery) Order(o ...OrderFunc) *EntCourseQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// QueryTodo chains the current query on the "todo" edge.
func (ecq *EntCourseQuery) QueryTodo() *EntTodoQuery {
	query := &EntTodoQuery{config: ecq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entcourse.Table, entcourse.FieldID, selector),
			sqlgraph.To(enttodo.Table, enttodo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, entcourse.TodoTable, entcourse.TodoColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAttendance chains the current query on the "attendance" edge.
func (ecq *EntCourseQuery) QueryAttendance() *EntAttendanceQuery {
	query := &EntAttendanceQuery{config: ecq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entcourse.Table, entcourse.FieldID, selector),
			sqlgraph.To(entattendance.Table, entattendance.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, entcourse.AttendanceTable, entcourse.AttendanceColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPost chains the current query on the "post" edge.
func (ecq *EntCourseQuery) QueryPost() *EntPostQuery {
	query := &EntPostQuery{config: ecq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entcourse.Table, entcourse.FieldID, selector),
			sqlgraph.To(entpost.Table, entpost.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, entcourse.PostTable, entcourse.PostColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwnedBy chains the current query on the "ownedBy" edge.
func (ecq *EntCourseQuery) QueryOwnedBy() *EntUserQuery {
	query := &EntUserQuery{config: ecq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entcourse.Table, entcourse.FieldID, selector),
			sqlgraph.To(entuser.Table, entuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, entcourse.OwnedByTable, entcourse.OwnedByPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryJoinedBy chains the current query on the "joinedBy" edge.
func (ecq *EntCourseQuery) QueryJoinedBy() *EntUserQuery {
	query := &EntUserQuery{config: ecq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entcourse.Table, entcourse.FieldID, selector),
			sqlgraph.To(entuser.Table, entuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, entcourse.JoinedByTable, entcourse.JoinedByPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EntCourse entity from the query.
// Returns a *NotFoundError when no EntCourse was found.
func (ecq *EntCourseQuery) First(ctx context.Context) (*EntCourse, error) {
	nodes, err := ecq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entcourse.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *EntCourseQuery) FirstX(ctx context.Context) *EntCourse {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EntCourse ID from the query.
// Returns a *NotFoundError when no EntCourse ID was found.
func (ecq *EntCourseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entcourse.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *EntCourseQuery) FirstIDX(ctx context.Context) int {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EntCourse entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EntCourse entity is found.
// Returns a *NotFoundError when no EntCourse entities are found.
func (ecq *EntCourseQuery) Only(ctx context.Context) (*EntCourse, error) {
	nodes, err := ecq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entcourse.Label}
	default:
		return nil, &NotSingularError{entcourse.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *EntCourseQuery) OnlyX(ctx context.Context) *EntCourse {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EntCourse ID in the query.
// Returns a *NotSingularError when more than one EntCourse ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *EntCourseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entcourse.Label}
	default:
		err = &NotSingularError{entcourse.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *EntCourseQuery) OnlyIDX(ctx context.Context) int {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EntCourses.
func (ecq *EntCourseQuery) All(ctx context.Context) ([]*EntCourse, error) {
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ecq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ecq *EntCourseQuery) AllX(ctx context.Context) []*EntCourse {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EntCourse IDs.
func (ecq *EntCourseQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ecq.Select(entcourse.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *EntCourseQuery) IDsX(ctx context.Context) []int {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *EntCourseQuery) Count(ctx context.Context) (int, error) {
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ecq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *EntCourseQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *EntCourseQuery) Exist(ctx context.Context) (bool, error) {
	if err := ecq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ecq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ecq *EntCourseQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EntCourseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *EntCourseQuery) Clone() *EntCourseQuery {
	if ecq == nil {
		return nil
	}
	return &EntCourseQuery{
		config:         ecq.config,
		limit:          ecq.limit,
		offset:         ecq.offset,
		order:          append([]OrderFunc{}, ecq.order...),
		predicates:     append([]predicate.EntCourse{}, ecq.predicates...),
		withTodo:       ecq.withTodo.Clone(),
		withAttendance: ecq.withAttendance.Clone(),
		withPost:       ecq.withPost.Clone(),
		withOwnedBy:    ecq.withOwnedBy.Clone(),
		withJoinedBy:   ecq.withJoinedBy.Clone(),
		// clone intermediate query.
		sql:    ecq.sql.Clone(),
		path:   ecq.path,
		unique: ecq.unique,
	}
}

// WithTodo tells the query-builder to eager-load the nodes that are connected to
// the "todo" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *EntCourseQuery) WithTodo(opts ...func(*EntTodoQuery)) *EntCourseQuery {
	query := &EntTodoQuery{config: ecq.config}
	for _, opt := range opts {
		opt(query)
	}
	ecq.withTodo = query
	return ecq
}

// WithAttendance tells the query-builder to eager-load the nodes that are connected to
// the "attendance" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *EntCourseQuery) WithAttendance(opts ...func(*EntAttendanceQuery)) *EntCourseQuery {
	query := &EntAttendanceQuery{config: ecq.config}
	for _, opt := range opts {
		opt(query)
	}
	ecq.withAttendance = query
	return ecq
}

// WithPost tells the query-builder to eager-load the nodes that are connected to
// the "post" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *EntCourseQuery) WithPost(opts ...func(*EntPostQuery)) *EntCourseQuery {
	query := &EntPostQuery{config: ecq.config}
	for _, opt := range opts {
		opt(query)
	}
	ecq.withPost = query
	return ecq
}

// WithOwnedBy tells the query-builder to eager-load the nodes that are connected to
// the "ownedBy" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *EntCourseQuery) WithOwnedBy(opts ...func(*EntUserQuery)) *EntCourseQuery {
	query := &EntUserQuery{config: ecq.config}
	for _, opt := range opts {
		opt(query)
	}
	ecq.withOwnedBy = query
	return ecq
}

// WithJoinedBy tells the query-builder to eager-load the nodes that are connected to
// the "joinedBy" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *EntCourseQuery) WithJoinedBy(opts ...func(*EntUserQuery)) *EntCourseQuery {
	query := &EntUserQuery{config: ecq.config}
	for _, opt := range opts {
		opt(query)
	}
	ecq.withJoinedBy = query
	return ecq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EntCourse.Query().
//		GroupBy(entcourse.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ecq *EntCourseQuery) GroupBy(field string, fields ...string) *EntCourseGroupBy {
	grbuild := &EntCourseGroupBy{config: ecq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ecq.sqlQuery(ctx), nil
	}
	grbuild.label = entcourse.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//	}
//
//	client.EntCourse.Query().
//		Select(entcourse.FieldCreatedAt).
//		Scan(ctx, &v)
func (ecq *EntCourseQuery) Select(fields ...string) *EntCourseSelect {
	ecq.fields = append(ecq.fields, fields...)
	selbuild := &EntCourseSelect{EntCourseQuery: ecq}
	selbuild.label = entcourse.Label
	selbuild.flds, selbuild.scan = &ecq.fields, selbuild.Scan
	return selbuild
}

func (ecq *EntCourseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ecq.fields {
		if !entcourse.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ecq.path != nil {
		prev, err := ecq.path(ctx)
		if err != nil {
			return err
		}
		ecq.sql = prev
	}
	return nil
}

func (ecq *EntCourseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EntCourse, error) {
	var (
		nodes       = []*EntCourse{}
		_spec       = ecq.querySpec()
		loadedTypes = [5]bool{
			ecq.withTodo != nil,
			ecq.withAttendance != nil,
			ecq.withPost != nil,
			ecq.withOwnedBy != nil,
			ecq.withJoinedBy != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*EntCourse).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &EntCourse{config: ecq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ecq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ecq.withTodo; query != nil {
		if err := ecq.loadTodo(ctx, query, nodes,
			func(n *EntCourse) { n.Edges.Todo = []*EntTodo{} },
			func(n *EntCourse, e *EntTodo) { n.Edges.Todo = append(n.Edges.Todo, e) }); err != nil {
			return nil, err
		}
	}
	if query := ecq.withAttendance; query != nil {
		if err := ecq.loadAttendance(ctx, query, nodes,
			func(n *EntCourse) { n.Edges.Attendance = []*EntAttendance{} },
			func(n *EntCourse, e *EntAttendance) { n.Edges.Attendance = append(n.Edges.Attendance, e) }); err != nil {
			return nil, err
		}
	}
	if query := ecq.withPost; query != nil {
		if err := ecq.loadPost(ctx, query, nodes,
			func(n *EntCourse) { n.Edges.Post = []*EntPost{} },
			func(n *EntCourse, e *EntPost) { n.Edges.Post = append(n.Edges.Post, e) }); err != nil {
			return nil, err
		}
	}
	if query := ecq.withOwnedBy; query != nil {
		if err := ecq.loadOwnedBy(ctx, query, nodes,
			func(n *EntCourse) { n.Edges.OwnedBy = []*EntUser{} },
			func(n *EntCourse, e *EntUser) { n.Edges.OwnedBy = append(n.Edges.OwnedBy, e) }); err != nil {
			return nil, err
		}
	}
	if query := ecq.withJoinedBy; query != nil {
		if err := ecq.loadJoinedBy(ctx, query, nodes,
			func(n *EntCourse) { n.Edges.JoinedBy = []*EntUser{} },
			func(n *EntCourse, e *EntUser) { n.Edges.JoinedBy = append(n.Edges.JoinedBy, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ecq *EntCourseQuery) loadTodo(ctx context.Context, query *EntTodoQuery, nodes []*EntCourse, init func(*EntCourse), assign func(*EntCourse, *EntTodo)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*EntCourse)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.EntTodo(func(s *sql.Selector) {
		s.Where(sql.InValues(entcourse.TodoColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ent_course_todo
		if fk == nil {
			return fmt.Errorf(`foreign-key "ent_course_todo" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ent_course_todo" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (ecq *EntCourseQuery) loadAttendance(ctx context.Context, query *EntAttendanceQuery, nodes []*EntCourse, init func(*EntCourse), assign func(*EntCourse, *EntAttendance)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*EntCourse)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.EntAttendance(func(s *sql.Selector) {
		s.Where(sql.InValues(entcourse.AttendanceColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ent_course_attendance
		if fk == nil {
			return fmt.Errorf(`foreign-key "ent_course_attendance" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ent_course_attendance" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (ecq *EntCourseQuery) loadPost(ctx context.Context, query *EntPostQuery, nodes []*EntCourse, init func(*EntCourse), assign func(*EntCourse, *EntPost)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*EntCourse)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.EntPost(func(s *sql.Selector) {
		s.Where(sql.InValues(entcourse.PostColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ent_course_post
		if fk == nil {
			return fmt.Errorf(`foreign-key "ent_course_post" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ent_course_post" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (ecq *EntCourseQuery) loadOwnedBy(ctx context.Context, query *EntUserQuery, nodes []*EntCourse, init func(*EntCourse), assign func(*EntCourse, *EntUser)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*EntCourse)
	nids := make(map[int]map[*EntCourse]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(entcourse.OwnedByTable)
		s.Join(joinT).On(s.C(entuser.FieldID), joinT.C(entcourse.OwnedByPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(entcourse.OwnedByPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(entcourse.OwnedByPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*EntCourse]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "ownedBy" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (ecq *EntCourseQuery) loadJoinedBy(ctx context.Context, query *EntUserQuery, nodes []*EntCourse, init func(*EntCourse), assign func(*EntCourse, *EntUser)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*EntCourse)
	nids := make(map[int]map[*EntCourse]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(entcourse.JoinedByTable)
		s.Join(joinT).On(s.C(entuser.FieldID), joinT.C(entcourse.JoinedByPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(entcourse.JoinedByPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(entcourse.JoinedByPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*EntCourse]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "joinedBy" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (ecq *EntCourseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecq.querySpec()
	_spec.Node.Columns = ecq.fields
	if len(ecq.fields) > 0 {
		_spec.Unique = ecq.unique != nil && *ecq.unique
	}
	return sqlgraph.CountNodes(ctx, ecq.driver, _spec)
}

func (ecq *EntCourseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ecq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ecq *EntCourseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entcourse.Table,
			Columns: entcourse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entcourse.FieldID,
			},
		},
		From:   ecq.sql,
		Unique: true,
	}
	if unique := ecq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ecq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entcourse.FieldID)
		for i := range fields {
			if fields[i] != entcourse.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ecq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ecq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ecq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ecq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ecq *EntCourseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(entcourse.Table)
	columns := ecq.fields
	if len(columns) == 0 {
		columns = entcourse.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ecq.sql != nil {
		selector = ecq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ecq.unique != nil && *ecq.unique {
		selector.Distinct()
	}
	for _, p := range ecq.predicates {
		p(selector)
	}
	for _, p := range ecq.order {
		p(selector)
	}
	if offset := ecq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ecq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EntCourseGroupBy is the group-by builder for EntCourse entities.
type EntCourseGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *EntCourseGroupBy) Aggregate(fns ...AggregateFunc) *EntCourseGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ecgb *EntCourseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ecgb.path(ctx)
	if err != nil {
		return err
	}
	ecgb.sql = query
	return ecgb.sqlScan(ctx, v)
}

func (ecgb *EntCourseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ecgb.fields {
		if !entcourse.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ecgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ecgb *EntCourseGroupBy) sqlQuery() *sql.Selector {
	selector := ecgb.sql.Select()
	aggregation := make([]string, 0, len(ecgb.fns))
	for _, fn := range ecgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ecgb.fields)+len(ecgb.fns))
		for _, f := range ecgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ecgb.fields...)...)
}

// EntCourseSelect is the builder for selecting fields of EntCourse entities.
type EntCourseSelect struct {
	*EntCourseQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *EntCourseSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	ecs.sql = ecs.EntCourseQuery.sqlQuery(ctx)
	return ecs.sqlScan(ctx, v)
}

func (ecs *EntCourseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ecs.sql.Query()
	if err := ecs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}