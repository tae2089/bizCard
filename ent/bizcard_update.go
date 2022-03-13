// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bizCard/ent/bizcard"
	"bizCard/ent/predicate"
	"bizCard/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BizCardUpdate is the builder for updating BizCard entities.
type BizCardUpdate struct {
	config
	hooks    []Hook
	mutation *BizCardMutation
}

// Where appends a list predicates to the BizCardUpdate builder.
func (bcu *BizCardUpdate) Where(ps ...predicate.BizCard) *BizCardUpdate {
	bcu.mutation.Where(ps...)
	return bcu
}

// SetName sets the "name" field.
func (bcu *BizCardUpdate) SetName(s string) *BizCardUpdate {
	bcu.mutation.SetName(s)
	return bcu
}

// SetPhoneNumber sets the "phone_number" field.
func (bcu *BizCardUpdate) SetPhoneNumber(s string) *BizCardUpdate {
	bcu.mutation.SetPhoneNumber(s)
	return bcu
}

// SetEmail sets the "email" field.
func (bcu *BizCardUpdate) SetEmail(s string) *BizCardUpdate {
	bcu.mutation.SetEmail(s)
	return bcu
}

// SetAge sets the "age" field.
func (bcu *BizCardUpdate) SetAge(i int) *BizCardUpdate {
	bcu.mutation.ResetAge()
	bcu.mutation.SetAge(i)
	return bcu
}

// AddAge adds i to the "age" field.
func (bcu *BizCardUpdate) AddAge(i int) *BizCardUpdate {
	bcu.mutation.AddAge(i)
	return bcu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (bcu *BizCardUpdate) SetOwnerID(id int) *BizCardUpdate {
	bcu.mutation.SetOwnerID(id)
	return bcu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (bcu *BizCardUpdate) SetNillableOwnerID(id *int) *BizCardUpdate {
	if id != nil {
		bcu = bcu.SetOwnerID(*id)
	}
	return bcu
}

// SetOwner sets the "owner" edge to the User entity.
func (bcu *BizCardUpdate) SetOwner(u *User) *BizCardUpdate {
	return bcu.SetOwnerID(u.ID)
}

// Mutation returns the BizCardMutation object of the builder.
func (bcu *BizCardUpdate) Mutation() *BizCardMutation {
	return bcu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (bcu *BizCardUpdate) ClearOwner() *BizCardUpdate {
	bcu.mutation.ClearOwner()
	return bcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bcu *BizCardUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bcu.hooks) == 0 {
		if err = bcu.check(); err != nil {
			return 0, err
		}
		affected, err = bcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BizCardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bcu.check(); err != nil {
				return 0, err
			}
			bcu.mutation = mutation
			affected, err = bcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bcu.hooks) - 1; i >= 0; i-- {
			if bcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bcu *BizCardUpdate) SaveX(ctx context.Context) int {
	affected, err := bcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bcu *BizCardUpdate) Exec(ctx context.Context) error {
	_, err := bcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcu *BizCardUpdate) ExecX(ctx context.Context) {
	if err := bcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcu *BizCardUpdate) check() error {
	if v, ok := bcu.mutation.Age(); ok {
		if err := bizcard.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "BizCard.age": %w`, err)}
		}
	}
	return nil
}

func (bcu *BizCardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bizcard.Table,
			Columns: bizcard.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bizcard.FieldID,
			},
		},
	}
	if ps := bcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bizcard.FieldName,
		})
	}
	if value, ok := bcu.mutation.PhoneNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bizcard.FieldPhoneNumber,
		})
	}
	if value, ok := bcu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bizcard.FieldEmail,
		})
	}
	if value, ok := bcu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bizcard.FieldAge,
		})
	}
	if value, ok := bcu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bizcard.FieldAge,
		})
	}
	if bcu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bizcard.OwnerTable,
			Columns: []string{bizcard.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bizcard.OwnerTable,
			Columns: []string{bizcard.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bizcard.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BizCardUpdateOne is the builder for updating a single BizCard entity.
type BizCardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BizCardMutation
}

// SetName sets the "name" field.
func (bcuo *BizCardUpdateOne) SetName(s string) *BizCardUpdateOne {
	bcuo.mutation.SetName(s)
	return bcuo
}

// SetPhoneNumber sets the "phone_number" field.
func (bcuo *BizCardUpdateOne) SetPhoneNumber(s string) *BizCardUpdateOne {
	bcuo.mutation.SetPhoneNumber(s)
	return bcuo
}

// SetEmail sets the "email" field.
func (bcuo *BizCardUpdateOne) SetEmail(s string) *BizCardUpdateOne {
	bcuo.mutation.SetEmail(s)
	return bcuo
}

// SetAge sets the "age" field.
func (bcuo *BizCardUpdateOne) SetAge(i int) *BizCardUpdateOne {
	bcuo.mutation.ResetAge()
	bcuo.mutation.SetAge(i)
	return bcuo
}

// AddAge adds i to the "age" field.
func (bcuo *BizCardUpdateOne) AddAge(i int) *BizCardUpdateOne {
	bcuo.mutation.AddAge(i)
	return bcuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (bcuo *BizCardUpdateOne) SetOwnerID(id int) *BizCardUpdateOne {
	bcuo.mutation.SetOwnerID(id)
	return bcuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (bcuo *BizCardUpdateOne) SetNillableOwnerID(id *int) *BizCardUpdateOne {
	if id != nil {
		bcuo = bcuo.SetOwnerID(*id)
	}
	return bcuo
}

// SetOwner sets the "owner" edge to the User entity.
func (bcuo *BizCardUpdateOne) SetOwner(u *User) *BizCardUpdateOne {
	return bcuo.SetOwnerID(u.ID)
}

// Mutation returns the BizCardMutation object of the builder.
func (bcuo *BizCardUpdateOne) Mutation() *BizCardMutation {
	return bcuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (bcuo *BizCardUpdateOne) ClearOwner() *BizCardUpdateOne {
	bcuo.mutation.ClearOwner()
	return bcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bcuo *BizCardUpdateOne) Select(field string, fields ...string) *BizCardUpdateOne {
	bcuo.fields = append([]string{field}, fields...)
	return bcuo
}

// Save executes the query and returns the updated BizCard entity.
func (bcuo *BizCardUpdateOne) Save(ctx context.Context) (*BizCard, error) {
	var (
		err  error
		node *BizCard
	)
	if len(bcuo.hooks) == 0 {
		if err = bcuo.check(); err != nil {
			return nil, err
		}
		node, err = bcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BizCardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bcuo.check(); err != nil {
				return nil, err
			}
			bcuo.mutation = mutation
			node, err = bcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bcuo.hooks) - 1; i >= 0; i-- {
			if bcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bcuo *BizCardUpdateOne) SaveX(ctx context.Context) *BizCard {
	node, err := bcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bcuo *BizCardUpdateOne) Exec(ctx context.Context) error {
	_, err := bcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcuo *BizCardUpdateOne) ExecX(ctx context.Context) {
	if err := bcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcuo *BizCardUpdateOne) check() error {
	if v, ok := bcuo.mutation.Age(); ok {
		if err := bizcard.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "BizCard.age": %w`, err)}
		}
	}
	return nil
}

func (bcuo *BizCardUpdateOne) sqlSave(ctx context.Context) (_node *BizCard, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bizcard.Table,
			Columns: bizcard.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bizcard.FieldID,
			},
		},
	}
	id, ok := bcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BizCard.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bizcard.FieldID)
		for _, f := range fields {
			if !bizcard.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bizcard.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bizcard.FieldName,
		})
	}
	if value, ok := bcuo.mutation.PhoneNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bizcard.FieldPhoneNumber,
		})
	}
	if value, ok := bcuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bizcard.FieldEmail,
		})
	}
	if value, ok := bcuo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bizcard.FieldAge,
		})
	}
	if value, ok := bcuo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bizcard.FieldAge,
		})
	}
	if bcuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bizcard.OwnerTable,
			Columns: []string{bizcard.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bizcard.OwnerTable,
			Columns: []string{bizcard.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BizCard{config: bcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bizcard.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
