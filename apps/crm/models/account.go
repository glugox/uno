package models

import (
	"time"

	"github.com/glugox/uno/pkg/schema"
)

type Account struct {
	Id                schema.ObjectId `json:"id"`
	ParentId          schema.ObjectId `json:"parent_id"`
	ReportsTo         schema.ObjectId `json:"reports_to"`
	BillingAddress    string          `json:"billing_address"`
	BusinessHours     string          `json:"business_hours"`
	Contacts          []Contact       `json:"contacts"`
	CreatedAt         time.Time       `json:"-"`
	Description       string          `json:"description"`
	Email             string          `json:"email"`
	Industry          string          `json:"industry"`
	IsActive          bool            `json:"is_active"`
	Location          string          `json:"location"`
	Name              string          `json:"name"`
	NumberOfEmployees int             `json:"number_of_employees"`
	Phone             string          `json:"phone"`
	Picture           string          `json:"picture"`
	ShippingAddress   string          `json:"shipping_address"`
	UpdatedAt         time.Time       `json:"-"`
	Website           string          `json:"website"`
}
