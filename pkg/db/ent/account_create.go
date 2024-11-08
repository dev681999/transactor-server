// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"transactor-server/pkg/db/ent/account"
	"transactor-server/pkg/db/ent/transaction"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (ac *AccountCreate) SetCreateTime(t time.Time) *AccountCreate {
	ac.mutation.SetCreateTime(t)
	return ac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreateTime(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetCreateTime(*t)
	}
	return ac
}

// SetUpdateTime sets the "update_time" field.
func (ac *AccountCreate) SetUpdateTime(t time.Time) *AccountCreate {
	ac.mutation.SetUpdateTime(t)
	return ac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdateTime(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetUpdateTime(*t)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *AccountCreate) SetName(s string) *AccountCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetDocumentNumber sets the "document_number" field.
func (ac *AccountCreate) SetDocumentNumber(s string) *AccountCreate {
	ac.mutation.SetDocumentNumber(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AccountCreate) SetID(i int) *AccountCreate {
	ac.mutation.SetID(i)
	return ac
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (ac *AccountCreate) AddTransactionIDs(ids ...int) *AccountCreate {
	ac.mutation.AddTransactionIDs(ids...)
	return ac
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (ac *AccountCreate) AddTransactions(t ...*Transaction) *AccountCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTransactionIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() {
	if _, ok := ac.mutation.CreateTime(); !ok {
		v := account.DefaultCreateTime()
		ac.mutation.SetCreateTime(v)
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		v := account.DefaultUpdateTime()
		ac.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Account.create_time"`)}
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Account.update_time"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Account.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := account.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Account.name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.DocumentNumber(); !ok {
		return &ValidationError{Name: "document_number", err: errors.New(`ent: missing required field "Account.document_number"`)}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(account.Table, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreateTime(); ok {
		_spec.SetField(account.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := ac.mutation.UpdateTime(); ok {
		_spec.SetField(account.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(account.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.DocumentNumber(); ok {
		_spec.SetField(account.FieldDocumentNumber, field.TypeString, value)
		_node.DocumentNumber = value
	}
	if nodes := ac.mutation.TransactionsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ac *AccountCreate) OnConflict(opts ...sql.ConflictOption) *AccountUpsertOne {
	ac.conflict = opts
	return &AccountUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AccountCreate) OnConflictColumns(columns ...string) *AccountUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AccountUpsertOne{
		create: ac,
	}
}

type (
	// AccountUpsertOne is the builder for "upsert"-ing
	//  one Account node.
	AccountUpsertOne struct {
		create *AccountCreate
	}

	// AccountUpsert is the "OnConflict" setter.
	AccountUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *AccountUpsert) SetUpdateTime(v time.Time) *AccountUpsert {
	u.Set(account.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *AccountUpsert) UpdateUpdateTime() *AccountUpsert {
	u.SetExcluded(account.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *AccountUpsert) SetName(v string) *AccountUpsert {
	u.Set(account.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccountUpsert) UpdateName() *AccountUpsert {
	u.SetExcluded(account.FieldName)
	return u
}

// SetDocumentNumber sets the "document_number" field.
func (u *AccountUpsert) SetDocumentNumber(v string) *AccountUpsert {
	u.Set(account.FieldDocumentNumber, v)
	return u
}

// UpdateDocumentNumber sets the "document_number" field to the value that was provided on create.
func (u *AccountUpsert) UpdateDocumentNumber() *AccountUpsert {
	u.SetExcluded(account.FieldDocumentNumber)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(account.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AccountUpsertOne) UpdateNewValues() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(account.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(account.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AccountUpsertOne) Ignore() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccountUpsertOne) DoNothing() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccountCreate.OnConflict
// documentation for more info.
func (u *AccountUpsertOne) Update(set func(*AccountUpsert)) *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *AccountUpsertOne) SetUpdateTime(v time.Time) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateUpdateTime() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *AccountUpsertOne) SetName(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateName() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateName()
	})
}

// SetDocumentNumber sets the "document_number" field.
func (u *AccountUpsertOne) SetDocumentNumber(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetDocumentNumber(v)
	})
}

// UpdateDocumentNumber sets the "document_number" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateDocumentNumber() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateDocumentNumber()
	})
}

// Exec executes the query.
func (u *AccountUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccountCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccountUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AccountUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AccountUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	err      error
	builders []*AccountCreate
	conflict []sql.ConflictOption
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (acb *AccountCreateBulk) OnConflict(opts ...sql.ConflictOption) *AccountUpsertBulk {
	acb.conflict = opts
	return &AccountUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AccountCreateBulk) OnConflictColumns(columns ...string) *AccountUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AccountUpsertBulk{
		create: acb,
	}
}

// AccountUpsertBulk is the builder for "upsert"-ing
// a bulk of Account nodes.
type AccountUpsertBulk struct {
	create *AccountCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(account.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AccountUpsertBulk) UpdateNewValues() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(account.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(account.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AccountUpsertBulk) Ignore() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccountUpsertBulk) DoNothing() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccountCreateBulk.OnConflict
// documentation for more info.
func (u *AccountUpsertBulk) Update(set func(*AccountUpsert)) *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *AccountUpsertBulk) SetUpdateTime(v time.Time) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateUpdateTime() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *AccountUpsertBulk) SetName(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateName() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateName()
	})
}

// SetDocumentNumber sets the "document_number" field.
func (u *AccountUpsertBulk) SetDocumentNumber(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetDocumentNumber(v)
	})
}

// UpdateDocumentNumber sets the "document_number" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateDocumentNumber() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateDocumentNumber()
	})
}

// Exec executes the query.
func (u *AccountUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AccountCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccountCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccountUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
