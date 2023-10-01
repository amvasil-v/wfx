// SPDX-FileCopyrightText: The entgo authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by ent, DO NOT EDIT.

package history

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/siemens/wfx/generated/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.History {
	return predicate.History(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.History {
	return predicate.History(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.History {
	return predicate.History(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.History {
	return predicate.History(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.History {
	return predicate.History(sql.FieldLTE(FieldID, id))
}

// Mtime applies equality check predicate on the "mtime" field. It's identical to MtimeEQ.
func Mtime(v time.Time) predicate.History {
	return predicate.History(sql.FieldEQ(FieldMtime, v))
}

// MtimeEQ applies the EQ predicate on the "mtime" field.
func MtimeEQ(v time.Time) predicate.History {
	return predicate.History(sql.FieldEQ(FieldMtime, v))
}

// MtimeNEQ applies the NEQ predicate on the "mtime" field.
func MtimeNEQ(v time.Time) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldMtime, v))
}

// MtimeIn applies the In predicate on the "mtime" field.
func MtimeIn(vs ...time.Time) predicate.History {
	return predicate.History(sql.FieldIn(FieldMtime, vs...))
}

// MtimeNotIn applies the NotIn predicate on the "mtime" field.
func MtimeNotIn(vs ...time.Time) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldMtime, vs...))
}

// MtimeGT applies the GT predicate on the "mtime" field.
func MtimeGT(v time.Time) predicate.History {
	return predicate.History(sql.FieldGT(FieldMtime, v))
}

// MtimeGTE applies the GTE predicate on the "mtime" field.
func MtimeGTE(v time.Time) predicate.History {
	return predicate.History(sql.FieldGTE(FieldMtime, v))
}

// MtimeLT applies the LT predicate on the "mtime" field.
func MtimeLT(v time.Time) predicate.History {
	return predicate.History(sql.FieldLT(FieldMtime, v))
}

// MtimeLTE applies the LTE predicate on the "mtime" field.
func MtimeLTE(v time.Time) predicate.History {
	return predicate.History(sql.FieldLTE(FieldMtime, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.History {
	return predicate.History(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.History {
	return predicate.History(sql.FieldNotNull(FieldStatus))
}

// DefinitionIsNil applies the IsNil predicate on the "definition" field.
func DefinitionIsNil() predicate.History {
	return predicate.History(sql.FieldIsNull(FieldDefinition))
}

// DefinitionNotNil applies the NotNil predicate on the "definition" field.
func DefinitionNotNil() predicate.History {
	return predicate.History(sql.FieldNotNull(FieldDefinition))
}

// HasJob applies the HasEdge predicate on the "job" edge.
func HasJob() predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, JobTable, JobColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJobWith applies the HasEdge predicate on the "job" edge with a given conditions (other predicates).
func HasJobWith(preds ...predicate.Job) predicate.History {
	return predicate.History(func(s *sql.Selector) {
		step := newJobStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.History) predicate.History {
	return predicate.History(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.History) predicate.History {
	return predicate.History(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.History) predicate.History {
	return predicate.History(sql.NotPredicates(p))
}
