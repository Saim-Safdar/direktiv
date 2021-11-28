// Code generated by entc, DO NOT EDIT.

package varref

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Behaviour applies equality check predicate on the "behaviour" field. It's identical to BehaviourEQ.
func Behaviour(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBehaviour), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.VarRef {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VarRef(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.VarRef {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VarRef(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// BehaviourEQ applies the EQ predicate on the "behaviour" field.
func BehaviourEQ(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBehaviour), v))
	})
}

// BehaviourNEQ applies the NEQ predicate on the "behaviour" field.
func BehaviourNEQ(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBehaviour), v))
	})
}

// BehaviourIn applies the In predicate on the "behaviour" field.
func BehaviourIn(vs ...string) predicate.VarRef {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VarRef(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBehaviour), v...))
	})
}

// BehaviourNotIn applies the NotIn predicate on the "behaviour" field.
func BehaviourNotIn(vs ...string) predicate.VarRef {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VarRef(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBehaviour), v...))
	})
}

// BehaviourGT applies the GT predicate on the "behaviour" field.
func BehaviourGT(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBehaviour), v))
	})
}

// BehaviourGTE applies the GTE predicate on the "behaviour" field.
func BehaviourGTE(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBehaviour), v))
	})
}

// BehaviourLT applies the LT predicate on the "behaviour" field.
func BehaviourLT(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBehaviour), v))
	})
}

// BehaviourLTE applies the LTE predicate on the "behaviour" field.
func BehaviourLTE(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBehaviour), v))
	})
}

// BehaviourContains applies the Contains predicate on the "behaviour" field.
func BehaviourContains(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBehaviour), v))
	})
}

// BehaviourHasPrefix applies the HasPrefix predicate on the "behaviour" field.
func BehaviourHasPrefix(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBehaviour), v))
	})
}

// BehaviourHasSuffix applies the HasSuffix predicate on the "behaviour" field.
func BehaviourHasSuffix(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBehaviour), v))
	})
}

// BehaviourIsNil applies the IsNil predicate on the "behaviour" field.
func BehaviourIsNil() predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBehaviour)))
	})
}

// BehaviourNotNil applies the NotNil predicate on the "behaviour" field.
func BehaviourNotNil() predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBehaviour)))
	})
}

// BehaviourEqualFold applies the EqualFold predicate on the "behaviour" field.
func BehaviourEqualFold(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBehaviour), v))
	})
}

// BehaviourContainsFold applies the ContainsFold predicate on the "behaviour" field.
func BehaviourContainsFold(v string) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBehaviour), v))
	})
}

// HasVardata applies the HasEdge predicate on the "vardata" edge.
func HasVardata() predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VardataTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VardataTable, VardataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVardataWith applies the HasEdge predicate on the "vardata" edge with a given conditions (other predicates).
func HasVardataWith(preds ...predicate.VarData) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VardataInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VardataTable, VardataColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.Namespace) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkflow applies the HasEdge predicate on the "workflow" edge.
func HasWorkflow() predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowTable, WorkflowColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkflowWith applies the HasEdge predicate on the "workflow" edge with a given conditions (other predicates).
func HasWorkflowWith(preds ...predicate.Workflow) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkflowInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowTable, WorkflowColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInstance applies the HasEdge predicate on the "instance" edge.
func HasInstance() predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstanceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstanceTable, InstanceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstanceWith applies the HasEdge predicate on the "instance" edge with a given conditions (other predicates).
func HasInstanceWith(preds ...predicate.Instance) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstanceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstanceTable, InstanceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VarRef) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VarRef) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
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
func Not(p predicate.VarRef) predicate.VarRef {
	return predicate.VarRef(func(s *sql.Selector) {
		p(s.Not())
	})
}
