package greasel

import (
	"bytes"
	"fmt"
)

// Query builder DSL

// Query is the abstract intermediate representation of a sql query
type Query struct {
	From *Table
	Filter Filter
	Joins []Join

	Aliases map[*Table]string
}

// From creates a new query
func From(table *Table) *Query {
	return &Query{
		From: table,
		Aliases: make(map[*Table]string),
	}
}

func (q *Query) Where(predicates ...Filter) *Query {
	q.Filter = And(predicates...) // Fixme
	return q
}

func (q *Query) InnerJoin(binding *Table, on ...Filter) *Query {
	// Panic if binding used in another join
	for _, j := range q.Joins {
		if binding == j.JoinTable {
			panic("binding was already joined")
		}
	}

	q.Joins = append(q.Joins, Join{
		Type:      "INNER",
		JoinTable: binding,
		Filter:    And(on...),
	})

	return q
}

func (q *Query) SQL() string {
	buf := bytes.Buffer{}

	buf.WriteString(fmt.Sprintf("FROM %s %s ", q.From.source, q.alias(q.From)))

	if q.Joins != nil {
		for _, join := range q.Joins {
			buf.WriteString(join.SQL(q))
		}
	}

	if q.Filter != nil {
		buf.WriteString(fmt.Sprintf("WHERE %s ", q.Filter.SQL(q)))
	}

	return buf.String()
}

func (q *Query) alias(table *Table) string {
	if _, ok := q.Aliases[table]; !ok {
		q.Aliases[table] = fmt.Sprintf("%c%d", table.source[0], len(q.Aliases))
	}
	return q.Aliases[table]
}

func (q *Query) falias(field *Field) string {
	return fmt.Sprintf("%s.%s", q.alias(field.parent), field.name)
}