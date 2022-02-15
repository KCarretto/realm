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
