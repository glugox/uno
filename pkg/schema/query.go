package schema

import (
	"fmt"
	"strings"
)

// Query is used to filter db records
type Query struct {
	Table   string
	Fields  *FieldCol
	Where   string
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

func (q *Query) ToSQL() string {
	return fmt.Sprintf("SELECT %s FROM %s%s%s", q.Fields.ToSqlString(), q.Table, q.FullWhere(), q.FullOrderBy())
}

// FullWhere will return empty string if we did not set the Where,
// and if we did, it will add prefix "WHERE " is needed.
func (q *Query) FullWhere() string {
	if q.Where == "" {
		return ""
	}
	newW := q.Where
	if !strings.HasPrefix(q.Where, "WHERE") {
		newW = "WHERE " + q.Where
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
