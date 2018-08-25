package repositories

import "github.com/speedwheel/mvc/datasource"

type PropertyRepository interface {
	Insert()
}

func NewPropertyRepository(source datasource.Database) PropertyRepository {
	source.UseDatabase("test")
	source.UseCollection("properties")
	return &propertyDatabaseRepository{source: source}
}

type propertyDatabaseRepository struct {
	source datasource.Database
}

func (r *propertyDatabaseRepository) Insert() {
	r.source.Set("1", []interface{}{1, 2})
}
