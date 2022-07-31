package schema

import (
	"fmt"
	"reflect"
)

type Collection interface {
	// ModelReflect keeps data of underlying models.
	// This reflection is the meda of the model used to create
	// the collection itself, and always points to that value.
	// To genrate new Values, use CreateModel()
	ModelReflect() reflect.Value
	// CreateModel creates empty model struct for this collection
	CreateModel() reflect.Value
	// Append
	Append(Model)
	// AppendValue
	AppendValue(val reflect.Value)
	// Print
	Print()
	// Size
	Size() int
	// Items
	Items() []Model
	// Marshall to json only the wields we have
	Marshal() ([]byte, error)

	//ToReflection(r reflect.Value) Collection

	//GetStructName() string
}

// CollectionBase is the base implementation
// Collection interface
type CollectionBase struct {
	StructName string
	// reflection is the metadata of the model
	// we used to crate this collection
	reflection reflect.Value
	items      []Model
}

// NewCollectionBase creates new instance of basic Collection implementation
func NewCollectionBase(rfl reflect.Value) *CollectionBase {
	return &CollectionBase{
		StructName: rfl.Type().String(),
		reflection: rfl,
		items:      []Model{},
	}
}

// NewCollection creates new empty collection for the passed model
func NewCollection(m Model) Collection {
	return NewCollectionBase(reflect.ValueOf(m))
}

// ModelReflect implements Collection.ModelReflect
func (c *CollectionBase) ModelReflect() reflect.Value {
	return c.reflection
}

// CreateModel implements Collection.CreateModel
func (c *CollectionBase) CreateModel() reflect.Value {
	t := c.reflection.Elem().Type()
	n := reflect.New(t)
	return n
}

// Append implements Collection.Append
func (c *CollectionBase) Append(m Model) {
	c.items = append(c.items, m)
}

// Append implements Collection.Append
func (c *CollectionBase) AppendValue(val reflect.Value) {
	c.Append(val.Addr().Interface().(Model))
}

func (c *CollectionBase) Marshal() ([]byte, error) {

	fNames := []string{}
	if c.Size() > 0 {
		fNames = BaseFieldNames(c.items[0])
	}
	fmt.Printf("Fiald names to marshal: %s \n", fNames)
	return MarshalOnlyCollection(c, fNames)
}

/*func (c *CollectionBase) ToReflection(r reflect.Value) Collection {

	rc := NewCollectionBase(c.ModelReflect())
	// TODO:
	return rc
}*/

// Print
func (c *CollectionBase) Print() {
	fmt.Printf(c.StructName)
	fmt.Printf("[%s] {%d}: \n", c.StructName, c.Size())
	for _, i := range c.items {
		fmt.Printf(" |_ %s\n", ToString(i))
	}

}

// Size
func (c *CollectionBase) Size() int {
	return len(c.items)
}

// Items
func (c *CollectionBase) Items() []Model {
	return c.items
}
