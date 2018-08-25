package services

import (
	"github.com/speedwheel/mvc/repositories"
)

type PropertyService interface {
	Create(property interface{}) error
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
