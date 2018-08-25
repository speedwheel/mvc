package repositories

import (
	"github.com/speedwheel/mvc/datasource"
	"github.com/speedwheel/mvc/datasource/db/aerospike"
)

type PropertyRepository interface {
	Insert(key string, property aerospike.RowObj) error
}

func NewPropertyRepository(source datasource.Database) PropertyRepository {
	source.UseDatabase("test")
	source.UseCollection("properties")
	return &propertyDatabaseRepository{source: source}
}

type propertyDatabaseRepository struct {
	source datasource.Database
}

func (r *propertyDatabaseRepository) Insert(key string, property aerospike.RowObj) error {
	return r.source.Set(key, property)
}
