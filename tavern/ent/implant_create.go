// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/ent/implant"
	"github.com/kcarretto/realm/tavern/ent/implantconfig"
	"github.com/kcarretto/realm/tavern/ent/target"
)

// ImplantCreate is the builder for creating a Implant entity.
type ImplantCreate struct {
	config
	mutation *ImplantMutation
	hooks    []Hook
}

// SetSessionID sets the "sessionID" field.
func (ic *ImplantCreate) SetSessionID(s string) *ImplantCreate {
	ic.mutation.SetSessionID(s)
	return ic
}

// SetProcessName sets the "processName" field.
func (ic *ImplantCreate) SetProcessName(s string) *ImplantCreate {
	ic.mutation.SetProcessName(s)
	return ic
}

// SetNillableProcessName sets the "processName" field if the given value is not nil.
func (ic *ImplantCreate) SetNillableProcessName(s *string) *ImplantCreate {
	if s != nil {
		ic.SetProcessName(*s)
	}
	return ic
}

// SetTargetID sets the "target" edge to the Target entity by ID.
func (ic *ImplantCreate) SetTargetID(id int) *ImplantCreate {
	ic.mutation.SetTargetID(id)
	return ic
}

// SetTarget sets the "target" edge to the Target entity.
func (ic *ImplantCreate) SetTarget(t *Target) *ImplantCreate {
	return ic.SetTargetID(t.ID)
}

// SetConfigID sets the "config" edge to the ImplantConfig entity by ID.
func (ic *ImplantCreate) SetConfigID(id int) *ImplantCreate {
	ic.mutation.SetConfigID(id)
	return ic
}

// SetConfig sets the "config" edge to the ImplantConfig entity.
func (ic *ImplantCreate) SetConfig(i *ImplantConfig) *ImplantCreate {
	return ic.SetConfigID(i.ID)
}

// Mutation returns the ImplantMutation object of the builder.
func (ic *ImplantCreate) Mutation() *ImplantMutation {
	return ic.mutation
}

// Save creates the Implant in the database.
func (ic *ImplantCreate) Save(ctx context.Context) (*Implant, error) {
	var (
		err  error
		node *Implant
	)
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImplantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ImplantCreate) SaveX(ctx context.Context) *Implant {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ImplantCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ImplantCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ImplantCreate) check() error {
	if _, ok := ic.mutation.SessionID(); !ok {
		return &ValidationError{Name: "sessionID", err: errors.New(`ent: missing required field "Implant.sessionID"`)}
	}
	if v, ok := ic.mutation.SessionID(); ok {
		if err := implant.SessionIDValidator(v); err != nil {
			return &ValidationError{Name: "sessionID", err: fmt.Errorf(`ent: validator failed for field "Implant.sessionID": %w`, err)}
		}
	}
	if _, ok := ic.mutation.TargetID(); !ok {
		return &ValidationError{Name: "target", err: errors.New(`ent: missing required edge "Implant.target"`)}
	}
	if _, ok := ic.mutation.ConfigID(); !ok {
		return &ValidationError{Name: "config", err: errors.New(`ent: missing required edge "Implant.config"`)}
	}
	return nil
}

func (ic *ImplantCreate) sqlSave(ctx context.Context) (*Implant, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ic *ImplantCreate) createSpec() (*Implant, *sqlgraph.CreateSpec) {
	var (
		_node = &Implant{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: implant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: implant.FieldID,
			},
		}
	)
	if value, ok := ic.mutation.SessionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: implant.FieldSessionID,
		})
		_node.SessionID = value
	}
	if value, ok := ic.mutation.ProcessName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: implant.FieldProcessName,
		})
		_node.ProcessName = value
	}
	if nodes := ic.mutation.TargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   implant.TargetTable,
			Columns: []string{implant.TargetColumn},
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
		_node.implant_target = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   implant.ConfigTable,
			Columns: []string{implant.ConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: implantconfig.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.implant_config = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ImplantCreateBulk is the builder for creating many Implant entities in bulk.
type ImplantCreateBulk struct {
	config
	builders []*ImplantCreate
}

// Save creates the Implant entities in the database.
func (icb *ImplantCreateBulk) Save(ctx context.Context) ([]*Implant, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Implant, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImplantMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ImplantCreateBulk) SaveX(ctx context.Context) []*Implant {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ImplantCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ImplantCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
