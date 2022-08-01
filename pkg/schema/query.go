package schema

import (
	"fmt"
	"strings"
)

// Query is used to filter db records
type Query struct {
	Table   string
	Fields  *FieldCol
	Where   *Where
	OrderBy string
}

// NewQuery returns a new instance of the Query
func NewQuery(table string, fields *FieldCol) (q *Query) {
	q = &Query{
		Table:  table,
		Fields: fields,
	}
	return
}

// Where represents single key value assignment filter
// e.g. "name" = "Uno"
type WhereOne struct {
	Key string
	Val string
}

// NewWhere creates new Where with passed operator
func NewWhereOne(key string, val string) *WhereOne {
	return &WhereOne{
		Key: key,
		Val: val,
	}
}

// WhereContext represents a block of multiple Where filters
// that can be joined by operator
// e.g. ("name" = "Uno" AND "active" = "true")
type WhereContext struct {
	Items    []*WhereOne
	Operator string
}

// AppendOne appends new one where item to the where context
func (c *WhereContext) AppendOne(o *WhereOne) {
	c.Items = append(c.Items, o)
}

// NewWhereContext creates new WhereContext with passed operator
func NewWhereContext(op string) *WhereContext {
	return &WhereContext{
		Operator: op,
		Items:    []*WhereOne{},
	}
}

// ToString returns tring representation of WhereContext
func (w *WhereContext) ToString() string {
	s := ""
	sArr := []string{}
	for _, o := range w.Items {
		sArr = append(sArr, o.ToString())
	}
	if len(sArr) > 0 {
		s = strings.Join(sArr, (" " + w.Operator + " "))
	}
	return s
}

// WhereContext is final toplevel filter
type Where struct {
	Contexts []*WhereContext
	Operator string
}

// NewWhere creates new Where with passed operator
func NewWhere(op string) *Where {
	return &Where{
		Operator: op,
		Contexts: []*WhereContext{},
	}
}

// AppendContext appends new context to the main where
func (w *Where) AppendContext(c *WhereContext) {
	w.Contexts = append(w.Contexts, c)
}

// ToString returns string representation of main where
func (w *Where) ToString() string {
	s := ""
	sArr := []string{}
	for _, c := range w.Contexts {
		sArr = append(sArr, c.ToString())
	}
	if len(sArr) > 0 {
		s = strings.Join(sArr, (" " + w.Operator + " "))
	}
	return s
}

// ToString returns tring representation of WhereOne
func (w *WhereOne) ToString() string {
	return fmt.Sprintf("%s = %q", w.Key, w.Val)
}

func (q *Query) ToSQL() string {
	return fmt.Sprintf("SELECT %s FROM %s%s%s", q.Fields.ToSqlString(), q.Table, q.FullWhere(), q.FullOrderBy())
}

// AddWhere appends new where condition ( AND )
func (q *Query) AddWhere(key string, val string) {
	wOne := NewWhereOne(key, val)

	// Add one to context
	c := NewWhereContext("AND")
	c.AppendOne(wOne)

	// Add context to main where
	q.Where = NewWhere("AND")
	q.Where.AppendContext(c)
}

// FullWhere will return empty string if we did not set the Where,
// and if we did, it will add prefix "WHERE " is needed.
func (q *Query) FullWhere() string {

	if q.Where == nil {
		return ""
	}
	newW := q.Where.ToString()
	if newW == "" {
		return ""
	}
	if !strings.HasPrefix(newW, "WHERE") {
		newW = "WHERE " + newW
	}
	if !strings.HasPrefix(newW, " ") {
		newW = " " + newW
	}
	return newW
}

// FullWhere will return empty string if we did not set the Where,
// and if we did, it will add prefix "WHERE " is needed.
func (q *Query) FullOrderBy() string {
	if q.OrderBy == "" {
		return ""
	}
	// New OrderBy
	newOB := ""
	if !strings.HasPrefix(q.OrderBy, "ORDER BY") {
		newOB = "ORDER BY " + q.OrderBy
	}
	if !strings.HasSuffix(q.OrderBy, "ASC") && !strings.HasSuffix(q.OrderBy, "DESC") {
		newOB = newOB + " ASC"
	}
	return newOB
}
