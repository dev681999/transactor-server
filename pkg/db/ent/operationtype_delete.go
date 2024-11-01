// Code generated by ent, DO NOT EDIT.

package ent

import (
	"transactor-server/pkg/db/ent/operationtype"
	"transactor-server/pkg/db/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OperationTypeDelete is the builder for deleting a OperationType entity.
type OperationTypeDelete struct {
	config
	hooks    []Hook
	mutation *OperationTypeMutation
}

// Where appends a list predicates to the OperationTypeDelete builder.
func (otd *OperationTypeDelete) Where(ps ...predicate.OperationType) *OperationTypeDelete {
	otd.mutation.Where(ps...)
	return otd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (otd *OperationTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, otd.sqlExec, otd.mutation, otd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (otd *OperationTypeDelete) ExecX(ctx context.Context) int {
	n, err := otd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (otd *OperationTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(operationtype.Table, sqlgraph.NewFieldSpec(operationtype.FieldID, field.TypeInt))
	if ps := otd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, otd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	otd.mutation.done = true
	return affected, err
}

// OperationTypeDeleteOne is the builder for deleting a single OperationType entity.
type OperationTypeDeleteOne struct {
	otd *OperationTypeDelete
}

// Where appends a list predicates to the OperationTypeDelete builder.
func (otdo *OperationTypeDeleteOne) Where(ps ...predicate.OperationType) *OperationTypeDeleteOne {
	otdo.otd.mutation.Where(ps...)
	return otdo
}

// Exec executes the deletion query.
func (otdo *OperationTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := otdo.otd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{operationtype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (otdo *OperationTypeDeleteOne) ExecX(ctx context.Context) {
	if err := otdo.Exec(ctx); err != nil {
		panic(err)
	}
}