// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"ppgroup.i0sys.com/ent/listing"
	"ppgroup.i0sys.com/ent/predicate"
	"ppgroup.i0sys.com/ent/realtor"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RealtorUpdate is the builder for updating Realtor entities.
type RealtorUpdate struct {
	config
	hooks    []Hook
	mutation *RealtorMutation
}

// Where appends a list predicates to the RealtorUpdate builder.
func (ru *RealtorUpdate) Where(ps ...predicate.Realtor) *RealtorUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdateTime sets the "update_time" field.
func (ru *RealtorUpdate) SetUpdateTime(t time.Time) *RealtorUpdate {
	ru.mutation.SetUpdateTime(t)
	return ru
}

// SetFullName sets the "full_name" field.
func (ru *RealtorUpdate) SetFullName(s string) *RealtorUpdate {
	ru.mutation.SetFullName(s)
	return ru
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (ru *RealtorUpdate) SetNillableFullName(s *string) *RealtorUpdate {
	if s != nil {
		ru.SetFullName(*s)
	}
	return ru
}

// SetPhoto sets the "photo" field.
func (ru *RealtorUpdate) SetPhoto(m map[string]interface{}) *RealtorUpdate {
	ru.mutation.SetPhoto(m)
	return ru
}

// ClearPhoto clears the value of the "photo" field.
func (ru *RealtorUpdate) ClearPhoto() *RealtorUpdate {
	ru.mutation.ClearPhoto()
	return ru
}

// SetDescription sets the "description" field.
func (ru *RealtorUpdate) SetDescription(s string) *RealtorUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *RealtorUpdate) SetNillableDescription(s *string) *RealtorUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// ClearDescription clears the value of the "description" field.
func (ru *RealtorUpdate) ClearDescription() *RealtorUpdate {
	ru.mutation.ClearDescription()
	return ru
}

// SetPhone sets the "phone" field.
func (ru *RealtorUpdate) SetPhone(s string) *RealtorUpdate {
	ru.mutation.SetPhone(s)
	return ru
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (ru *RealtorUpdate) SetNillablePhone(s *string) *RealtorUpdate {
	if s != nil {
		ru.SetPhone(*s)
	}
	return ru
}

// SetEmail sets the "email" field.
func (ru *RealtorUpdate) SetEmail(s string) *RealtorUpdate {
	ru.mutation.SetEmail(s)
	return ru
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (ru *RealtorUpdate) SetNillableEmail(s *string) *RealtorUpdate {
	if s != nil {
		ru.SetEmail(*s)
	}
	return ru
}

// SetIsMvp sets the "is_mvp" field.
func (ru *RealtorUpdate) SetIsMvp(b bool) *RealtorUpdate {
	ru.mutation.SetIsMvp(b)
	return ru
}

// SetNillableIsMvp sets the "is_mvp" field if the given value is not nil.
func (ru *RealtorUpdate) SetNillableIsMvp(b *bool) *RealtorUpdate {
	if b != nil {
		ru.SetIsMvp(*b)
	}
	return ru
}

// AddListingIDs adds the "listings" edge to the Listing entity by IDs.
func (ru *RealtorUpdate) AddListingIDs(ids ...uuid.UUID) *RealtorUpdate {
	ru.mutation.AddListingIDs(ids...)
	return ru
}

// AddListings adds the "listings" edges to the Listing entity.
func (ru *RealtorUpdate) AddListings(l ...*Listing) *RealtorUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ru.AddListingIDs(ids...)
}

// Mutation returns the RealtorMutation object of the builder.
func (ru *RealtorUpdate) Mutation() *RealtorMutation {
	return ru.mutation
}

// ClearListings clears all "listings" edges to the Listing entity.
func (ru *RealtorUpdate) ClearListings() *RealtorUpdate {
	ru.mutation.ClearListings()
	return ru
}

// RemoveListingIDs removes the "listings" edge to Listing entities by IDs.
func (ru *RealtorUpdate) RemoveListingIDs(ids ...uuid.UUID) *RealtorUpdate {
	ru.mutation.RemoveListingIDs(ids...)
	return ru
}

// RemoveListings removes "listings" edges to Listing entities.
func (ru *RealtorUpdate) RemoveListings(l ...*Listing) *RealtorUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ru.RemoveListingIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RealtorUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RealtorUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RealtorUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RealtorUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RealtorUpdate) defaults() {
	if _, ok := ru.mutation.UpdateTime(); !ok {
		v := realtor.UpdateDefaultUpdateTime()
		ru.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RealtorUpdate) check() error {
	if v, ok := ru.mutation.FullName(); ok {
		if err := realtor.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "Realtor.full_name": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Description(); ok {
		if err := realtor.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Realtor.description": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Phone(); ok {
		if err := realtor.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Realtor.phone": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Email(); ok {
		if err := realtor.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Realtor.email": %w`, err)}
		}
	}
	return nil
}

func (ru *RealtorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(realtor.Table, realtor.Columns, sqlgraph.NewFieldSpec(realtor.FieldID, field.TypeUUID))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdateTime(); ok {
		_spec.SetField(realtor.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ru.mutation.FullName(); ok {
		_spec.SetField(realtor.FieldFullName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Photo(); ok {
		_spec.SetField(realtor.FieldPhoto, field.TypeJSON, value)
	}
	if ru.mutation.PhotoCleared() {
		_spec.ClearField(realtor.FieldPhoto, field.TypeJSON)
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.SetField(realtor.FieldDescription, field.TypeString, value)
	}
	if ru.mutation.DescriptionCleared() {
		_spec.ClearField(realtor.FieldDescription, field.TypeString)
	}
	if value, ok := ru.mutation.Phone(); ok {
		_spec.SetField(realtor.FieldPhone, field.TypeString, value)
	}
	if value, ok := ru.mutation.Email(); ok {
		_spec.SetField(realtor.FieldEmail, field.TypeString, value)
	}
	if value, ok := ru.mutation.IsMvp(); ok {
		_spec.SetField(realtor.FieldIsMvp, field.TypeBool, value)
	}
	if ru.mutation.ListingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   realtor.ListingsTable,
			Columns: []string{realtor.ListingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(listing.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedListingsIDs(); len(nodes) > 0 && !ru.mutation.ListingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   realtor.ListingsTable,
			Columns: []string{realtor.ListingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(listing.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ListingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   realtor.ListingsTable,
			Columns: []string{realtor.ListingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(listing.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{realtor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RealtorUpdateOne is the builder for updating a single Realtor entity.
type RealtorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RealtorMutation
}

// SetUpdateTime sets the "update_time" field.
func (ruo *RealtorUpdateOne) SetUpdateTime(t time.Time) *RealtorUpdateOne {
	ruo.mutation.SetUpdateTime(t)
	return ruo
}

// SetFullName sets the "full_name" field.
func (ruo *RealtorUpdateOne) SetFullName(s string) *RealtorUpdateOne {
	ruo.mutation.SetFullName(s)
	return ruo
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (ruo *RealtorUpdateOne) SetNillableFullName(s *string) *RealtorUpdateOne {
	if s != nil {
		ruo.SetFullName(*s)
	}
	return ruo
}

// SetPhoto sets the "photo" field.
func (ruo *RealtorUpdateOne) SetPhoto(m map[string]interface{}) *RealtorUpdateOne {
	ruo.mutation.SetPhoto(m)
	return ruo
}

// ClearPhoto clears the value of the "photo" field.
func (ruo *RealtorUpdateOne) ClearPhoto() *RealtorUpdateOne {
	ruo.mutation.ClearPhoto()
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RealtorUpdateOne) SetDescription(s string) *RealtorUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *RealtorUpdateOne) SetNillableDescription(s *string) *RealtorUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// ClearDescription clears the value of the "description" field.
func (ruo *RealtorUpdateOne) ClearDescription() *RealtorUpdateOne {
	ruo.mutation.ClearDescription()
	return ruo
}

// SetPhone sets the "phone" field.
func (ruo *RealtorUpdateOne) SetPhone(s string) *RealtorUpdateOne {
	ruo.mutation.SetPhone(s)
	return ruo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (ruo *RealtorUpdateOne) SetNillablePhone(s *string) *RealtorUpdateOne {
	if s != nil {
		ruo.SetPhone(*s)
	}
	return ruo
}

// SetEmail sets the "email" field.
func (ruo *RealtorUpdateOne) SetEmail(s string) *RealtorUpdateOne {
	ruo.mutation.SetEmail(s)
	return ruo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (ruo *RealtorUpdateOne) SetNillableEmail(s *string) *RealtorUpdateOne {
	if s != nil {
		ruo.SetEmail(*s)
	}
	return ruo
}

// SetIsMvp sets the "is_mvp" field.
func (ruo *RealtorUpdateOne) SetIsMvp(b bool) *RealtorUpdateOne {
	ruo.mutation.SetIsMvp(b)
	return ruo
}

// SetNillableIsMvp sets the "is_mvp" field if the given value is not nil.
func (ruo *RealtorUpdateOne) SetNillableIsMvp(b *bool) *RealtorUpdateOne {
	if b != nil {
		ruo.SetIsMvp(*b)
	}
	return ruo
}

// AddListingIDs adds the "listings" edge to the Listing entity by IDs.
func (ruo *RealtorUpdateOne) AddListingIDs(ids ...uuid.UUID) *RealtorUpdateOne {
	ruo.mutation.AddListingIDs(ids...)
	return ruo
}

// AddListings adds the "listings" edges to the Listing entity.
func (ruo *RealtorUpdateOne) AddListings(l ...*Listing) *RealtorUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ruo.AddListingIDs(ids...)
}

// Mutation returns the RealtorMutation object of the builder.
func (ruo *RealtorUpdateOne) Mutation() *RealtorMutation {
	return ruo.mutation
}

// ClearListings clears all "listings" edges to the Listing entity.
func (ruo *RealtorUpdateOne) ClearListings() *RealtorUpdateOne {
	ruo.mutation.ClearListings()
	return ruo
}

// RemoveListingIDs removes the "listings" edge to Listing entities by IDs.
func (ruo *RealtorUpdateOne) RemoveListingIDs(ids ...uuid.UUID) *RealtorUpdateOne {
	ruo.mutation.RemoveListingIDs(ids...)
	return ruo
}

// RemoveListings removes "listings" edges to Listing entities.
func (ruo *RealtorUpdateOne) RemoveListings(l ...*Listing) *RealtorUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ruo.RemoveListingIDs(ids...)
}

// Where appends a list predicates to the RealtorUpdate builder.
func (ruo *RealtorUpdateOne) Where(ps ...predicate.Realtor) *RealtorUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RealtorUpdateOne) Select(field string, fields ...string) *RealtorUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Realtor entity.
func (ruo *RealtorUpdateOne) Save(ctx context.Context) (*Realtor, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RealtorUpdateOne) SaveX(ctx context.Context) *Realtor {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RealtorUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RealtorUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RealtorUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateTime(); !ok {
		v := realtor.UpdateDefaultUpdateTime()
		ruo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RealtorUpdateOne) check() error {
	if v, ok := ruo.mutation.FullName(); ok {
		if err := realtor.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "Realtor.full_name": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Description(); ok {
		if err := realtor.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Realtor.description": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Phone(); ok {
		if err := realtor.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Realtor.phone": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Email(); ok {
		if err := realtor.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Realtor.email": %w`, err)}
		}
	}
	return nil
}

func (ruo *RealtorUpdateOne) sqlSave(ctx context.Context) (_node *Realtor, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(realtor.Table, realtor.Columns, sqlgraph.NewFieldSpec(realtor.FieldID, field.TypeUUID))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Realtor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, realtor.FieldID)
		for _, f := range fields {
			if !realtor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != realtor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdateTime(); ok {
		_spec.SetField(realtor.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.FullName(); ok {
		_spec.SetField(realtor.FieldFullName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Photo(); ok {
		_spec.SetField(realtor.FieldPhoto, field.TypeJSON, value)
	}
	if ruo.mutation.PhotoCleared() {
		_spec.ClearField(realtor.FieldPhoto, field.TypeJSON)
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.SetField(realtor.FieldDescription, field.TypeString, value)
	}
	if ruo.mutation.DescriptionCleared() {
		_spec.ClearField(realtor.FieldDescription, field.TypeString)
	}
	if value, ok := ruo.mutation.Phone(); ok {
		_spec.SetField(realtor.FieldPhone, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Email(); ok {
		_spec.SetField(realtor.FieldEmail, field.TypeString, value)
	}
	if value, ok := ruo.mutation.IsMvp(); ok {
		_spec.SetField(realtor.FieldIsMvp, field.TypeBool, value)
	}
	if ruo.mutation.ListingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   realtor.ListingsTable,
			Columns: []string{realtor.ListingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(listing.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedListingsIDs(); len(nodes) > 0 && !ruo.mutation.ListingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   realtor.ListingsTable,
			Columns: []string{realtor.ListingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(listing.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ListingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   realtor.ListingsTable,
			Columns: []string{realtor.ListingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(listing.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Realtor{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{realtor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
