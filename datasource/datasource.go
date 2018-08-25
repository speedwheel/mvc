package datasource

import (
	aero "github.com/aerospike/aerospike-client-go"
	"github.com/speedwheel/mvc/datasource/db/aerospike"
)

type Engine uint32

const (
	Mongo Engine = iota
	Aerospike
)

type Database interface {
	Set(key string, value interface{}) error
	Get(key string) (*aero.Record, error)
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
