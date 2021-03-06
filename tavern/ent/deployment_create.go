// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/ent/deployment"
	"github.com/kcarretto/realm/tavern/ent/deploymentconfig"
	"github.com/kcarretto/realm/tavern/ent/target"
)

// DeploymentCreate is the builder for creating a Deployment entity.
type DeploymentCreate struct {
	config
	mutation *DeploymentMutation
	hooks    []Hook
}

// SetOutput sets the "output" field.
func (dc *DeploymentCreate) SetOutput(s string) *DeploymentCreate {
	dc.mutation.SetOutput(s)
	return dc
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableOutput(s *string) *DeploymentCreate {
	if s != nil {
		dc.SetOutput(*s)
	}
	return dc
}

// SetError sets the "error" field.
func (dc *DeploymentCreate) SetError(s string) *DeploymentCreate {
	dc.mutation.SetError(s)
	return dc
}

// SetNillableError sets the "error" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableError(s *string) *DeploymentCreate {
	if s != nil {
		dc.SetError(*s)
	}
	return dc
}

// SetQueuedAt sets the "queuedAt" field.
func (dc *DeploymentCreate) SetQueuedAt(t time.Time) *DeploymentCreate {
	dc.mutation.SetQueuedAt(t)
	return dc
}

// SetNillableQueuedAt sets the "queuedAt" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableQueuedAt(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetQueuedAt(*t)
	}
	return dc
}

// SetLastModifiedAt sets the "lastModifiedAt" field.
func (dc *DeploymentCreate) SetLastModifiedAt(t time.Time) *DeploymentCreate {
	dc.mutation.SetLastModifiedAt(t)
	return dc
}

// SetNillableLastModifiedAt sets the "lastModifiedAt" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableLastModifiedAt(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetLastModifiedAt(*t)
	}
	return dc
}

// SetStartedAt sets the "startedAt" field.
func (dc *DeploymentCreate) SetStartedAt(t time.Time) *DeploymentCreate {
	dc.mutation.SetStartedAt(t)
	return dc
}

// SetNillableStartedAt sets the "startedAt" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableStartedAt(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetStartedAt(*t)
	}
	return dc
}

// SetFinishedAt sets the "finishedAt" field.
func (dc *DeploymentCreate) SetFinishedAt(t time.Time) *DeploymentCreate {
	dc.mutation.SetFinishedAt(t)
	return dc
}

// SetNillableFinishedAt sets the "finishedAt" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableFinishedAt(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetFinishedAt(*t)
	}
	return dc
}

// SetConfigID sets the "config" edge to the DeploymentConfig entity by ID.
func (dc *DeploymentCreate) SetConfigID(id int) *DeploymentCreate {
	dc.mutation.SetConfigID(id)
	return dc
}

// SetConfig sets the "config" edge to the DeploymentConfig entity.
func (dc *DeploymentCreate) SetConfig(d *DeploymentConfig) *DeploymentCreate {
	return dc.SetConfigID(d.ID)
}

// SetTargetID sets the "target" edge to the Target entity by ID.
func (dc *DeploymentCreate) SetTargetID(id int) *DeploymentCreate {
	dc.mutation.SetTargetID(id)
	return dc
}

// SetTarget sets the "target" edge to the Target entity.
func (dc *DeploymentCreate) SetTarget(t *Target) *DeploymentCreate {
	return dc.SetTargetID(t.ID)
}

// Mutation returns the DeploymentMutation object of the builder.
func (dc *DeploymentCreate) Mutation() *DeploymentMutation {
	return dc.mutation
}

// Save creates the Deployment in the database.
func (dc *DeploymentCreate) Save(ctx context.Context) (*Deployment, error) {
	var (
		err  error
		node *Deployment
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeploymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeploymentCreate) SaveX(ctx context.Context) *Deployment {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeploymentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeploymentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeploymentCreate) defaults() {
	if _, ok := dc.mutation.Output(); !ok {
		v := deployment.DefaultOutput
		dc.mutation.SetOutput(v)
	}
	if _, ok := dc.mutation.Error(); !ok {
		v := deployment.DefaultError
		dc.mutation.SetError(v)
	}
	if _, ok := dc.mutation.QueuedAt(); !ok {
		v := deployment.DefaultQueuedAt()
		dc.mutation.SetQueuedAt(v)
	}
	if _, ok := dc.mutation.LastModifiedAt(); !ok {
		v := deployment.DefaultLastModifiedAt()
		dc.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeploymentCreate) check() error {
	if _, ok := dc.mutation.Output(); !ok {
		return &ValidationError{Name: "output", err: errors.New(`ent: missing required field "Deployment.output"`)}
	}
	if _, ok := dc.mutation.Error(); !ok {
		return &ValidationError{Name: "error", err: errors.New(`ent: missing required field "Deployment.error"`)}
	}
	if _, ok := dc.mutation.QueuedAt(); !ok {
		return &ValidationError{Name: "queuedAt", err: errors.New(`ent: missing required field "Deployment.queuedAt"`)}
	}
	if _, ok := dc.mutation.LastModifiedAt(); !ok {
		return &ValidationError{Name: "lastModifiedAt", err: errors.New(`ent: missing required field "Deployment.lastModifiedAt"`)}
	}
	if _, ok := dc.mutation.ConfigID(); !ok {
		return &ValidationError{Name: "config", err: errors.New(`ent: missing required edge "Deployment.config"`)}
	}
	if _, ok := dc.mutation.TargetID(); !ok {
		return &ValidationError{Name: "target", err: errors.New(`ent: missing required edge "Deployment.target"`)}
	}
	return nil
}

func (dc *DeploymentCreate) sqlSave(ctx context.Context) (*Deployment, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dc *DeploymentCreate) createSpec() (*Deployment, *sqlgraph.CreateSpec) {
	var (
		_node = &Deployment{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deployment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deployment.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Output(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deployment.FieldOutput,
		})
		_node.Output = value
	}
	if value, ok := dc.mutation.Error(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deployment.FieldError,
		})
		_node.Error = value
	}
	if value, ok := dc.mutation.QueuedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deployment.FieldQueuedAt,
		})
		_node.QueuedAt = value
	}
	if value, ok := dc.mutation.LastModifiedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deployment.FieldLastModifiedAt,
		})
		_node.LastModifiedAt = value
	}
	if value, ok := dc.mutation.StartedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deployment.FieldStartedAt,
		})
		_node.StartedAt = value
	}
	if value, ok := dc.mutation.FinishedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deployment.FieldFinishedAt,
		})
		_node.FinishedAt = value
	}
	if nodes := dc.mutation.ConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.ConfigTable,
			Columns: []string{deployment.ConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deploymentconfig.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.deployment_config = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.TargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.TargetTable,
			Columns: []string{deployment.TargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: target.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.deployment_target = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeploymentCreateBulk is the builder for creating many Deployment entities in bulk.
type DeploymentCreateBulk struct {
	config
	builders []*DeploymentCreate
}

// Save creates the Deployment entities in the database.
func (dcb *DeploymentCreateBulk) Save(ctx context.Context) ([]*Deployment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Deployment, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeploymentMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeploymentCreateBulk) SaveX(ctx context.Context) []*Deployment {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeploymentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeploymentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
