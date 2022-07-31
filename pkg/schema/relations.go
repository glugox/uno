package schema

import (
	"encoding/json"
	"fmt"
)

type RelationType string

const (
	OneToOne   RelationType = "OneToOne"
	OneToMany  RelationType = "OneToMany"
	ManyToMany RelationType = "ManyToMany"
)

// Relation defines association between 2 fields
// from 2 tables
type Relation struct {
	Name  string
	Type  RelationType
	Table string
	Field string
}

// NewRelation creates new Relation
func NewRelation(name string) *Relation {
	return &Relation{
		Name: name,
	}
}

// RelationCol is collection of Fields
type RelationCol struct {
	items []*Relation
}

// NewFieldCol returns new instance of RelationCol
func NewRelationCol() *RelationCol {
	return &RelationCol{
		items: []*Relation{},
	}
}

// Append adds one Relation item to the relation collection
func (col *RelationCol) Append(r *Relation) {
	col.items = append(col.items, r)
}

// Items
func (col *RelationCol) Items() []*Relation {
	return col.items
}

// Size
func (col *RelationCol) Size() int {
	return len(col.items)
}

// At returns relation item At specified index
func (col *RelationCol) At(index int) (*Relation, error) {
	if col.Size() > index {
		return col.items[index], nil
	}
	return nil, fmt.Errorf("could not find Relation at index: %d", index)
}

// MarshalJSON
func (col *RelationCol) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["items"] = col.items
	return json.Marshal(m)
}
