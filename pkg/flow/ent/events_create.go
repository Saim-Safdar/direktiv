// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/events"
	"github.com/direktiv/direktiv/pkg/flow/ent/eventswait"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// EventsCreate is the builder for creating a Events entity.
type EventsCreate struct {
	config
	mutation *EventsMutation
	hooks    []Hook
}

// SetEvents sets the "events" field.
func (ec *EventsCreate) SetEvents(m []map[string]interface{}) *EventsCreate {
	ec.mutation.SetEvents(m)
	return ec
}

// SetCorrelations sets the "correlations" field.
func (ec *EventsCreate) SetCorrelations(s []string) *EventsCreate {
	ec.mutation.SetCorrelations(s)
	return ec
}

// SetSignature sets the "signature" field.
func (ec *EventsCreate) SetSignature(b []byte) *EventsCreate {
	ec.mutation.SetSignature(b)
	return ec
}

// SetCount sets the "count" field.
func (ec *EventsCreate) SetCount(i int) *EventsCreate {
	ec.mutation.SetCount(i)
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EventsCreate) SetCreatedAt(t time.Time) *EventsCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EventsCreate) SetNillableCreatedAt(t *time.Time) *EventsCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EventsCreate) SetUpdatedAt(t time.Time) *EventsCreate {
	ec.mutation.SetUpdatedAt(t)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *EventsCreate) SetNillableUpdatedAt(t *time.Time) *EventsCreate {
	if t != nil {
		ec.SetUpdatedAt(*t)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *EventsCreate) SetID(u uuid.UUID) *EventsCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (ec *EventsCreate) SetWorkflowID(id uuid.UUID) *EventsCreate {
	ec.mutation.SetWorkflowID(id)
	return ec
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (ec *EventsCreate) SetWorkflow(w *Workflow) *EventsCreate {
	return ec.SetWorkflowID(w.ID)
}

// AddWfeventswaitIDs adds the "wfeventswait" edge to the EventsWait entity by IDs.
func (ec *EventsCreate) AddWfeventswaitIDs(ids ...uuid.UUID) *EventsCreate {
	ec.mutation.AddWfeventswaitIDs(ids...)
	return ec
}

// AddWfeventswait adds the "wfeventswait" edges to the EventsWait entity.
func (ec *EventsCreate) AddWfeventswait(e ...*EventsWait) *EventsCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddWfeventswaitIDs(ids...)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (ec *EventsCreate) SetInstanceID(id uuid.UUID) *EventsCreate {
	ec.mutation.SetInstanceID(id)
	return ec
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (ec *EventsCreate) SetNillableInstanceID(id *uuid.UUID) *EventsCreate {
	if id != nil {
		ec = ec.SetInstanceID(*id)
	}
	return ec
}

// SetInstance sets the "instance" edge to the Instance entity.
func (ec *EventsCreate) SetInstance(i *Instance) *EventsCreate {
	return ec.SetInstanceID(i.ID)
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (ec *EventsCreate) SetNamespaceID(id uuid.UUID) *EventsCreate {
	ec.mutation.SetNamespaceID(id)
	return ec
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (ec *EventsCreate) SetNamespace(n *Namespace) *EventsCreate {
	return ec.SetNamespaceID(n.ID)
}

// Mutation returns the EventsMutation object of the builder.
func (ec *EventsCreate) Mutation() *EventsMutation {
	return ec.mutation
}

// Save creates the Events in the database.
func (ec *EventsCreate) Save(ctx context.Context) (*Events, error) {
	var (
		err  error
		node *Events
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventsCreate) SaveX(ctx context.Context) *Events {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventsCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventsCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EventsCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := events.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		v := events.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		v := events.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventsCreate) check() error {
	if _, ok := ec.mutation.Events(); !ok {
		return &ValidationError{Name: "events", err: errors.New(`ent: missing required field "events"`)}
	}
	if _, ok := ec.mutation.Correlations(); !ok {
		return &ValidationError{Name: "correlations", err: errors.New(`ent: missing required field "correlations"`)}
	}
	if _, ok := ec.mutation.Count(); !ok {
		return &ValidationError{Name: "count", err: errors.New(`ent: missing required field "count"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := ec.mutation.WorkflowID(); !ok {
		return &ValidationError{Name: "workflow", err: errors.New("ent: missing required edge \"workflow\"")}
	}
	if _, ok := ec.mutation.NamespaceID(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New("ent: missing required edge \"namespace\"")}
	}
	return nil
}

func (ec *EventsCreate) sqlSave(ctx context.Context) (*Events, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (ec *EventsCreate) createSpec() (*Events, *sqlgraph.CreateSpec) {
	var (
		_node = &Events{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: events.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: events.FieldID,
			},
		}
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.Events(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: events.FieldEvents,
		})
		_node.Events = value
	}
	if value, ok := ec.mutation.Correlations(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: events.FieldCorrelations,
		})
		_node.Correlations = value
	}
	if value, ok := ec.mutation.Signature(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: events.FieldSignature,
		})
		_node.Signature = value
	}
	if value, ok := ec.mutation.Count(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: events.FieldCount,
		})
		_node.Count = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: events.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: events.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ec.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.WorkflowTable,
			Columns: []string{events.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workflow_wfevents = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.WfeventswaitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   events.WfeventswaitTable,
			Columns: []string{events.WfeventswaitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: eventswait.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.InstanceTable,
			Columns: []string{events.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.instance_eventlisteners = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   events.NamespaceTable,
			Columns: []string{events.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.namespace_namespacelisteners = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EventsCreateBulk is the builder for creating many Events entities in bulk.
type EventsCreateBulk struct {
	config
	builders []*EventsCreate
}

// Save creates the Events entities in the database.
func (ecb *EventsCreateBulk) Save(ctx context.Context) ([]*Events, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Events, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventsMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventsCreateBulk) SaveX(ctx context.Context) []*Events {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventsCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventsCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
