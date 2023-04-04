// Code generated by ent, DO NOT EDIT.

package pubsubmessage

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// MessageID applies equality check predicate on the "message_id" field. It's identical to MessageIDEQ.
func MessageID(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessageID), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// ResponseToID applies equality check predicate on the "response_to_id" field. It's identical to ResponseToIDEQ.
func ResponseToID(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResponseToID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// MessageIDEQ applies the EQ predicate on the "message_id" field.
func MessageIDEQ(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessageID), v))
	})
}

// MessageIDNEQ applies the NEQ predicate on the "message_id" field.
func MessageIDNEQ(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessageID), v))
	})
}

// MessageIDIn applies the In predicate on the "message_id" field.
func MessageIDIn(vs ...string) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMessageID), v...))
	})
}

// MessageIDNotIn applies the NotIn predicate on the "message_id" field.
func MessageIDNotIn(vs ...string) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMessageID), v...))
	})
}

// MessageIDGT applies the GT predicate on the "message_id" field.
func MessageIDGT(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessageID), v))
	})
}

// MessageIDGTE applies the GTE predicate on the "message_id" field.
func MessageIDGTE(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessageID), v))
	})
}

// MessageIDLT applies the LT predicate on the "message_id" field.
func MessageIDLT(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessageID), v))
	})
}

// MessageIDLTE applies the LTE predicate on the "message_id" field.
func MessageIDLTE(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessageID), v))
	})
}

// MessageIDContains applies the Contains predicate on the "message_id" field.
func MessageIDContains(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessageID), v))
	})
}

// MessageIDHasPrefix applies the HasPrefix predicate on the "message_id" field.
func MessageIDHasPrefix(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessageID), v))
	})
}

// MessageIDHasSuffix applies the HasSuffix predicate on the "message_id" field.
func MessageIDHasSuffix(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessageID), v))
	})
}

// MessageIDEqualFold applies the EqualFold predicate on the "message_id" field.
func MessageIDEqualFold(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessageID), v))
	})
}

// MessageIDContainsFold applies the ContainsFold predicate on the "message_id" field.
func MessageIDContainsFold(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessageID), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// ResponseToIDEQ applies the EQ predicate on the "response_to_id" field.
func ResponseToIDEQ(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResponseToID), v))
	})
}

// ResponseToIDNEQ applies the NEQ predicate on the "response_to_id" field.
func ResponseToIDNEQ(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResponseToID), v))
	})
}

// ResponseToIDIn applies the In predicate on the "response_to_id" field.
func ResponseToIDIn(vs ...uuid.UUID) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldResponseToID), v...))
	})
}

// ResponseToIDNotIn applies the NotIn predicate on the "response_to_id" field.
func ResponseToIDNotIn(vs ...uuid.UUID) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldResponseToID), v...))
	})
}

// ResponseToIDGT applies the GT predicate on the "response_to_id" field.
func ResponseToIDGT(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResponseToID), v))
	})
}

// ResponseToIDGTE applies the GTE predicate on the "response_to_id" field.
func ResponseToIDGTE(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResponseToID), v))
	})
}

// ResponseToIDLT applies the LT predicate on the "response_to_id" field.
func ResponseToIDLT(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResponseToID), v))
	})
}

// ResponseToIDLTE applies the LTE predicate on the "response_to_id" field.
func ResponseToIDLTE(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResponseToID), v))
	})
}

// ResponseToIDIsNil applies the IsNil predicate on the "response_to_id" field.
func ResponseToIDIsNil() predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldResponseToID)))
	})
}

// ResponseToIDNotNil applies the NotNil predicate on the "response_to_id" field.
func ResponseToIDNotNil() predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldResponseToID)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PubsubMessage) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PubsubMessage) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PubsubMessage) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		p(s.Not())
	})
}
