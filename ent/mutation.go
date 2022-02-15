// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/kcarretto/realm/ent/credential"
	"github.com/kcarretto/realm/ent/predicate"
	"github.com/kcarretto/realm/ent/target"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCredential = "Credential"
	TypeTarget     = "Target"
)

// CredentialMutation represents an operation that mutates the Credential nodes in the graph.
type CredentialMutation struct {
	config
	op            Op
	typ           string
	id            *int
	principal     *string
	secret        *string
	kind          *credential.Kind
	clearedFields map[string]struct{}
	target        *int
	clearedtarget bool
	done          bool
	oldValue      func(context.Context) (*Credential, error)
	predicates    []predicate.Credential
}

var _ ent.Mutation = (*CredentialMutation)(nil)

// credentialOption allows management of the mutation configuration using functional options.
type credentialOption func(*CredentialMutation)

// newCredentialMutation creates new mutation for the Credential entity.
func newCredentialMutation(c config, op Op, opts ...credentialOption) *CredentialMutation {
	m := &CredentialMutation{
		config:        c,
		op:            op,
		typ:           TypeCredential,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCredentialID sets the ID field of the mutation.
func withCredentialID(id int) credentialOption {
	return func(m *CredentialMutation) {
		var (
			err   error
			once  sync.Once
			value *Credential
		)
		m.oldValue = func(ctx context.Context) (*Credential, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Credential.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCredential sets the old Credential of the mutation.
func withCredential(node *Credential) credentialOption {
	return func(m *CredentialMutation) {
		m.oldValue = func(context.Context) (*Credential, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CredentialMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CredentialMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CredentialMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CredentialMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Credential.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetPrincipal sets the "principal" field.
func (m *CredentialMutation) SetPrincipal(s string) {
	m.principal = &s
}

// Principal returns the value of the "principal" field in the mutation.
func (m *CredentialMutation) Principal() (r string, exists bool) {
	v := m.principal
	if v == nil {
		return
	}
	return *v, true
}

// OldPrincipal returns the old "principal" field's value of the Credential entity.
// If the Credential object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CredentialMutation) OldPrincipal(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPrincipal is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPrincipal requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPrincipal: %w", err)
	}
	return oldValue.Principal, nil
}

// ResetPrincipal resets all changes to the "principal" field.
func (m *CredentialMutation) ResetPrincipal() {
	m.principal = nil
}

// SetSecret sets the "secret" field.
func (m *CredentialMutation) SetSecret(s string) {
	m.secret = &s
}

// Secret returns the value of the "secret" field in the mutation.
func (m *CredentialMutation) Secret() (r string, exists bool) {
	v := m.secret
	if v == nil {
		return
	}
	return *v, true
}

// OldSecret returns the old "secret" field's value of the Credential entity.
// If the Credential object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CredentialMutation) OldSecret(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSecret is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSecret requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSecret: %w", err)
	}
	return oldValue.Secret, nil
}

// ResetSecret resets all changes to the "secret" field.
func (m *CredentialMutation) ResetSecret() {
	m.secret = nil
}

// SetKind sets the "kind" field.
func (m *CredentialMutation) SetKind(c credential.Kind) {
	m.kind = &c
}

// Kind returns the value of the "kind" field in the mutation.
func (m *CredentialMutation) Kind() (r credential.Kind, exists bool) {
	v := m.kind
	if v == nil {
		return
	}
	return *v, true
}

// OldKind returns the old "kind" field's value of the Credential entity.
// If the Credential object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CredentialMutation) OldKind(ctx context.Context) (v credential.Kind, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldKind is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldKind requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldKind: %w", err)
	}
	return oldValue.Kind, nil
}

// ResetKind resets all changes to the "kind" field.
func (m *CredentialMutation) ResetKind() {
	m.kind = nil
}

// SetTargetID sets the "target" edge to the Target entity by id.
func (m *CredentialMutation) SetTargetID(id int) {
	m.target = &id
}

// ClearTarget clears the "target" edge to the Target entity.
func (m *CredentialMutation) ClearTarget() {
	m.clearedtarget = true
}

// TargetCleared reports if the "target" edge to the Target entity was cleared.
func (m *CredentialMutation) TargetCleared() bool {
	return m.clearedtarget
}

// TargetID returns the "target" edge ID in the mutation.
func (m *CredentialMutation) TargetID() (id int, exists bool) {
	if m.target != nil {
		return *m.target, true
	}
	return
}

// TargetIDs returns the "target" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// TargetID instead. It exists only for internal usage by the builders.
func (m *CredentialMutation) TargetIDs() (ids []int) {
	if id := m.target; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetTarget resets all changes to the "target" edge.
func (m *CredentialMutation) ResetTarget() {
	m.target = nil
	m.clearedtarget = false
}

// Where appends a list predicates to the CredentialMutation builder.
func (m *CredentialMutation) Where(ps ...predicate.Credential) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CredentialMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Credential).
func (m *CredentialMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CredentialMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.principal != nil {
		fields = append(fields, credential.FieldPrincipal)
	}
	if m.secret != nil {
		fields = append(fields, credential.FieldSecret)
	}
	if m.kind != nil {
		fields = append(fields, credential.FieldKind)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CredentialMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case credential.FieldPrincipal:
		return m.Principal()
	case credential.FieldSecret:
		return m.Secret()
	case credential.FieldKind:
		return m.Kind()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CredentialMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case credential.FieldPrincipal:
		return m.OldPrincipal(ctx)
	case credential.FieldSecret:
		return m.OldSecret(ctx)
	case credential.FieldKind:
		return m.OldKind(ctx)
	}
	return nil, fmt.Errorf("unknown Credential field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CredentialMutation) SetField(name string, value ent.Value) error {
	switch name {
	case credential.FieldPrincipal:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPrincipal(v)
		return nil
	case credential.FieldSecret:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSecret(v)
		return nil
	case credential.FieldKind:
		v, ok := value.(credential.Kind)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetKind(v)
		return nil
	}
	return fmt.Errorf("unknown Credential field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CredentialMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CredentialMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CredentialMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Credential numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CredentialMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CredentialMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CredentialMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Credential nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CredentialMutation) ResetField(name string) error {
	switch name {
	case credential.FieldPrincipal:
		m.ResetPrincipal()
		return nil
	case credential.FieldSecret:
		m.ResetSecret()
		return nil
	case credential.FieldKind:
		m.ResetKind()
		return nil
	}
	return fmt.Errorf("unknown Credential field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CredentialMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.target != nil {
		edges = append(edges, credential.EdgeTarget)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CredentialMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case credential.EdgeTarget:
		if id := m.target; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CredentialMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CredentialMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CredentialMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtarget {
		edges = append(edges, credential.EdgeTarget)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CredentialMutation) EdgeCleared(name string) bool {
	switch name {
	case credential.EdgeTarget:
		return m.clearedtarget
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CredentialMutation) ClearEdge(name string) error {
	switch name {
	case credential.EdgeTarget:
		m.ClearTarget()
		return nil
	}
	return fmt.Errorf("unknown Credential unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CredentialMutation) ResetEdge(name string) error {
	switch name {
	case credential.EdgeTarget:
		m.ResetTarget()
		return nil
	}
	return fmt.Errorf("unknown Credential edge %s", name)
}

// TargetMutation represents an operation that mutates the Target nodes in the graph.
type TargetMutation struct {
	config
	op                 Op
	typ                string
	id                 *int
	name               *string
	forwardConnectIP   *string
	clearedFields      map[string]struct{}
	credentials        map[int]struct{}
	removedcredentials map[int]struct{}
	clearedcredentials bool
	done               bool
	oldValue           func(context.Context) (*Target, error)
	predicates         []predicate.Target
}

var _ ent.Mutation = (*TargetMutation)(nil)

// targetOption allows management of the mutation configuration using functional options.
type targetOption func(*TargetMutation)

// newTargetMutation creates new mutation for the Target entity.
func newTargetMutation(c config, op Op, opts ...targetOption) *TargetMutation {
	m := &TargetMutation{
		config:        c,
		op:            op,
		typ:           TypeTarget,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTargetID sets the ID field of the mutation.
func withTargetID(id int) targetOption {
	return func(m *TargetMutation) {
		var (
			err   error
			once  sync.Once
			value *Target
		)
		m.oldValue = func(ctx context.Context) (*Target, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Target.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTarget sets the old Target of the mutation.
func withTarget(node *Target) targetOption {
	return func(m *TargetMutation) {
		m.oldValue = func(context.Context) (*Target, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TargetMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TargetMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TargetMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TargetMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Target.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *TargetMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TargetMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Target entity.
// If the Target object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TargetMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TargetMutation) ResetName() {
	m.name = nil
}

// SetForwardConnectIP sets the "forwardConnectIP" field.
func (m *TargetMutation) SetForwardConnectIP(s string) {
	m.forwardConnectIP = &s
}

// ForwardConnectIP returns the value of the "forwardConnectIP" field in the mutation.
func (m *TargetMutation) ForwardConnectIP() (r string, exists bool) {
	v := m.forwardConnectIP
	if v == nil {
		return
	}
	return *v, true
}

// OldForwardConnectIP returns the old "forwardConnectIP" field's value of the Target entity.
// If the Target object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TargetMutation) OldForwardConnectIP(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldForwardConnectIP is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldForwardConnectIP requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldForwardConnectIP: %w", err)
	}
	return oldValue.ForwardConnectIP, nil
}

// ResetForwardConnectIP resets all changes to the "forwardConnectIP" field.
func (m *TargetMutation) ResetForwardConnectIP() {
	m.forwardConnectIP = nil
}

// AddCredentialIDs adds the "credentials" edge to the Credential entity by ids.
func (m *TargetMutation) AddCredentialIDs(ids ...int) {
	if m.credentials == nil {
		m.credentials = make(map[int]struct{})
	}
	for i := range ids {
		m.credentials[ids[i]] = struct{}{}
	}
}

// ClearCredentials clears the "credentials" edge to the Credential entity.
func (m *TargetMutation) ClearCredentials() {
	m.clearedcredentials = true
}

// CredentialsCleared reports if the "credentials" edge to the Credential entity was cleared.
func (m *TargetMutation) CredentialsCleared() bool {
	return m.clearedcredentials
}

// RemoveCredentialIDs removes the "credentials" edge to the Credential entity by IDs.
func (m *TargetMutation) RemoveCredentialIDs(ids ...int) {
	if m.removedcredentials == nil {
		m.removedcredentials = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.credentials, ids[i])
		m.removedcredentials[ids[i]] = struct{}{}
	}
}

// RemovedCredentials returns the removed IDs of the "credentials" edge to the Credential entity.
func (m *TargetMutation) RemovedCredentialsIDs() (ids []int) {
	for id := range m.removedcredentials {
		ids = append(ids, id)
	}
	return
}

// CredentialsIDs returns the "credentials" edge IDs in the mutation.
func (m *TargetMutation) CredentialsIDs() (ids []int) {
	for id := range m.credentials {
		ids = append(ids, id)
	}
	return
}

// ResetCredentials resets all changes to the "credentials" edge.
func (m *TargetMutation) ResetCredentials() {
	m.credentials = nil
	m.clearedcredentials = false
	m.removedcredentials = nil
}

// Where appends a list predicates to the TargetMutation builder.
func (m *TargetMutation) Where(ps ...predicate.Target) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *TargetMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Target).
func (m *TargetMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TargetMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, target.FieldName)
	}
	if m.forwardConnectIP != nil {
		fields = append(fields, target.FieldForwardConnectIP)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TargetMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case target.FieldName:
		return m.Name()
	case target.FieldForwardConnectIP:
		return m.ForwardConnectIP()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TargetMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case target.FieldName:
		return m.OldName(ctx)
	case target.FieldForwardConnectIP:
		return m.OldForwardConnectIP(ctx)
	}
	return nil, fmt.Errorf("unknown Target field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TargetMutation) SetField(name string, value ent.Value) error {
	switch name {
	case target.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case target.FieldForwardConnectIP:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetForwardConnectIP(v)
		return nil
	}
	return fmt.Errorf("unknown Target field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TargetMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TargetMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TargetMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Target numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TargetMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TargetMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TargetMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Target nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TargetMutation) ResetField(name string) error {
	switch name {
	case target.FieldName:
		m.ResetName()
		return nil
	case target.FieldForwardConnectIP:
		m.ResetForwardConnectIP()
		return nil
	}
	return fmt.Errorf("unknown Target field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TargetMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.credentials != nil {
		edges = append(edges, target.EdgeCredentials)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TargetMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case target.EdgeCredentials:
		ids := make([]ent.Value, 0, len(m.credentials))
		for id := range m.credentials {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TargetMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedcredentials != nil {
		edges = append(edges, target.EdgeCredentials)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TargetMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case target.EdgeCredentials:
		ids := make([]ent.Value, 0, len(m.removedcredentials))
		for id := range m.removedcredentials {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TargetMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcredentials {
		edges = append(edges, target.EdgeCredentials)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TargetMutation) EdgeCleared(name string) bool {
	switch name {
	case target.EdgeCredentials:
		return m.clearedcredentials
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TargetMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Target unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TargetMutation) ResetEdge(name string) error {
	switch name {
	case target.EdgeCredentials:
		m.ResetCredentials()
		return nil
	}
	return fmt.Errorf("unknown Target edge %s", name)
}
