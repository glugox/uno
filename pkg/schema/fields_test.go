package schema

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// DummyModel
type DummyModel struct {
	Id        ObjectId  `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
}

func (u *DummyModel) Meta() TableMeta {
	return NewTableMeta("users")
}

func TestFields(t *testing.T) {
	m := &DummyModel{}
	fields := QueryFields(m, []string{"Id", "FirstName"})
	assert.Equal(t, 2, fields.Size())
}

// BaseFieldNames implements Model.BaseFieldNames
func (m *DummyModel) BaseFieldNames() []string {
	return []string{"Id", "FirstName", "LastName", "Email"}
}

// Collection ctrates new Collection based on Model
func (m *DummyModel) Collection() Collection {
	rfl := reflect.ValueOf(m)
	return NewCollectionBase(rfl)
}

// ToString implements Model.ToString
func (m *DummyModel) ToString() string {
	tmp := make(map[string]interface{})
	o := reflect.ValueOf(m)
	for _, f := range m.BaseFieldNames() {
		tmp[f] = o.Elem().FieldByName(f)
	}
	return fmt.Sprintf("[DummyModel %v]", tmp)
}

func TestByName(t *testing.T) {
	m := &DummyModel{}

	fields := QueryFields(m, []string{"Id", "FirstName"})
	fFN, err := fields.ByName("FirstName")
	assert.NoError(t, err)
	assert.Equal(t, "FirstName", fFN.Name)

	fields = QueryFields(m, []string{"Email"})
	fFN, err = fields.ByName("FirstName")
	assert.EqualError(t, err, "could not find Field by name: FirstName")
	assert.Nil(t, fFN)
}
