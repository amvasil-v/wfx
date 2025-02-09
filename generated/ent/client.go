// SPDX-FileCopyrightText: The entgo authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/siemens/wfx/generated/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/siemens/wfx/generated/ent/history"
	"github.com/siemens/wfx/generated/ent/job"
	"github.com/siemens/wfx/generated/ent/tag"
	"github.com/siemens/wfx/generated/ent/workflow"

	stdsql "database/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// History is the client for interacting with the History builders.
	History *HistoryClient
	// Job is the client for interacting with the Job builders.
	Job *JobClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// Workflow is the client for interacting with the Workflow builders.
	Workflow *WorkflowClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.History = NewHistoryClient(c.config)
	c.Job = NewJobClient(c.config)
	c.Tag = NewTagClient(c.config)
	c.Workflow = NewWorkflowClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		History:  NewHistoryClient(cfg),
		Job:      NewJobClient(cfg),
		Tag:      NewTagClient(cfg),
		Workflow: NewWorkflowClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		History:  NewHistoryClient(cfg),
		Job:      NewJobClient(cfg),
		Tag:      NewTagClient(cfg),
		Workflow: NewWorkflowClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		History.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.History.Use(hooks...)
	c.Job.Use(hooks...)
	c.Tag.Use(hooks...)
	c.Workflow.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.History.Intercept(interceptors...)
	c.Job.Intercept(interceptors...)
	c.Tag.Intercept(interceptors...)
	c.Workflow.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *HistoryMutation:
		return c.History.mutate(ctx, m)
	case *JobMutation:
		return c.Job.mutate(ctx, m)
	case *TagMutation:
		return c.Tag.mutate(ctx, m)
	case *WorkflowMutation:
		return c.Workflow.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// HistoryClient is a client for the History schema.
type HistoryClient struct {
	config
}

// NewHistoryClient returns a client for the History from the given config.
func NewHistoryClient(c config) *HistoryClient {
	return &HistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `history.Hooks(f(g(h())))`.
func (c *HistoryClient) Use(hooks ...Hook) {
	c.hooks.History = append(c.hooks.History, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `history.Intercept(f(g(h())))`.
func (c *HistoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.History = append(c.inters.History, interceptors...)
}

// Create returns a builder for creating a History entity.
func (c *HistoryClient) Create() *HistoryCreate {
	mutation := newHistoryMutation(c.config, OpCreate)
	return &HistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of History entities.
func (c *HistoryClient) CreateBulk(builders ...*HistoryCreate) *HistoryCreateBulk {
	return &HistoryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *HistoryClient) MapCreateBulk(slice any, setFunc func(*HistoryCreate, int)) *HistoryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &HistoryCreateBulk{err: fmt.Errorf("calling to HistoryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*HistoryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &HistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for History.
func (c *HistoryClient) Update() *HistoryUpdate {
	mutation := newHistoryMutation(c.config, OpUpdate)
	return &HistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HistoryClient) UpdateOne(h *History) *HistoryUpdateOne {
	mutation := newHistoryMutation(c.config, OpUpdateOne, withHistory(h))
	return &HistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HistoryClient) UpdateOneID(id int) *HistoryUpdateOne {
	mutation := newHistoryMutation(c.config, OpUpdateOne, withHistoryID(id))
	return &HistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for History.
func (c *HistoryClient) Delete() *HistoryDelete {
	mutation := newHistoryMutation(c.config, OpDelete)
	return &HistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *HistoryClient) DeleteOne(h *History) *HistoryDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *HistoryClient) DeleteOneID(id int) *HistoryDeleteOne {
	builder := c.Delete().Where(history.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HistoryDeleteOne{builder}
}

// Query returns a query builder for History.
func (c *HistoryClient) Query() *HistoryQuery {
	return &HistoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeHistory},
		inters: c.Interceptors(),
	}
}

// Get returns a History entity by its id.
func (c *HistoryClient) Get(ctx context.Context, id int) (*History, error) {
	return c.Query().Where(history.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HistoryClient) GetX(ctx context.Context, id int) *History {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryJob queries the job edge of a History.
func (c *HistoryClient) QueryJob(h *History) *JobQuery {
	query := (&JobClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(history.Table, history.FieldID, id),
			sqlgraph.To(job.Table, job.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, history.JobTable, history.JobColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *HistoryClient) Hooks() []Hook {
	return c.hooks.History
}

// Interceptors returns the client interceptors.
func (c *HistoryClient) Interceptors() []Interceptor {
	return c.inters.History
}

func (c *HistoryClient) mutate(ctx context.Context, m *HistoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&HistoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&HistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&HistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&HistoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown History mutation op: %q", m.Op())
	}
}

// JobClient is a client for the Job schema.
type JobClient struct {
	config
}

// NewJobClient returns a client for the Job from the given config.
func NewJobClient(c config) *JobClient {
	return &JobClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `job.Hooks(f(g(h())))`.
func (c *JobClient) Use(hooks ...Hook) {
	c.hooks.Job = append(c.hooks.Job, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `job.Intercept(f(g(h())))`.
func (c *JobClient) Intercept(interceptors ...Interceptor) {
	c.inters.Job = append(c.inters.Job, interceptors...)
}

// Create returns a builder for creating a Job entity.
func (c *JobClient) Create() *JobCreate {
	mutation := newJobMutation(c.config, OpCreate)
	return &JobCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Job entities.
func (c *JobClient) CreateBulk(builders ...*JobCreate) *JobCreateBulk {
	return &JobCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *JobClient) MapCreateBulk(slice any, setFunc func(*JobCreate, int)) *JobCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &JobCreateBulk{err: fmt.Errorf("calling to JobClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*JobCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &JobCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Job.
func (c *JobClient) Update() *JobUpdate {
	mutation := newJobMutation(c.config, OpUpdate)
	return &JobUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *JobClient) UpdateOne(j *Job) *JobUpdateOne {
	mutation := newJobMutation(c.config, OpUpdateOne, withJob(j))
	return &JobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *JobClient) UpdateOneID(id string) *JobUpdateOne {
	mutation := newJobMutation(c.config, OpUpdateOne, withJobID(id))
	return &JobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Job.
func (c *JobClient) Delete() *JobDelete {
	mutation := newJobMutation(c.config, OpDelete)
	return &JobDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *JobClient) DeleteOne(j *Job) *JobDeleteOne {
	return c.DeleteOneID(j.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *JobClient) DeleteOneID(id string) *JobDeleteOne {
	builder := c.Delete().Where(job.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &JobDeleteOne{builder}
}

// Query returns a query builder for Job.
func (c *JobClient) Query() *JobQuery {
	return &JobQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeJob},
		inters: c.Interceptors(),
	}
}

// Get returns a Job entity by its id.
func (c *JobClient) Get(ctx context.Context, id string) (*Job, error) {
	return c.Query().Where(job.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *JobClient) GetX(ctx context.Context, id string) *Job {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWorkflow queries the workflow edge of a Job.
func (c *JobClient) QueryWorkflow(j *Job) *WorkflowQuery {
	query := (&WorkflowClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := j.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(job.Table, job.FieldID, id),
			sqlgraph.To(workflow.Table, workflow.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, job.WorkflowTable, job.WorkflowColumn),
		)
		fromV = sqlgraph.Neighbors(j.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHistory queries the history edge of a Job.
func (c *JobClient) QueryHistory(j *Job) *HistoryQuery {
	query := (&HistoryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := j.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(job.Table, job.FieldID, id),
			sqlgraph.To(history.Table, history.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, job.HistoryTable, job.HistoryColumn),
		)
		fromV = sqlgraph.Neighbors(j.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTags queries the tags edge of a Job.
func (c *JobClient) QueryTags(j *Job) *TagQuery {
	query := (&TagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := j.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(job.Table, job.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, job.TagsTable, job.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(j.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *JobClient) Hooks() []Hook {
	return c.hooks.Job
}

// Interceptors returns the client interceptors.
func (c *JobClient) Interceptors() []Interceptor {
	return c.inters.Job
}

func (c *JobClient) mutate(ctx context.Context, m *JobMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&JobCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&JobUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&JobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&JobDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Job mutation op: %q", m.Op())
	}
}

// TagClient is a client for the Tag schema.
type TagClient struct {
	config
}

// NewTagClient returns a client for the Tag from the given config.
func NewTagClient(c config) *TagClient {
	return &TagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tag.Hooks(f(g(h())))`.
func (c *TagClient) Use(hooks ...Hook) {
	c.hooks.Tag = append(c.hooks.Tag, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `tag.Intercept(f(g(h())))`.
func (c *TagClient) Intercept(interceptors ...Interceptor) {
	c.inters.Tag = append(c.inters.Tag, interceptors...)
}

// Create returns a builder for creating a Tag entity.
func (c *TagClient) Create() *TagCreate {
	mutation := newTagMutation(c.config, OpCreate)
	return &TagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tag entities.
func (c *TagClient) CreateBulk(builders ...*TagCreate) *TagCreateBulk {
	return &TagCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TagClient) MapCreateBulk(slice any, setFunc func(*TagCreate, int)) *TagCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TagCreateBulk{err: fmt.Errorf("calling to TagClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TagCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tag.
func (c *TagClient) Update() *TagUpdate {
	mutation := newTagMutation(c.config, OpUpdate)
	return &TagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TagClient) UpdateOne(t *Tag) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTag(t))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TagClient) UpdateOneID(id int) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTagID(id))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tag.
func (c *TagClient) Delete() *TagDelete {
	mutation := newTagMutation(c.config, OpDelete)
	return &TagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TagClient) DeleteOne(t *Tag) *TagDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TagClient) DeleteOneID(id int) *TagDeleteOne {
	builder := c.Delete().Where(tag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TagDeleteOne{builder}
}

// Query returns a query builder for Tag.
func (c *TagClient) Query() *TagQuery {
	return &TagQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTag},
		inters: c.Interceptors(),
	}
}

// Get returns a Tag entity by its id.
func (c *TagClient) Get(ctx context.Context, id int) (*Tag, error) {
	return c.Query().Where(tag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TagClient) GetX(ctx context.Context, id int) *Tag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryJobs queries the jobs edge of a Tag.
func (c *TagClient) QueryJobs(t *Tag) *JobQuery {
	query := (&JobClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(job.Table, job.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, tag.JobsTable, tag.JobsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TagClient) Hooks() []Hook {
	return c.hooks.Tag
}

// Interceptors returns the client interceptors.
func (c *TagClient) Interceptors() []Interceptor {
	return c.inters.Tag
}

func (c *TagClient) mutate(ctx context.Context, m *TagMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TagCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TagUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TagDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Tag mutation op: %q", m.Op())
	}
}

// WorkflowClient is a client for the Workflow schema.
type WorkflowClient struct {
	config
}

// NewWorkflowClient returns a client for the Workflow from the given config.
func NewWorkflowClient(c config) *WorkflowClient {
	return &WorkflowClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `workflow.Hooks(f(g(h())))`.
func (c *WorkflowClient) Use(hooks ...Hook) {
	c.hooks.Workflow = append(c.hooks.Workflow, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `workflow.Intercept(f(g(h())))`.
func (c *WorkflowClient) Intercept(interceptors ...Interceptor) {
	c.inters.Workflow = append(c.inters.Workflow, interceptors...)
}

// Create returns a builder for creating a Workflow entity.
func (c *WorkflowClient) Create() *WorkflowCreate {
	mutation := newWorkflowMutation(c.config, OpCreate)
	return &WorkflowCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Workflow entities.
func (c *WorkflowClient) CreateBulk(builders ...*WorkflowCreate) *WorkflowCreateBulk {
	return &WorkflowCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *WorkflowClient) MapCreateBulk(slice any, setFunc func(*WorkflowCreate, int)) *WorkflowCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &WorkflowCreateBulk{err: fmt.Errorf("calling to WorkflowClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*WorkflowCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &WorkflowCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Workflow.
func (c *WorkflowClient) Update() *WorkflowUpdate {
	mutation := newWorkflowMutation(c.config, OpUpdate)
	return &WorkflowUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WorkflowClient) UpdateOne(w *Workflow) *WorkflowUpdateOne {
	mutation := newWorkflowMutation(c.config, OpUpdateOne, withWorkflow(w))
	return &WorkflowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WorkflowClient) UpdateOneID(id int) *WorkflowUpdateOne {
	mutation := newWorkflowMutation(c.config, OpUpdateOne, withWorkflowID(id))
	return &WorkflowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Workflow.
func (c *WorkflowClient) Delete() *WorkflowDelete {
	mutation := newWorkflowMutation(c.config, OpDelete)
	return &WorkflowDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *WorkflowClient) DeleteOne(w *Workflow) *WorkflowDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *WorkflowClient) DeleteOneID(id int) *WorkflowDeleteOne {
	builder := c.Delete().Where(workflow.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WorkflowDeleteOne{builder}
}

// Query returns a query builder for Workflow.
func (c *WorkflowClient) Query() *WorkflowQuery {
	return &WorkflowQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeWorkflow},
		inters: c.Interceptors(),
	}
}

// Get returns a Workflow entity by its id.
func (c *WorkflowClient) Get(ctx context.Context, id int) (*Workflow, error) {
	return c.Query().Where(workflow.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WorkflowClient) GetX(ctx context.Context, id int) *Workflow {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryJobs queries the jobs edge of a Workflow.
func (c *WorkflowClient) QueryJobs(w *Workflow) *JobQuery {
	query := (&JobClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workflow.Table, workflow.FieldID, id),
			sqlgraph.To(job.Table, job.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workflow.JobsTable, workflow.JobsColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WorkflowClient) Hooks() []Hook {
	return c.hooks.Workflow
}

// Interceptors returns the client interceptors.
func (c *WorkflowClient) Interceptors() []Interceptor {
	return c.inters.Workflow
}

func (c *WorkflowClient) mutate(ctx context.Context, m *WorkflowMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&WorkflowCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&WorkflowUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&WorkflowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&WorkflowDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Workflow mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		History, Job, Tag, Workflow []ent.Hook
	}
	inters struct {
		History, Job, Tag, Workflow []ent.Interceptor
	}
)

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
