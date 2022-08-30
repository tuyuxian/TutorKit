// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/entcomment"
	"backend/ent/entpost"
	"backend/ent/entuser"
	"backend/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntCommentQuery is the builder for querying EntComment entities.
type EntCommentQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.EntComment
	withBelongsTo *EntPostQuery
	withOwnedBy   *EntUserQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EntCommentQuery builder.
func (ecq *EntCommentQuery) Where(ps ...predicate.EntComment) *EntCommentQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit adds a limit step to the query.
func (ecq *EntCommentQuery) Limit(limit int) *EntCommentQuery {
	ecq.limit = &limit
	return ecq
}

// Offset adds an offset step to the query.
func (ecq *EntCommentQuery) Offset(offset int) *EntCommentQuery {
	ecq.offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *EntCommentQuery) Unique(unique bool) *EntCommentQuery {
	ecq.unique = &unique
	return ecq
}

// Order adds an order step to the query.
func (ecq *EntCommentQuery) Order(o ...OrderFunc) *EntCommentQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// QueryBelongsTo chains the current query on the "belongsTo" edge.
func (ecq *EntCommentQuery) QueryBelongsTo() *EntPostQuery {
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
			sqlgraph.From(entcomment.Table, entcomment.FieldID, selector),
			sqlgraph.To(entpost.Table, entpost.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entcomment.BelongsToTable, entcomment.BelongsToColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwnedBy chains the current query on the "ownedBy" edge.
func (ecq *EntCommentQuery) QueryOwnedBy() *EntUserQuery {
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
			sqlgraph.From(entcomment.Table, entcomment.FieldID, selector),
			sqlgraph.To(entuser.Table, entuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entcomment.OwnedByTable, entcomment.OwnedByColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EntComment entity from the query.
// Returns a *NotFoundError when no EntComment was found.
func (ecq *EntCommentQuery) First(ctx context.Context) (*EntComment, error) {
	nodes, err := ecq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entcomment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *EntCommentQuery) FirstX(ctx context.Context) *EntComment {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EntComment ID from the query.
// Returns a *NotFoundError when no EntComment ID was found.
func (ecq *EntCommentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entcomment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *EntCommentQuery) FirstIDX(ctx context.Context) int {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EntComment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EntComment entity is found.
// Returns a *NotFoundError when no EntComment entities are found.
func (ecq *EntCommentQuery) Only(ctx context.Context) (*EntComment, error) {
	nodes, err := ecq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entcomment.Label}
	default:
		return nil, &NotSingularError{entcomment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *EntCommentQuery) OnlyX(ctx context.Context) *EntComment {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EntComment ID in the query.
// Returns a *NotSingularError when more than one EntComment ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *EntCommentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entcomment.Label}
	default:
		err = &NotSingularError{entcomment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *EntCommentQuery) OnlyIDX(ctx context.Context) int {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EntComments.
func (ecq *EntCommentQuery) All(ctx context.Context) ([]*EntComment, error) {
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ecq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ecq *EntCommentQuery) AllX(ctx context.Context) []*EntComment {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EntComment IDs.
func (ecq *EntCommentQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ecq.Select(entcomment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *EntCommentQuery) IDsX(ctx context.Context) []int {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *EntCommentQuery) Count(ctx context.Context) (int, error) {
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ecq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *EntCommentQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *EntCommentQuery) Exist(ctx context.Context) (bool, error) {
	if err := ecq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ecq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ecq *EntCommentQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EntCommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *EntCommentQuery) Clone() *EntCommentQuery {
	if ecq == nil {
		return nil
	}
	return &EntCommentQuery{
		config:        ecq.config,
		limit:         ecq.limit,
		offset:        ecq.offset,
		order:         append([]OrderFunc{}, ecq.order...),
		predicates:    append([]predicate.EntComment{}, ecq.predicates...),
		withBelongsTo: ecq.withBelongsTo.Clone(),
		withOwnedBy:   ecq.withOwnedBy.Clone(),
		// clone intermediate query.
		sql:    ecq.sql.Clone(),
		path:   ecq.path,
		unique: ecq.unique,
	}
}

// WithBelongsTo tells the query-builder to eager-load the nodes that are connected to
// the "belongsTo" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *EntCommentQuery) WithBelongsTo(opts ...func(*EntPostQuery)) *EntCommentQuery {
	query := &EntPostQuery{config: ecq.config}
	for _, opt := range opts {
		opt(query)
	}
	ecq.withBelongsTo = query
	return ecq
}

// WithOwnedBy tells the query-builder to eager-load the nodes that are connected to
// the "ownedBy" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *EntCommentQuery) WithOwnedBy(opts ...func(*EntUserQuery)) *EntCommentQuery {
	query := &EntUserQuery{config: ecq.config}
	for _, opt := range opts {
		opt(query)
	}
	ecq.withOwnedBy = query
	return ecq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Timestamp time.Time `json:"timestamp,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EntComment.Query().
//		GroupBy(entcomment.FieldTimestamp).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ecq *EntCommentQuery) GroupBy(field string, fields ...string) *EntCommentGroupBy {
	grbuild := &EntCommentGroupBy{config: ecq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ecq.sqlQuery(ctx), nil
	}
	grbuild.label = entcomment.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Timestamp time.Time `json:"timestamp,omitempty"`
//	}
//
//	client.EntComment.Query().
//		Select(entcomment.FieldTimestamp).
//		Scan(ctx, &v)
func (ecq *EntCommentQuery) Select(fields ...string) *EntCommentSelect {
	ecq.fields = append(ecq.fields, fields...)
	selbuild := &EntCommentSelect{EntCommentQuery: ecq}
	selbuild.label = entcomment.Label
	selbuild.flds, selbuild.scan = &ecq.fields, selbuild.Scan
	return selbuild
}

func (ecq *EntCommentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ecq.fields {
		if !entcomment.ValidColumn(f) {
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

func (ecq *EntCommentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EntComment, error) {
	var (
		nodes       = []*EntComment{}
		withFKs     = ecq.withFKs
		_spec       = ecq.querySpec()
		loadedTypes = [2]bool{
			ecq.withBelongsTo != nil,
			ecq.withOwnedBy != nil,
		}
	)
	if ecq.withBelongsTo != nil || ecq.withOwnedBy != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, entcomment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*EntComment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &EntComment{config: ecq.config}
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
	if query := ecq.withBelongsTo; query != nil {
		if err := ecq.loadBelongsTo(ctx, query, nodes, nil,
			func(n *EntComment, e *EntPost) { n.Edges.BelongsTo = e }); err != nil {
			return nil, err
		}
	}
	if query := ecq.withOwnedBy; query != nil {
		if err := ecq.loadOwnedBy(ctx, query, nodes, nil,
			func(n *EntComment, e *EntUser) { n.Edges.OwnedBy = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ecq *EntCommentQuery) loadBelongsTo(ctx context.Context, query *EntPostQuery, nodes []*EntComment, init func(*EntComment), assign func(*EntComment, *EntPost)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*EntComment)
	for i := range nodes {
		if nodes[i].ent_post_comment == nil {
			continue
		}
		fk := *nodes[i].ent_post_comment
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(entpost.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ent_post_comment" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ecq *EntCommentQuery) loadOwnedBy(ctx context.Context, query *EntUserQuery, nodes []*EntComment, init func(*EntComment), assign func(*EntComment, *EntUser)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*EntComment)
	for i := range nodes {
		if nodes[i].ent_user_comment == nil {
			continue
		}
		fk := *nodes[i].ent_user_comment
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(entuser.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ent_user_comment" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ecq *EntCommentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecq.querySpec()
	_spec.Node.Columns = ecq.fields
	if len(ecq.fields) > 0 {
		_spec.Unique = ecq.unique != nil && *ecq.unique
	}
	return sqlgraph.CountNodes(ctx, ecq.driver, _spec)
}

func (ecq *EntCommentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ecq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ecq *EntCommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entcomment.Table,
			Columns: entcomment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entcomment.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, entcomment.FieldID)
		for i := range fields {
			if fields[i] != entcomment.FieldID {
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

func (ecq *EntCommentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(entcomment.Table)
	columns := ecq.fields
	if len(columns) == 0 {
		columns = entcomment.Columns
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

// EntCommentGroupBy is the group-by builder for EntComment entities.
type EntCommentGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *EntCommentGroupBy) Aggregate(fns ...AggregateFunc) *EntCommentGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ecgb *EntCommentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ecgb.path(ctx)
	if err != nil {
		return err
	}
	ecgb.sql = query
	return ecgb.sqlScan(ctx, v)
}

func (ecgb *EntCommentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ecgb.fields {
		if !entcomment.ValidColumn(f) {
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

func (ecgb *EntCommentGroupBy) sqlQuery() *sql.Selector {
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

// EntCommentSelect is the builder for selecting fields of EntComment entities.
type EntCommentSelect struct {
	*EntCommentQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *EntCommentSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	ecs.sql = ecs.EntCommentQuery.sqlQuery(ctx)
	return ecs.sqlScan(ctx, v)
}

func (ecs *EntCommentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ecs.sql.Query()
	if err := ecs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
