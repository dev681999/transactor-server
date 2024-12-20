// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"transactor-server/pkg/db/ent/predicate"
	"transactor-server/pkg/db/ent/transaction"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TransactionUpdate is the builder for updating Transaction entities.
type TransactionUpdate struct {
	config
	hooks     []Hook
	mutation  *TransactionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tu *TransactionUpdate) Where(ps ...predicate.Transaction) *TransactionUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdateTime sets the "update_time" field.
func (tu *TransactionUpdate) SetUpdateTime(t time.Time) *TransactionUpdate {
	tu.mutation.SetUpdateTime(t)
	return tu
}

// SetBalance sets the "balance" field.
func (tu *TransactionUpdate) SetBalance(f float64) *TransactionUpdate {
	tu.mutation.ResetBalance()
	tu.mutation.SetBalance(f)
	return tu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableBalance(f *float64) *TransactionUpdate {
	if f != nil {
		tu.SetBalance(*f)
	}
	return tu
}

// AddBalance adds f to the "balance" field.
func (tu *TransactionUpdate) AddBalance(f float64) *TransactionUpdate {
	tu.mutation.AddBalance(f)
	return tu
}

// Mutation returns the TransactionMutation object of the builder.
func (tu *TransactionUpdate) Mutation() *TransactionMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TransactionUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TransactionUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TransactionUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TransactionUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TransactionUpdate) defaults() {
	if _, ok := tu.mutation.UpdateTime(); !ok {
		v := transaction.UpdateDefaultUpdateTime()
		tu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TransactionUpdate) check() error {
	if tu.mutation.AccountCleared() && len(tu.mutation.AccountIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Transaction.account"`)
	}
	if tu.mutation.OperationTypeCleared() && len(tu.mutation.OperationTypeIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Transaction.operation_type"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TransactionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TransactionUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TransactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(transaction.Table, transaction.Columns, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdateTime(); ok {
		_spec.SetField(transaction.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Balance(); ok {
		_spec.SetField(transaction.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedBalance(); ok {
		_spec.AddField(transaction.FieldBalance, field.TypeFloat64, value)
	}
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TransactionUpdateOne is the builder for updating a single Transaction entity.
type TransactionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TransactionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (tuo *TransactionUpdateOne) SetUpdateTime(t time.Time) *TransactionUpdateOne {
	tuo.mutation.SetUpdateTime(t)
	return tuo
}

// SetBalance sets the "balance" field.
func (tuo *TransactionUpdateOne) SetBalance(f float64) *TransactionUpdateOne {
	tuo.mutation.ResetBalance()
	tuo.mutation.SetBalance(f)
	return tuo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableBalance(f *float64) *TransactionUpdateOne {
	if f != nil {
		tuo.SetBalance(*f)
	}
	return tuo
}

// AddBalance adds f to the "balance" field.
func (tuo *TransactionUpdateOne) AddBalance(f float64) *TransactionUpdateOne {
	tuo.mutation.AddBalance(f)
	return tuo
}

// Mutation returns the TransactionMutation object of the builder.
func (tuo *TransactionUpdateOne) Mutation() *TransactionMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tuo *TransactionUpdateOne) Where(ps ...predicate.Transaction) *TransactionUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TransactionUpdateOne) Select(field string, fields ...string) *TransactionUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Transaction entity.
func (tuo *TransactionUpdateOne) Save(ctx context.Context) (*Transaction, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TransactionUpdateOne) SaveX(ctx context.Context) *Transaction {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TransactionUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TransactionUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TransactionUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdateTime(); !ok {
		v := transaction.UpdateDefaultUpdateTime()
		tuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TransactionUpdateOne) check() error {
	if tuo.mutation.AccountCleared() && len(tuo.mutation.AccountIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Transaction.account"`)
	}
	if tuo.mutation.OperationTypeCleared() && len(tuo.mutation.OperationTypeIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Transaction.operation_type"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TransactionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TransactionUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TransactionUpdateOne) sqlSave(ctx context.Context) (_node *Transaction, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(transaction.Table, transaction.Columns, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Transaction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transaction.FieldID)
		for _, f := range fields {
			if !transaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdateTime(); ok {
		_spec.SetField(transaction.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Balance(); ok {
		_spec.SetField(transaction.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedBalance(); ok {
		_spec.AddField(transaction.FieldBalance, field.TypeFloat64, value)
	}
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Transaction{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
