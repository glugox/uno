package schema

import (
	"fmt"
	"reflect"
	"time"
)

// User
type User struct {
	Id        ObjectId  `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Roles     []Role    `json:"roles"`
}

func (u *User) Meta() TableMeta {
	return NewTableMeta("users")
}

func NewUser() (u *User) {
	u = &User{}
	return
}

// BaseFieldNames implements Model.BaseFieldNames
func (m *User) BaseFieldNames() []string {
	return []string{"Id", "FirstName", "LastName", "Email"}
}

// Collection ctrates new Collection based on Model
func (m *User) Collection() Collection {
	rfl := reflect.ValueOf(m)
	return NewCollectionBase(rfl)
}

// ToString implements Model.ToString
func (m *User) ToString() string {
	tmp := make(map[string]interface{})
	o := reflect.ValueOf(m)
	for _, f := range m.BaseFieldNames() {
		tmp[f] = o.Elem().FieldByName(f)
	}
	return fmt.Sprintf("[User %v]", tmp)
}

// Role
type Role struct {
	Id   ObjectId `json:"id"`
	Name string   `json:"name"`
}

func (*Role) Meta() TableMeta {
	return NewTableMeta("roles")
}

func NewRole() (u *Role) {
	u = &Role{}
	return
}

// BaseFieldNames implements Model.BaseFieldNames
func (m *Role) BaseFieldNames() []string {
	return []string{"Id", "FirstName", "LastName", "Email"}
}

// Collection ctrates new Collection based on Model
func (m *Role) Collection() Collection {
	rfl := reflect.ValueOf(m)
	return NewCollectionBase(rfl)
}

// ToString implements Model.ToString
func (m *Role) ToString() string {
	tmp := make(map[string]interface{})
	o := reflect.ValueOf(m)
	for _, f := range m.BaseFieldNames() {
		tmp[f] = o.Elem().FieldByName(f)
	}
	return fmt.Sprintf("[Role %v]", tmp)
}

// Menu
type Menu struct {
	Id        ObjectId   `json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	Label     string     `json:"label,omitempty"`
	Items     []MenuItem `json:"items" rel:"menu_items"`
}

// Name implements schema.Namer interface
func (m *Menu) Name() string {
	return "menus"
}

// Menu Item

type MenuItem struct {
	Id        ObjectId  `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	MenuId    ObjectId  `json:"menu_id"`
	ParentId  ObjectId  `json:"parent_id"`
	Label     string    `json:"label"`
	Path      string    `json:"path"`
}

// Name implements schema.Namer interface
func (m *MenuItem) Name() string {
	return "menu_items"
}
