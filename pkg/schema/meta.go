package schema

// TableMeta holds information about Model
type TableMeta struct {
	Name string
}

// NewTableMeta creates new TableMeta
func NewTableMeta(name string) TableMeta {
	return TableMeta{
		Name: name,
	}
}
