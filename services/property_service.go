package services

import (
	aerospike "github.com/aerospike/aerospike-client-go"
	"github.com/speedwheel/mvc/repositories"
)

type PropertyService interface {
	Create(property interface{}) error
	Select(key string) (*aerospike.Record, error)
}

func NewPropertyService(repo repositories.PropertyRepository) PropertyService {
	return &propertyService{
		repo: repo,
	}
}

type propertyService struct {
	repo repositories.PropertyRepository
}

func (s *propertyService) Create(property interface{}) error {
	return s.repo.Insert("1", property)
}

func (s *propertyService) Select(key string) (*aerospike.Record, error) {
	return s.repo.Select(key)
}
