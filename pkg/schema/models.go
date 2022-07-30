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

// Meta implements Model.Meta
func (m *Menu) Meta() TableMeta {
	return NewTableMeta("menus")
}

// BaseFieldNames implements Model.BaseFieldNames
func (m *Menu) BaseFieldNames() []string {
	return []string{"Id", "Label", "Items"}
}

// BeforeInsert implements Model.BeforeInsert
func (m *Menu) BeforeInsert() {
	//
}

// BeforeUpdate implements Model.BeforeUpdate
func (m *Menu) BeforeUpdate() {
	//
}

// Validate implements Model.Validate
func (m *Menu) Validate() bool {
	return true
}

// ToString implements Model.ToString
func (m *Menu) ToString() string {
	tmp := make(map[string]interface{})
	o := reflect.ValueOf(m)
	for _, f := range m.BaseFieldNames() {
		tmp[f] = o.Elem().FieldByName(f)
	}
	return fmt.Sprintf("[Menu %v]", tmp)
}

// Collection ctrates new Collection based on Model
func (m *Menu) Collection() Collection {
	rfl := reflect.ValueOf(m)
	return NewCollectionBase(rfl)
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

// Meta implements Model.Meta
func (m *MenuItem) Meta() TableMeta {
	return NewTableMeta("menu_items")
}

// BaseFieldNames implements Model.BaseFieldNames
func (m *MenuItem) BaseFieldNames() []string {
	return []string{"Id", "Label", "Path", "ParentId"}
}

// Collection ctrates new Collection based on Model
func (m *MenuItem) Collection() Collection {
	fmt.Println("menu Item Collection called!")
	rfl := reflect.ValueOf(m)
	return NewCollectionBase(rfl)
}

// ToString implements Model.ToString
func (m *MenuItem) ToString() string {
	tmp := make(map[string]interface{})
	o := reflect.ValueOf(m)
	for _, f := range m.BaseFieldNames() {
		tmp[f] = o.Elem().FieldByName(f)
	}
	return fmt.Sprintf("[MenuItem %v]", tmp)
}
