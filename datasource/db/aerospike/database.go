package aerospike

import "github.com/aerospike/aerospike-client-go"

type Database struct {
	Service        *aerospike.Client
	DatabaseName   string
	CollectionName string
}

func New(hostname string, port int) (*Database, error) {
	service, err := aerospike.NewClient(hostname, port)
	db := &Database{Service: service}
	return db, err
}

func (db *Database) UseDatabase(databaseName string) {
	db.DatabaseName = databaseName
}

func (db *Database) UseCollection(collectionName string) {
	db.CollectionName = collectionName
}

func (db *Database) Set(key string, value interface{}) error {
	newKey, err := aerospike.NewKey(db.DatabaseName, db.CollectionName, key)
	if err != nil {
		return err
	}
	/*bins := aerospike.BinMap{
		"bin1": 42,
		"bin2": "An elephant is a mouse with an operating system",
		"bin3": []interface{}{"Go", 2009},
	}*/
	err = db.Service.PutObject(nil, newKey, value)
	return err
}

func (db *Database) Close() {
	defer db.Service.Close()
}