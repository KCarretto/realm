// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CredentialQuery) CollectFields(ctx context.Context, satisfies ...string) *CredentialQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		c = c.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return c
}

func (c *CredentialQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *CredentialQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "target":
			c = c.WithTarget(func(query *TargetQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return c
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (d *DeploymentQuery) CollectFields(ctx context.Context, satisfies ...string) *DeploymentQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		d = d.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return d
}

func (d *DeploymentQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *DeploymentQuery {
	return d
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (dc *DeploymentConfigQuery) CollectFields(ctx context.Context, satisfies ...string) *DeploymentConfigQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		dc = dc.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return dc
}

func (dc *DeploymentConfigQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *DeploymentConfigQuery {
	return dc
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (f *FileQuery) CollectFields(ctx context.Context, satisfies ...string) *FileQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		f = f.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return f
}

func (f *FileQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *FileQuery {
	return f
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (i *ImplantQuery) CollectFields(ctx context.Context, satisfies ...string) *ImplantQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		i = i.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return i
}

func (i *ImplantQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ImplantQuery {
	return i
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (icc *ImplantCallbackConfigQuery) CollectFields(ctx context.Context, satisfies ...string) *ImplantCallbackConfigQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		icc = icc.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return icc
}

func (icc *ImplantCallbackConfigQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ImplantCallbackConfigQuery {
	return icc
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ic *ImplantConfigQuery) CollectFields(ctx context.Context, satisfies ...string) *ImplantConfigQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		ic = ic.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return ic
}

func (ic *ImplantConfigQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ImplantConfigQuery {
	return ic
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (isc *ImplantServiceConfigQuery) CollectFields(ctx context.Context, satisfies ...string) *ImplantServiceConfigQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		isc = isc.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return isc
}

func (isc *ImplantServiceConfigQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ImplantServiceConfigQuery {
	return isc
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TagQuery) CollectFields(ctx context.Context, satisfies ...string) *TagQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return t
}

func (t *TagQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *TagQuery {
	return t
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TargetQuery) CollectFields(ctx context.Context, satisfies ...string) *TargetQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return t
}

func (t *TargetQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *TargetQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "credentials":
			t = t.WithCredentials(func(query *CredentialQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return t
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	return u
}
