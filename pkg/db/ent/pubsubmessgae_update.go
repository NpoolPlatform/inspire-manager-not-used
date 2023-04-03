// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/pubsubmessgae"
	"github.com/google/uuid"
)

// PubsubMessgaeUpdate is the builder for updating PubsubMessgae entities.
type PubsubMessgaeUpdate struct {
	config
	hooks     []Hook
	mutation  *PubsubMessgaeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PubsubMessgaeUpdate builder.
func (pmu *PubsubMessgaeUpdate) Where(ps ...predicate.PubsubMessgae) *PubsubMessgaeUpdate {
	pmu.mutation.Where(ps...)
	return pmu
}

// SetCreatedAt sets the "created_at" field.
func (pmu *PubsubMessgaeUpdate) SetCreatedAt(u uint32) *PubsubMessgaeUpdate {
	pmu.mutation.ResetCreatedAt()
	pmu.mutation.SetCreatedAt(u)
	return pmu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pmu *PubsubMessgaeUpdate) SetNillableCreatedAt(u *uint32) *PubsubMessgaeUpdate {
	if u != nil {
		pmu.SetCreatedAt(*u)
	}
	return pmu
}

// AddCreatedAt adds u to the "created_at" field.
func (pmu *PubsubMessgaeUpdate) AddCreatedAt(u int32) *PubsubMessgaeUpdate {
	pmu.mutation.AddCreatedAt(u)
	return pmu
}

// SetUpdatedAt sets the "updated_at" field.
func (pmu *PubsubMessgaeUpdate) SetUpdatedAt(u uint32) *PubsubMessgaeUpdate {
	pmu.mutation.ResetUpdatedAt()
	pmu.mutation.SetUpdatedAt(u)
	return pmu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (pmu *PubsubMessgaeUpdate) AddUpdatedAt(u int32) *PubsubMessgaeUpdate {
	pmu.mutation.AddUpdatedAt(u)
	return pmu
}

// SetDeletedAt sets the "deleted_at" field.
func (pmu *PubsubMessgaeUpdate) SetDeletedAt(u uint32) *PubsubMessgaeUpdate {
	pmu.mutation.ResetDeletedAt()
	pmu.mutation.SetDeletedAt(u)
	return pmu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pmu *PubsubMessgaeUpdate) SetNillableDeletedAt(u *uint32) *PubsubMessgaeUpdate {
	if u != nil {
		pmu.SetDeletedAt(*u)
	}
	return pmu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (pmu *PubsubMessgaeUpdate) AddDeletedAt(u int32) *PubsubMessgaeUpdate {
	pmu.mutation.AddDeletedAt(u)
	return pmu
}

// SetUniqueID sets the "unique_id" field.
func (pmu *PubsubMessgaeUpdate) SetUniqueID(u uuid.UUID) *PubsubMessgaeUpdate {
	pmu.mutation.SetUniqueID(u)
	return pmu
}

// SetMessageID sets the "message_id" field.
func (pmu *PubsubMessgaeUpdate) SetMessageID(s string) *PubsubMessgaeUpdate {
	pmu.mutation.SetMessageID(s)
	return pmu
}

// SetSender sets the "sender" field.
func (pmu *PubsubMessgaeUpdate) SetSender(s string) *PubsubMessgaeUpdate {
	pmu.mutation.SetSender(s)
	return pmu
}

// SetBody sets the "body" field.
func (pmu *PubsubMessgaeUpdate) SetBody(b []byte) *PubsubMessgaeUpdate {
	pmu.mutation.SetBody(b)
	return pmu
}

// SetState sets the "state" field.
func (pmu *PubsubMessgaeUpdate) SetState(s string) *PubsubMessgaeUpdate {
	pmu.mutation.SetState(s)
	return pmu
}

// SetResponseID sets the "response_id" field.
func (pmu *PubsubMessgaeUpdate) SetResponseID(u uuid.UUID) *PubsubMessgaeUpdate {
	pmu.mutation.SetResponseID(u)
	return pmu
}

// SetErrorMessage sets the "error_message" field.
func (pmu *PubsubMessgaeUpdate) SetErrorMessage(s string) *PubsubMessgaeUpdate {
	pmu.mutation.SetErrorMessage(s)
	return pmu
}

// SetNillableErrorMessage sets the "error_message" field if the given value is not nil.
func (pmu *PubsubMessgaeUpdate) SetNillableErrorMessage(s *string) *PubsubMessgaeUpdate {
	if s != nil {
		pmu.SetErrorMessage(*s)
	}
	return pmu
}

// ClearErrorMessage clears the value of the "error_message" field.
func (pmu *PubsubMessgaeUpdate) ClearErrorMessage() *PubsubMessgaeUpdate {
	pmu.mutation.ClearErrorMessage()
	return pmu
}

// Mutation returns the PubsubMessgaeMutation object of the builder.
func (pmu *PubsubMessgaeUpdate) Mutation() *PubsubMessgaeMutation {
	return pmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pmu *PubsubMessgaeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := pmu.defaults(); err != nil {
		return 0, err
	}
	if len(pmu.hooks) == 0 {
		affected, err = pmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PubsubMessgaeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pmu.mutation = mutation
			affected, err = pmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pmu.hooks) - 1; i >= 0; i-- {
			if pmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pmu *PubsubMessgaeUpdate) SaveX(ctx context.Context) int {
	affected, err := pmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pmu *PubsubMessgaeUpdate) Exec(ctx context.Context) error {
	_, err := pmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmu *PubsubMessgaeUpdate) ExecX(ctx context.Context) {
	if err := pmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pmu *PubsubMessgaeUpdate) defaults() error {
	if _, ok := pmu.mutation.UpdatedAt(); !ok {
		if pubsubmessgae.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized pubsubmessgae.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := pubsubmessgae.UpdateDefaultUpdatedAt()
		pmu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pmu *PubsubMessgaeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PubsubMessgaeUpdate {
	pmu.modifiers = append(pmu.modifiers, modifiers...)
	return pmu
}

func (pmu *PubsubMessgaeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pubsubmessgae.Table,
			Columns: pubsubmessgae.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pubsubmessgae.FieldID,
			},
		},
	}
	if ps := pmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessgae.FieldCreatedAt,
		})
	}
	if value, ok := pmu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessgae.FieldCreatedAt,
		})
	}
	if value, ok := pmu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessgae.FieldUpdatedAt,
		})
	}
	if value, ok := pmu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessgae.FieldUpdatedAt,
		})
	}
	if value, ok := pmu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessgae.FieldDeletedAt,
		})
	}
	if value, ok := pmu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessgae.FieldDeletedAt,
		})
	}
	if value, ok := pmu.mutation.UniqueID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pubsubmessgae.FieldUniqueID,
		})
	}
	if value, ok := pmu.mutation.MessageID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pubsubmessgae.FieldMessageID,
		})
	}
	if value, ok := pmu.mutation.Sender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pubsubmessgae.FieldSender,
		})
	}
	if value, ok := pmu.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: pubsubmessgae.FieldBody,
		})
	}
	if value, ok := pmu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pubsubmessgae.FieldState,
		})
	}
	if value, ok := pmu.mutation.ResponseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pubsubmessgae.FieldResponseID,
		})
	}
	if value, ok := pmu.mutation.ErrorMessage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pubsubmessgae.FieldErrorMessage,
		})
	}
	if pmu.mutation.ErrorMessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pubsubmessgae.FieldErrorMessage,
		})
	}
	_spec.Modifiers = pmu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, pmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pubsubmessgae.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PubsubMessgaeUpdateOne is the builder for updating a single PubsubMessgae entity.
type PubsubMessgaeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PubsubMessgaeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (pmuo *PubsubMessgaeUpdateOne) SetCreatedAt(u uint32) *PubsubMessgaeUpdateOne {
	pmuo.mutation.ResetCreatedAt()
	pmuo.mutation.SetCreatedAt(u)
	return pmuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pmuo *PubsubMessgaeUpdateOne) SetNillableCreatedAt(u *uint32) *PubsubMessgaeUpdateOne {
	if u != nil {
		pmuo.SetCreatedAt(*u)
	}
	return pmuo
}

// AddCreatedAt adds u to the "created_at" field.
func (pmuo *PubsubMessgaeUpdateOne) AddCreatedAt(u int32) *PubsubMessgaeUpdateOne {
	pmuo.mutation.AddCreatedAt(u)
	return pmuo
}

// SetUpdatedAt sets the "updated_at" field.
func (pmuo *PubsubMessgaeUpdateOne) SetUpdatedAt(u uint32) *PubsubMessgaeUpdateOne {
	pmuo.mutation.ResetUpdatedAt()
	pmuo.mutation.SetUpdatedAt(u)
	return pmuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (pmuo *PubsubMessgaeUpdateOne) AddUpdatedAt(u int32) *PubsubMessgaeUpdateOne {
	pmuo.mutation.AddUpdatedAt(u)
	return pmuo
}

// SetDeletedAt sets the "deleted_at" field.
func (pmuo *PubsubMessgaeUpdateOne) SetDeletedAt(u uint32) *PubsubMessgaeUpdateOne {
	pmuo.mutation.ResetDeletedAt()
	pmuo.mutation.SetDeletedAt(u)
	return pmuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pmuo *PubsubMessgaeUpdateOne) SetNillableDeletedAt(u *uint32) *PubsubMessgaeUpdateOne {
	if u != nil {
		pmuo.SetDeletedAt(*u)
	}
	return pmuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (pmuo *PubsubMessgaeUpdateOne) AddDeletedAt(u int32) *PubsubMessgaeUpdateOne {
	pmuo.mutation.AddDeletedAt(u)
	return pmuo
}

// SetUniqueID sets the "unique_id" field.
func (pmuo *PubsubMessgaeUpdateOne) SetUniqueID(u uuid.UUID) *PubsubMessgaeUpdateOne {
	pmuo.mutation.SetUniqueID(u)
	return pmuo
}

// SetMessageID sets the "message_id" field.
func (pmuo *PubsubMessgaeUpdateOne) SetMessageID(s string) *PubsubMessgaeUpdateOne {
	pmuo.mutation.SetMessageID(s)
	return pmuo
}

// SetSender sets the "sender" field.
func (pmuo *PubsubMessgaeUpdateOne) SetSender(s string) *PubsubMessgaeUpdateOne {
	pmuo.mutation.SetSender(s)
	return pmuo
}

// SetBody sets the "body" field.
func (pmuo *PubsubMessgaeUpdateOne) SetBody(b []byte) *PubsubMessgaeUpdateOne {
	pmuo.mutation.SetBody(b)
	return pmuo
}

// SetState sets the "state" field.
func (pmuo *PubsubMessgaeUpdateOne) SetState(s string) *PubsubMessgaeUpdateOne {
	pmuo.mutation.SetState(s)
	return pmuo
}

// SetResponseID sets the "response_id" field.
func (pmuo *PubsubMessgaeUpdateOne) SetResponseID(u uuid.UUID) *PubsubMessgaeUpdateOne {
	pmuo.mutation.SetResponseID(u)
	return pmuo
}

// SetErrorMessage sets the "error_message" field.
func (pmuo *PubsubMessgaeUpdateOne) SetErrorMessage(s string) *PubsubMessgaeUpdateOne {
	pmuo.mutation.SetErrorMessage(s)
	return pmuo
}

// SetNillableErrorMessage sets the "error_message" field if the given value is not nil.
func (pmuo *PubsubMessgaeUpdateOne) SetNillableErrorMessage(s *string) *PubsubMessgaeUpdateOne {
	if s != nil {
		pmuo.SetErrorMessage(*s)
	}
	return pmuo
}

// ClearErrorMessage clears the value of the "error_message" field.
func (pmuo *PubsubMessgaeUpdateOne) ClearErrorMessage() *PubsubMessgaeUpdateOne {
	pmuo.mutation.ClearErrorMessage()
	return pmuo
}

// Mutation returns the PubsubMessgaeMutation object of the builder.
func (pmuo *PubsubMessgaeUpdateOne) Mutation() *PubsubMessgaeMutation {
	return pmuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pmuo *PubsubMessgaeUpdateOne) Select(field string, fields ...string) *PubsubMessgaeUpdateOne {
	pmuo.fields = append([]string{field}, fields...)
	return pmuo
}

// Save executes the query and returns the updated PubsubMessgae entity.
func (pmuo *PubsubMessgaeUpdateOne) Save(ctx context.Context) (*PubsubMessgae, error) {
	var (
		err  error
		node *PubsubMessgae
	)
	if err := pmuo.defaults(); err != nil {
		return nil, err
	}
	if len(pmuo.hooks) == 0 {
		node, err = pmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PubsubMessgaeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pmuo.mutation = mutation
			node, err = pmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pmuo.hooks) - 1; i >= 0; i-- {
			if pmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pmuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pmuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PubsubMessgae)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PubsubMessgaeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pmuo *PubsubMessgaeUpdateOne) SaveX(ctx context.Context) *PubsubMessgae {
	node, err := pmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pmuo *PubsubMessgaeUpdateOne) Exec(ctx context.Context) error {
	_, err := pmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmuo *PubsubMessgaeUpdateOne) ExecX(ctx context.Context) {
	if err := pmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pmuo *PubsubMessgaeUpdateOne) defaults() error {
	if _, ok := pmuo.mutation.UpdatedAt(); !ok {
		if pubsubmessgae.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized pubsubmessgae.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := pubsubmessgae.UpdateDefaultUpdatedAt()
		pmuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pmuo *PubsubMessgaeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PubsubMessgaeUpdateOne {
	pmuo.modifiers = append(pmuo.modifiers, modifiers...)
	return pmuo
}

func (pmuo *PubsubMessgaeUpdateOne) sqlSave(ctx context.Context) (_node *PubsubMessgae, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pubsubmessgae.Table,
			Columns: pubsubmessgae.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pubsubmessgae.FieldID,
			},
		},
	}
	id, ok := pmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PubsubMessgae.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pubsubmessgae.FieldID)
		for _, f := range fields {
			if !pubsubmessgae.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pubsubmessgae.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessgae.FieldCreatedAt,
		})
	}
	if value, ok := pmuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessgae.FieldCreatedAt,
		})
	}
	if value, ok := pmuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessgae.FieldUpdatedAt,
		})
	}
	if value, ok := pmuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessgae.FieldUpdatedAt,
		})
	}
	if value, ok := pmuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessgae.FieldDeletedAt,
		})
	}
	if value, ok := pmuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessgae.FieldDeletedAt,
		})
	}
	if value, ok := pmuo.mutation.UniqueID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pubsubmessgae.FieldUniqueID,
		})
	}
	if value, ok := pmuo.mutation.MessageID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pubsubmessgae.FieldMessageID,
		})
	}
	if value, ok := pmuo.mutation.Sender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pubsubmessgae.FieldSender,
		})
	}
	if value, ok := pmuo.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: pubsubmessgae.FieldBody,
		})
	}
	if value, ok := pmuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pubsubmessgae.FieldState,
		})
	}
	if value, ok := pmuo.mutation.ResponseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pubsubmessgae.FieldResponseID,
		})
	}
	if value, ok := pmuo.mutation.ErrorMessage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pubsubmessgae.FieldErrorMessage,
		})
	}
	if pmuo.mutation.ErrorMessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pubsubmessgae.FieldErrorMessage,
		})
	}
	_spec.Modifiers = pmuo.modifiers
	_node = &PubsubMessgae{config: pmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pubsubmessgae.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
