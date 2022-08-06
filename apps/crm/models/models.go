package models

import "github.com/glugox/uno/pkg/schema"

// AllModels from CRM applications
func AllModels() []schema.Model {
	return []schema.Model{
		&Contact{},
		&Account{},
	}
}
