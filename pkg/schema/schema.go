package schema

type Model interface {
	// Table
	//Meta() TableMeta
	// BaseFieldNames returns slice of field name
	// that are by default loadable from database.
	//BaseFieldNames() []string
	// ExtraFields accepts slice of field names we want
	// to return merged with BaseFields as FieldCol
	//ExtraFields(qF []string) *FieldCol
	//BeforeInsert()
	//BeforeUpdate()
	//Validate() bool
	// Collection
	//Collection() Collection
	//ToString() string
}

type Namer interface {
	Name() string
}
