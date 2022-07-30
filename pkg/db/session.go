package db

import (
	"github.com/glugox/uno/pkg/schema"
)

type Session struct {
	DB    *DB
	model schema.Model
}

func NewSession(db *DB, model schema.Model) *Session {
	return &Session{
		DB:    db,
		model: model,
	}
}

func NewEntitySession(db *DB, model schema.Model) *Session {
	return &Session{
		DB:    db,
		model: model,
	}
}

// All returns all models from the database
func (s *Session) All() (schema.Collection, error) {
	query := schema.NewQuery(s.model.Meta().Name, schema.QueryFields(s.model, s.model.BaseFieldNames()))
	col, err := dbGet(s, query)
	if err != nil {
		return nil, err
	}
	return col, nil
}

func dbGet(s *Session, query *schema.Query) (schema.Collection, error) {
	// Create empty collection for our models
	db := s.DB
	m := s.model
	col := m.Collection()

	// Fill col collection with records from DB
	err := db.Adapter.ScanCollection(col, query)
	if err != nil {
		return nil, err
	}

	// Fill col collection with records from DB
	err = db.Adapter.ScanRelations(db.Schema, col, query)
	if err != nil {
		return nil, err
	}

	return col, nil
}
