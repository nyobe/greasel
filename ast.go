package greasel

import (
	"fmt"
	"strings"
)

// Intermediate ast representing a sql query

// Filters

type Expression interface {
	SQL(q *Query) string
	Params() []interface{}
}

type Filter interface {
	Expression
}

type CombinedFilter struct {
	Sep     string // AND, OR
	Filters []Filter
}

func And(filters ...Filter) Filter {
	return CombinedFilter{
		Sep: " AND ",
		Filters: filters,
	}
}

func Or(filters ...Filter) Filter {
	return CombinedFilter{
		Sep: " OR ",
		Filters: filters,
	}
}

func (f CombinedFilter) SQL(q *Query) string {
	clauses := make([]string, len(f.Filters))
	for i, filter := range f.Filters {
		clauses[i] = filter.SQL(q)
	}
	return "(" + strings.Join(clauses, f.Sep) + ") "
}

func (f CombinedFilter) Params() []interface{} {
	params := make([]interface{}, 0, len(f.Filters))
	for _, filter := range f.Filters {
		params = append(params, filter.Params()...)
	}
	return params
}

type ValueComparisonFilter struct {
	Lhs *Field
	Op  string
	Rhs interface{} // value, field, or expression(?)
}

func (f ValueComparisonFilter) SQL(q *Query) string {
	return fmt.Sprintf("%s %s ? ", q.falias(f.Lhs), f.Op)
}

func (f ValueComparisonFilter) Params() []interface{} {
	return []interface{}{f.Rhs}
}


type FieldComparisonFilter struct {
	Lhs *Field
	Op  string
	Rhs *Field
}

func (f FieldComparisonFilter) SQL(q *Query) string {
	return fmt.Sprintf("%s %s %s ", q.falias(f.Lhs), f.Op, q.falias(f.Rhs))
}

func (f FieldComparisonFilter) Params() []interface{} {
	return []interface{}{}
}

// Joins

type Join struct {
	Type string // Inner, Left
	//Field *Field  ??
	JoinTable *Table
	Filter    Filter
}

func (j Join) SQL(q *Query) string {
	return fmt.Sprintf("%s JOIN %s AS %s ON %s ", j.Type, j.JoinTable.source, q.alias(j.JoinTable), j.Filter.SQL(q))
}