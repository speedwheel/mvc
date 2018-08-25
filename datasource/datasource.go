package datasource

import "github.com/speedwheel/mvc/datasource/db/aerospike"

type Engine uint32

const (
	Mongo Engine = iota
	Aerospike
)

type Database interface {
	Set(key string, value aerospike.RowObj) error
	UseDatabase(databaseName string)
	UseCollection(collectionName string)
	Close()
}

func NewDatabase(engine Engine) (Database, error) {
	if engine == Aerospike {
		return aerospike.New("127.0.0.1", 3000)
	}
	return nil, nil
}
