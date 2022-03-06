// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/kcarretto/realm/ent"

	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query a policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutate a  policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The CredentialQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CredentialQueryRuleFunc func(context.Context, *ent.CredentialQuery) error

// EvalQuery return f(ctx, q).
func (f CredentialQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CredentialQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CredentialQuery", q)
}

// The CredentialMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CredentialMutationRuleFunc func(context.Context, *ent.CredentialMutation) error

// EvalMutation calls f(ctx, m).
func (f CredentialMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CredentialMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CredentialMutation", m)
}

// The DeploymentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DeploymentQueryRuleFunc func(context.Context, *ent.DeploymentQuery) error

// EvalQuery return f(ctx, q).
func (f DeploymentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DeploymentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DeploymentQuery", q)
}

// The DeploymentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DeploymentMutationRuleFunc func(context.Context, *ent.DeploymentMutation) error

// EvalMutation calls f(ctx, m).
func (f DeploymentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DeploymentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DeploymentMutation", m)
}

// The DeploymentConfigQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DeploymentConfigQueryRuleFunc func(context.Context, *ent.DeploymentConfigQuery) error

// EvalQuery return f(ctx, q).
func (f DeploymentConfigQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DeploymentConfigQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DeploymentConfigQuery", q)
}

// The DeploymentConfigMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DeploymentConfigMutationRuleFunc func(context.Context, *ent.DeploymentConfigMutation) error

// EvalMutation calls f(ctx, m).
func (f DeploymentConfigMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DeploymentConfigMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DeploymentConfigMutation", m)
}

// The FileQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FileQueryRuleFunc func(context.Context, *ent.FileQuery) error

// EvalQuery return f(ctx, q).
func (f FileQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FileQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.FileQuery", q)
}

// The FileMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FileMutationRuleFunc func(context.Context, *ent.FileMutation) error

// EvalMutation calls f(ctx, m).
func (f FileMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.FileMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.FileMutation", m)
}

// The ImplantQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ImplantQueryRuleFunc func(context.Context, *ent.ImplantQuery) error

// EvalQuery return f(ctx, q).
func (f ImplantQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ImplantQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ImplantQuery", q)
}

// The ImplantMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ImplantMutationRuleFunc func(context.Context, *ent.ImplantMutation) error

// EvalMutation calls f(ctx, m).
func (f ImplantMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ImplantMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ImplantMutation", m)
}

// The ImplantCallbackConfigQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ImplantCallbackConfigQueryRuleFunc func(context.Context, *ent.ImplantCallbackConfigQuery) error

// EvalQuery return f(ctx, q).
func (f ImplantCallbackConfigQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ImplantCallbackConfigQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ImplantCallbackConfigQuery", q)
}

// The ImplantCallbackConfigMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ImplantCallbackConfigMutationRuleFunc func(context.Context, *ent.ImplantCallbackConfigMutation) error

// EvalMutation calls f(ctx, m).
func (f ImplantCallbackConfigMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ImplantCallbackConfigMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ImplantCallbackConfigMutation", m)
}

// The ImplantConfigQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ImplantConfigQueryRuleFunc func(context.Context, *ent.ImplantConfigQuery) error

// EvalQuery return f(ctx, q).
func (f ImplantConfigQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ImplantConfigQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ImplantConfigQuery", q)
}

// The ImplantConfigMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ImplantConfigMutationRuleFunc func(context.Context, *ent.ImplantConfigMutation) error

// EvalMutation calls f(ctx, m).
func (f ImplantConfigMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ImplantConfigMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ImplantConfigMutation", m)
}

// The ImplantServiceConfigQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ImplantServiceConfigQueryRuleFunc func(context.Context, *ent.ImplantServiceConfigQuery) error

// EvalQuery return f(ctx, q).
func (f ImplantServiceConfigQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ImplantServiceConfigQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ImplantServiceConfigQuery", q)
}

// The ImplantServiceConfigMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ImplantServiceConfigMutationRuleFunc func(context.Context, *ent.ImplantServiceConfigMutation) error

// EvalMutation calls f(ctx, m).
func (f ImplantServiceConfigMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ImplantServiceConfigMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ImplantServiceConfigMutation", m)
}

// The TargetQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TargetQueryRuleFunc func(context.Context, *ent.TargetQuery) error

// EvalQuery return f(ctx, q).
func (f TargetQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TargetQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TargetQuery", q)
}

// The TargetMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TargetMutationRuleFunc func(context.Context, *ent.TargetMutation) error

// EvalMutation calls f(ctx, m).
func (f TargetMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TargetMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TargetMutation", m)
}
