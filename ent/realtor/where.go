// Code generated by ent, DO NOT EDIT.

package realtor

import (
	"time"

	"ppgroup.i0sys.com/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Realtor {
	return predicate.Realtor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Realtor {
	return predicate.Realtor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Realtor {
	return predicate.Realtor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Realtor {
	return predicate.Realtor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Realtor {
	return predicate.Realtor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Realtor {
	return predicate.Realtor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Realtor {
	return predicate.Realtor(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldUpdateTime, v))
}

// FullName applies equality check predicate on the "full_name" field. It's identical to FullNameEQ.
func FullName(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldFullName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldDescription, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldPhone, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldEmail, v))
}

// IsMvp applies equality check predicate on the "is_mvp" field. It's identical to IsMvpEQ.
func IsMvp(v bool) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldIsMvp, v))
}

// HireDate applies equality check predicate on the "hire_date" field. It's identical to HireDateEQ.
func HireDate(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldHireDate, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldLTE(FieldUpdateTime, v))
}

// FullNameEQ applies the EQ predicate on the "full_name" field.
func FullNameEQ(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldFullName, v))
}

// FullNameNEQ applies the NEQ predicate on the "full_name" field.
func FullNameNEQ(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldNEQ(FieldFullName, v))
}

// FullNameIn applies the In predicate on the "full_name" field.
func FullNameIn(vs ...string) predicate.Realtor {
	return predicate.Realtor(sql.FieldIn(FieldFullName, vs...))
}

// FullNameNotIn applies the NotIn predicate on the "full_name" field.
func FullNameNotIn(vs ...string) predicate.Realtor {
	return predicate.Realtor(sql.FieldNotIn(FieldFullName, vs...))
}

// FullNameGT applies the GT predicate on the "full_name" field.
func FullNameGT(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldGT(FieldFullName, v))
}

// FullNameGTE applies the GTE predicate on the "full_name" field.
func FullNameGTE(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldGTE(FieldFullName, v))
}

// FullNameLT applies the LT predicate on the "full_name" field.
func FullNameLT(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldLT(FieldFullName, v))
}

// FullNameLTE applies the LTE predicate on the "full_name" field.
func FullNameLTE(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldLTE(FieldFullName, v))
}

// FullNameContains applies the Contains predicate on the "full_name" field.
func FullNameContains(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldContains(FieldFullName, v))
}

// FullNameHasPrefix applies the HasPrefix predicate on the "full_name" field.
func FullNameHasPrefix(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldHasPrefix(FieldFullName, v))
}

// FullNameHasSuffix applies the HasSuffix predicate on the "full_name" field.
func FullNameHasSuffix(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldHasSuffix(FieldFullName, v))
}

// FullNameEqualFold applies the EqualFold predicate on the "full_name" field.
func FullNameEqualFold(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldEqualFold(FieldFullName, v))
}

// FullNameContainsFold applies the ContainsFold predicate on the "full_name" field.
func FullNameContainsFold(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldContainsFold(FieldFullName, v))
}

// PhotoIsNil applies the IsNil predicate on the "photo" field.
func PhotoIsNil() predicate.Realtor {
	return predicate.Realtor(sql.FieldIsNull(FieldPhoto))
}

// PhotoNotNil applies the NotNil predicate on the "photo" field.
func PhotoNotNil() predicate.Realtor {
	return predicate.Realtor(sql.FieldNotNull(FieldPhoto))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Realtor {
	return predicate.Realtor(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Realtor {
	return predicate.Realtor(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Realtor {
	return predicate.Realtor(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Realtor {
	return predicate.Realtor(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldContainsFold(FieldDescription, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Realtor {
	return predicate.Realtor(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Realtor {
	return predicate.Realtor(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldContainsFold(FieldPhone, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Realtor {
	return predicate.Realtor(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Realtor {
	return predicate.Realtor(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Realtor {
	return predicate.Realtor(sql.FieldContainsFold(FieldEmail, v))
}

// IsMvpEQ applies the EQ predicate on the "is_mvp" field.
func IsMvpEQ(v bool) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldIsMvp, v))
}

// IsMvpNEQ applies the NEQ predicate on the "is_mvp" field.
func IsMvpNEQ(v bool) predicate.Realtor {
	return predicate.Realtor(sql.FieldNEQ(FieldIsMvp, v))
}

// HireDateEQ applies the EQ predicate on the "hire_date" field.
func HireDateEQ(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldEQ(FieldHireDate, v))
}

// HireDateNEQ applies the NEQ predicate on the "hire_date" field.
func HireDateNEQ(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldNEQ(FieldHireDate, v))
}

// HireDateIn applies the In predicate on the "hire_date" field.
func HireDateIn(vs ...time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldIn(FieldHireDate, vs...))
}

// HireDateNotIn applies the NotIn predicate on the "hire_date" field.
func HireDateNotIn(vs ...time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldNotIn(FieldHireDate, vs...))
}

// HireDateGT applies the GT predicate on the "hire_date" field.
func HireDateGT(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldGT(FieldHireDate, v))
}

// HireDateGTE applies the GTE predicate on the "hire_date" field.
func HireDateGTE(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldGTE(FieldHireDate, v))
}

// HireDateLT applies the LT predicate on the "hire_date" field.
func HireDateLT(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldLT(FieldHireDate, v))
}

// HireDateLTE applies the LTE predicate on the "hire_date" field.
func HireDateLTE(v time.Time) predicate.Realtor {
	return predicate.Realtor(sql.FieldLTE(FieldHireDate, v))
}

// HasListings applies the HasEdge predicate on the "listings" edge.
func HasListings() predicate.Realtor {
	return predicate.Realtor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ListingsTable, ListingsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasListingsWith applies the HasEdge predicate on the "listings" edge with a given conditions (other predicates).
func HasListingsWith(preds ...predicate.Listing) predicate.Realtor {
	return predicate.Realtor(func(s *sql.Selector) {
		step := newListingsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Realtor) predicate.Realtor {
	return predicate.Realtor(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Realtor) predicate.Realtor {
	return predicate.Realtor(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Realtor) predicate.Realtor {
	return predicate.Realtor(sql.NotPredicates(p))
}
