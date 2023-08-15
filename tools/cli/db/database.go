package db

import memdb "go-mem-db/lib"

var DB *memdb.KeyValueDB

func Init(dbname string) {
	DB = memdb.NewKeyValueDB(dbname)
}
