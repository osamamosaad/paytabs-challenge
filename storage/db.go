package storage

import (
	"errors"

	"github.com/google/uuid"
	"github.com/osamamosaad/paytabs/entities"
)

type Storage struct {
	db        map[string]map[string]interface{}
	tableName string
}

var DBMemory map[string]map[string]interface{}

func New() *Storage {
	return &Storage{db: DBMemory}
}

func (s Storage) SetTableName(tableName string) *Storage {
	s.tableName = tableName
	return &s
}

func (s *Storage) FindById(ID string) (interface{}, error) {
	entity, ok := s.db[s.tableName][ID]

	if ok {
		return entity, nil
	}
	return nil, errors.New(s.tableName + " not found")
}

func (s *Storage) FindAll() (map[string]interface{}, error) {
	accountsList, ok := s.db[s.tableName]
	if ok {
		return accountsList, nil
	}
	return nil, errors.New(s.tableName + " not found")
}

func (s *Storage) Store(entity entities.EntityInterface) (interface{}, error) {

	if entity.GetID() == "" {
		entity.SetID(uuid.New().String())
	}

	_, ok := s.db[s.tableName]
	if !ok {
		s.db[s.tableName] = map[string]interface{}{}
	}

	s.db[s.tableName][entity.GetID()] = entity

	return s.FindById(entity.GetID())
}
