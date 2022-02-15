// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/ent/credential"
	"github.com/kcarretto/realm/ent/target"
)

// TargetCreate is the builder for creating a Target entity.
type TargetCreate struct {
	config
	mutation *TargetMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TargetCreate) SetName(s string) *TargetCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetForwardConnectIP sets the "forwardConnectIP" field.
func (tc *TargetCreate) SetForwardConnectIP(s string) *TargetCreate {
	tc.mutation.SetForwardConnectIP(s)
	return tc
}

// AddCredentialIDs adds the "credentials" edge to the Credential entity by IDs.
func (tc *TargetCreate) AddCredentialIDs(ids ...int) *TargetCreate {
	tc.mutation.AddCredentialIDs(ids...)
	return tc
}

// AddCredentials adds the "credentials" edges to the Credential entity.
func (tc *TargetCreate) AddCredentials(c ...*Credential) *TargetCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tc.AddCredentialIDs(ids...)
}

// Mutation returns the TargetMutation object of the builder.
func (tc *TargetCreate) Mutation() *TargetMutation {
	return tc.mutation
}

// Save creates the Target in the database.
func (tc *TargetCreate) Save(ctx context.Context) (*Target, error) {
	var (
		err  error
		node *Target
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TargetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TargetCreate) SaveX(ctx context.Context) *Target {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TargetCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TargetCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TargetCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Target.name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := target.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Target.name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.ForwardConnectIP(); !ok {
		return &ValidationError{Name: "forwardConnectIP", err: errors.New(`ent: missing required field "Target.forwardConnectIP"`)}
	}
	if v, ok := tc.mutation.ForwardConnectIP(); ok {
		if err := target.ForwardConnectIPValidator(v); err != nil {
			return &ValidationError{Name: "forwardConnectIP", err: fmt.Errorf(`ent: validator failed for field "Target.forwardConnectIP": %w`, err)}
		}
	}
	return nil
}

func (tc *TargetCreate) sqlSave(ctx context.Context) (*Target, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TargetCreate) createSpec() (*Target, *sqlgraph.CreateSpec) {
	var (
		_node = &Target{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: target.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: target.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: target.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.ForwardConnectIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: target.FieldForwardConnectIP,
		})
		_node.ForwardConnectIP = value
	}
	if nodes := tc.mutation.CredentialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   target.CredentialsTable,
			Columns: []string{target.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: credential.FieldID,
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

// TargetCreateBulk is the builder for creating many Target entities in bulk.
type TargetCreateBulk struct {
	config
	builders []*TargetCreate
}

// Save creates the Target entities in the database.
func (tcb *TargetCreateBulk) Save(ctx context.Context) ([]*Target, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Target, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TargetMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TargetCreateBulk) SaveX(ctx context.Context) []*Target {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TargetCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TargetCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
