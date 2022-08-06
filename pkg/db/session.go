package db

import (
	"fmt"

	"github.com/glugox/uno/pkg/schema"
)

type Session struct {
	DB    *DB
	model schema.Model
	query *schema.Query
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
	query := schema.NewQuery(schema.TableName(s.model), schema.QueryFields(s.model, schema.BaseFieldNames(s.model)))
	col, err := dbGet(s, query)
	if err != nil {
		return nil, err
	}
	return col, nil
}

// Where returns back the same session object
// and assigns query with passed where filter
func (s *Session) Where(key string, val string) *Session {
	if s.query == nil {
		s.query = schema.NewQuery(schema.TableName(s.model), schema.QueryFields(s.model, schema.BaseFieldNames(s.model)))
		s.query.AddWhere(key, val)
	}
	return s
}

// All returns all models from the database
func (s *Session) Get() (schema.Collection, error) {
	col, err := dbGet(s, s.query)
	if err != nil {
		return nil, err
	}
	return col, nil
}

// All returns all models from the database
func (s *Session) First() (schema.Model, error) {
	col, err := dbGet(s, s.query)
	if err != nil {
		return nil, err
	}
	return col.First(), nil
}

// Save inserts or updates model item into the database
func (s *Session) Save() error {
	return nil
}

func dbGet(s *Session, query *schema.Query) (schema.Collection, error) {
	// Create empty collection for our models
	db := s.DB
	m := s.model
	col := schema.NewCollection(m)
	// Fill col collection with records from DB
	err := db.Adapter.ScanCollection(col, query)
	if err != nil {
		return nil, err
	}

	s.DB.Logger.Info("Scaned collection with size: %d", col.Size())

	// Fill col collection with records from DB
	err = db.Adapter.ScanRelations(db.Schema, col, query)
	fmt.Println("Scanned collection: ")
	fmt.Printf("%+v \n", col)
	if err != nil {
		return nil, err
	}

	return col, nil
}
