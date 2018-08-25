package services

import (
	"github.com/speedwheel/mvc/datasource/db/aerospike"
	"github.com/speedwheel/mvc/repositories"
)

type PropertyService interface {
	Create(property aerospike.RowObj) error
}

func NewPropertyService(repo repositories.PropertyRepository) PropertyService {
	return &propertyService{
		repo: repo,
	}
}

type propertyService struct {
	repo repositories.PropertyRepository
}

func (s *propertyService) Create(property aerospike.RowObj) error {
	return s.repo.Insert("1", property)
}
