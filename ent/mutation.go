// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"main/ent/bizcard"
	"main/ent/predicate"
	"sync"

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
	TypeBizCard = "BizCard"
)

// BizCardMutation represents an operation that mutates the BizCard nodes in the graph.
type BizCardMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	phone_number  *string
	email         *string
	age           *int
	addage        *int
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*BizCard, error)
	predicates    []predicate.BizCard
}

var _ ent.Mutation = (*BizCardMutation)(nil)

// bizcardOption allows management of the mutation configuration using functional options.
type bizcardOption func(*BizCardMutation)

// newBizCardMutation creates new mutation for the BizCard entity.
func newBizCardMutation(c config, op Op, opts ...bizcardOption) *BizCardMutation {
	m := &BizCardMutation{
		config:        c,
		op:            op,
		typ:           TypeBizCard,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withBizCardID sets the ID field of the mutation.
func withBizCardID(id int) bizcardOption {
	return func(m *BizCardMutation) {
		var (
			err   error
			once  sync.Once
			value *BizCard
		)
		m.oldValue = func(ctx context.Context) (*BizCard, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().BizCard.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withBizCard sets the old BizCard of the mutation.
func withBizCard(node *BizCard) bizcardOption {
	return func(m *BizCardMutation) {
		m.oldValue = func(context.Context) (*BizCard, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m BizCardMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m BizCardMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *BizCardMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *BizCardMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().BizCard.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *BizCardMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *BizCardMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the BizCard entity.
// If the BizCard object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BizCardMutation) OldName(ctx context.Context) (v string, err error) {
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
func (m *BizCardMutation) ResetName() {
	m.name = nil
}

// SetPhoneNumber sets the "phone_number" field.
func (m *BizCardMutation) SetPhoneNumber(s string) {
	m.phone_number = &s
}

// PhoneNumber returns the value of the "phone_number" field in the mutation.
func (m *BizCardMutation) PhoneNumber() (r string, exists bool) {
	v := m.phone_number
	if v == nil {
		return
	}
	return *v, true
}

// OldPhoneNumber returns the old "phone_number" field's value of the BizCard entity.
// If the BizCard object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BizCardMutation) OldPhoneNumber(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPhoneNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPhoneNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhoneNumber: %w", err)
	}
	return oldValue.PhoneNumber, nil
}

// ResetPhoneNumber resets all changes to the "phone_number" field.
func (m *BizCardMutation) ResetPhoneNumber() {
	m.phone_number = nil
}

// SetEmail sets the "email" field.
func (m *BizCardMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *BizCardMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the BizCard entity.
// If the BizCard object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BizCardMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *BizCardMutation) ResetEmail() {
	m.email = nil
}

// SetAge sets the "age" field.
func (m *BizCardMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *BizCardMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the BizCard entity.
// If the BizCard object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BizCardMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *BizCardMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *BizCardMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *BizCardMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// Where appends a list predicates to the BizCardMutation builder.
func (m *BizCardMutation) Where(ps ...predicate.BizCard) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *BizCardMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (BizCard).
func (m *BizCardMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *BizCardMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, bizcard.FieldName)
	}
	if m.phone_number != nil {
		fields = append(fields, bizcard.FieldPhoneNumber)
	}
	if m.email != nil {
		fields = append(fields, bizcard.FieldEmail)
	}
	if m.age != nil {
		fields = append(fields, bizcard.FieldAge)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *BizCardMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case bizcard.FieldName:
		return m.Name()
	case bizcard.FieldPhoneNumber:
		return m.PhoneNumber()
	case bizcard.FieldEmail:
		return m.Email()
	case bizcard.FieldAge:
		return m.Age()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *BizCardMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case bizcard.FieldName:
		return m.OldName(ctx)
	case bizcard.FieldPhoneNumber:
		return m.OldPhoneNumber(ctx)
	case bizcard.FieldEmail:
		return m.OldEmail(ctx)
	case bizcard.FieldAge:
		return m.OldAge(ctx)
	}
	return nil, fmt.Errorf("unknown BizCard field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BizCardMutation) SetField(name string, value ent.Value) error {
	switch name {
	case bizcard.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case bizcard.FieldPhoneNumber:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhoneNumber(v)
		return nil
	case bizcard.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case bizcard.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	}
	return fmt.Errorf("unknown BizCard field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *BizCardMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, bizcard.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *BizCardMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case bizcard.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BizCardMutation) AddField(name string, value ent.Value) error {
	switch name {
	case bizcard.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown BizCard numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *BizCardMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *BizCardMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *BizCardMutation) ClearField(name string) error {
	return fmt.Errorf("unknown BizCard nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *BizCardMutation) ResetField(name string) error {
	switch name {
	case bizcard.FieldName:
		m.ResetName()
		return nil
	case bizcard.FieldPhoneNumber:
		m.ResetPhoneNumber()
		return nil
	case bizcard.FieldEmail:
		m.ResetEmail()
		return nil
	case bizcard.FieldAge:
		m.ResetAge()
		return nil
	}
	return fmt.Errorf("unknown BizCard field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *BizCardMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *BizCardMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *BizCardMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *BizCardMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *BizCardMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *BizCardMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *BizCardMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown BizCard unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *BizCardMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown BizCard edge %s", name)
}