package repositories

import (
	aerospike "github.com/aerospike/aerospike-client-go"
	"github.com/speedwheel/mvc/datasource"
)

type PropertyRepository interface {
	Insert(key string, property interface{}) error
	Select(key string) (*aerospike.Record, error)
}

func NewPropertyRepository(source datasource.Database) PropertyRepository {
	source.UseDatabase("test")
	source.UseCollection("properties")
	return &propertyDatabaseRepository{source: source}
}

type propertyDatabaseRepository struct {
	source datasource.Database
}

func (r *propertyDatabaseRepository) Insert(key string, property interface{}) error {
	return r.source.Set(key, property)
}

func (r *propertyDatabaseRepository) Select(key string) (*aerospike.Record, error) {
	return r.source.Get(key)
}
