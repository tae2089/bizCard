// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"main/ent/bizcard"
	"main/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BizCardDelete is the builder for deleting a BizCard entity.
type BizCardDelete struct {
	config
	hooks    []Hook
	mutation *BizCardMutation
}

// Where appends a list predicates to the BizCardDelete builder.
func (bcd *BizCardDelete) Where(ps ...predicate.BizCard) *BizCardDelete {
	bcd.mutation.Where(ps...)
	return bcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bcd *BizCardDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bcd.hooks) == 0 {
		affected, err = bcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BizCardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bcd.mutation = mutation
			affected, err = bcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bcd.hooks) - 1; i >= 0; i-- {
			if bcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcd *BizCardDelete) ExecX(ctx context.Context) int {
	n, err := bcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bcd *BizCardDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: bizcard.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bizcard.FieldID,
			},
		},
	}
	if ps := bcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bcd.driver, _spec)
}

// BizCardDeleteOne is the builder for deleting a single BizCard entity.
type BizCardDeleteOne struct {
	bcd *BizCardDelete
}

// Exec executes the deletion query.
func (bcdo *BizCardDeleteOne) Exec(ctx context.Context) error {
	n, err := bcdo.bcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bizcard.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdo *BizCardDeleteOne) ExecX(ctx context.Context) {
	bcdo.bcd.ExecX(ctx)
}