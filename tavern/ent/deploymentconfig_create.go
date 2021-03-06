// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/ent/deployment"
	"github.com/kcarretto/realm/tavern/ent/deploymentconfig"
	"github.com/kcarretto/realm/tavern/ent/file"
	"github.com/kcarretto/realm/tavern/ent/implantconfig"
)

// DeploymentConfigCreate is the builder for creating a DeploymentConfig entity.
type DeploymentConfigCreate struct {
	config
	mutation *DeploymentConfigMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dcc *DeploymentConfigCreate) SetName(s string) *DeploymentConfigCreate {
	dcc.mutation.SetName(s)
	return dcc
}

// SetCmd sets the "cmd" field.
func (dcc *DeploymentConfigCreate) SetCmd(s string) *DeploymentConfigCreate {
	dcc.mutation.SetCmd(s)
	return dcc
}

// SetNillableCmd sets the "cmd" field if the given value is not nil.
func (dcc *DeploymentConfigCreate) SetNillableCmd(s *string) *DeploymentConfigCreate {
	if s != nil {
		dcc.SetCmd(*s)
	}
	return dcc
}

// SetStartCmd sets the "startCmd" field.
func (dcc *DeploymentConfigCreate) SetStartCmd(b bool) *DeploymentConfigCreate {
	dcc.mutation.SetStartCmd(b)
	return dcc
}

// SetNillableStartCmd sets the "startCmd" field if the given value is not nil.
func (dcc *DeploymentConfigCreate) SetNillableStartCmd(b *bool) *DeploymentConfigCreate {
	if b != nil {
		dcc.SetStartCmd(*b)
	}
	return dcc
}

// SetFileDst sets the "fileDst" field.
func (dcc *DeploymentConfigCreate) SetFileDst(s string) *DeploymentConfigCreate {
	dcc.mutation.SetFileDst(s)
	return dcc
}

// SetNillableFileDst sets the "fileDst" field if the given value is not nil.
func (dcc *DeploymentConfigCreate) SetNillableFileDst(s *string) *DeploymentConfigCreate {
	if s != nil {
		dcc.SetFileDst(*s)
	}
	return dcc
}

// AddDeploymentIDs adds the "deployments" edge to the Deployment entity by IDs.
func (dcc *DeploymentConfigCreate) AddDeploymentIDs(ids ...int) *DeploymentConfigCreate {
	dcc.mutation.AddDeploymentIDs(ids...)
	return dcc
}

// AddDeployments adds the "deployments" edges to the Deployment entity.
func (dcc *DeploymentConfigCreate) AddDeployments(d ...*Deployment) *DeploymentConfigCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dcc.AddDeploymentIDs(ids...)
}

// SetFileID sets the "file" edge to the File entity by ID.
func (dcc *DeploymentConfigCreate) SetFileID(id int) *DeploymentConfigCreate {
	dcc.mutation.SetFileID(id)
	return dcc
}

// SetNillableFileID sets the "file" edge to the File entity by ID if the given value is not nil.
func (dcc *DeploymentConfigCreate) SetNillableFileID(id *int) *DeploymentConfigCreate {
	if id != nil {
		dcc = dcc.SetFileID(*id)
	}
	return dcc
}

// SetFile sets the "file" edge to the File entity.
func (dcc *DeploymentConfigCreate) SetFile(f *File) *DeploymentConfigCreate {
	return dcc.SetFileID(f.ID)
}

// SetImplantConfigID sets the "implantConfig" edge to the ImplantConfig entity by ID.
func (dcc *DeploymentConfigCreate) SetImplantConfigID(id int) *DeploymentConfigCreate {
	dcc.mutation.SetImplantConfigID(id)
	return dcc
}

// SetNillableImplantConfigID sets the "implantConfig" edge to the ImplantConfig entity by ID if the given value is not nil.
func (dcc *DeploymentConfigCreate) SetNillableImplantConfigID(id *int) *DeploymentConfigCreate {
	if id != nil {
		dcc = dcc.SetImplantConfigID(*id)
	}
	return dcc
}

// SetImplantConfig sets the "implantConfig" edge to the ImplantConfig entity.
func (dcc *DeploymentConfigCreate) SetImplantConfig(i *ImplantConfig) *DeploymentConfigCreate {
	return dcc.SetImplantConfigID(i.ID)
}

// Mutation returns the DeploymentConfigMutation object of the builder.
func (dcc *DeploymentConfigCreate) Mutation() *DeploymentConfigMutation {
	return dcc.mutation
}

// Save creates the DeploymentConfig in the database.
func (dcc *DeploymentConfigCreate) Save(ctx context.Context) (*DeploymentConfig, error) {
	var (
		err  error
		node *DeploymentConfig
	)
	dcc.defaults()
	if len(dcc.hooks) == 0 {
		if err = dcc.check(); err != nil {
			return nil, err
		}
		node, err = dcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeploymentConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dcc.check(); err != nil {
				return nil, err
			}
			dcc.mutation = mutation
			if node, err = dcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dcc.hooks) - 1; i >= 0; i-- {
			if dcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dcc *DeploymentConfigCreate) SaveX(ctx context.Context) *DeploymentConfig {
	v, err := dcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcc *DeploymentConfigCreate) Exec(ctx context.Context) error {
	_, err := dcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcc *DeploymentConfigCreate) ExecX(ctx context.Context) {
	if err := dcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dcc *DeploymentConfigCreate) defaults() {
	if _, ok := dcc.mutation.Cmd(); !ok {
		v := deploymentconfig.DefaultCmd
		dcc.mutation.SetCmd(v)
	}
	if _, ok := dcc.mutation.StartCmd(); !ok {
		v := deploymentconfig.DefaultStartCmd
		dcc.mutation.SetStartCmd(v)
	}
	if _, ok := dcc.mutation.FileDst(); !ok {
		v := deploymentconfig.DefaultFileDst
		dcc.mutation.SetFileDst(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dcc *DeploymentConfigCreate) check() error {
	if _, ok := dcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "DeploymentConfig.name"`)}
	}
	if v, ok := dcc.mutation.Name(); ok {
		if err := deploymentconfig.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "DeploymentConfig.name": %w`, err)}
		}
	}
	if _, ok := dcc.mutation.Cmd(); !ok {
		return &ValidationError{Name: "cmd", err: errors.New(`ent: missing required field "DeploymentConfig.cmd"`)}
	}
	if _, ok := dcc.mutation.StartCmd(); !ok {
		return &ValidationError{Name: "startCmd", err: errors.New(`ent: missing required field "DeploymentConfig.startCmd"`)}
	}
	if _, ok := dcc.mutation.FileDst(); !ok {
		return &ValidationError{Name: "fileDst", err: errors.New(`ent: missing required field "DeploymentConfig.fileDst"`)}
	}
	return nil
}

func (dcc *DeploymentConfigCreate) sqlSave(ctx context.Context) (*DeploymentConfig, error) {
	_node, _spec := dcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dcc *DeploymentConfigCreate) createSpec() (*DeploymentConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &DeploymentConfig{config: dcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deploymentconfig.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deploymentconfig.FieldID,
			},
		}
	)
	if value, ok := dcc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentconfig.FieldName,
		})
		_node.Name = value
	}
	if value, ok := dcc.mutation.Cmd(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentconfig.FieldCmd,
		})
		_node.Cmd = value
	}
	if value, ok := dcc.mutation.StartCmd(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: deploymentconfig.FieldStartCmd,
		})
		_node.StartCmd = value
	}
	if value, ok := dcc.mutation.FileDst(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deploymentconfig.FieldFileDst,
		})
		_node.FileDst = value
	}
	if nodes := dcc.mutation.DeploymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   deploymentconfig.DeploymentsTable,
			Columns: []string{deploymentconfig.DeploymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dcc.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deploymentconfig.FileTable,
			Columns: []string{deploymentconfig.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.deployment_config_file = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dcc.mutation.ImplantConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deploymentconfig.ImplantConfigTable,
			Columns: []string{deploymentconfig.ImplantConfigColumn},
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
		_node.deployment_config_implant_config = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeploymentConfigCreateBulk is the builder for creating many DeploymentConfig entities in bulk.
type DeploymentConfigCreateBulk struct {
	config
	builders []*DeploymentConfigCreate
}

// Save creates the DeploymentConfig entities in the database.
func (dccb *DeploymentConfigCreateBulk) Save(ctx context.Context) ([]*DeploymentConfig, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dccb.builders))
	nodes := make([]*DeploymentConfig, len(dccb.builders))
	mutators := make([]Mutator, len(dccb.builders))
	for i := range dccb.builders {
		func(i int, root context.Context) {
			builder := dccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeploymentConfigMutation)
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
					_, err = mutators[i+1].Mutate(root, dccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dccb *DeploymentConfigCreateBulk) SaveX(ctx context.Context) []*DeploymentConfig {
	v, err := dccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dccb *DeploymentConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := dccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dccb *DeploymentConfigCreateBulk) ExecX(ctx context.Context) {
	if err := dccb.Exec(ctx); err != nil {
		panic(err)
	}
}
