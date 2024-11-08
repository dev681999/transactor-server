// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"transactor-server/pkg/db/ent/account"
	"transactor-server/pkg/db/ent/predicate"
	"transactor-server/pkg/db/ent/transaction"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks     []Hook
	mutation  *AccountMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdateTime sets the "update_time" field.
func (au *AccountUpdate) SetUpdateTime(t time.Time) *AccountUpdate {
	au.mutation.SetUpdateTime(t)
	return au
}

// SetName sets the "name" field.
func (au *AccountUpdate) SetName(s string) *AccountUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AccountUpdate) SetNillableName(s *string) *AccountUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetDocumentNumber sets the "document_number" field.
func (au *AccountUpdate) SetDocumentNumber(s string) *AccountUpdate {
	au.mutation.SetDocumentNumber(s)
	return au
}

// SetNillableDocumentNumber sets the "document_number" field if the given value is not nil.
func (au *AccountUpdate) SetNillableDocumentNumber(s *string) *AccountUpdate {
	if s != nil {
		au.SetDocumentNumber(*s)
	}
	return au
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (au *AccountUpdate) AddTransactionIDs(ids ...int) *AccountUpdate {
	au.mutation.AddTransactionIDs(ids...)
	return au
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (au *AccountUpdate) AddTransactions(t ...*Transaction) *AccountUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTransactionIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearTransactions clears all "transactions" edges to the Transaction entity.
func (au *AccountUpdate) ClearTransactions() *AccountUpdate {
	au.mutation.ClearTransactions()
	return au
}

// RemoveTransactionIDs removes the "transactions" edge to Transaction entities by IDs.
func (au *AccountUpdate) RemoveTransactionIDs(ids ...int) *AccountUpdate {
	au.mutation.RemoveTransactionIDs(ids...)
	return au
}

// RemoveTransactions removes "transactions" edges to Transaction entities.
func (au *AccountUpdate) RemoveTransactions(t ...*Transaction) *AccountUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTransactionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() {
	if _, ok := au.mutation.UpdateTime(); !ok {
		v := account.UpdateDefaultUpdateTime()
		au.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := account.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Account.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *AccountUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AccountUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdateTime(); ok {
		_spec.SetField(account.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(account.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.DocumentNumber(); ok {
		_spec.SetField(account.FieldDocumentNumber, field.TypeString, value)
	}
	if au.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TransactionsTable,
			Columns: []string{account.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTransactionsIDs(); len(nodes) > 0 && !au.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TransactionsTable,
			Columns: []string{account.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TransactionsTable,
			Columns: []string{account.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AccountMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (auo *AccountUpdateOne) SetUpdateTime(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdateTime(t)
	return auo
}

// SetName sets the "name" field.
func (auo *AccountUpdateOne) SetName(s string) *AccountUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableName(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetDocumentNumber sets the "document_number" field.
func (auo *AccountUpdateOne) SetDocumentNumber(s string) *AccountUpdateOne {
	auo.mutation.SetDocumentNumber(s)
	return auo
}

// SetNillableDocumentNumber sets the "document_number" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableDocumentNumber(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetDocumentNumber(*s)
	}
	return auo
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (auo *AccountUpdateOne) AddTransactionIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.AddTransactionIDs(ids...)
	return auo
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (auo *AccountUpdateOne) AddTransactions(t ...*Transaction) *AccountUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTransactionIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearTransactions clears all "transactions" edges to the Transaction entity.
func (auo *AccountUpdateOne) ClearTransactions() *AccountUpdateOne {
	auo.mutation.ClearTransactions()
	return auo
}

// RemoveTransactionIDs removes the "transactions" edge to Transaction entities by IDs.
func (auo *AccountUpdateOne) RemoveTransactionIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.RemoveTransactionIDs(ids...)
	return auo
}

// RemoveTransactions removes "transactions" edges to Transaction entities.
func (auo *AccountUpdateOne) RemoveTransactions(t ...*Transaction) *AccountUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTransactionIDs(ids...)
}

// Where appends a list predicates to the AccountUpdate builder.
func (auo *AccountUpdateOne) Where(ps ...predicate.Account) *AccountUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdateTime(); !ok {
		v := account.UpdateDefaultUpdateTime()
		auo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := account.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Account.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *AccountUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AccountUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdateTime(); ok {
		_spec.SetField(account.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(account.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.DocumentNumber(); ok {
		_spec.SetField(account.FieldDocumentNumber, field.TypeString, value)
	}
	if auo.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TransactionsTable,
			Columns: []string{account.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTransactionsIDs(); len(nodes) > 0 && !auo.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TransactionsTable,
			Columns: []string{account.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TransactionsTable,
			Columns: []string{account.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(auo.modifiers...)
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
