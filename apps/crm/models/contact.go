package models

import (
	"time"

	"github.com/glugox/uno/pkg/schema"
)

type Contact struct {
	Id          schema.ObjectId `json:"id"`
	Birthdate   time.Time       `json:"birthdate"`
	CreatedAt   time.Time       `json:"-"`
	Department  string          `json:"department"`
	Description string          `json:"description"`
	Email       string          `json:"email"`
	FirstName   string          `json:"first_name"`
	LastName    string          `json:"last_name"`
	Languages   string          `json:"languages"`
	Phone       string          `json:"phone"`
	Picture     string          `json:"picture"`
	ReportsTo   schema.ObjectId `json:"reports_to"`
	UpdatedAt   time.Time       `json:"-"`
}
